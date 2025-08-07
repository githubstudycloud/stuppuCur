package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"github.com/company/go-enterprise-template/pkg/database"
)

// HealthHandler 健康检查处理器
type HealthHandler struct{}

// NewHealthHandler 创建健康检查处理器实例
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthResponse 健康检查响应
type HealthResponse struct {
	Status   string            `json:"status"`
	Service  string            `json:"service"`
	Version  string            `json:"version"`
	Checks   map[string]string `json:"checks"`
}

// Health 基础健康检查
// @Summary 健康检查
// @Description 检查服务基本状态
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=HealthResponse} "服务正常"
// @Router /health [get]
func (h *HealthHandler) Health(c *gin.Context) {
	Success(c, HealthResponse{
		Status:  "ok",
		Service: "go-enterprise-template",
		Version: "1.0.0",
	})
}

// Ready 就绪检查
// @Summary 就绪检查
// @Description 检查服务是否就绪（包括依赖服务）
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=HealthResponse} "服务就绪"
// @Failure 503 {object} Response "服务未就绪"
// @Router /ready [get]
func (h *HealthHandler) Ready(c *gin.Context) {
	checks := make(map[string]string)
	allHealthy := true

	// 检查MySQL连接
	if err := database.HealthCheck(); err != nil {
		checks["mysql"] = "unhealthy: " + err.Error()
		allHealthy = false
	} else {
		checks["mysql"] = "healthy"
	}

	// 检查Redis连接
	if err := database.RedisHealthCheck(); err != nil {
		checks["redis"] = "unhealthy: " + err.Error()
		allHealthy = false
	} else {
		checks["redis"] = "healthy"
	}

	response := HealthResponse{
		Service: "go-enterprise-template",
		Version: "1.0.0",
		Checks:  checks,
	}

	if allHealthy {
		response.Status = "ready"
		Success(c, response)
	} else {
		response.Status = "not ready"
		c.JSON(http.StatusServiceUnavailable, Response{
			Code:    http.StatusServiceUnavailable,
			Message: "service not ready",
			Data:    response,
		})
	}
}

// Live 存活检查
// @Summary 存活检查
// @Description 检查服务是否存活
// @Tags 健康检查
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=HealthResponse} "服务存活"
// @Router /live [get]
func (h *HealthHandler) Live(c *gin.Context) {
	Success(c, HealthResponse{
		Status:  "alive",
		Service: "go-enterprise-template",
		Version: "1.0.0",
	})
}