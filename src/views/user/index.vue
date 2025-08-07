<template>
  <div class="user-management">
    <!-- 搜索和操作栏 -->
    <el-card class="search-card" shadow="never">
      <el-row :gutter="20">
        <el-col :span="6">
          <el-input
            v-model="searchForm.keyword"
            placeholder="搜索用户名/邮箱"
            prefix-icon="Search"
            clearable
            @keyup.enter="handleSearch"
          />
        </el-col>
        <el-col :span="4">
          <el-select v-model="searchForm.status" placeholder="状态" clearable>
            <el-option label="启用" value="active" />
            <el-option label="禁用" value="inactive" />
          </el-select>
        </el-col>
        <el-col :span="4">
          <el-select v-model="searchForm.role" placeholder="角色" clearable>
            <el-option label="管理员" value="admin" />
            <el-option label="普通用户" value="user" />
            <el-option label="访客" value="guest" />
          </el-select>
        </el-col>
        <el-col :span="10">
          <el-button type="primary" @click="handleSearch">搜索</el-button>
          <el-button @click="handleReset">重置</el-button>
          <el-button type="success" @click="handleAdd">新增用户</el-button>
        </el-col>
      </el-row>
    </el-card>

    <!-- 用户列表 -->
    <el-card class="table-card" shadow="never">
      <el-table
        :data="userList"
        v-loading="loading"
        stripe
        border
        style="width: 100%"
      >
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="用户名" width="120" />
        <el-table-column prop="email" label="邮箱" width="200" />
        <el-table-column prop="role" label="角色" width="100">
          <template #default="{ row }">
            <el-tag :type="getRoleType(row.role)">{{ getRoleLabel(row.role) }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-switch
              v-model="row.status"
              :active-value="'active'"
              :inactive-value="'inactive'"
              @change="handleStatusChange(row)"
            />
          </template>
        </el-table-column>
        <el-table-column prop="createTime" label="创建时间" width="180" />
        <el-table-column prop="lastLoginTime" label="最后登录" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button type="primary" size="small" @click="handleEdit(row)">编辑</el-button>
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

    <!-- 用户表单对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="500px"
      @close="handleDialogClose"
    >
      <el-form
        ref="userFormRef"
        :model="userForm"
        :rules="userRules"
        label-width="80px"
      >
        <el-form-item label="用户名" prop="username">
          <el-input v-model="userForm.username" placeholder="请输入用户名" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userForm.email" placeholder="请输入邮箱" />
        </el-form-item>
        <el-form-item label="密码" prop="password" v-if="!userForm.id">
          <el-input
            v-model="userForm.password"
            type="password"
            placeholder="请输入密码"
            show-password
          />
        </el-form-item>
        <el-form-item label="角色" prop="role">
          <el-select v-model="userForm.role" placeholder="请选择角色">
            <el-option label="管理员" value="admin" />
            <el-option label="普通用户" value="user" />
            <el-option label="访客" value="guest" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="userForm.status">
            <el-radio label="active">启用</el-radio>
            <el-radio label="inactive">禁用</el-radio>
          </el-radio-group>
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

// 搜索表单
const searchForm = reactive({
  keyword: '',
  status: '',
  role: ''
})

// 用户列表
const userList = ref([
  {
    id: 1,
    username: 'admin',
    email: 'admin@example.com',
    role: 'admin',
    status: 'active',
    createTime: '2024-01-01 10:00:00',
    lastLoginTime: '2024-01-15 15:30:00'
  },
  {
    id: 2,
    username: 'user1',
    email: 'user1@example.com',
    role: 'user',
    status: 'active',
    createTime: '2024-01-02 11:00:00',
    lastLoginTime: '2024-01-15 14:20:00'
  },
  {
    id: 3,
    username: 'guest1',
    email: 'guest1@example.com',
    role: 'guest',
    status: 'inactive',
    createTime: '2024-01-03 12:00:00',
    lastLoginTime: '2024-01-14 16:45:00'
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
const dialogTitle = ref('新增用户')
const submitLoading = ref(false)

// 用户表单
const userFormRef = ref<FormInstance>()
const userForm = reactive({
  id: '',
  username: '',
  email: '',
  password: '',
  role: '',
  status: 'active'
})

// 表单验证规则
const userRules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度在 3 到 20 个字符', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, max: 20, message: '密码长度在 6 到 20 个字符', trigger: 'blur' }
  ],
  role: [
    { required: true, message: '请选择角色', trigger: 'change' }
  ]
}

// 获取角色类型
const getRoleType = (role: string) => {
  const types: Record<string, string> = {
    admin: 'danger',
    user: 'primary',
    guest: 'info'
  }
  return types[role] || 'info'
}

// 获取角色标签
const getRoleLabel = (role: string) => {
  const labels: Record<string, string> = {
    admin: '管理员',
    user: '普通用户',
    guest: '访客'
  }
  return labels[role] || role
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
    role: ''
  })
  handleSearch()
}

// 新增用户
const handleAdd = () => {
  dialogTitle.value = '新增用户'
  Object.assign(userForm, {
    id: '',
    username: '',
    email: '',
    password: '',
    role: '',
    status: 'active'
  })
  dialogVisible.value = true
}

// 编辑用户
const handleEdit = (row: any) => {
  dialogTitle.value = '编辑用户'
  Object.assign(userForm, row)
  dialogVisible.value = true
}

// 删除用户
const handleDelete = async (row: any) => {
  try {
    await ElMessageBox.confirm(`确定要删除用户 "${row.username}" 吗？`, '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    
    // 模拟删除操作
    const index = userList.value.findIndex(item => item.id === row.id)
    if (index > -1) {
      userList.value.splice(index, 1)
      ElMessage.success('删除成功')
    }
  } catch {
    // 用户取消
  }
}

// 状态变更
const handleStatusChange = (row: any) => {
  ElMessage.success(`用户 ${row.username} 状态已${row.status === 'active' ? '启用' : '禁用'}`)
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
  if (!userFormRef.value) return
  
  try {
    await userFormRef.value.validate()
    submitLoading.value = true
    
    // 模拟提交
    setTimeout(() => {
      if (userForm.id) {
        // 编辑
        const index = userList.value.findIndex(item => item.id === userForm.id)
        if (index > -1) {
          Object.assign(userList.value[index], userForm)
        }
        ElMessage.success('编辑成功')
      } else {
        // 新增
        const newUser = {
          ...userForm,
          id: Date.now(),
          createTime: new Date().toLocaleString(),
          lastLoginTime: '-'
        }
        userList.value.unshift(newUser)
        ElMessage.success('新增成功')
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
  userFormRef.value?.resetFields()
}

// 初始化
onMounted(() => {
  handleSearch()
})
</script>

<style lang="scss" scoped>
.user-management {
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