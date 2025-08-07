import { get, post, put, del } from '@/utils/request'
import type { UserInfo, PaginationParams, PaginationResponse } from '@/types'

// 用户登录
export const login = (data: { username: string; password: string }) => {
  return post<{ token: string; userInfo: UserInfo }>('/auth/login', data)
}

// 用户登出
export const logout = () => {
  return post('/auth/logout')
}

// 获取用户信息
export const getUserInfo = () => {
  return get<UserInfo>('/user/info')
}

// 获取用户列表
export const getUserList = (params: PaginationParams & {
  keyword?: string
  status?: string
  role?: string
}) => {
  return get<PaginationResponse<UserInfo>>('/user/list', params)
}

// 创建用户
export const createUser = (data: Partial<UserInfo>) => {
  return post<UserInfo>('/user/create', data)
}

// 更新用户
export const updateUser = (id: number, data: Partial<UserInfo>) => {
  return put<UserInfo>(`/user/${id}`, data)
}

// 删除用户
export const deleteUser = (id: number) => {
  return del(`/user/${id}`)
}

// 更新用户状态
export const updateUserStatus = (id: number, status: string) => {
  return put(`/user/${id}/status`, { status })
}

// 重置用户密码
export const resetUserPassword = (id: number) => {
  return post(`/user/${id}/reset-password`)
}

// 修改密码
export const changePassword = (data: {
  oldPassword: string
  newPassword: string
}) => {
  return post('/user/change-password', data)
}

// 上传头像
export const uploadAvatar = (file: File) => {
  const formData = new FormData()
  formData.append('avatar', file)
  return post<{ url: string }>('/user/upload-avatar', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}