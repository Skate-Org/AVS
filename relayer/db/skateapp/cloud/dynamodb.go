package cloud

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/relayer/db/skateapp"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
)

const SEPARATOR = "#"

var logger = logging.NewLoggerWithConsoleWriter()

type (
	SignedTask    = skateapp.SignedTask
	CompletedTask = skateapp.CompletedTask
)

type DynamoDBService struct {
	*dynamodb.Client
}

func NewDynamoDBService() *DynamoDBService {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		logger.Fatal("Unable to load SDK config", "error", err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	return &DynamoDBService{svc}
}

func Initialize() {
	svc := NewDynamoDBService()
	svc.createSignedTaskTable()
	svc.createCompletedTaskTable()
}

type DynamoDBStorage[T any] interface {
	Underlying() T
}

// //////////////////////////////////////////////////////////////

// /////////////////// SignedTask schema ////////////////////////
// //////////////////////////////////////////////////////////////
type DynamoDBSignedTask struct {
	SignedTask
	CompositeKey string
}

func (d *DynamoDBSignedTask) Underlying() SignedTask {
	return d.SignedTask
}

var _ DynamoDBStorage[SignedTask] = (*DynamoDBSignedTask)(nil)

func (svc *DynamoDBService) createSignedTaskTable() {
	_, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("SignedTask"),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("chainType"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("compositeKey"), KeyType: types.KeyTypeRange},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("chainType"), AttributeType: types.ScalarAttributeTypeN},
			{AttributeName: aws.String("compositeKey"), AttributeType: types.ScalarAttributeTypeS},
		},
		BillingMode: types.BillingModePayPerRequest, // Use on-demand pricing
	})
	if err != nil {
		logger.Fatal("Failed to create SignedTask table", "error", err)
	}
	logger.Info("SignedTask table created successfully!")
}

// InsertSignedTask insert a `SignedTask` object to the `SignedTask` schema.
// The primary key is determined by [chainType, (chainId, taskId, operator)].
// This function will override entry with duplicate primary key.
func (svc *DynamoDBService) InsertSignedTask(task SignedTask) error {
	compositeKey := CompositeKey(task.ChainId, task.TaskId, task.Operator)

	// Marshalling the task struct to a map
	item, err := attributevalue.MarshalMap(task)
	if err != nil {
		logger.Error("Failed to marshal input", "error", errors.Wrap(err, "InsertSignedTask"))
		return err
	}

	// Adding the composite key to the item map
	item["compositeKey"] = &types.AttributeValueMemberS{Value: compositeKey}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("SignedTask"),
		Item:      item,
	}

	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		logger.Error("Failed to put to db", "error", errors.Wrap(err, "InsertSignedTask"))
		return err
	}
	return nil
}

// QuerySignedTaskByChainId returns mapping of `chainId -> []SignedTask`
func (svc *DynamoDBService) QuerySignedTaskByChainId(chainType, chainId uint32) (map[uint32][]SignedTask, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("SignedTask"),
		KeyConditionExpression: aws.String("chainType = :ct AND begins_with(compositeKey, :ck)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ct": &types.AttributeValueMemberN{Value: CompositeKey(chainType)},
			":ck": &types.AttributeValueMemberS{Value: CompositeKey(chainId)},
		},
	}

	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		logger.Error("SVC query failed", "error", errors.Wrap(err, "QuerySignedTaskByChainId"))
		return nil, err
	}

	var dynamoDBTasks []DynamoDBSignedTask
	err = attributevalue.UnmarshalListOfMaps(result.Items, &dynamoDBTasks)
	if err != nil {
		logger.Error("Failed to unmarshal results: %s", "error", errors.Wrap(err, "QuerySignedTaskByChainId"))
		return nil, err
	}

	// Extract the underlying items of type T
	signedTasks := make([]SignedTask, 0, len(result.Items))
	for _, item := range dynamoDBTasks {
		signedTasks = append(signedTasks, item.Underlying())
	}

	// Group results by taskId
	tasksGroupByTaskId := make(map[uint32][]SignedTask)
	for _, task := range signedTasks {
		tasksGroupByTaskId[task.TaskId] = append(tasksGroupByTaskId[task.TaskId], task)
	}

	return tasksGroupByTaskId, nil
}

