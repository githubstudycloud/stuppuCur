# Go Enterprise Template Makefile

# 变量定义
APP_NAME = go-enterprise-template
VERSION = 1.0.0
BUILD_TIME = $(shell date +%Y-%m-%d\ %H:%M:%S)
COMMIT_SHA = $(shell git rev-parse --short HEAD)

# Go相关变量
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOMOD = $(GOCMD) mod
BINARY_DIR = bin
BINARY_NAME = $(APP_NAME)
MAIN_PATH = cmd/server/main.go

# Docker相关变量
DOCKER_IMAGE = $(APP_NAME)
DOCKER_TAG = $(VERSION)

# 默认目标
.PHONY: all
all: clean test build

# 清理构建文件
.PHONY: clean
clean:
	$(GOCLEAN)
	rm -rf $(BINARY_DIR)
	rm -rf logs
	rm -rf tmp

# 下载依赖
.PHONY: deps
deps:
	$(GOMOD) download
	$(GOMOD) tidy

# 代码格式化
.PHONY: fmt
fmt:
	$(GOCMD) fmt ./...

# 代码检查
.PHONY: vet
vet:
	$(GOCMD) vet ./...

# 运行测试
.PHONY: test
test:
	$(GOTEST) -v ./...

# 运行测试（带覆盖率）
.PHONY: test-coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# 构建应用
.PHONY: build
build:
	mkdir -p $(BINARY_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) \
		-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.commitSHA=$(COMMIT_SHA)" \
		-o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# 本地构建
.PHONY: build-local
build-local:
	mkdir -p $(BINARY_DIR)
	$(GOBUILD) \
		-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.commitSHA=$(COMMIT_SHA)" \
		-o $(BINARY_DIR)/$(BINARY_NAME) $(MAIN_PATH)

# 运行应用
.PHONY: run
run:
	$(GOCMD) run $(MAIN_PATH)

# 运行应用（使用开发配置）
.PHONY: run-dev
run-dev:
	$(GOCMD) run $(MAIN_PATH) -config=configs/config.dev.yaml

# 生成API文档
.PHONY: docs
docs:
	swag init -g $(MAIN_PATH) -o docs

# Docker构建
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .
	docker tag $(DOCKER_IMAGE):$(DOCKER_TAG) $(DOCKER_IMAGE):latest

# Docker运行
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

# Docker Compose启动
.PHONY: docker-up
docker-up:
	docker-compose up -d

# Docker Compose停止
.PHONY: docker-down
docker-down:
	docker-compose down

# Docker Compose重启
.PHONY: docker-restart
docker-restart: docker-down docker-up

# 查看Docker Compose日志
.PHONY: docker-logs
docker-logs:
	docker-compose logs -f

# 数据库迁移
.PHONY: migrate-up
migrate-up:
	migrate -path ./migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/enterprise_db" up

# 数据库回滚
.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations -database "mysql://root:rootpassword@tcp(localhost:3306)/enterprise_db" down

# 生成模拟数据
.PHONY: mock
mock:
	mockgen -source=internal/repository/user_repository.go -destination=mocks/mock_user_repository.go
	mockgen -source=internal/service/user_service.go -destination=mocks/mock_user_service.go

# 性能测试
.PHONY: bench
bench:
	$(GOTEST) -bench=. -benchmem ./...

# 代码质量检查
.PHONY: lint
lint:
	golangci-lint run

# 安全检查
.PHONY: security
security:
	gosec ./...

# 完整的CI检查
.PHONY: ci
ci: deps fmt vet lint security test

# 生产环境部署
.PHONY: deploy
deploy: ci build docker-build
	@echo "Ready for deployment"

# 开发环境设置
.PHONY: dev-setup
dev-setup:
	@echo "Setting up development environment..."
	$(GOMOD) download
	@echo "Installing development tools..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Development environment setup complete"

# 帮助信息
.PHONY: help
help:
	@echo "Go Enterprise Template Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all             Clean, test and build"
	@echo "  clean           Clean build files"
	@echo "  deps            Download dependencies"
	@echo "  fmt             Format code"
	@echo "  vet             Run go vet"
	@echo "  test            Run tests"
	@echo "  test-coverage   Run tests with coverage"
	@echo "  build           Build application"
	@echo "  build-local     Build for local OS"
	@echo "  run             Run application"
	@echo "  run-dev         Run with dev config"
	@echo "  docs            Generate API docs"
	@echo "  docker-build    Build Docker image"
	@echo "  docker-run      Run Docker container"
	@echo "  docker-up       Start docker-compose"
	@echo "  docker-down     Stop docker-compose"
	@echo "  docker-restart  Restart docker-compose"
	@echo "  docker-logs     View docker-compose logs"
	@echo "  lint            Run linter"
	@echo "  security        Run security checks"
	@echo "  ci              Run all CI checks"
	@echo "  deploy          Deploy to production"
	@echo "  dev-setup       Setup development environment"
	@echo "  help            Show this help message"