package model

import (
	"log"
	"time"

	"github.com/zhangyiming748/FastStorage/storage"
)

type M_Example struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Content   string    `xorm:"comment('内容') VARCHAR(512)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (m *M_Example) Sync() {
	log.Printf("开始同步表结构")
	if err := storage.GetMysql().Sync2(M_Example{}); err != nil {
		log.Printf("同步表结构失败: %v", err)
	}
	log.Printf("同步表结构完成")
}
