package storage

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite() *gorm.DB {
	// 创建数据目录
	dbPath := filepath.Join(".", "data", "faststorage.db")
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	// 连接到 SQLite 数据库，如果文件不存在则会自动创建
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite database:", err)
	}

	log.Println("Successfully connected to SQLite database with GORM.")
	gormDB = db
	return db
}

func GetSqlite() *gorm.DB {
	return gormDB
}
