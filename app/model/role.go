package model

// Role 表模型role表
type Role struct {
	Name      string `json:"name"`
	Enable    uint   `json:"enable"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Role) TableName() string {
	return dbPrefix() + "role"
}
