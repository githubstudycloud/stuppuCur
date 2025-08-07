package service

import (
	"context"
	"fmt"

	"github.com/company/go-enterprise-template/internal/model"
	"github.com/company/go-enterprise-template/internal/repository"
	"github.com/company/go-enterprise-template/pkg/utils"
)

// UserService 用户服务接口
type UserService interface {
	Register(ctx context.Context, req *model.UserCreateRequest) (*model.UserResponse, error)
	Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error)
	GetProfile(ctx context.Context, userID uint) (*model.UserResponse, error)
	UpdateProfile(ctx context.Context, userID uint, req *model.UserUpdateRequest) (*model.UserResponse, error)
	GetUserList(ctx context.Context, pagination *model.PaginationRequest) (*model.PaginationResponse, error)
	DeleteUser(ctx context.Context, userID uint) error
}

// userService 用户服务实现
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Register 用户注册
func (s *userService) Register(ctx context.Context, req *model.UserCreateRequest) (*model.UserResponse, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to check username: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("username already exists")
	}

	// 检查邮箱是否已存在
	existingUser, err = s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Phone:    req.Phone,
		Status:   model.UserStatusActive,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user.ToResponse(), nil
}

// Login 用户登录
func (s *userService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.UserLoginResponse, error) {
	// 根据用户名或邮箱查找用户
	var user *model.User
	var err error

	// 首先尝试用户名
	user, err = s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	// 如果用户名未找到，尝试邮箱
	if user == nil {
		user, err = s.userRepo.GetByEmail(ctx, req.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to get user by email: %w", err)
		}
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// 检查用户状态
	if user.Status != model.UserStatusActive {
		return nil, fmt.Errorf("user account is not active")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// 更新最后登录时间
	if err := s.userRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		// 登录时间更新失败不影响登录流程，只记录日志
		// logger.Errorf("failed to update last login time: %v", err)
	}

	return &model.UserLoginResponse{
		User:  user.ToResponse(),
		Token: token,
	}, nil
}

// GetProfile 获取用户资料
func (s *userService) GetProfile(ctx context.Context, userID uint) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return user.ToResponse(), nil
}

// UpdateProfile 更新用户资料
func (s *userService) UpdateProfile(ctx context.Context, userID uint, req *model.UserUpdateRequest) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	// 更新字段
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user.ToResponse(), nil
}

// GetUserList 获取用户列表
func (s *userService) GetUserList(ctx context.Context, pagination *model.PaginationRequest) (*model.PaginationResponse, error) {
	pagination.GetDefaultPagination()

	users, total, err := s.userRepo.List(ctx, pagination.Page, pagination.PageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}

	// 转换为响应格式
	userResponses := make([]*model.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = user.ToResponse()
	}

	response := &model.PaginationResponse{
		List:     userResponses,
		Total:    total,
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
	}
	response.CalculateTotalPages()

	return response, nil
}

// DeleteUser 删除用户
func (s *userService) DeleteUser(ctx context.Context, userID uint) error {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return fmt.Errorf("user not found")
	}

	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}