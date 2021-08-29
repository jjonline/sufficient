package model

// Menu 表模型menu表
type Menu struct {
	Name      string `json:"name"`
	Pid       int    `json:"pid"`
	Level     int    `json:"level"`
	Frontend  string `json:"frontend"`
	Icon      string `json:"icon"`
	Sort      int    `json:"sort"`
	Remark    string `json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Menu) TableName() string {
	return dbPrefix() + "menu"
}
