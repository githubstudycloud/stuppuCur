// 通用响应类型
export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

// 分页参数类型
export interface PaginationParams {
  current: number
  size: number
  total?: number
}

// 分页响应类型
export interface PaginationResponse<T = any> {
  list: T[]
  total: number
  current: number
  size: number
}

// 用户相关类型
export interface UserInfo {
  id: number
  username: string
  email: string
  avatar?: string
  role: string
  permissions: string[]
  createTime?: string
  lastLoginTime?: string
}

// 项目相关类型
export interface Project {
  id: number
  name: string
  description: string
  manager: string
  status: 'active' | 'completed' | 'paused' | 'cancelled'
  priority: 'high' | 'medium' | 'low'
  progress: number
  startDate: string
  endDate: string
  createTime?: string
  updateTime?: string
}

// 路由元信息类型
export interface RouteMeta {
  title?: string
  icon?: string
  requiresAuth?: boolean
  permissions?: string[]
  hidden?: boolean
  keepAlive?: boolean
}

// 菜单项类型
export interface MenuItem {
  id: string
  title: string
  icon?: string
  path?: string
  children?: MenuItem[]
  meta?: RouteMeta
}

// 表格列配置类型
export interface TableColumn {
  prop: string
  label: string
  width?: number | string
  minWidth?: number | string
  fixed?: boolean | 'left' | 'right'
  sortable?: boolean
  align?: 'left' | 'center' | 'right'
  showOverflowTooltip?: boolean
  formatter?: (row: any, column: any, cellValue: any, index: number) => string
}

// 表单字段类型
export interface FormField {
  prop: string
  label: string
  type: 'input' | 'select' | 'textarea' | 'date' | 'datetime' | 'switch' | 'radio' | 'checkbox'
  placeholder?: string
  required?: boolean
  rules?: any[]
  options?: Array<{ label: string; value: any }>
  props?: Record<string, any>
}

// 文件上传类型
export interface UploadFile {
  uid: string
  name: string
  status: 'uploading' | 'done' | 'error'
  url?: string
  response?: any
}

// 通知消息类型
export interface Notification {
  id: string
  title: string
  content: string
  type: 'info' | 'success' | 'warning' | 'error'
  timestamp: number
  read: boolean
}

// 系统配置类型
export interface SystemConfig {
  theme: 'light' | 'dark'
  language: 'zh-CN' | 'en-US'
  sidebarCollapsed: boolean
  showBreadcrumb: boolean
  showTags: boolean
}