func (svc *DynamoDBService) QuerySignedTask(chainType, chainId, taskId uint32) ([]SignedTask, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("SignedTask"),
		KeyConditionExpression: aws.String("chainType = :ct AND begins_with(compositeKey, :ck)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ct": &types.AttributeValueMemberN{Value: CompositeKey(chainType)},
			":ck": &types.AttributeValueMemberS{Value: CompositeKey(chainId, taskId)},
		},
	}

	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		logger.Error("SVC query failed", "error", errors.Wrap(err, "QuerySignedTask"))
		return nil, err
	}
	var dynamoDBTasks []DynamoDBSignedTask
	err = attributevalue.UnmarshalListOfMaps(result.Items, &dynamoDBTasks)
	if err != nil {
		logger.Error("Failed to unmarshal results", "error", errors.Wrap(err, "QuerySignedTask"))
		return nil, err
	}

	// Extract the underlying items of type T
	signedTasks := make([]SignedTask, 0, len(result.Items))
	for _, item := range dynamoDBTasks {
		signedTasks = append(signedTasks, item.Underlying())
	}

	return signedTasks, nil
}

// //////////////////////////////////////////////////////////////

// /////////////////// CompletedTask schema /////////////////////
// //////////////////////////////////////////////////////////////
type DynamoDBCompletedTask struct {
	CompletedTask
	CompositeKey string
}

var _ DynamoDBStorage[CompletedTask] = (*DynamoDBCompletedTask)(nil)

func (d *DynamoDBCompletedTask) Underlying() CompletedTask {
	return d.CompletedTask
}

func (svc *DynamoDBService) createCompletedTaskTable() {
	_, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String("CompletedTask"),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("chainType"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("compositeKey"), KeyType: types.KeyTypeRange},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("chainType"), AttributeType: types.ScalarAttributeTypeN},
			{AttributeName: aws.String("compositeKey"), AttributeType: types.ScalarAttributeTypeS},
		},
		BillingMode: types.BillingModePayPerRequest, // Use on-demand pricing
	})
	if err != nil {
		logger.Fatal("Failed to create CompletedTask table", "error", err)
	}
	logger.Info("CompletedTask table created successfully!")
}

// InsertCompletedTask insert a `CompletedTask` object to the `CompletedTask` schema.
// The primary key is determined by [chainType, (chainId, taskId)].
// This function will override entry with duplicate primary key.
func (svc *DynamoDBService) InsertCompletedTask(task CompletedTask) error {
	compositeKey := CompositeKey(task.ChainId, task.TaskId)

	// Marshalling the task struct to a map
	item, err := attributevalue.MarshalMap(task)
	if err != nil {
		logger.Error("Failed to unmarshal input", "error", errors.Wrap(err, "InsertCompletedTask"))
		return err
	}

	// Adding the composite key to the item map
	item["compositeKey"] = &types.AttributeValueMemberS{Value: compositeKey}

	input := &dynamodb.PutItemInput{
		TableName: aws.String("SignedTask"),
		Item:      item,
	}

	_, err = svc.PutItem(context.TODO(), input)
	if err != nil {
		logger.Error("Failed to put to db", "error", errors.Wrap(err, "InsertCompletedTask"))
		return err
	}
	return nil
}

