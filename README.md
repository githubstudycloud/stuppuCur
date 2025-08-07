# Vue Enterprise Base

一个现代化的企业级Vue项目基座，专为大厂多项目管理而设计。

## 🚀 特性

- **Vue 3 + TypeScript** - 使用最新的Vue 3 Composition API和TypeScript
- **Vite** - 快速的构建工具和开发服务器
- **Element Plus** - 企业级UI组件库
- **Pinia** - 现代化的状态管理
- **Vue Router 4** - 官方路由管理器
- **完整的开发工具链** - ESLint, Prettier, Husky等
- **模块化架构** - 清晰的目录结构和代码组织
- **权限管理** - 基于角色的访问控制
- **响应式设计** - 支持多设备适配
- **国际化支持** - 多语言支持架构
- **主题定制** - 支持明暗主题切换

## 📦 技术栈

- **前端框架**: Vue 3.4+
- **开发语言**: TypeScript 5.3+
- **构建工具**: Vite 5.0+
- **UI组件库**: Element Plus 2.4+
- **状态管理**: Pinia 2.1+
- **路由管理**: Vue Router 4.2+
- **HTTP客户端**: Axios 1.6+
- **工具库**: VueUse, Lodash-es, Day.js
- **代码规范**: ESLint, Prettier
- **Git钩子**: Husky, lint-staged
- **测试框架**: Vitest

## 🏗️ 项目结构

```
vue-enterprise-base/
├── public/                 # 静态资源
├── src/
│   ├── api/               # API接口
│   ├── assets/            # 静态资源
│   ├── components/        # 公共组件
│   ├── layout/            # 布局组件
│   ├── router/            # 路由配置
│   ├── stores/            # 状态管理
│   ├── styles/            # 全局样式
│   ├── types/             # TypeScript类型定义
│   ├── utils/             # 工具函数
│   ├── views/             # 页面组件
│   ├── App.vue            # 根组件
│   └── main.ts            # 入口文件
├── .eslintrc.cjs          # ESLint配置
├── .prettierrc            # Prettier配置
├── index.html             # HTML模板
├── package.json           # 项目配置
├── tsconfig.json          # TypeScript配置
├── vite.config.ts         # Vite配置
└── README.md              # 项目文档
```

## 🚀 快速开始

### 环境要求

- Node.js >= 18.0.0
- npm >= 9.0.0 或 yarn >= 1.22.0

### 安装依赖

```bash
npm install
# 或
yarn install
```

### 开发环境

```bash
npm run dev
# 或
yarn dev
```

### 构建生产版本

```bash
npm run build
# 或
yarn build
```

### 预览生产版本

```bash
npm run preview
# 或
yarn preview
```

## 📋 可用脚本

- `npm run dev` - 启动开发服务器
- `npm run build` - 构建生产版本
- `npm run preview` - 预览生产版本
- `npm run lint` - 代码检查和修复
- `npm run type-check` - TypeScript类型检查
- `npm run test` - 运行测试
- `npm run test:ui` - 运行测试UI
- `npm run test:coverage` - 生成测试覆盖率报告

## 🔧 配置说明

### 环境变量

创建 `.env` 文件：

```env
VITE_APP_TITLE=Vue Enterprise Base
VITE_API_BASE_URL=http://localhost:8080/api
VITE_APP_ENV=development
```

### Vite配置

主要配置项：

- **别名配置**: 支持 `@` 等路径别名
- **代理配置**: 开发环境API代理
- **插件配置**: Vue、自动导入、组件注册等
- **构建配置**: 输出目录、文件命名等

### TypeScript配置

- 严格模式启用
- 路径映射配置
- 类型检查配置

## 🎨 主题定制

### 颜色变量

```scss
:root {
  --el-color-primary: #409eff;
  --el-color-success: #67c23a;
  --el-color-warning: #e6a23c;
  --el-color-danger: #f56c6c;
  --el-color-info: #909399;
}
```

### 暗色主题

```scss
html.dark {
  --el-bg-color: #141414;
  --el-bg-color-page: #0a0a0a;
  --el-text-color-primary: #e5eaf3;
  // ... 更多变量
}
```

## 🔐 权限管理

### 角色定义

- **admin**: 管理员，拥有所有权限
- **user**: 普通用户，拥有基本权限
- **guest**: 访客，只有查看权限

### 权限检查

```typescript
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()
const hasPermission = userStore.hasPermission('read')
```

## 📱 响应式设计

### 断点配置

- **xs**: < 768px (手机)
- **sm**: >= 768px (平板)
- **md**: >= 992px (桌面)
- **lg**: >= 1200px (大屏)
- **xl**: >= 1920px (超大屏)

### 工具类

```scss
.hidden-mobile { /* 移动端隐藏 */ }
.hidden-desktop { /* 桌面端隐藏 */ }
```

## 🌍 国际化

### 语言配置

支持中文和英文：

```typescript
const messages = {
  'zh-CN': zhCN,
  'en-US': enUS
}
```

### 使用示例

```vue
<template>
  <div>{{ $t('common.welcome') }}</div>
</template>
```

## 🧪 测试

### 单元测试

使用Vitest进行单元测试：

```bash
npm run test
```

### 测试覆盖率

```bash
npm run test:coverage
```

## 📦 部署

### 构建

```bash
npm run build
```

### 部署到服务器

将 `dist` 目录部署到Web服务器即可。

### Docker部署

```dockerfile
FROM nginx:alpine
COPY dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- [Vue.js](https://vuejs.org/) - 渐进式JavaScript框架
- [Element Plus](https://element-plus.org/) - 企业级UI组件库
- [Vite](https://vitejs.dev/) - 下一代前端构建工具
- [Pinia](https://pinia.vuejs.org/) - Vue的状态管理库

## 📞 联系方式

如有问题或建议，请通过以下方式联系：

- 邮箱: your-email@example.com
- 项目地址: https://github.com/your-username/vue-enterprise-base

---

**Vue Enterprise Base** - 让企业级Vue开发更简单、更高效！