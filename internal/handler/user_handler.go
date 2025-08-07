package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	
	"github.com/company/go-enterprise-template/internal/middleware"
	"github.com/company/go-enterprise-template/internal/model"
	"github.com/company/go-enterprise-template/internal/service"
)

// UserHandler 用户处理器
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器实例
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register 用户注册
// @Summary 用户注册
// @Description 用户注册接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body model.UserCreateRequest true "注册信息"
// @Success 200 {object} Response{data=model.UserResponse} "注册成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req model.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationError(c, err)
		return
	}

	user, err := h.userService.Register(c.Request.Context(), &req)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	SuccessWithMessage(c, "Registration successful", user)
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body model.UserLoginRequest true "登录信息"
// @Success 200 {object} Response{data=model.UserLoginResponse} "登录成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "认证失败"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req model.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationError(c, err)
		return
	}

	loginResp, err := h.userService.Login(c.Request.Context(), &req)
	if err != nil {
		Unauthorized(c, err.Error())
		return
	}

	SuccessWithMessage(c, "Login successful", loginResp)
}

// GetProfile 获取当前用户资料
// @Summary 获取用户资料
// @Description 获取当前登录用户的资料信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} Response{data=model.UserResponse} "获取成功"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "用户不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/user/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		Unauthorized(c, "User not authenticated")
		return
	}

	user, err := h.userService.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if err.Error() == "user not found" {
			NotFound(c, "User not found")
			return
		}
		InternalServerError(c, err.Error())
		return
	}

	Success(c, user)
}

// UpdateProfile 更新用户资料
// @Summary 更新用户资料
// @Description 更新当前登录用户的资料信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.UserUpdateRequest true "更新信息"
// @Success 200 {object} Response{data=model.UserResponse} "更新成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "用户不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/user/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		Unauthorized(c, "User not authenticated")
		return
	}

	var req model.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ValidationError(c, err)
		return
	}

	user, err := h.userService.UpdateProfile(c.Request.Context(), userID, &req)
	if err != nil {
		if err.Error() == "user not found" {
			NotFound(c, "User not found")
			return
		}
		InternalServerError(c, err.Error())
		return
	}

	SuccessWithMessage(c, "Profile updated successfully", user)
}

// GetUsers 获取用户列表（管理员功能）
// @Summary 获取用户列表
// @Description 分页获取用户列表
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} Response{data=model.PaginationResponse} "获取成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	var pagination model.PaginationRequest
	if err := c.ShouldBindQuery(&pagination); err != nil {
		ValidationError(c, err)
		return
	}

	users, err := h.userService.GetUserList(c.Request.Context(), &pagination)
	if err != nil {
		InternalServerError(c, err.Error())
		return
	}

	Success(c, users)
}

// GetUser 根据ID获取用户信息
// @Summary 获取指定用户信息
// @Description 根据用户ID获取用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} Response{data=model.UserResponse} "获取成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "用户不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		BadRequest(c, "Invalid user ID")
		return
	}

	user, err := h.userService.GetProfile(c.Request.Context(), uint(id))
	if err != nil {
		if err.Error() == "user not found" {
			NotFound(c, "User not found")
			return
		}
		InternalServerError(c, err.Error())
		return
	}

	Success(c, user)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} Response "删除成功"
// @Failure 400 {object} Response "参数错误"
// @Failure 401 {object} Response "未授权"
// @Failure 404 {object} Response "用户不存在"
// @Failure 500 {object} Response "服务器错误"
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		BadRequest(c, "Invalid user ID")
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), uint(id)); err != nil {
		if err.Error() == "user not found" {
			NotFound(c, "User not found")
			return
		}
		InternalServerError(c, err.Error())
		return
	}

	SuccessWithMessage(c, "User deleted successfully", nil)
}