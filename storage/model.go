package storage

import (
	"errors"
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

// GetBySrc根据 Src 获取记录 写回到c
func (c *Cache) GetBySrc(src string) (bool, error) {
	result := GetSqlite().Where("src = ?", src).First(&c)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
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
