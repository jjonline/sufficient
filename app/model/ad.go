package model

import "fmt"

type Ad struct {
	// ID 主键
	ID uint32 `json:"id" gorm:"primaryKey"`
	// Title 名称
	Type uint8 `json:"type"`
	// 标题
	Title string `json:"title"`
	model
}

func (m Ad) tableName() string {
	fmt.Println("ec_ad")
	return "ec_ad"
}