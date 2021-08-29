package model

// 管理员所属的部门 --- 一个管理员可以属于多个部门

// ManagerDept 表模型ManagerDept表
type ManagerDept struct {
	BaseField // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model     // 引入基础通用方法
}

// TableName 返回表名称方法
func (t ManagerDept) TableName() string {
	return dbPrefix() + "manager"
}
