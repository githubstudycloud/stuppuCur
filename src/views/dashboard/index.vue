<template>
  <div class="dashboard">
    <!-- 统计卡片 -->
    <el-row :gutter="20" class="stats-row">
      <el-col :span="6" v-for="stat in stats" :key="stat.title">
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

    <!-- 图表区域 -->
    <el-row :gutter="20" class="charts-row">
      <el-col :span="12">
        <el-card class="chart-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>项目统计</span>
              <el-button type="text">查看更多</el-button>
            </div>
          </template>
          <div class="chart-placeholder">
            <el-icon size="48" color="#409eff"><TrendCharts /></el-icon>
            <p>图表区域 - 可集成 ECharts 等图表库</p>
          </div>
        </el-card>
      </el-col>
      
      <el-col :span="12">
        <el-card class="chart-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>用户活跃度</span>
              <el-button type="text">查看更多</el-button>
            </div>
          </template>
          <div class="chart-placeholder">
            <el-icon size="48" color="#67c23a"><DataLine /></el-icon>
            <p>图表区域 - 可集成 ECharts 等图表库</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 最近活动 -->
    <el-row :gutter="20" class="activity-row">
      <el-col :span="16">
        <el-card class="activity-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>最近活动</span>
              <el-button type="text">查看全部</el-button>
            </div>
          </template>
          <el-timeline>
            <el-timeline-item
              v-for="activity in activities"
              :key="activity.id"
              :timestamp="activity.time"
              :type="activity.type"
            >
              {{ activity.content }}
            </el-timeline-item>
          </el-timeline>
        </el-card>
      </el-col>
      
      <el-col :span="8">
        <el-card class="quick-actions-card" shadow="hover">
          <template #header>
            <div class="card-header">
              <span>快捷操作</span>
            </div>
          </template>
          <div class="quick-actions">
            <el-button
              v-for="action in quickActions"
              :key="action.name"
              :type="action.type"
              :icon="action.icon"
              class="action-btn"
              @click="handleQuickAction(action)"
            >
              {{ action.name }}
            </el-button>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'

const router = useRouter()

// 统计数据
const stats = ref([
  {
    title: '总用户数',
    value: '1,234',
    icon: 'User',
    color: '#409eff'
  },
  {
    title: '项目数量',
    value: '56',
    icon: 'Folder',
    color: '#67c23a'
  },
  {
    title: '今日访问',
    value: '892',
    icon: 'View',
    color: '#e6a23c'
  },
  {
    title: '系统消息',
    value: '12',
    icon: 'Bell',
    color: '#f56c6c'
  }
])

// 最近活动
const activities = ref([
  {
    id: 1,
    content: '用户张三创建了新项目 "电商平台"',
    time: '2024-01-15 10:30',
    type: 'primary'
  },
  {
    id: 2,
    content: '项目 "CRM系统" 更新了版本 v2.1.0',
    time: '2024-01-15 09:15',
    type: 'success'
  },
  {
    id: 3,
    content: '用户李四加入了项目 "数据分析平台"',
    time: '2024-01-15 08:45',
    type: 'info'
  },
  {
    id: 4,
    content: '系统维护完成，所有服务正常运行',
    time: '2024-01-15 08:00',
    type: 'warning'
  }
])

// 快捷操作
const quickActions = ref([
  {
    name: '新建项目',
    icon: 'Plus',
    type: 'primary',
    action: () => router.push('/project/create')
  },
  {
    name: '用户管理',
    icon: 'User',
    type: 'success',
    action: () => router.push('/user')
  },
  {
    name: '系统设置',
    icon: 'Setting',
    type: 'info',
    action: () => router.push('/settings')
  },
  {
    name: '数据备份',
    icon: 'Download',
    type: 'warning',
    action: () => ElMessage.info('数据备份功能开发中...')
  }
])

// 处理快捷操作
const handleQuickAction = (action: any) => {
  if (action.action) {
    action.action()
  }
}
</script>

<style lang="scss" scoped>
.dashboard {
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

  .charts-row {
    margin-bottom: 20px;
  }

  .chart-card {
    .chart-placeholder {
      height: 300px;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      color: #999;

      p {
        margin-top: 16px;
        font-size: 14px;
      }
    }
  }

  .activity-row {
    .activity-card {
      .el-timeline {
        padding: 0;
      }
    }

    .quick-actions-card {
      .quick-actions {
        display: flex;
        flex-direction: column;
        gap: 12px;

        .action-btn {
          width: 100%;
          justify-content: flex-start;
        }
      }
    }
  }

  .card-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
}
</style>