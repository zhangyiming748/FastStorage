package main

import (
	"log"

	"github.com/zhangyiming748/FastStorage/model"
	"github.com/zhangyiming748/FastStorage/storage"
)

func main() {
	// 初始化 SQLite 数据库连接
	storage.SetSqlite()

	// 同步表结构
	t := new(model.S_Example)
	t.Sync()

	// 测试 CRUD 操作
	testCRUD()
}

func testCRUD() {
	// 创建测试数据
	example := model.S_Example{
		Content: "这是一个测试内容",
	}

	// 创建记录
	if err := example.Create(); err != nil {
		log.Printf("创建记录失败: %v", err)
	} else {
		log.Printf("成功创建记录，ID: %d", example.Id)
	}

	// 查询记录
	var retrieved model.S_Example
	if err := retrieved.GetByID(example.Id); err != nil {
		log.Printf("查询记录失败: %v", err)
	} else {
		log.Printf("成功查询记录: %+v", retrieved)
	}

	// 更新记录
	retrieved.Content = "这是更新后的内容"
	if err := retrieved.Update(); err != nil {
		log.Printf("更新记录失败: %v", err)
	} else {
		log.Printf("成功更新记录")
	}

	// 查询所有记录
	allExamples, err := retrieved.GetAll()
	if err != nil {
		log.Printf("查询所有记录失败: %v", err)
	} else {
		log.Printf("查询到 %d 条记录", len(allExamples))
		for _, ex := range allExamples {
			log.Printf("记录: %+v", ex)
		}
	}

	// 删除记录
	if err := retrieved.Delete(); err != nil {
		log.Printf("删除记录失败: %v", err)
	} else {
		log.Printf("成功删除记录")
	}
}
