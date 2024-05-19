package metrics

import (
	"fmt"
	"net/http"

	"github.com/Skate-Org/AVS/lib/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type BaseMetrics struct {
	httpPort string
	Logger   *logging.Logger
	Registry *prometheus.Registry
}

func NewMetrics(httpPort string, logger *logging.Logger) *BaseMetrics {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	reg.MustRegister(collectors.NewGoCollector())

	metrics := &BaseMetrics{
		Registry: reg,
		httpPort: httpPort,
		Logger:   logger,
	}
	return metrics
}

func (g *BaseMetrics) Start() {
	g.Logger.Info("Starting metrics server at ", "port", g.httpPort)
	addr := fmt.Sprintf(":%s", g.httpPort)
	go func() {
		log := g.Logger
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(
			g.Registry,
			promhttp.HandlerOpts{},
		))
		err := http.ListenAndServe(addr, mux)
		log.Error("Prometheus server failed", "err", err)
	}()
}