// QueryCompletedTaskByChain returns a mapping of `chainId -> []completedTasks`
func (svc *DynamoDBService) QueryCompletedTaskByChain(chainType, chainId uint32) (map[uint32][]CompletedTask, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("CompletedTask"),
		KeyConditionExpression: aws.String("chainType = :ct AND begins_with(compositeKey, :ck)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ct": &types.AttributeValueMemberN{Value: CompositeKey(chainType)},
			":ck": &types.AttributeValueMemberS{Value: CompositeKey(chainId)},
		},
	}

	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		logger.Error("SVC query failed", "error", errors.Wrap(err, "QueryCompletedTaskByChain"))
		return nil, err
	}

	var dynamoDBCompletedTasks []DynamoDBCompletedTask
	err = attributevalue.UnmarshalListOfMaps(result.Items, &dynamoDBCompletedTasks)
	if err != nil {
		logger.Error("Failed to unmarshal results", "error", errors.Wrap(err, "QueryCompletedTaskByChain"))
		return nil, err
	}

	// Extract the underlying items of type T
	completedTasks := make([]CompletedTask, 0, len(result.Items))
	for _, item := range dynamoDBCompletedTasks {
		completedTasks = append(completedTasks, item.Underlying())
	}

	// Group results by taskId
	completedTasksGroupById := make(map[uint32][]CompletedTask)
	for _, task := range completedTasks {
		completedTasksGroupById[task.TaskId] = append(completedTasksGroupById[task.TaskId], task)
	}

	return completedTasksGroupById, nil
}

// QueryCompletedTask returns exact completed task.
func (svc *DynamoDBService) QueryCompletedTask(chainType, chainId, taskId uint32) ([]CompletedTask, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String("CompletedTask"),
		KeyConditionExpression: aws.String("chainType = :ct AND begins_with(compositeKey, :ck)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ct": &types.AttributeValueMemberN{Value: CompositeKey(chainType)},
			":ck": &types.AttributeValueMemberS{Value: CompositeKey(chainId, taskId)},
		},
	}

	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		logger.Error("SVC query failed", "error", errors.Wrap(err, "QueryCompletedTask"))
		return nil, err
	}

	var dynamoDBTasks []DynamoDBCompletedTask
	err = attributevalue.UnmarshalListOfMaps(result.Items, &dynamoDBTasks)
	if err != nil {
		logger.Error("Failed to unmarshal results", "error", errors.Wrap(err, "QueryCompletedTask"))
		return nil, err
	}

	// Extract the underlying items of type T
	completedTasks := make([]CompletedTask, 0, len(result.Items))
	for _, item := range dynamoDBTasks {
		completedTasks = append(completedTasks, item.Underlying())
	}

	return completedTasks, nil
}

// ///////////////////////////////////////////////////////////////
// /////////////////// Composite attr helper /////////////////////

// CompositeKey constructs a composite key from string or uint32 inputs
// Calling on single arguments, give formatted version of itself.
func CompositeKey(params ...interface{}) string {
	var parts []string

	for _, param := range params {
		switch v := param.(type) {
		case uint32:
			parts = append(parts, fmt.Sprintf("%d", v))
		case string:
			parts = append(parts, v)
		default:
			// Optionally handle or ignore unsupported types
			continue
		}
	}

	return strings.Join(parts, SEPARATOR)
}

// ParseCompletedTaskCompositeKey splits the composite key back into chainType and chainId.
func ParseCompletedTaskCompositeKey(compositeKey string) (uint32, uint32, error) {
	parts := strings.Split(compositeKey, SEPARATOR)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid composite key format")
	}
	chainId, err1 := strconv.ParseUint(parts[0], 10, 32)
	taskId, err2 := strconv.ParseUint(parts[1], 10, 32)
	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing composite key elements")
	}
	return uint32(chainId), uint32(taskId), nil
}

// ParseSignedTaskCompositeKey splits the composite key back into chainType, chainId, and taskId.
func ParseSignedTaskCompositeKey(compositeKey string) (uint32, uint32, string, error) {
	parts := strings.Split(compositeKey, SEPARATOR)
	if len(parts) != 3 {
		return 0, 0, "", fmt.Errorf("invalid composite key format")
	}
	chainId, err1 := strconv.ParseUint(parts[0], 10, 32)
	taskId, err2 := strconv.ParseUint(parts[1], 10, 32)
	operator := parts[2]
	if err1 != nil || err2 != nil {
		return 0, 0, "", fmt.Errorf("error parsing composite key elements")
	}
	return uint32(chainId), uint32(taskId), operator, nil
}
