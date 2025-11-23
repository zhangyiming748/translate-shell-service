package storage

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Cache struct {
	Id        int64          `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Src       string         `gorm:"size:512;comment:原文"`
	Dst       string         `gorm:"size:512;comment:译文"`
	CreatedAt time.Time      // GORM 会自动管理这些时间字段
	UpdatedAt time.Time      // GORM 会自动管理这些时间字段
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除支持
}

func (c *Cache) Sync() {
	log.Printf("开始同步表结构")

	// 使用 GORM 自动迁移功能创建表
	if err := GetSqlite().AutoMigrate(c); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}

// Create 创建一个新的 Example 记录
func (c *Cache) Create() error {
	result := GetSqlite().Create(&c)
	return result.Error
}

// GetByID 根据 ID 获取 Example 记录
func (c *Cache) GetBySrc(id int64) error {
	result := GetSqlite().First(&c, id)
	return result.Error
}

// Update 更新 Example 记录
func (c *Cache) Update() error {
	result := GetSqlite().Save(&c)
	return result.Error
}

// Delete 删除 Example 记录
func (c *Cache) Delete() error {
	result := GetSqlite().Delete(&c)
	return result.Error
}

// GetAll 获取所有 Example 记录
func (c *Cache) GetAll() ([]Cache, error) {
	var caches []Cache
	result := GetSqlite().Find(&caches)
	return caches, result.Error
}
