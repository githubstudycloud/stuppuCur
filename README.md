# React Monorepo Foundation

Modern React monorepo foundation for enterprise multi-project management. Built with TypeScript, Next.js, and modern tooling for optimal developer experience and scalability.

## 🚀 Features

- **🏗️ Modern Monorepo**: Workspace-based architecture with pnpm and Turbo
- **⚡ Fast Builds**: Turborepo for intelligent build caching and parallelization
- **🎨 Design System**: Shared UI components with Tailwind CSS and Radix UI
- **📝 TypeScript**: End-to-end type safety with shared type definitions
- **🔧 Developer Experience**: ESLint, Prettier, Husky, and VS Code integration
- **🧪 Testing Ready**: Jest setup for unit testing
- **📚 Documentation**: Storybook for component documentation
- **🚀 CI/CD**: GitHub Actions workflow for automated testing and deployment

## 📁 Project Structure

```
├── apps/                    # Applications
│   ├── web/                # Main Next.js web application
│   ├── admin/              # Admin dashboard (placeholder)
│   └── mobile/             # Mobile app (placeholder)
├── packages/               # Shared packages
│   ├── ui/                 # UI component library
│   ├── utils/              # Utility functions
│   ├── config/             # Shared configuration
│   └── types/              # TypeScript type definitions
├── .github/                # GitHub Actions workflows
├── .vscode/                # VS Code workspace settings
└── docs/                   # Documentation
```

## 🛠️ Technology Stack

### Core Technologies
- **React 18** - Modern React with concurrent features
- **Next.js 14** - Full-stack React framework with App Router
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS** - Utility-first CSS framework

### Build Tools
- **Turbo** - High-performance build system
- **pnpm** - Fast, disk space efficient package manager
- **tsup** - TypeScript bundler powered by esbuild

### Development Tools
- **ESLint** - Code linting
- **Prettier** - Code formatting
- **Husky** - Git hooks
- **lint-staged** - Run linters on staged files

### UI Components
- **Radix UI** - Unstyled, accessible components
- **Class Variance Authority** - Component variants
- **Lucide React** - Beautiful icons

## 🚀 Quick Start

### Prerequisites

- Node.js 18+
- pnpm 8+

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd react-monorepo
   ```

2. **Install dependencies**
   ```bash
   pnpm install
   ```

3. **Start development**
   ```bash
   pnpm dev
   ```

This will start all applications in development mode:
- Web app: http://localhost:3000

## 📦 Packages Overview

### @company/ui
Shared UI component library built with React, TypeScript, and Tailwind CSS.

**Features:**
- Modern, accessible components
- Type-safe props with TypeScript
- Customizable with CSS variables
- Built with Radix UI primitives

**Usage:**
```tsx
import { Button, Card } from '@company/ui';

function MyComponent() {
  return (
    <Card>
      <Button variant="primary">Click me</Button>
    </Card>
  );
}
```

### @company/utils
Collection of utility functions for common operations.

**Features:**
- Date and number formatting
- Form validation with Zod
- HTTP client wrapper
- Type-safe storage utilities

**Usage:**
```tsx
import { formatCurrency, createApiClient } from '@company/utils';

const api = createApiClient('https://api.example.com');
const price = formatCurrency(1299.99); // $1,299.99
```

### @company/config
Shared configuration and constants.

**Features:**
- Environment-specific configurations
- Feature flags
- API endpoints and routes
- Validation schemas

**Usage:**
```tsx
import { API_ENDPOINTS, config } from '@company/config';

const apiUrl = config.api.baseUrl + API_ENDPOINTS.USERS;
```

### @company/types
Shared TypeScript type definitions.

**Features:**
- API response types
- Entity interfaces
- Form and UI types
- Utility types

**Usage:**
```tsx
import type { User, ApiResponse } from '@company/types';

const user: User = {
  id: '1',
  name: 'John Doe',
  email: 'john@example.com',
  // ...
};
```

## 🏗️ Available Scripts

### Root Level Commands

```bash
# Install dependencies
pnpm install

# Start all apps in development mode
pnpm dev

# Build all packages and apps
pnpm build

# Run linting
pnpm lint

# Run type checking
pnpm type-check

# Format code
pnpm format

# Clean all build artifacts
pnpm clean
```

### Package-Specific Commands

```bash
# Build specific package
pnpm --filter @company/ui build

# Run specific app
pnpm --filter @company/web dev

