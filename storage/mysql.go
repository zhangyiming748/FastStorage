package storage

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	// 分别定义数据库连接信息
	var (
		user     string = "root"      // 请替换为你的数据库用户名
		password string = "163453"    // 请替换为你的数据库密码
		host     string = "127.0.0.1" // 请替换为你的数据库主机地址
		port     int    = 3306        // 请替换为你的数据库端口号
		dbName   string = "test"      // 请替换为你的数据库名称
	)
	//CREATE DATABASE `test` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_unicode_ci';
	// 构建 DSN 字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Asia%%2FShanghai", user, password, host, port, dbName)

	var err error
	// 使用构建的 DSN 连接数据库
	engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}
	if err := engine.Ping(); err != nil {
		log.Fatalf("连接数据库出错:%v\n", err)
	}
	// 可选：显示 SQL 语句
	engine.ShowSQL(true)

	// 可选：设置连接池
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(100)

	log.Println("Successfully connected to database!")
}

func GetMysql() *xorm.Engine {
	return engine
}
