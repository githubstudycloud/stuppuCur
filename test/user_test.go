package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
	"github.com/company/go-enterprise-template/internal/model"
	"github.com/company/go-enterprise-template/internal/service"
)

// MockUserRepository 模拟用户仓储
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) List(ctx context.Context, page, pageSize int) ([]*model.User, int64, error) {
	args := m.Called(ctx, page, pageSize)
	return args.Get(0).([]*model.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserRepository) UpdateLastLogin(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestUserService_Register(t *testing.T) {
	// 创建模拟仓储
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	// 测试数据
	req := &model.UserCreateRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Nickname: "Test User",
	}

	// 设置模拟期望
	mockRepo.On("GetByUsername", mock.Anything, req.Username).Return(nil, nil)
	mockRepo.On("GetByEmail", mock.Anything, req.Email).Return(nil, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*model.User")).Return(nil)

	// 执行测试
	ctx := context.Background()
	result, err := userService.Register(ctx, req)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Username, result.Username)
	assert.Equal(t, req.Email, result.Email)
	assert.Equal(t, req.Nickname, result.Nickname)

	// 验证模拟调用
	mockRepo.AssertExpectations(t)
}

func TestUserService_Login(t *testing.T) {
	// 创建模拟仓储
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	// 测试数据
	testUser := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "$2a$10$N9qo8uLOickgx2ZMRZoMye1S4c8J0fB8ksVfJ0iKdJ8g3k8k9F9iW", // password123 的哈希
		Status:   model.UserStatusActive,
	}
	testUser.ID = 1

	req := &model.UserLoginRequest{
		Username: "testuser",
		Password: "password123",
	}

	// 设置模拟期望
	mockRepo.On("GetByUsername", mock.Anything, req.Username).Return(testUser, nil)
	mockRepo.On("UpdateLastLogin", mock.Anything, testUser.ID).Return(nil)

	// 执行测试
	ctx := context.Background()
	result, err := userService.Login(ctx, req)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Token)
	assert.Equal(t, testUser.Username, result.User.Username)

	// 验证模拟调用
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetProfile(t *testing.T) {
	// 创建模拟仓储
	mockRepo := new(MockUserRepository)
	userService := service.NewUserService(mockRepo)

	// 测试数据
	userID := uint(1)
	testUser := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Nickname: "Test User",
		Status:   model.UserStatusActive,
	}
	testUser.ID = userID

	// 设置模拟期望
	mockRepo.On("GetByID", mock.Anything, userID).Return(testUser, nil)

	// 执行测试
	ctx := context.Background()
	result, err := userService.GetProfile(ctx, userID)

	// 断言
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testUser.Username, result.Username)
	assert.Equal(t, testUser.Email, result.Email)

	// 验证模拟调用
	mockRepo.AssertExpectations(t)
}