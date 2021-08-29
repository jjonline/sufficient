package model

// 菜单关联的路由api--一个菜单可以有多个route，一个route可被多个菜单关联（route多条记录）

// MenuRoute 表模型MenuRoute表
type MenuRoute struct {
	MenuID    uint   `json:"menu_id"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t MenuRoute) TableName() string {
	return dbPrefix() + "manager"
}
