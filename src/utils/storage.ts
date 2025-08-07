// 本地存储工具类
class Storage {
  private prefix = 'vue_enterprise_'

  // 设置存储
  set(key: string, value: any, expire?: number): void {
    const data = {
      value,
      expire: expire ? Date.now() + expire * 1000 : null
    }
    localStorage.setItem(this.prefix + key, JSON.stringify(data))
  }

  // 获取存储
  get<T = any>(key: string): T | null {
    const item = localStorage.getItem(this.prefix + key)
    if (!item) return null

    try {
      const data = JSON.parse(item)
      if (data.expire && Date.now() > data.expire) {
        this.remove(key)
        return null
      }
      return data.value
    } catch {
      return null
    }
  }

  // 删除存储
  remove(key: string): void {
    localStorage.removeItem(this.prefix + key)
  }

  // 清空所有存储
  clear(): void {
    const keys = Object.keys(localStorage)
    keys.forEach(key => {
      if (key.startsWith(this.prefix)) {
        localStorage.removeItem(key)
      }
    })
  }

  // 检查是否存在
  has(key: string): boolean {
    return this.get(key) !== null
  }
}

// 会话存储工具类
class SessionStorage {
  private prefix = 'vue_enterprise_'

  // 设置存储
  set(key: string, value: any): void {
    sessionStorage.setItem(this.prefix + key, JSON.stringify(value))
  }

  // 获取存储
  get<T = any>(key: string): T | null {
    const item = sessionStorage.getItem(this.prefix + key)
    if (!item) return null

    try {
      return JSON.parse(item)
    } catch {
      return null
    }
  }

  // 删除存储
  remove(key: string): void {
    sessionStorage.removeItem(this.prefix + key)
  }

  // 清空所有存储
  clear(): void {
    const keys = Object.keys(sessionStorage)
    keys.forEach(key => {
      if (key.startsWith(this.prefix)) {
        sessionStorage.removeItem(key)
      }
    })
  }

  // 检查是否存在
  has(key: string): boolean {
    return this.get(key) !== null
  }
}

// 导出实例
export const storage = new Storage()
export const sessionStorage = new SessionStorage()

// 常用存储键名
export const STORAGE_KEYS = {
  TOKEN: 'token',
  USER_INFO: 'user_info',
  THEME: 'theme',
  LANGUAGE: 'language',
  SIDEBAR_COLLAPSED: 'sidebar_collapsed',
  SYSTEM_CONFIG: 'system_config'
} as const