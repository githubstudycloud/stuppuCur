package main

import (
	"flag"
	"log"

	"github.com/company/go-enterprise-template/internal/app"
	"github.com/company/go-enterprise-template/internal/config"
)

// @title Go Enterprise Template API
// @version 1.0
// @description 企业级Go项目模板API文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer token格式: "Bearer {token}"

func main() {
	// 解析命令行参数
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建应用实例
	application := app.NewApp(cfg)

	// 初始化应用
	if err := application.Initialize(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// 运行应用
	if err := application.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}