package models

import "time"

// CommonField 表公共字段
type CommonField struct {
	// 创建时间戳
	CreateTime *time.Time `gorm:"column:create_time;default:null" json:"create_time"`
	// 修改时间戳
	UpdateTime *time.Time `gorm:"column:create_time;default:null" json:"update_time"`
}
