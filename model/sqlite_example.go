package model

import (
	"log"
	"time"

	"github.com/zhangyiming748/FastStorage/storage"
	"gorm.io/gorm"
)

type S_Example struct {
	Id        int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Content   string         `gorm:"size:512;comment:内容"`
	CreatedAt time.Time      // GORM 会自动管理这些时间字段
	UpdatedAt time.Time      // GORM 会自动管理这些时间字段
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除支持
}

func (e S_Example) Sync() {
	log.Printf("开始同步表结构")

	// 使用 GORM 自动迁移功能创建表
	if err := storage.GetSqlite().AutoMigrate(&S_Example{}); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}

// Create 创建一个新的 Example 记录
func (s S_Example) Create() error {
	result := storage.GetSqlite().Create(&s)
	return result.Error
}

// GetByID 根据 ID 获取 Example 记录
func (s *S_Example) GetByID(id int64) error {
	result := storage.GetSqlite().First(&s, id)
	return result.Error
}

// Update 更新 Example 记录
func (s *S_Example) Update() error {
	result := storage.GetSqlite().Save(&s)
	return result.Error
}

// Delete 删除 Example 记录
func (s *S_Example) Delete() error {
	result := storage.GetSqlite().Delete(&s)
	return result.Error
}

// GetAll 获取所有 Example 记录
func (s *S_Example) GetAll() ([]S_Example, error) {
	var examples []S_Example
	result := storage.GetSqlite().Find(&examples)
	return examples, result.Error
}
