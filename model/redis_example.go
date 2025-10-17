package model

import (
	"context"
	"log"
	"time"

	"github.com/zhangyiming748/FastStorage/storage"
	"gorm.io/gorm"
)

type R_Example struct {
	Id        int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Content   string         `gorm:"size:512;comment:内容"`
	CreatedAt time.Time      // GORM 会自动管理这些时间字段
	UpdatedAt time.Time      // GORM 会自动管理这些时间字段
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除支持
}

func (r *R_Example) Sync() {
	log.Printf("开始同步表结构")
	
	// 使用 GORM 自动迁移功能创建表
	if err := storage.GetSqlite().AutoMigrate(&R_Example{}); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}

// Redis 相关操作示例

/*
这里仿照orm的方法实现redis的get和set方法，实际使用中可以根据需求自行扩展
关于 context 的说明：
1. context.Background() - 返回一个空的 Context，通常用作初始 Context 或在 main、init 等函数中
2. context 用于控制请求的超时、取消等操作
3. 在实际应用中，可以使用 context.WithTimeout 或 context.WithCancel 创建具有特定功能的 Context
*/

// 一些常用的 Redis 操作示例

// SetWithTimeout 设置键值对，带超时控制
func SetWithTimeout(key string, value interface{}, timeout time.Duration) error {
	// 创建一个具有超时控制的 context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // 确保在函数结束时释放资源
	
	return storage.GetRedis().Set(ctx, key, value, 0).Err()
}

// Set 设置键值对
func Set(key string, value interface{}) error {
	// context.Background() 创建一个空的 Context
	// 它是所有 Context 的根，不会被取消，没有截止时间，不能携带值
	ctx := context.Background()
	return storage.GetRedis().Set(ctx, key, value, 0).Err()
}

// Get 获取键对应的值
func Get(key string) (string, error) {
	ctx := context.Background()
	return storage.GetRedis().Get(ctx, key).Result()
}

// GetWithTimeout 获取键对应的值，带超时控制
func GetWithTimeout(key string, timeout time.Duration) (string, error) {
	// 创建一个具有超时控制的 context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // 确保在函数结束时释放资源
	
	return storage.GetRedis().Get(ctx, key).Result()
}

// Delete 删除键
func Delete(key string) error {
	ctx := context.Background()
	return storage.GetRedis().Del(ctx, key).Err()
}

// Exists 检查键是否存在
func Exists(key string) (bool, error) {
	ctx := context.Background()
	count, err := storage.GetRedis().Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}