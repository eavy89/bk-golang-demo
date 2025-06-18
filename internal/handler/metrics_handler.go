package handler

import (
	"backend-go-demo/internal/middleware/prometheus"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type MetricsHandler struct {
	metrics *prometheus.Metrics
}

func NewMetricsHandler(m *prometheus.Metrics) *MetricsHandler {
	return &MetricsHandler{metrics: m}
}

func (h *MetricsHandler) Handle() gin.HandlerFunc {
	return gin.WrapH(promhttp.HandlerFor(h.metrics.Registry, promhttp.HandlerOpts{}))
}
