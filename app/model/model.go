package model

import (
	"fmt"
	"github.com/jjonline/golang-backend/client"
)

type Model interface {
	tableName() string
}

type model struct {}

func (m model) tableName() string {
	fmt.Println("model")
	return "1"
}

// One 主键ID查询1条记录
// 参数：主键ID值、返回值model指针引用
func (m model) One(ID uint32, target interface{}) (err error) {
	return client.DB.Table((&m).tableName()).Find(target, ID).Error
}

func (m model) List(ID uint32, target interface{}) (err error) {
	return client.DB.Model(&m).Find(target, ID).Error
}
