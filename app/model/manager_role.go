package model

// 管理员所属的角色，一个管理员可以有多个角色

// ManagerRole 表模型ManagerRole表
type ManagerRole struct {
	ManagerID uint   `json:"manager_id"`
	RoleID    uint   `json:"role_id"`
	Sort      uint   `json:"sort"`
	Remark    string ` json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t ManagerRole) TableName() string {
	return dbPrefix() + "manager"
}
