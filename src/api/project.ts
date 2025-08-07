import { get, post, put, del } from '@/utils/request'
import type { Project, PaginationParams, PaginationResponse } from '@/types'

// 获取项目列表
export const getProjectList = (params: PaginationParams & {
  keyword?: string
  status?: string
  priority?: string
  manager?: string
}) => {
  return get<PaginationResponse<Project>>('/project/list', params)
}

// 获取项目详情
export const getProjectDetail = (id: number) => {
  return get<Project>(`/project/${id}`)
}

// 创建项目
export const createProject = (data: Partial<Project>) => {
  return post<Project>('/project/create', data)
}

// 更新项目
export const updateProject = (id: number, data: Partial<Project>) => {
  return put<Project>(`/project/${id}`, data)
}

// 删除项目
export const deleteProject = (id: number) => {
  return del(`/project/${id}`)
}

// 更新项目状态
export const updateProjectStatus = (id: number, status: string) => {
  return put(`/project/${id}/status`, { status })
}

// 更新项目进度
export const updateProjectProgress = (id: number, progress: number) => {
  return put(`/project/${id}/progress`, { progress })
}

// 获取项目统计
export const getProjectStats = () => {
  return get<{
    total: number
    active: number
    completed: number
    paused: number
    cancelled: number
  }>('/project/stats')
}

// 获取项目成员
export const getProjectMembers = (id: number) => {
  return get<any[]>(`/project/${id}/members`)
}

// 添加项目成员
export const addProjectMember = (id: number, userId: number, role: string) => {
  return post(`/project/${id}/members`, { userId, role })
}

// 移除项目成员
export const removeProjectMember = (id: number, userId: number) => {
  return del(`/project/${id}/members/${userId}`)
}

// 获取项目文件
export const getProjectFiles = (id: number) => {
  return get<any[]>(`/project/${id}/files`)
}

// 上传项目文件
export const uploadProjectFile = (id: number, file: File) => {
  const formData = new FormData()
  formData.append('file', file)
  return post<{ url: string }>(`/project/${id}/files`, formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  })
}

// 删除项目文件
export const deleteProjectFile = (id: number, fileId: number) => {
  return del(`/project/${id}/files/${fileId}`)
}