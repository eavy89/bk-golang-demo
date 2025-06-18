package prometheus

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func (m *Metrics) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		status := strconv.Itoa(c.Writer.Status())

		m.HTTPRequestsTotal.WithLabelValues(
			c.Request.Method,
			path,
			status,
		).Inc()

		m.HTTPRequestDuration.WithLabelValues(
			c.Request.Method,
			path,
			status,
		).Observe(time.Since(start).Seconds())
	}
}
