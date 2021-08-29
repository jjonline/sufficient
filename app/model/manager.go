package model

// Manager 表模型manager表
type Manager struct {
	BaseField // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model     // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Manager) TableName() string {
	return dbPrefix() + "manager"
}
