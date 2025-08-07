#!/bin/bash

# Vue Enterprise Base 部署脚本
# 使用方法: ./scripts/deploy.sh [production|staging]

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查环境
check_environment() {
    log_info "检查环境..."
    
    # 检查Node.js
    if ! command -v node &> /dev/null; then
        log_error "Node.js 未安装"
        exit 1
    fi
    
    # 检查npm
    if ! command -v npm &> /dev/null; then
        log_error "npm 未安装"
        exit 1
    fi
    
    # 检查Docker（可选）
    if command -v docker &> /dev/null; then
        log_info "Docker 已安装"
    else
        log_warn "Docker 未安装，将跳过容器化部署"
    fi
    
    log_info "环境检查完成"
}

# 安装依赖
install_dependencies() {
    log_info "安装依赖..."
    npm ci
    log_info "依赖安装完成"
}

# 代码检查
lint_code() {
    log_info "代码检查..."
    npm run lint
    log_info "代码检查完成"
}

# 类型检查
type_check() {
    log_info "TypeScript类型检查..."
    npm run type-check
    log_info "类型检查完成"
}

# 运行测试
run_tests() {
    log_info "运行测试..."
    npm run test
    log_info "测试完成"
}

# 构建项目
build_project() {
    local env=$1
    log_info "构建项目 (环境: $env)..."
    
    # 设置环境变量
    export NODE_ENV=$env
    
    # 构建
    npm run build
    
    log_info "项目构建完成"
}

# Docker构建
build_docker() {
    if ! command -v docker &> /dev/null; then
        log_warn "跳过Docker构建"
        return
    fi
    
    log_info "构建Docker镜像..."
    docker build -t vue-enterprise-base:latest .
    log_info "Docker镜像构建完成"
}

# 部署到服务器
deploy_to_server() {
    local env=$1
    log_info "部署到服务器 (环境: $env)..."
    
    # 这里可以添加实际的部署逻辑
    # 例如：rsync、scp、kubectl等
    
    log_info "部署完成"
}

# 主函数
main() {
    local env=${1:-production}
    
    log_info "开始部署 Vue Enterprise Base..."
    log_info "部署环境: $env"
    
    # 检查环境
    check_environment
    
    # 安装依赖
    install_dependencies
    
    # 代码检查
    lint_code
    
    # 类型检查
    type_check
    
    # 运行测试
    run_tests
    
    # 构建项目
    build_project $env
    
    # Docker构建
    build_docker
    
    # 部署到服务器
    deploy_to_server $env
    
    log_info "部署完成！"
}

# 执行主函数
main "$@"