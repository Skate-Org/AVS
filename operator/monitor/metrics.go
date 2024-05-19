package monitor

import (
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type TaskStatus string

const (
  // Tasks that are detected from Skate App contracts
	TaskStatus_DETECTED      TaskStatus = "DETECTED"
  // Tasks that are detected then saved to local db
	TaskStatus_SAVED         TaskStatus = "SAVED"
  // Tasks that are detected but failed to save to local db
	TaskStatus_SAVE_FAILED   TaskStatus = "SAVE_FAILED"
  // Tasks that are detected then successfully sent to relayer for verification
	TaskStatus_VERIFIED      TaskStatus = "VERIFIED"
  // Tasks that are detected but failed to send to relayer for verification
	TaskStatus_VERIFY_FAILED TaskStatus = "VERIFY_FAILED"
)

type Metrics struct {
	*metrics.BaseMetrics
	TaskProcessed *prometheus.CounterVec
}

const NAME_SPACE = "OPERATOR"

func NewMetrics(httpPort string, logger *logging.Logger) *Metrics {
	logger = logger.With("process", NAME_SPACE)
	baseMetrics := metrics.NewMetrics(httpPort, logger)
  reg := baseMetrics.Registry

	m := &Metrics{
		BaseMetrics: baseMetrics,
		TaskProcessed: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: NAME_SPACE,
				Name:      "monitor_tasks",
				Help:      "Number of task processed by operator's monitor server",
			},
			[]string{"status"}, // status can be: DETECTED | SAVED | SAVE_FAILED | VERIFIED | VERIFY_FAILED
		),
		// TODO: also include task that successfully went through AVS that this operator participated it (status=VERIFIED)
	}

	return m
}

// IncreaseTaskProcessed increments the counter for the given TaskStatus.
func (m *Metrics) IncreaseTaskProcessed(status TaskStatus) {
	m.TaskProcessed.WithLabelValues(string(status)).Inc()
}

func IncreaseTaskProcessed(m *Metrics, status TaskStatus) {
	if m == nil {
		return
	}
	m.IncreaseTaskProcessed(status)
}
