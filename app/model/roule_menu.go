package model

// 角色下的菜单

// RoleMenu 表模型RoleMenu表
type RoleMenu struct {
	RoleID    uint `json:"role_id"`
	MenuID    uint `json:"menu_id"`
	BaseField      // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model          // 引入基础通用方法
}

// TableName 返回表名称方法
func (t RoleMenu) TableName() string {
	return dbPrefix() + "manager"
}
