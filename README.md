# Go Enterprise Template

一个现代化的企业级Go项目基座，采用分层架构设计，集成了最新的技术栈和最佳实践。

## 🚀 特性

- ⚡ **高性能**: 基于Gin框架，优化的中间件和路由设计
- 🏗️ **分层架构**: Repository-Service-Handler三层架构，职责清晰
- 🔐 **安全认证**: JWT令牌认证，密码加密存储
- 📊 **监控可观测**: Prometheus指标收集，Grafana可视化，Jaeger链路追踪
- 🗄️ **数据存储**: MySQL主数据库，Redis缓存，GORM ORM
- 📝 **结构化日志**: 基于logrus的结构化日志，支持多种输出格式
- 🐳 **容器化**: Docker和Docker Compose支持
- 🔄 **CI/CD**: GitHub Actions自动化构建、测试、部署
- 📖 **API文档**: Swagger自动生成API文档
- 🧪 **测试覆盖**: 单元测试、集成测试、性能测试
- ⚙️ **配置管理**: 多环境配置，环境变量支持

## 📁 项目结构

```
├── cmd/                    # 应用程序入口
│   └── server/            # 服务器启动代码
├── internal/              # 私有应用代码
│   ├── app/               # 应用程序配置和初始化
│   ├── config/            # 配置管理
│   ├── handler/           # HTTP处理器（Controller层）
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   ├── repository/        # 数据访问层
│   ├── service/           # 业务逻辑层
│   └── utils/             # 内部工具函数
├── pkg/                   # 可共享的库代码
│   ├── database/          # 数据库连接和管理
│   ├── logger/            # 日志工具
│   └── validator/         # 验证器
├── api/                   # API定义文件
├── configs/               # 配置文件
├── deployments/           # 部署配置
├── docs/                  # 文档
├── scripts/               # 脚本文件
├── test/                  # 测试文件
├── .github/workflows/     # GitHub Actions工作流
├── docker-compose.yml     # Docker Compose配置
├── Dockerfile            # Docker镜像构建文件
└── Makefile              # 构建和开发工具
```

## 🛠️ 技术栈

### 后端框架
- **Gin**: 高性能HTTP Web框架
- **GORM**: 强大的ORM库
- **Viper**: 配置管理
- **Logrus**: 结构化日志

### 数据库
- **MySQL**: 主数据库
- **Redis**: 缓存和会话存储

### 监控和可观测性
- **Prometheus**: 指标收集
- **Grafana**: 监控面板
- **Jaeger**: 分布式链路追踪

### 开发工具
- **Swag**: API文档生成
- **golangci-lint**: 代码质量检查
- **gosec**: 安全扫描
- **testify**: 测试框架

## 🚀 快速开始

### 环境要求

- Go 1.22+
- Docker & Docker Compose
- MySQL 8.0+
- Redis 7+

### 安装和运行

1. **克隆项目**
```bash
git clone <repository-url>
cd go-enterprise-template
```

2. **设置开发环境**
```bash
make dev-setup
```

3. **启动服务（使用Docker Compose）**
```bash
make docker-up
```

4. **本地开发运行**
```bash
# 复制配置文件
cp configs/config.dev.yaml configs/config.yaml

# 运行应用
make run-dev
```

### 访问服务

- **应用程序**: http://localhost:8080
- **API文档**: http://localhost:8080/swagger/index.html
- **健康检查**: http://localhost:8080/health
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin123)
- **Jaeger**: http://localhost:16686

## 📚 API文档

### 认证相关

#### 用户注册
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "nickname": "测试用户"
  }'
```

#### 用户登录
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### 用户管理

#### 获取用户资料
```bash
curl -X GET http://localhost:8080/api/v1/user/profile \
  -H "Authorization: Bearer <your-token>"
```

#### 更新用户资料
```bash
curl -X PUT http://localhost:8080/api/v1/user/profile \
  -H "Authorization: Bearer <your-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "nickname": "新昵称"
  }'
```

## 🔧 配置说明

### 环境变量

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `APP_SERVER_PORT` | 服务端口 | `8080` |
| `APP_SERVER_MODE` | 运行模式 | `debug` |
| `APP_DATABASE_HOST` | 数据库主机 | `localhost` |
| `APP_DATABASE_PORT` | 数据库端口 | `3306` |
| `APP_REDIS_HOST` | Redis主机 | `localhost` |
| `APP_REDIS_PORT` | Redis端口 | `6379` |

### 配置文件

配置文件位于 `configs/` 目录下：
- `config.yaml`: 生产环境配置
- `config.dev.yaml`: 开发环境配置

## 🧪 测试

### 运行测试
```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 运行性能测试
make bench
```

### 代码质量检查
```bash
# 代码格式化
make fmt

# 代码检查
make vet

# 代码质量检查
make lint

# 安全检查
make security

# 完整CI检查
make ci
```

## 🚀 部署

### Docker部署

1. **构建镜像**
```bash
make docker-build
```

2. **运行容器**
```bash
make docker-run
```

### Docker Compose部署

```bash
# 启动所有服务
make docker-up

# 查看日志
make docker-logs

# 停止服务
make docker-down
```

### 生产环境部署

```bash
# 完整的生产环境构建
make deploy
```

## 📊 监控和可观测性

### 监控指标

- **HTTP请求指标**: 请求数量、延迟、状态码分布
- **系统指标**: CPU、内存、磁盘使用率
- **数据库指标**: 连接数、查询性能
- **业务指标**: 用户注册、登录统计

### 日志管理

- **结构化日志**: JSON格式，便于解析和搜索
- **日志级别**: DEBUG、INFO、WARN、ERROR、FATAL
- **日志轮转**: 按大小和时间自动轮转
- **链路追踪**: 分布式系统请求追踪

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 添加必要的注释和文档
- 编写测试用例
- 确保所有测试通过

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 支持

- 📧 Email: support@example.com
- 💬 Issues: [GitHub Issues](https://github.com/your-repo/issues)
- 📖 文档: [项目文档](https://your-docs-site.com)

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者和开源社区。

---

**注意**: 这是一个企业级Go项目模板，旨在提供一个结构良好、功能完整的起始点。请根据具体需求进行定制和扩展。