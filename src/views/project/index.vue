<template>
  <div class="project-management">
    <!-- 项目统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6" v-for="stat in projectStats" :key="stat.title">
        <el-card class="stat-card" shadow="hover">
          <div class="stat-content">
            <div class="stat-icon" :style="{ backgroundColor: stat.color }">
              <el-icon><component :is="stat.icon" /></el-icon>
            </div>
            <div class="stat-info">
              <div class="stat-value">{{ stat.value }}</div>
              <div class="stat-title">{{ stat.title }}</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 搜索和操作栏 -->
    <el-card class="search-card" shadow="never">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-input
            v-model="searchForm.keyword"
            placeholder="搜索项目名称/描述"
            prefix-icon="Search"
            clearable
            @keyup.enter="handleSearch"
          />
        </el-col>
        <el-col :span="4">
          <el-select v-model="searchForm.status" placeholder="状态" clearable>
            <el-option label="进行中" value="active" />
            <el-option label="已完成" value="completed" />
            <el-option label="已暂停" value="paused" />
            <el-option label="已取消" value="cancelled" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="searchForm.priority" placeholder="优先级" clearable>
            <el-option label="高" value="high" />
            <el-option label="中" value="medium" />
            <el-option label="低" value="low" />
          </el-select>
        </el-col>
        <el-col :span="10">
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleAdd">新建项目</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 项目列表 -->
    <el-card class="table-card" shadow="never">
      <el-table
        :data="projectList"
        v-loading="loading"
        stripe
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="项目名称" width="200" />
        <el-table-column prop="description" label="描述" width="300" show-overflow-tooltip />
        <el-table-column prop="manager" label="项目经理" width="120" />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="getStatusType(row.status)">{{ getStatusLabel(row.status) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="priority" label="优先级" width="100">
          <template #default="{ row }">
            <el-tag :type="getPriorityType(row.priority)" size="small">
              {{ getPriorityLabel(row.priority) }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="progress" label="进度" width="150">
          <template #default="{ row }">
            <el-progress :percentage="row.progress" :status="getProgressStatus(row.progress)" />
          </template>
        </el-table-column>
        <el-table-column prop="startDate" label="开始时间" width="120" />
        <el-table-column prop="endDate" label="结束时间" width="120" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button type="success" size="small" @click="handleView(row)">查看</el-button>
            <el-button type="danger" size="small" @click="handleDelete(row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页 -->
      <div class="pagination-wrapper">
        <el-pagination
          v-model:current-page="pagination.current"
          v-model:page-size="pagination.size"
          :page-sizes="[10, 20, 50, 100]"
          :total="pagination.total"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </el-card>

    <!-- 项目表单对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="600px"
      @close="handleDialogClose"
    >
      <el-form
        ref="projectFormRef"
        :model="projectForm"
        :rules="projectRules"
        label-width="100px"
      >
        <el-form-item label="项目名称" prop="name">
          <el-input v-model="projectForm.name" placeholder="请输入项目名称" />
        </el-form-item>
        <el-form-item label="项目描述" prop="description">
          <el-input
            v-model="projectForm.description"
            type="textarea"
            :rows="3"
            placeholder="请输入项目描述"
          />
        </el-form-item>
        <el-form-item label="项目经理" prop="manager">
          <el-input v-model="projectForm.manager" placeholder="请输入项目经理" />
        </el-form-item>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="开始时间" prop="startDate">
              <el-date-picker
                v-model="projectForm.startDate"
                type="date"
                placeholder="选择开始时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="结束时间" prop="endDate">
              <el-date-picker
                v-model="projectForm.endDate"
                type="date"
                placeholder="选择结束时间"
                style="width: 100%"
              />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row :gutter="20">
          <el-col :span="12">
            <el-form-item label="项目状态" prop="status">
              <el-select v-model="projectForm.status" placeholder="请选择状态" style="width: 100%">
                <el-option label="进行中" value="active" />
                <el-option label="已完成" value="completed" />
                <el-option label="已暂停" value="paused" />
                <el-option label="已取消" value="cancelled" />
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="优先级" prop="priority">
              <el-select v-model="projectForm.priority" placeholder="请选择优先级" style="width: 100%">
                <el-option label="高" value="high" />
                <el-option label="中" value="medium" />
                <el-option label="低" value="low" />
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>
        <el-form-item label="项目进度" prop="progress">
          <el-slider
            v-model="projectForm.progress"
            :min="0"
            :max="100"
            :step="5"
            show-input
            input-size="small"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="submitLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'

// 项目统计数据
const projectStats = ref([
  {
    title: '总项目数',
    value: '56',
    icon: 'Folder',
    color: '#409eff'
  },
  {
    title: '进行中',
    value: '23',
    icon: 'Loading',
    color: '#67c23a'
  },
  {
    title: '已完成',
    value: '28',
    icon: 'CircleCheck',
    color: '#e6a23c'
  },
  {
    title: '已暂停',
    value: '5',
    icon: 'Warning',
    color: '#f56c6c'
  }
])

// 搜索表单
const searchForm = reactive({
  keyword: '',
  status: '',
  priority: ''
})

// 项目列表
const projectList = ref([
  {
    id: 1,
    name: '电商平台开发',
    description: '基于Vue3和Node.js的现代化电商平台',
    manager: '张三',
    status: 'active',
    priority: 'high',
    progress: 75,
    startDate: '2024-01-01',
    endDate: '2024-06-30'
  },
  {
    id: 2,
    name: 'CRM系统升级',
    description: '客户关系管理系统的功能升级和优化',
    manager: '李四',
    status: 'completed',
    priority: 'medium',
    progress: 100,
    startDate: '2023-10-01',
    endDate: '2024-01-15'
  },
  {
    id: 3,
    name: '数据分析平台',
    description: '大数据分析和可视化平台建设',
    manager: '王五',
    status: 'paused',
    priority: 'low',
    progress: 30,
    startDate: '2024-02-01',
    endDate: '2024-12-31'
  }
])

// 分页
const pagination = reactive({
  current: 1,
  size: 10,
  total: 100
})

// 加载状态
const loading = ref(false)

// 对话框
const dialogVisible = ref(false)
const dialogTitle = ref('新建项目')
const submitLoading = ref(false)

// 项目表单
const projectFormRef = ref<FormInstance>()
const projectForm = reactive({
  id: '',
  name: '',
  description: '',
  manager: '',
  status: 'active',
  priority: 'medium',
  progress: 0,
  startDate: '',
  endDate: ''
})

// 表单验证规则
const projectRules: FormRules = {
  name: [
    { required: true, message: '请输入项目名称', trigger: 'blur' },
    { min: 2, max: 50, message: '项目名称长度在 2 到 50 个字符', trigger: 'blur' }
  ],
  description: [
    { required: true, message: '请输入项目描述', trigger: 'blur' }
  ],
  manager: [
    { required: true, message: '请输入项目经理', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择项目状态', trigger: 'change' }
  ],
  priority: [
    { required: true, message: '请选择优先级', trigger: 'change' }
  ]
}

// 获取状态类型
const getStatusType = (status: string) => {
  const types: Record<string, string> = {
    active: 'primary',
    completed: 'success',
    paused: 'warning',
    cancelled: 'danger'
  }
  return types[status] || 'info'
}

// 获取状态标签
const getStatusLabel = (status: string) => {
  const labels: Record<string, string> = {
    active: '进行中',
    completed: '已完成',
    paused: '已暂停',
    cancelled: '已取消'
  }
  return labels[status] || status
}

// 获取优先级类型
const getPriorityType = (priority: string) => {
  const types: Record<string, string> = {
    high: 'danger',
    medium: 'warning',
    low: 'info'
  }
  return types[priority] || 'info'
}

// 获取优先级标签
const getPriorityLabel = (priority: string) => {
  const labels: Record<string, string> = {
    high: '高',
    medium: '中',
    low: '低'
  }
  return labels[priority] || priority
}

// 获取进度状态
const getProgressStatus = (progress: number) => {
  if (progress >= 100) return 'success'
  if (progress >= 80) return 'warning'
  return ''
}

// 搜索
const handleSearch = () => {
  loading.value = true
  // 模拟API调用
  setTimeout(() => {
    loading.value = false
    ElMessage.success('搜索完成')
  }, 1000)
}

// 重置搜索
const handleReset = () => {
  Object.assign(searchForm, {
    keyword: '',
    status: '',
    priority: ''
  })
  handleSearch()
}

// 新建项目
const handleAdd = () => {
  dialogTitle.value = '新建项目'
  Object.assign(projectForm, {
    id: '',
    name: '',
    description: '',
    manager: '',
    status: 'active',
    priority: 'medium',
    progress: 0,
    startDate: '',
    endDate: ''
  })
  dialogVisible.value = true
}

// 编辑项目
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑项目'
  Object.assign(projectForm, row)
  dialogVisible.value = true
}

// 查看项目
const handleView = (row: any) => {
  ElMessage.info(`查看项目：${row.name}`)
  // 这里可以跳转到项目详情页面
}

// 删除项目
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除项目 "${row.name}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    // 模拟删除操作
    const index = projectList.value.findIndex(item => item.id === row.id)
    if (index > -1) {
      projectList.value.splice(index, 1)
      ElMessage.success('删除成功')
    }
  } catch {
    // 用户取消
  }
}

// 分页大小变更
const handleSizeChange = (size: number) => {
  pagination.size = size
  handleSearch()
}

// 当前页变更
const handleCurrentChange = (current: number) => {
  pagination.current = current
  handleSearch()
}

// 提交表单
const handleSubmit = async () => {
  if (!projectFormRef.value) return
  
  try {
    await projectFormRef.value.validate()
    submitLoading.value = true
    
    // 模拟提交
    setTimeout(() => {
      if (projectForm.id) {
        // 编辑
        const index = projectList.value.findIndex(item => item.id === projectForm.id)
        if (index > -1) {
          Object.assign(projectList.value[index], projectForm)
        }
        ElMessage.success('编辑成功')
      } else {
        // 新增
        const newProject = {
          ...projectForm,
          id: Date.now()
        }
        projectList.value.unshift(newProject)
        ElMessage.success('新建成功')
      }
      
      submitLoading.value = false
      dialogVisible.value = false
    }, 1000)
  } catch (error) {
    console.error('Form validation error:', error)
  }
}

// 对话框关闭
const handleDialogClose = () => {
  projectFormRef.value?.resetFields()
}

// 初始化
onMounted(() => {
  handleSearch()
})
</script>

<style lang="scss" scoped>
.project-management {
  .stats-row {
    margin-bottom: 20px;
  }

  .stat-card {
    .stat-content {
      display: flex;
      align-items: center;
      gap: 16px;
    }

    .stat-icon {
      width: 48px;
      height: 48px;
      border-radius: 8px;
      display: flex;
      align-items: center;
      justify-content: center;
      color: #fff;
      font-size: 24px;
    }

    .stat-info {
      flex: 1;

      .stat-value {
        font-size: 24px;
        font-weight: bold;
        color: #333;
        line-height: 1;
        margin-bottom: 4px;
      }

      .stat-title {
        font-size: 14px;
        color: #666;
      }
    }
  }

  .search-card {
    margin-bottom: 20px;
  }

  .table-card {
    .pagination-wrapper {
      margin-top: 20px;
      display: flex;
      justify-content: flex-end;
    }
  }
}
</style>