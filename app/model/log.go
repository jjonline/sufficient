package model

type Approval struct {
	// ID 主键
	ID uint32 `json:"id" gorm:"primaryKey"`
	// Title 名称
	AppID string `json:"type"`
	// 标题
	Title string `json:"title"`
	model
}

func (m Approval) tableName() string {
	return "ec_ad"
}