package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	
	"github.com/company/go-enterprise-template/internal/config"
	"github.com/company/go-enterprise-template/internal/handler"
	"github.com/company/go-enterprise-template/internal/repository"
	"github.com/company/go-enterprise-template/internal/service"
	"github.com/company/go-enterprise-template/pkg/database"
	"github.com/company/go-enterprise-template/pkg/logger"
)

// App 应用程序结构体
type App struct {
	config     *config.Config
	server     *http.Server
	router     *gin.Engine
}

// NewApp 创建应用实例
func NewApp(cfg *config.Config) *App {
	return &App{
		config: cfg,
	}
}

// Initialize 初始化应用程序
func (a *App) Initialize() error {
	// 初始化日志系统
	if err := logger.Init(&a.config.Log); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	// 初始化数据库
	if err := database.InitMySQL(&a.config.Database); err != nil {
		return fmt.Errorf("failed to initialize MySQL: %w", err)
	}

	if err := database.InitRedis(&a.config.Redis); err != nil {
		return fmt.Errorf("failed to initialize Redis: %w", err)
	}

	// 自动迁移数据库表
	if err := a.migrateDatabase(); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	// 设置Gin模式
	gin.SetMode(a.config.Server.Mode)

	// 初始化依赖注入
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()

	// 设置路由
	a.router = SetupRoutes(userHandler, healthHandler)

	// 创建HTTP服务器
	a.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", a.config.Server.Port),
		Handler:      a.router,
		ReadTimeout:  time.Duration(a.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(a.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(a.config.Server.IdleTimeout) * time.Second,
	}

	logger.Info("Application initialized successfully")
	return nil
}

// Run 运行应用程序
func (a *App) Run() error {
	// 启动服务器
	go func() {
		logger.Infof("Server starting on port %d", a.config.Server.Port)
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// 优雅关闭
	return a.Shutdown()
}

// Shutdown 优雅关闭应用程序
func (a *App) Shutdown() error {
	// 设置5秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭HTTP服务器
	if err := a.server.Shutdown(ctx); err != nil {
		logger.Errorf("Server forced to shutdown: %v", err)
		return err
	}

	// 关闭数据库连接
	if err := database.Close(); err != nil {
		logger.Errorf("Failed to close database: %v", err)
	}

	if err := database.CloseRedis(); err != nil {
		logger.Errorf("Failed to close Redis: %v", err)
	}

	logger.Info("Server exited")
	return nil
}

// migrateDatabase 自动迁移数据库表
func (a *App) migrateDatabase() error {
	// 这里可以添加自动迁移的代码
	// 例如使用GORM的AutoMigrate功能
	// return database.GetDB().AutoMigrate(&model.User{})
	
	logger.Info("Database migration skipped (implement as needed)")
	return nil
}

// GetRouter 获取路由器（用于测试）
func (a *App) GetRouter() *gin.Engine {
	return a.router
}