package app

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	
	"github.com/company/go-enterprise-template/internal/handler"
	"github.com/company/go-enterprise-template/internal/middleware"
)

// SetupRoutes 设置路由
func SetupRoutes(userHandler *handler.UserHandler, healthHandler *handler.HealthHandler) *gin.Engine {
	r := gin.New()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.GinLoggerMiddleware())
	r.Use(middleware.PrometheusMiddleware())
	r.Use(gin.Recovery())

	// 健康检查端点（无需认证）
	r.GET("/health", healthHandler.Health)
	r.GET("/ready", healthHandler.Ready)
	r.GET("/live", healthHandler.Live)

	// Prometheus监控端点
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API版本分组
	v1 := r.Group("/api/v1")

	// 认证相关路由（无需token）
	auth := v1.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
	}

	// 用户相关路由（需要token）
	user := v1.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", userHandler.GetProfile)
		user.PUT("/profile", userHandler.UpdateProfile)
	}

	// 用户管理路由（管理员功能，需要token）
	users := v1.Group("/users")
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}