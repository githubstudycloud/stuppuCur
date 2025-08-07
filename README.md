# 企业级Spring Boot 3.x多项目管理基座

基于Spring Boot 3.x构建的企业级多项目管理基座，采用现代化的技术栈和最佳实践，为大厂多项目管理提供完整的解决方案。

## 🚀 技术栈

### 核心框架
- **Spring Boot 3.2.0** - 主框架
- **Spring Security 6.x** - 安全框架
- **Spring Data JPA** - 数据访问层
- **MyBatis Plus 3.5.4** - ORM框架

### 数据库
- **MySQL 8.0** - 主数据库
- **Druid 1.2.20** - 数据库连接池
- **H2 Database** - 测试数据库

### 安全认证
- **JWT** - 无状态认证
- **BCrypt** - 密码加密

### API文档
- **Knife4j 4.3.0** - API文档生成

### 工具库
- **Lombok** - 代码简化
- **MapStruct** - 对象映射
- **Hutool** - 工具类库
- **FastJSON2** - JSON处理

### 开发工具
- **Maven** - 项目管理
- **Docker** - 容器化部署
- **Docker Compose** - 容器编排

## 📁 项目结构

```
spring-boot-enterprise-platform/
├── enterprise-common/          # 通用模块
│   ├── src/main/java/com/enterprise/common/
│   │   ├── core/              # 核心组件
│   │   ├── exception/         # 异常处理
│   │   └── utils/             # 工具类
│   └── pom.xml
├── enterprise-security/        # 安全模块
│   ├── src/main/java/com/enterprise/security/
│   │   ├── config/            # 安全配置
│   │   ├── filter/            # 安全过滤器
│   │   ├── handler/           # 安全处理器
│   │   └── utils/             # 安全工具
│   └── pom.xml
├── enterprise-system/          # 系统模块
│   ├── src/main/java/com/enterprise/system/
│   │   ├── entity/            # 实体类
│   │   ├── mapper/            # 数据访问层
│   │   ├── service/           # 业务逻辑层
│   │   └── resources/mapper/  # MyBatis映射文件
│   └── pom.xml
├── enterprise-admin/           # 管理后台
│   ├── src/main/java/com/enterprise/admin/
│   │   ├── config/            # 配置类
│   │   ├── controller/        # 控制器
│   │   └── EnterpriseAdminApplication.java
│   ├── src/main/resources/
│   │   ├── application.yml    # 应用配置
│   │   └── db/                # 数据库脚本
│   └── pom.xml
├── pom.xml                     # 父级POM
├── Dockerfile                  # Docker镜像构建
├── docker-compose.yml          # Docker编排
└── README.md                   # 项目文档
```

## 🛠️ 快速开始

### 环境要求

- **JDK 17+**
- **Maven 3.8+**
- **MySQL 8.0+**
- **Docker & Docker Compose** (可选)

### 本地开发

1. **克隆项目**
   ```bash
   git clone https://github.com/your-repo/spring-boot-enterprise-platform.git
   cd spring-boot-enterprise-platform
   ```

2. **配置数据库**
   ```bash
   # 创建数据库
   mysql -u root -p
   CREATE DATABASE enterprise DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   
   # 执行初始化脚本
   mysql -u root -p enterprise < enterprise-admin/src/main/resources/db/init.sql
   ```

3. **修改配置**
   ```yaml
   # enterprise-admin/src/main/resources/application.yml
   spring:
     datasource:
       url: jdbc:mysql://localhost:3306/enterprise?useUnicode=true&characterEncoding=utf8&zeroDateTimeBehavior=convertToNull&useSSL=true&serverTimezone=GMT%2B8
       username: your-username
       password: your-password
   ```

4. **启动应用**
   ```bash
   mvn clean install
   cd enterprise-admin
   mvn spring-boot:run
   ```

5. **访问应用**
   - 应用地址: http://localhost:8080/api
   - API文档: http://localhost:8080/api/doc.html
   - Druid监控: http://localhost:8080/api/druid

### Docker部署

1. **使用Docker Compose启动**
   ```bash
   docker-compose up -d
   ```

2. **查看服务状态**
   ```bash
   docker-compose ps
   ```

3. **查看日志**
   ```bash
   docker-compose logs -f enterprise-admin
   ```

## 🔐 默认账户

- **管理员账户**: admin / admin123
- **测试账户**: test / admin123

## 📚 API文档

启动应用后，访问以下地址查看API文档：

- **Knife4j文档**: http://localhost:8080/api/doc.html
- **Swagger文档**: http://localhost:8080/api/swagger-ui/index.html

## 🔧 主要功能

### 用户管理
- 用户注册、登录、注销
- 用户信息管理
- 密码修改、重置
- 用户状态管理

### 权限管理
- 基于JWT的无状态认证
- 角色权限控制
- 菜单权限管理
- 细粒度权限控制

### 系统管理
- 统一异常处理
- 统一响应格式
- 参数校验
- 日志管理

### 数据库管理
- 数据库连接池监控
- SQL性能监控
- 数据库备份恢复

## 🏗️ 架构设计

### 分层架构
- **Controller层**: 处理HTTP请求，参数校验
- **Service层**: 业务逻辑处理
- **Mapper层**: 数据访问层
- **Entity层**: 数据实体类

### 模块化设计
- **enterprise-common**: 通用组件，可被其他模块依赖
- **enterprise-security**: 安全组件，提供认证授权功能
- **enterprise-system**: 系统组件，提供核心业务功能
- **enterprise-admin**: 管理后台，提供Web接口

### 安全设计
- JWT无状态认证
- 密码BCrypt加密
- 统一异常处理
- 参数校验和防注入

## 🚀 部署指南

### 生产环境部署

1. **环境准备**
   ```bash
   # 安装JDK 17
   sudo apt-get update
   sudo apt-get install openjdk-17-jdk
   
   # 安装MySQL
   sudo apt-get install mysql-server
   
   # 安装Redis
   sudo apt-get install redis-server
   ```

2. **数据库配置**
   ```sql
   CREATE DATABASE enterprise DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
   CREATE USER 'enterprise'@'%' IDENTIFIED BY 'your-password';
   GRANT ALL PRIVILEGES ON enterprise.* TO 'enterprise'@'%';
   FLUSH PRIVILEGES;
   ```

3. **应用配置**
   ```yaml
   # application-prod.yml
   spring:
     datasource:
       url: jdbc:mysql://your-mysql-host:3306/enterprise
       username: enterprise
       password: your-password
     redis:
       host: your-redis-host
       port: 6379
   ```

4. **启动应用**
   ```bash
   java -jar -Dspring.profiles.active=prod enterprise-admin-1.0.0.jar
   ```

### 性能优化

1. **JVM调优**
   ```bash
   java -Xms2g -Xmx4g -XX:+UseG1GC -jar enterprise-admin-1.0.0.jar
   ```

2. **数据库优化**
   - 配置合适的连接池大小
   - 优化SQL查询
   - 添加适当的索引

3. **缓存优化**
   - 使用Redis缓存热点数据
   - 配置合适的缓存策略

## 🤝 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 📞 联系我们

- 项目地址: https://github.com/your-repo/spring-boot-enterprise-platform
- 问题反馈: https://github.com/your-repo/spring-boot-enterprise-platform/issues
- 邮箱: enterprise@example.com

## 🙏 致谢

感谢所有为这个项目做出贡献的开发者和开源社区。