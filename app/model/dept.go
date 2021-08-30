package model

import "context"

// Dept 表模型Dept表
type Dept struct {
	Name      string `json:"name"`
	Pid       uint   `json:"pid"`
	Level     uint8  `json:"level"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Dept) TableName() string {
	return dbPrefix() + "dept"
}

// DeptByPid pid查询父部门
func (t Dept) DeptByPid(ctx context.Context, pid uint, dept *Dept) error {
	return t.model.FindByPrimary(ctx, pid, dept)
}
