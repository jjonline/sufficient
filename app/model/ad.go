package model

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
	return "ec_ad"
}