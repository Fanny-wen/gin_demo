package Middlewares

import (
	"fmt"
	mylogrus "gin_demo/Services/logrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LogMiddlewareDemo() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Warning级别日志
		mylogrus.GinLogInstance.WithFields(logrus.Fields{
			"Method": context.Request.Method,
		}).Warning("Warning级别日志")

		// Error级别日志
		mylogrus.GinLogInstance.WithFields(logrus.Fields{
			"param-name": context.DefaultQuery("name", ""),
		}).Error("Error级别日志")

		// info级别日志
		mylogrus.GinLogInstance.WithFields(logrus.Fields{
			"code":   context.Writer.Status(),
			"url":    context.Request.URL.Path,
			"method": context.Request.Method,
		}).Info("Info级别日志")
		context.Next()
	}
}

// LogMiddleware Gin中间件函数，记录请求日志
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := fmt.Sprintf("%6v", endTime.Sub(startTime))
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		mylogrus.GinLogInstance.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info()
	}
}
