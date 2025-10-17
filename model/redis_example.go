package model

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zhangyiming748/FastStorage/storage"
)

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

/*这里的key为string,filed和value为json(golang中表现为结构体形式)中的每一组key和value
比如
```json
{
  "name":"zhangsan",
  "age":18
}```
其中json手动设定一个key
第一组filed为name,value为zhangsan
第二组filed为age,value为18
*/
// HSet 将结构体数据存储到Redis哈希表中
func HSet(key string, value interface{}) error {
	// 将结构体序列化为JSON，然后解析为map[string]interface{}
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	// 使用HSet将所有字段存储到Redis哈希表中
	ctx := context.Background()
	return storage.GetRedis().HSet(ctx, key, fields).Err()
}

// HGet 获取哈希表中指定字段的值
func HGet(key, field string) (string, error) {
	ctx := context.Background()
	return storage.GetRedis().HGet(ctx, key, field).Result()
}

// HGetAll 获取哈希表中所有的字段和值
func HGetAll(key string) (map[string]string, error) {
	ctx := context.Background()
	return storage.GetRedis().HGetAll(ctx, key).Result()
}
