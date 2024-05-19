package retrieve

import (
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type RetrieveStatus string

const (
	// Tasks that are submitted by operators with valid credentials (is an operator and signature is valid)
	RetrieveStatus_VALID RetrieveStatus = "VALID"
	// Tasks that are submitted by operators but the destination chain of the task is invalid.
	RetrieveStatus_INVALID_CHAIN RetrieveStatus = "INVALID_CHAIN"
	// Tasks that are submitted by operators but the submitter is not an operator.
	RetrieveStatus_INVALID_OPERATOR RetrieveStatus = "INVALID_OPERATOR"
	// Tasks that are submitted by operators but the signature payload is invalid.
	RetrieveStatus_INVALID_SIGNATURE RetrieveStatus = "INVALID_SIGNATURE"
	// Valid submitted tasks that are successfully saved to db
	RetrieveStatus_SAVED RetrieveStatus = "SAVED"
	// Valid submitted tasks but failed to save to db
	RetrieveStatus_SAVE_FAILED RetrieveStatus = "SAVE_FAILED"
)

type Metrics struct {
	*metrics.BaseMetrics
	TaskRetrieved *prometheus.CounterVec
}

const NAME_SPACE = "RELAYER"

func NewMetrics(httpPort string, logger *logging.Logger) *Metrics {
	logger = logger.With("process", NAME_SPACE)
	baseMetrics := metrics.NewMetrics(httpPort, logger)
	reg := baseMetrics.Registry

	m := &Metrics{
		BaseMetrics: baseMetrics,
		TaskRetrieved: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: NAME_SPACE,
				Name:      "retrieve_tasks",
				Help:      "the number of on chain tasks collected from this operator",
			},
			// see RetrieveStatus enum
			[]string{"retrieve_status"},
		),
	}

	return m
}

// IncreaseTaskRetrieved increments the counter for the given RetrieveStatus.
func (m *Metrics) IncreaseTaskRetrieved(status RetrieveStatus) {
	m.TaskRetrieved.WithLabelValues(string(status)).Inc()
}

func IncreaseTaskRetrieved(m *Metrics, status RetrieveStatus) {
	if m == nil {
		return
	}
	m.IncreaseTaskRetrieved(status)
}
