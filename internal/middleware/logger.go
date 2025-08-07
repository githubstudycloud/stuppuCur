package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	
	"github.com/company/go-enterprise-template/pkg/logger"
)

// LoggerMiddleware 请求日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 使用结构化日志记录请求信息
		logger.WithFields(logrus.Fields{
			"timestamp":   param.TimeStamp.Format(time.RFC3339),
			"status":      param.StatusCode,
			"latency":     param.Latency,
			"client_ip":   param.ClientIP,
			"method":      param.Method,
			"path":        param.Path,
			"user_agent":  param.Request.UserAgent(),
			"error":       param.ErrorMessage,
		}).Info("HTTP Request")
		
		return ""
	})
}

// GinLoggerMiddleware Gin框架专用的日志中间件
func GinLoggerMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		// 计算延迟时间
		latency := time.Since(start)

		// 获取状态码
		status := c.Writer.Status()

		// 构建日志字段
		fields := logrus.Fields{
			"status":     status,
			"method":     c.Request.Method,
			"path":       path,
			"query":      raw,
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
			"latency":    latency,
			"size":       c.Writer.Size(),
		}

		// 如果有用户信息，添加到日志中
		if userID, exists := c.Get("user_id"); exists {
			fields["user_id"] = userID
		}
		if username, exists := c.Get("username"); exists {
			fields["username"] = username
		}

		// 根据状态码选择日志级别
		switch {
		case status >= 500:
			logger.WithFields(fields).Error("HTTP Request")
		case status >= 400:
			logger.WithFields(fields).Warn("HTTP Request")
		default:
			logger.WithFields(fields).Info("HTTP Request")
		}
	})
}