package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

// Metrics registers Prometheus metrics and middleware.
type Metrics struct {
	httpRequestsTotal   *prometheus.CounterVec
	registry            *prometheus.Registry
	httpRequestDuration *prometheus.HistogramVec
}

// Handler returns the Gin middleware for tracking HTTP metrics.
func (m *Metrics) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip metrics endpoint to avoid double-counting
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next() // process request

		// Get the route pattern if available
		path := c.Request.URL.Path
		if c.FullPath() != "" {
			path = c.FullPath() // This gives the route pattern (e.g., "/user/:id")
		}

		status := strconv.Itoa(c.Writer.Status()) // Get numeric status as string

		// Record metrics after request is processed
		m.httpRequestsTotal.WithLabelValues(
			c.Request.Method,
			path,
			status,
		).Inc()

		duration := time.Since(start).Seconds()
		m.httpRequestDuration.WithLabelValues(
			c.Request.Method,
			path,
			status,
		).Observe(duration) // Records the duration in the histogram
	}
}

// MetricsHandler returns the `/metrics` endpoint handler.
func (m *Metrics) MetricsHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{}))
}

// NewMetrics creates a new Metrics instance with default Prometheus metrics.
func NewMetrics() *Metrics {
	reg := prometheus.NewRegistry()

	// register default metrics (Go runtime and process metrics)
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	// custom metrics
	httpRequestsTotal := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	reg.MustRegister(httpRequestsTotal)

	// Create the HistogramVec
	httpRequestDuration := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: []float64{0.1, 0.3, 0.5, 1, 2, 5}, // Custom buckets (in seconds)
		},
		[]string{"method", "path", "status"}, // Labels
	)
	reg.MustRegister(httpRequestDuration)

	return &Metrics{
		httpRequestsTotal:   httpRequestsTotal,
		httpRequestDuration: httpRequestDuration,
		registry:            reg,
	}
}
