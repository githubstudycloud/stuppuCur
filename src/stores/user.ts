import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface UserInfo {
  id: number
  username: string
  email: string
  avatar?: string
  role: string
  permissions: string[]
}

export const useUserStore = defineStore('user', () => {
  // 状态
  const token = ref<string>(localStorage.getItem('token') || '')
  const userInfo = ref<UserInfo | null>(null)
  const isLoggedIn = computed(() => !!token.value)

  // 设置token
  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  // 设置用户信息
  const setUserInfo = (info: UserInfo) => {
    userInfo.value = info
  }

  // 登录
  const login = async (credentials: { username: string; password: string }) => {
    try {
      // 这里应该调用实际的登录API
      const mockToken = 'mock-jwt-token'
      const mockUserInfo: UserInfo = {
        id: 1,
        username: credentials.username,
        email: `${credentials.username}@example.com`,
        role: 'admin',
        permissions: ['read', 'write', 'delete']
      }

      setToken(mockToken)
      setUserInfo(mockUserInfo)
      return { success: true }
    } catch (error) {
      return { success: false, error }
    }
  }

  // 登出
  const logout = () => {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
  }

  // 检查权限
  const hasPermission = (permission: string) => {
    return userInfo.value?.permissions.includes(permission) || false
  }

  return {
    token,
    userInfo,
    isLoggedIn,
    setToken,
    setUserInfo,
    login,
    logout,
    hasPermission
  }
})