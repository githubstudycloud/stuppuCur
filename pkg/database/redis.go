package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	
	"github.com/company/go-enterprise-template/internal/config"
	"github.com/company/go-enterprise-template/pkg/logger"
)

var RDB *redis.Client

// InitRedis 初始化Redis连接
func InitRedis(cfg *config.RedisConfig) error {
	RDB = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  4 * time.Second,
		IdleTimeout:  300 * time.Second,
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to ping redis: %w", err)
	}

	logger.Infof("Redis connected successfully, ping result: %s", pong)
	return nil
}

// GetRedis 获取Redis客户端
func GetRedis() *redis.Client {
	return RDB
}

// CloseRedis 关闭Redis连接
func CloseRedis() error {
	if RDB != nil {
		return RDB.Close()
	}
	return nil
}

// Set 设置键值对
func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return RDB.Set(ctx, key, value, expiration).Err()
}

// Get 获取值
func Get(ctx context.Context, key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

// Del 删除键
func Del(ctx context.Context, keys ...string) error {
	return RDB.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func Exists(ctx context.Context, keys ...string) (int64, error) {
	return RDB.Exists(ctx, keys...).Result()
}

// Expire 设置键的过期时间
func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return RDB.Expire(ctx, key, expiration).Err()
}

// TTL 获取键的剩余生存时间
func TTL(ctx context.Context, key string) (time.Duration, error) {
	return RDB.TTL(ctx, key).Result()
}

// HSet 设置哈希字段
func HSet(ctx context.Context, key string, values ...interface{}) error {
	return RDB.HSet(ctx, key, values...).Err()
}

// HGet 获取哈希字段值
func HGet(ctx context.Context, key, field string) (string, error) {
	return RDB.HGet(ctx, key, field).Result()
}

// HGetAll 获取哈希的所有字段和值
func HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return RDB.HGetAll(ctx, key).Result()
}

// HDel 删除哈希字段
func HDel(ctx context.Context, key string, fields ...string) error {
	return RDB.HDel(ctx, key, fields...).Err()
}

// LPush 从列表左侧推入元素
func LPush(ctx context.Context, key string, values ...interface{}) error {
	return RDB.LPush(ctx, key, values...).Err()
}

// RPush 从列表右侧推入元素
func RPush(ctx context.Context, key string, values ...interface{}) error {
	return RDB.RPush(ctx, key, values...).Err()
}

// LPop 从列表左侧弹出元素
func LPop(ctx context.Context, key string) (string, error) {
	return RDB.LPop(ctx, key).Result()
}

// RPop 从列表右侧弹出元素
func RPop(ctx context.Context, key string) (string, error) {
	return RDB.RPop(ctx, key).Result()
}

// LLen 获取列表长度
func LLen(ctx context.Context, key string) (int64, error) {
	return RDB.LLen(ctx, key).Result()
}

// SAdd 向集合添加成员
func SAdd(ctx context.Context, key string, members ...interface{}) error {
	return RDB.SAdd(ctx, key, members...).Err()
}

// SMembers 获取集合的所有成员
func SMembers(ctx context.Context, key string) ([]string, error) {
	return RDB.SMembers(ctx, key).Result()
}

// SIsMember 检查成员是否在集合中
func SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	return RDB.SIsMember(ctx, key, member).Result()
}

// SRem 从集合中移除成员
func SRem(ctx context.Context, key string, members ...interface{}) error {
	return RDB.SRem(ctx, key, members...).Err()
}

// ZAdd 向有序集合添加成员
func ZAdd(ctx context.Context, key string, members ...*redis.Z) error {
	return RDB.ZAdd(ctx, key, members...).Err()
}

// ZRange 获取有序集合指定范围的成员
func ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return RDB.ZRange(ctx, key, start, stop).Result()
}

// ZRangeWithScores 获取有序集合指定范围的成员和分数
func ZRangeWithScores(ctx context.Context, key string, start, stop int64) ([]redis.Z, error) {
	return RDB.ZRangeWithScores(ctx, key, start, stop).Result()
}

// ZRem 从有序集合中移除成员
func ZRem(ctx context.Context, key string, members ...interface{}) error {
	return RDB.ZRem(ctx, key, members...).Err()
}

// ZCard 获取有序集合的成员数量
func ZCard(ctx context.Context, key string) (int64, error) {
	return RDB.ZCard(ctx, key).Result()
}

// RedisHealthCheck Redis健康检查
func RedisHealthCheck() error {
	if RDB == nil {
		return fmt.Errorf("redis client is nil")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis ping failed: %w", err)
	}

	return nil
}