package publish

import (
	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/Skate-Org/AVS/lib/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type PublishStatus string

const (
	// AVS quorums that are sent for verification
	PublishStatus_SENT PublishStatus = "SENT"
	// Sent that returns with confirmed tx receipt
	PublishStatus_CONFIRMED PublishStatus = "CONFIRMED"
)

type Metrics struct {
	*metrics.BaseMetrics
	AVSPublished     *prometheus.CounterVec
	GatewayPublished *prometheus.CounterVec
}

const NAME_SPACE = "RELAYER"

func NewMetrics(httpPort string, logger *logging.Logger) *Metrics {
	logger = logger.With("process", NAME_SPACE)
	baseMetrics := metrics.NewMetrics(httpPort, logger)
	reg := baseMetrics.Registry

	m := &Metrics{
		BaseMetrics: baseMetrics,
		AVSPublished: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: NAME_SPACE,
				Name:      "publish_to_avs",
				Help:      "Number of verified quorums that this relayer send to AVS",
			},
			[]string{"status"}, // SENT | CONFIRMED
		),
		GatewayPublished: promauto.With(reg).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: NAME_SPACE,
				Name:      "publish_to_gateway",
				Help:      "Number of messages need to be relayed to destination gateway contracts",
			},
			[]string{"gateway_status"}, //
		),
	}

	return m
}

// IncreaseAVSPublished increments the counter for the given PublishStatus.
func (m *Metrics) IncreaseAVSPublished(status PublishStatus) {
	m.AVSPublished.WithLabelValues(string(status)).Inc()
}

// Wrapper for `Metrics.IncreaseAVSPublished`
func IncreaseAVSPublished(m *Metrics, status PublishStatus) {
	if m == nil {
		return
	}
	m.IncreaseAVSPublished(status)
}

// AddAVSPublished adds the counter for the given PublishStatus by arbitrary non-negative value.
func (m *Metrics) AddAVSPublished(status PublishStatus, by float64) {
	m.AVSPublished.WithLabelValues(string(status)).Add(by)
}

// Wrapper for `Metrics.AddAVSPublished`
func AddAVSPublished(m *Metrics, status PublishStatus, by float64) {
	if m == nil {
		return
	}
	m.AddAVSPublished(status, by)
}

// IncreaseGatewayPublished increments the counter for the given PublishStatus.
func (m *Metrics) IncreaseGatewayPublished(status PublishStatus) {
	m.GatewayPublished.WithLabelValues(string(status)).Inc()
}

// Wrapper for `Metrics.IncreaseGatewayPublished`
func IncreaseGatewayPublished(m *Metrics, status PublishStatus) {
	if m == nil {
		return
	}
	m.IncreaseAVSPublished(status)
}

// AddGatewayPublished adds the counter for the given PublishStatus by arbitrary non-negative value.
func (m *Metrics) AddGatewayPublished(status PublishStatus, by float64) {
	m.GatewayPublished.WithLabelValues(string(status)).Add(by)
}

// Wrapper for `Metrics.AddGatewayPublished`
func AddGatewayPublished(m *Metrics, status PublishStatus, by float64) {
	if m == nil {
		return
	}
	m.AddGatewayPublished(status, by)
}
