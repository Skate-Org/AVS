############################# GLOBAL COMMANDs #############################
extract_vcs_info: # Extract version control system information
	$(eval gitCommitHash=$(shell git rev-parse HEAD))
	$(eval gitCommitHashShort=$(shell git rev-parse --short=7 HEAD))
	$(eval gitTimestamp=$(shell git log -1 --format=%cI))

############################# HELP MESSAGE #############################
# Make sure the help command stays first, so that it's printed by default when `make` is called without arguments
.PHONY: help tests
help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


-----------------------------: ## 
___CONTRACTS___: ## Go binding for smart contracts
gen-contracts-bindings: ## generates contracts bindings
	cd contracts && ./gen_bindings.sh

-----------------------------: ## 
___API___: ## Go binding for GRPC server
gen-api-bindings: ## generate protobuf binding for grpc server
	cd api && ./gen_pb.sh


-----------------------------: ## 
___DOCKER___: ## Docker package management

# Variables
DOCKER_REGISTRY = ghcr.io/skate-org
KMS_IMAGE_NAME = skate-kms
OPERATOR_IMAGE_NAME = skate-operator
KMS_VERSION ?= 1.0.0
OPERATOR_VERSION ?= 0.1.0
LATEST_RELEASE ?= false

docker-login: ## authenticate to GitHub Container Registry
	@echo "Logging into GitHub Container Registry..."
	@echo $${GITHUB_TOKEN} | docker login $(DOCKER_REGISTRY) -u $${GITHUB_USERNAME} --password-stdin

docker-build-and-publish-all: docker-login docker-build-kms docker-publish-kms docker-clean-kms \ ## build and publish all docker images
	docker-build-operator docker-publish-operator docker-clean-operator

# KMS Docker targets
docker-build-kms: extract_vcs_info ## build KMS Docker Image with VCS info
	@echo "Building KMS Docker Image..."
	docker build --build-arg GIT_COMMIT=${gitCommitHash} --build-arg GIT_TIMESTAMP=${gitTimestamp} \
	-t $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${gitCommitHashShort} -f kms/Dockerfile .

docker-publish-kms: docker-build-kms ## publish KMS Docker Image
	@echo "Publishing KMS Docker Image..."
	docker tag $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${gitCommitHashShort} $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${KMS_VERSION}
	docker push $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${gitCommitHashShort}
	docker push $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${KMS_VERSION}
ifeq ($(LATEST_RELEASE),true)
	docker tag $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${gitCommitHashShort} $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):latest
endif

docker-clean-kms: ## clean KMS Docker Image
	@echo "Removing local KMS Docker Image..."
	docker rmi $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${gitCommitHashShort} $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):${KMS_VERSION}
ifeq ($(LATEST_RELEASE),true)
	docker rmi $(DOCKER_REGISTRY)/$(KMS_IMAGE_NAME):latest
endif

# Operator Docker targets
docker-build-operator: extract_vcs_info ## build Operator Docker Image with VCS info
	@echo "Building Operator Docker Image..."
	docker build --build-arg GIT_COMMIT=${gitCommitHash} --build-arg GIT_TIMESTAMP=${gitTimestamp} \
	-t $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${gitCommitHashShort} -f operator/Dockerfile .

docker-publish-operator: docker-build-operator ## publish Operator Docker Image
	@echo "Publishing Operator Docker Image..."
	docker tag $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${gitCommitHashShort} $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${OPERATOR_VERSION}
	docker push $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${gitCommitHashShort}
	docker push $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${OPERATOR_VERSION}
ifeq ($(LATEST_RELEASE),true)
	docker tag $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${gitCommitHashShort} $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):latest
	docker push $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):latest
endif

docker-clean-operator: ## clean Operator Docker Image
	@echo "Removing local Operator Docker Image..."
	docker rmi $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):${gitCommitHashShort}
ifeq ($(LATEST_RELEASE),true)
	docker rmi $(DOCKER_REGISTRY)/$(OPERATOR_IMAGE_NAME):latest
endif


-----------------------------: ## 
___BINARY___: ## Build binary from sources
build-all: build-operator build-relayer build-kms

build-operator: extract_vcs_info ## build the operator binary with version control metadata
	@go build -ldflags "-X 'main.Commit=${gitCommitHash}' -X 'main.Timestamp=${gitTimestamp}'" -o bin/operator ./operator
	@echo "Operator binary built in ./bin/operator"

build-relayer: extract_vcs_info ## build the relayer binary with version control metadata
	@go build -ldflags "-X 'main.Commit=${gitCommitHash}' -X 'main.Timestamp=${gitTimestamp}'" -o bin/relayer ./relayer
	@echo "Relayer binary built in ./bin/relayer"

build-kms: extract_vcs_info ## build the kms binary with version control metadata
	@go build -ldflags "-X 'main.Commit=${gitCommitHash}' -X 'main.Timestamp=${gitTimestamp}'" -o bin/kms ./kms
	@echo "Key Management Service (KMS) binary built in ./bin/kms"


-----------------------------: ## 
___PROCESSES___: ## Start process using go run from source
start-operators: ## start all skate operators
	cd operator && ./start.sh

start-relayer: ## start skate relayer
	cd relayer && ./start.sh


-----------------------------: ## 
_____HELPER_____: ## Utitiles for dev
mocks: ## generates mocks for tests
	go install go.uber.org/mock/mockgen@v0.4.0
	go generate ./...

run-unit-tests: ## runs all unit tests
	go test $$(go list ./... | grep -v /bindings | grep -v /pb) -coverprofile=coverage.out -covermode=atomic --timeout 15s
	go tool cover -html=coverage.out -o coverage.html

run-integration-test: ## runs integration test
	# TODO:
