# 贡献指南

感谢您对本项目的关注！我们欢迎所有形式的贡献，包括但不限于：

- 🐛 Bug修复
- ✨ 新功能
- 📚 文档改进
- 🧪 测试用例
- 💡 功能建议

## 🚀 快速开始

### 环境准备

1. Fork并克隆仓库
```bash
git clone https://github.com/your-username/enterprise-python-template.git
cd enterprise-python-template
```

2. 安装依赖
```bash
./scripts/setup.sh
```

3. 创建开发分支
```bash
git checkout -b feature/your-feature-name
```

## 📝 开发规范

### 代码风格

我们使用以下工具确保代码质量：

- **Black**: 代码格式化
- **isort**: 导入排序
- **flake8**: 代码规范检查
- **mypy**: 类型检查

运行代码格式化：
```bash
poetry run python -m src.cli format-code
```

运行代码检查：
```bash
poetry run python -m src.cli lint
```

### 提交规范

我们使用 [Conventional Commits](https://www.conventionalcommits.org/) 规范：

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

**类型说明：**
- `feat`: 新功能
- `fix`: Bug修复
- `docs`: 文档更新
- `style`: 代码格式（不影响功能）
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

**示例：**
```
feat(api): add user authentication endpoint

- Add JWT token generation
- Add password hashing
- Add login/logout endpoints

Closes #123
```

### 分支命名

- `feature/功能名` - 新功能
- `fix/问题描述` - Bug修复
- `docs/文档更新` - 文档更新
- `refactor/重构描述` - 代码重构

## 🧪 测试

### 编写测试

- 所有新功能必须包含单元测试
- 使用pytest框架
- 测试覆盖率应保持在90%以上

```bash
# 运行测试
poetry run pytest

# 运行特定测试
poetry run pytest tests/test_users.py

# 生成覆盖率报告
poetry run pytest --cov=src --cov-report=html
```

### 测试分类

使用pytest标记对测试进行分类：

```python
import pytest

@pytest.mark.unit
def test_user_creation():
    """单元测试"""
    pass

@pytest.mark.integration
def test_database_integration():
    """集成测试"""
    pass

@pytest.mark.slow
def test_performance():
    """性能测试"""
    pass
```

运行特定类型的测试：
```bash
# 只运行单元测试
poetry run pytest -m unit

# 跳过慢速测试
poetry run pytest -m "not slow"
```

## 📖 文档

### API文档

- 所有API端点必须包含完整的docstring
- 使用Pydantic模型定义请求/响应格式
- 在FastAPI自动生成的文档中验证

### 代码文档

```python
def create_user(user_data: UserCreate) -> User:
    """
    创建新用户。
    
    Args:
        user_data: 用户创建数据
        
    Returns:
        User: 创建的用户对象
        
    Raises:
        ValueError: 当用户已存在时
        
    Example:
        >>> user_data = UserCreate(email="test@example.com", username="test")
        >>> user = create_user(user_data)
        >>> print(user.id)
        1
    """
    pass
```

## 🔍 代码审查

### 提交PR前检查清单

- [ ] 代码通过所有测试
- [ ] 代码通过linting检查
- [ ] 添加了必要的测试用例
- [ ] 更新了相关文档
- [ ] 提交信息符合规范
- [ ] 功能在本地环境正常工作

### PR模板

```markdown
## 变更类型
- [ ] Bug修复
- [ ] 新功能
- [ ] 文档更新
- [ ] 重构
- [ ] 其他

## 变更描述
简要描述本次变更的内容和原因。

## 测试
描述如何测试这些变更。

## 检查清单
- [ ] 代码通过所有测试
- [ ] 添加了测试用例
- [ ] 更新了文档
- [ ] 遵循代码规范
```

## 🐛 报告Bug

### Bug报告模板

```markdown
**Bug描述**
简洁清晰地描述bug。

**复现步骤**
1. 执行'...'
2. 点击'....'
3. 滚动到'....'
4. 看到错误

**期望行为**
清晰简洁地描述你期望发生什么。

**实际行为**
描述实际发生了什么。

**环境信息**
- OS: [e.g. Ubuntu 20.04]
- Python版本: [e.g. 3.11.0]
- 项目版本: [e.g. 0.1.0]

**额外信息**
添加任何其他有关问题的上下文。
```

## 💡 功能请求

### 功能请求模板

```markdown
**问题描述**
简洁清晰地描述问题。例如：我总是感到沫暴 [...]

**期望的解决方案**
清晰简洁地描述你希望发生什么。

**可考虑的替代方案**
清晰简洁地描述你考虑过的任何替代解决方案或功能。

**额外信息**
添加任何其他有关功能请求的上下文或屏幕截图。
```

## 🏗️ 架构决策

### 新增依赖

在添加新依赖前，请考虑：

1. **必要性**: 是否真的需要这个依赖？
2. **维护状态**: 依赖是否得到积极维护？
3. **安全性**: 依赖是否有已知的安全问题？
4. **大小**: 依赖的大小是否合理？
5. **许可证**: 许可证是否兼容？

### 性能考虑

- 使用异步编程模式
- 避免N+1查询问题
- 合理使用缓存
- 监控内存使用
- 优化数据库查询

## 📞 获取帮助

如果您在贡献过程中遇到问题，可以通过以下方式获取帮助：

- 💬 [GitHub Discussions](https://github.com/your-org/enterprise-python-template/discussions)
- 📧 发送邮件到 dev@yourcompany.com
- 🐛 在 [GitHub Issues](https://github.com/your-org/enterprise-python-template/issues) 提问

## 🎉 贡献者

感谢所有为这个项目贡献的开发者！

<!-- 这里可以添加贡献者列表 -->

---

再次感谢您的贡献！🙏