package storage

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite(root string) {
	// 创建数据目录
	err := os.MkdirAll(root, 0755)
	if err != nil {
		log.Fatal("无法创建数据目录:", err)
	}
	// 使用纯Go SQLite驱动连接数据库
	dbFile := filepath.Join(root, "translate.db")
	db, err := gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatal("无法连接到数据库:", err)
	}
	gormDB = db
	log.Println("成功连接到SQLite数据库")
}

func GetSqlite() *gorm.DB {
	return gormDB
}
