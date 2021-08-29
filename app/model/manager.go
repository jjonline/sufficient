package model

// Manager 表模型manager表
type Manager struct {
	Name      string `json:"name"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	IsRoot    uint8  `json:"is_root"`
	Enable    uint8  `json:"enable"`
	Remark    string `json:"remark"`
	BaseField        // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model            // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Manager) TableName() string {
	return dbPrefix() + "manager"
}