# Test specific package
pnpm --filter @company/utils test
```

## 🎨 Component Development

### Creating New Components

1. **Add to UI package**
   ```bash
   # Create component file
   touch packages/ui/src/components/my-component.tsx
   ```

2. **Component structure**
   ```tsx
   import * as React from 'react';
   import { cva, type VariantProps } from 'class-variance-authority';
   import { cn } from '../lib/utils';

   const myComponentVariants = cva(
     'base-classes',
     {
       variants: {
         variant: {
           default: 'default-classes',
           secondary: 'secondary-classes',
         },
       },
       defaultVariants: {
         variant: 'default',
       },
     }
   );

   export interface MyComponentProps
     extends React.HTMLAttributes<HTMLDivElement>,
       VariantProps<typeof myComponentVariants> {}

   const MyComponent = React.forwardRef<HTMLDivElement, MyComponentProps>(
     ({ className, variant, ...props }, ref) => {
       return (
         <div
           className={cn(myComponentVariants({ variant, className }))}
           ref={ref}
           {...props}
         />
       );
     }
   );
   MyComponent.displayName = 'MyComponent';

   export { MyComponent };
   ```

3. **Export from index**
   ```tsx
   // packages/ui/src/index.ts
   export * from './components/my-component';
   ```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
pnpm test

# Run tests for specific package
pnpm --filter @company/utils test

# Run tests in watch mode
pnpm --filter @company/utils test --watch
```

### Writing Tests

```tsx
// packages/utils/src/__tests__/formatting.test.ts
import { formatCurrency } from '../formatting';

describe('formatCurrency', () => {
  it('formats currency correctly', () => {
    expect(formatCurrency(1299.99)).toBe('$1,299.99');
  });
});
```

## 🎯 Adding New Applications

### Creating a New App

1. **Create app directory**
   ```bash
   mkdir apps/my-new-app
   cd apps/my-new-app
   ```

2. **Initialize with Next.js (or other framework)**
   ```bash
   npx create-next-app@latest . --typescript --tailwind --app
   ```

3. **Update package.json**
   ```json
   {
     "name": "@company/my-new-app",
     "dependencies": {
       "@company/ui": "workspace:*",
       "@company/utils": "workspace:*",
       "@company/config": "workspace:*",
       "@company/types": "workspace:*"
     }
   }
   ```

4. **Configure Turbo transpilation**
   ```js
   // next.config.js
   const nextConfig = {
     transpilePackages: ['@company/ui', '@company/utils', '@company/config', '@company/types'],
   };
   ```

## 🔧 Configuration

### Environment Variables

Create `.env.local` files in app directories:

```bash
# apps/web/.env.local
NEXT_PUBLIC_API_URL=http://localhost:3001
NEXT_PUBLIC_APP_NAME=Company Web App
```

### Tailwind Configuration

Each app should extend the base Tailwind config:

```js
// apps/web/tailwind.config.js
module.exports = {
  content: [
    './src/**/*.{js,ts,jsx,tsx,mdx}',
    '../../packages/ui/src/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      // Custom theme extensions
    },
  },
  plugins: [],
};
```

## 🚀 Deployment

### Build for Production

```bash
# Build all packages and apps
pnpm build

# Build specific app
pnpm --filter @company/web build
```

### Deployment Platforms

- **Vercel**: Recommended for Next.js apps
- **Netlify**: Alternative for static sites
- **Docker**: Container-based deployment

### GitHub Actions

The repository includes a CI/CD workflow that:
- Runs linting and type checking
- Executes tests
- Builds all packages and applications
- Caches dependencies for faster builds

## 📚 Best Practices

### Code Organization

1. **Keep packages focused**: Each package should have a single responsibility
2. **Use barrel exports**: Export everything through index files
3. **Follow naming conventions**: Use consistent naming across packages
4. **Document public APIs**: Add JSDoc comments for exported functions

### TypeScript

1. **Use strict mode**: Enable strict TypeScript checking
2. **Define interfaces**: Create interfaces for all data structures
3. **Leverage utility types**: Use built-in and custom utility types
4. **Avoid any**: Use proper typing instead of `any`

### Components

1. **Use forwardRef**: For components that need ref forwarding
2. **Add displayName**: For better debugging experience
3. **Support variants**: Use CVA for component variants
4. **Make accessible**: Follow ARIA guidelines

### Performance

1. **Use Turbo cache**: Leverage Turbo's intelligent caching
2. **Optimize imports**: Use tree-shaking friendly imports
3. **Code splitting**: Split large applications into chunks
4. **Monitor bundle size**: Keep track of bundle sizes

## 🤝 Contributing

### Development Workflow

1. **Create feature branch**
   ```bash
   git checkout -b feature/my-feature
   ```

2. **Make changes and test**
   ```bash
   pnpm dev
   pnpm test
   pnpm lint
   ```

3. **Commit with conventional commits**
   ```bash
   git commit -m "feat: add new component"
   ```

4. **Push and create PR**
   ```bash
   git push origin feature/my-feature
   ```

### Code Style

- Follow ESLint and Prettier configurations
- Use conventional commit messages
- Write meaningful test cases
- Document new features

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

- **Documentation**: Check the `/docs` directory for detailed guides
- **Issues**: Report bugs and request features via GitHub Issues
- **Discussions**: Use GitHub Discussions for questions and ideas

---

Built with ❤️ for modern web development