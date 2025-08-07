package model

import (
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Username    string     `gorm:"uniqueIndex;size:50;not null" json:"username" binding:"required,min=3,max=50"`
	Email       string     `gorm:"uniqueIndex;size:100;not null" json:"email" binding:"required,email"`
	Password    string     `gorm:"size:255;not null" json:"-"`
	Nickname    string     `gorm:"size:100" json:"nickname"`
	Avatar      string     `gorm:"size:255" json:"avatar"`
	Phone       string     `gorm:"size:20" json:"phone"`
	Status      UserStatus `gorm:"default:1" json:"status"`
	LastLoginAt *time.Time `json:"last_login_at"`
	LoginCount  int        `gorm:"default:0" json:"login_count"`
}

// UserStatus 用户状态枚举
type UserStatus int

const (
	UserStatusInactive UserStatus = 0 // 未激活
	UserStatusActive   UserStatus = 1 // 激活
	UserStatusBanned   UserStatus = 2 // 被封禁
)

// UserCreateRequest 用户创建请求
type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
	Nickname string `json:"nickname" binding:"max=100"`
	Phone    string `json:"phone" binding:"max=20"`
}

// UserUpdateRequest 用户更新请求
type UserUpdateRequest struct {
	Nickname string `json:"nickname" binding:"max=100"`
	Avatar   string `json:"avatar" binding:"max=255"`
	Phone    string `json:"phone" binding:"max=20"`
}

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

// UserResponse 用户响应
type UserResponse struct {
	ID          uint       `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Nickname    string     `json:"nickname"`
	Avatar      string     `json:"avatar"`
	Phone       string     `json:"phone"`
	Status      UserStatus `json:"status"`
	LastLoginAt *time.Time `json:"last_login_at"`
	LoginCount  int        `json:"login_count"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// ToResponse 转换为响应格式
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Nickname:    u.Nickname,
		Avatar:      u.Avatar,
		Phone:       u.Phone,
		Status:      u.Status,
		LastLoginAt: u.LastLoginAt,
		LoginCount:  u.LoginCount,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// TableName 表名
func (User) TableName() string {
	return "users"
}