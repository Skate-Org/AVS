package monitor

import (
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type TaskStatus string

const (
	TaskStatus_COLLECTED     TaskStatus = "COLLECTED"
	TaskStatus_SAVED         TaskStatus = "SAVED"
	TaskStatus_SAVE_FAILED   TaskStatus = "SAVE_FAILED"
	TaskStatus_VERIFIED      TaskStatus = "VERIFIED"
	TaskStatus_VERIFY_FAILED TaskStatus = "VERIFY_FAILED"
)

type Metrics struct {
	*metrics.BaseMetrics
	TaskCollected *prometheus.CounterVec
}

const NAME_SPACE = "OPERATOR_monitor"

func NewMetrics(httpPort string, logger logging.Logger) *Metrics {
	logger = logger.With("process", NAME_SPACE)
	baseMetrics := metrics.NewMetrics(httpPort, logger)
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	return &Metrics{
		BaseMetrics: baseMetrics,
		TaskCollected: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: NAME_SPACE,
				Name:      "skateapp_tasks",
				Help:      "the number of on chain tasks collected from this operator",
			},
			[]string{"status"}, // status can be: COLLECTED | SAVED | SAVE_FAILED | VERIFIED | VERIFY_FAILED
		),
		// TODO: also include task that successfully went through AVS that this operator participated it (status=VERIFIED)
	}
}

// IncreaseTaskCollected increments the counter for the given TaskStatus.
func (m *Metrics) IncreaseTaskCollected(status TaskStatus) {
	m.TaskCollected.WithLabelValues(string(status)).Inc()
}
