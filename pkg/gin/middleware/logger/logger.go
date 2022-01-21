package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vatusa/api2/pkg/vatlog"
)

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	c.Next()

	end := time.Now()
	latency := end.Sub(start)
	size := c.Writer.Size()
	status := c.Writer.Status()
	clientIP := c.ClientIP()
	method := c.Request.Method
	userAgent := c.Request.UserAgent()

	if vatlog.Format == "json" {
		l := vatlog.Logger.WithFields(logrus.Fields{
			"component": "gin",

			"status":     status,
			"method":     method,
			"path":       path,
			"ip":         clientIP,
			"latency":    latency,
			"size":       size,
			"user-agent": userAgent,
		})
		if len(c.Errors) > 0 {
			l.Error(c.Errors.String())
		} else {
			l.Info()
		}
	} else {
		l := vatlog.Logger.WithField("component", "gin")
		msg := fmt.Sprintf("%d %s %s %s %s %d %s",
			status,
			method,
			path,
			clientIP,
			latency,
			size,
			userAgent,
		)
		if len(c.Errors) > 0 {
			l.Errorf("%s - %s", msg, c.Errors.String())
		} else {
			l.Info(msg)
		}
	}
}
