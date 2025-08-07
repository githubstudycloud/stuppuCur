# Enterprise Python Template

[![CI/CD Pipeline](https://github.com/your-org/enterprise-python-template/workflows/CI/CD%20Pipeline/badge.svg)](https://github.com/your-org/enterprise-python-template/actions)
[![codecov](https://codecov.io/gh/your-org/enterprise-python-template/branch/main/graph/badge.svg)](https://codecov.io/gh/your-org/enterprise-python-template)
[![Python 3.11+](https://img.shields.io/badge/python-3.11+-blue.svg)](https://www.python.org/downloads/)
[![Code style: black](https://img.shields.io/badge/code%20style-black-000000.svg)](https://github.com/psf/black)

现代化的Python项目基座，专为大厂多项目管理设计。包含完整的开发工具链、CI/CD流水线、监控系统和部署配置。

## ✨ 特性

### 🏗️ 现代化架构
- **FastAPI** - 高性能异步Web框架
- **SQLAlchemy 2.0** - 现代ORM与异步支持
- **Pydantic V2** - 数据验证与序列化
- **Poetry** - 现代依赖管理
- **Typer** - 现代CLI框架

### 🔧 开发工具
- **代码质量**: Black, isort, flake8, mypy
- **测试框架**: pytest + pytest-asyncio
- **Git Hooks**: pre-commit
- **文档生成**: MkDocs + Material theme

### 🚀 部署与运维
- **容器化**: Docker + Docker Compose
- **CI/CD**: GitHub Actions
- **监控**: Prometheus + Grafana
- **日志**: Structured logging with Rich
- **反向代理**: Nginx配置

### 🛡️ 企业级特性
- **安全性**: JWT认证，密码加密，CORS配置
- **可观测性**: 指标收集，链路追踪，健康检查
- **可扩展性**: 异步架构，数据库连接池
- **配置管理**: 环境变量，多环境支持

## 🚀 快速开始

### 前置要求

- Python 3.11+
- Poetry 1.7+
- Docker (可选)
- Git

### 安装与设置

1. **克隆项目**
```bash
git clone https://github.com/your-org/enterprise-python-template.git
cd enterprise-python-template
```

2. **运行设置脚本**
```bash
./scripts/setup.sh
```

3. **启动开发服务器**
```bash
poetry run python -m src.cli serve --reload
```

4. **访问应用**
- API文档: http://localhost:8000/docs
- 应用主页: http://localhost:8000
- 健康检查: http://localhost:8000/health

### 使用Docker

```bash
# 构建并启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f app

# 停止服务
docker-compose down
```

## 📚 项目结构

```
enterprise-python-template/
├── src/                          # 源代码
│   ├── api/                      # API路由
│   │   └── v1/                   # API版本1
│   ├── core/                     # 核心模块
│   │   ├── config.py             # 配置管理
│   │   ├── database.py           # 数据库配置
│   │   └── logging.py            # 日志配置
│   ├── models/                   # 数据模型
│   ├── schemas/                  # Pydantic模式
│   ├── services/                 # 业务逻辑
│   ├── utils/                    # 工具函数
│   ├── cli.py                    # 命令行界面
│   └── main.py                   # 应用入口
├── tests/                        # 测试代码
├── docs/                         # 文档
├── scripts/                      # 脚本文件
├── config/                       # 配置文件
├── deployments/                  # 部署配置
│   ├── docker/                   # Docker相关
│   ├── k8s/                      # Kubernetes配置
│   ├── nginx/                    # Nginx配置
│   └── monitoring/               # 监控配置
├── .github/                      # GitHub Actions
├── pyproject.toml               # 项目配置
├── docker-compose.yml           # Docker编排
└── README.md                    # 项目文档
```

## 🔧 开发指南

### 环境配置

复制并编辑环境变量文件：
```bash
cp .env.example .env
```

主要配置项：
- `ENVIRONMENT`: 运行环境 (development/staging/production)
- `DATABASE_URL`: 数据库连接字符串
- `SECRET_KEY`: JWT加密密钥
- `LOG_LEVEL`: 日志级别

### 命令行工具

```bash
# 查看所有可用命令
poetry run python -m src.cli --help

# 启动开发服务器
poetry run python -m src.cli serve --reload

# 运行数据库迁移
poetry run python -m src.cli migrate

# 创建用户
poetry run python -m src.cli create-user \
  --email admin@example.com \
  --username admin \
  --password admin123 \
  --superuser

# 运行测试
poetry run python -m src.cli test --coverage

# 代码格式化
poetry run python -m src.cli format-code

# 查看应用信息
poetry run python -m src.cli info
```

### 代码质量

```bash
# 格式化代码
poetry run black src tests
poetry run isort src tests

# 运行linting
poetry run flake8 src tests
poetry run mypy src

# 运行安全检查
poetry run bandit -r src/
poetry run safety check

# 运行所有检查
poetry run python -m src.cli lint
```

### 测试

```bash
# 运行所有测试
poetry run pytest

# 运行特定标记的测试
poetry run pytest -m "not slow"

# 生成覆盖率报告
poetry run pytest --cov=src --cov-report=html

# 运行性能测试
poetry run pytest -m slow
```

## 🏗️ API文档

### 认证

使用JWT Bearer Token认证：

```bash
# 获取访问令牌
curl -X POST "http://localhost:8000/api/v1/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@example.com", "password": "admin123"}'

# 使用令牌访问受保护的API
curl -X GET "http://localhost:8000/api/v1/users/" \
  -H "Authorization: Bearer <your_token>"
```

### 主要端点

- `GET /` - 应用主页
- `GET /health` - 健康检查
- `GET /docs` - API文档 (Swagger UI)
- `GET /metrics` - Prometheus指标
- `POST /api/v1/users/` - 创建用户
- `GET /api/v1/users/` - 获取用户列表
- `GET /api/v1/users/{id}` - 获取特定用户

## 📊 监控与运维

### 健康检查

```bash
# 基础健康检查
curl http://localhost:8000/health

# 详细健康检查 (包含数据库状态)
curl http://localhost:8000/api/v1/health/detailed
```

### 指标监控

- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)

主要指标：
- HTTP请求次数和延迟
- 数据库连接池状态
- 应用程序内存使用
- 错误率和响应时间

### 日志

结构化日志输出，支持JSON和Pretty格式：

```bash
# 查看应用日志
docker-compose logs -f app

# 查看特定级别的日志
grep "ERROR" logs/app.log
```

## 🚀 部署

### Docker部署

```bash
# 构建镜像
docker build -t enterprise-python-template .

# 运行容器
docker run -p 8000:8000 \
  -e DATABASE_URL="postgresql://user:pass@host/db" \
  enterprise-python-template
```

### Kubernetes部署

```bash
# 部署到Kubernetes
kubectl apply -f deployments/k8s/

# 查看部署状态
kubectl get pods -l app=enterprise-python-template
```

### 生产环境配置

1. **环境变量**
```bash
export ENVIRONMENT=production
export SECRET_KEY="your-production-secret"
export DATABASE_URL="postgresql://..."
```

2. **数据库迁移**
```bash
poetry run python -m src.cli migrate
```

3. **启动应用**
```bash
poetry run python -m src.cli serve --workers 4
```

## 🔒 安全

### 最佳实践

- ✅ 密码使用bcrypt加密
- ✅ JWT Token过期机制
- ✅ CORS配置
- ✅ SQL注入防护
- ✅ 输入验证
- ✅ 安全HTTP头
- ✅ 依赖安全扫描

### 安全配置

```python
# 生产环境安全配置
CORS_ORIGINS = ["https://yourdomain.com"]
ALLOWED_HOSTS = ["yourdomain.com"]
SECRET_KEY = "complex-random-string"
```

## 🤝 贡献指南

1. Fork项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开Pull Request

### 开发规范

- 使用Black进行代码格式化
- 编写类型注解
- 添加docstring文档
- 编写单元测试
- 遵循语义化版本

## 📋 TODO

- [ ] 添加Redis缓存层
- [ ] 实现WebSocket支持
- [ ] 添加后台任务队列
- [ ] 集成OpenAPI 3.1
- [ ] 添加GraphQL支持
- [ ] 实现多租户架构

## 📄 许可证

本项目使用 [MIT License](LICENSE) 许可证。

## 🙏 致谢

- [FastAPI](https://fastapi.tiangolo.com/) - 现代化Web框架
- [Poetry](https://python-poetry.org/) - 依赖管理工具
- [Pydantic](https://pydantic-docs.helpmanual.io/) - 数据验证库
- [SQLAlchemy](https://www.sqlalchemy.org/) - Python ORM

## 📞 支持

如有问题或建议，请通过以下方式联系：

- 📧 Email: dev@yourcompany.com
- 🐛 Issues: [GitHub Issues](https://github.com/your-org/enterprise-python-template/issues)
- 📖 文档: [项目文档](https://your-org.github.io/enterprise-python-template)

---

⭐ 如果这个项目对您有帮助，请给个Star！