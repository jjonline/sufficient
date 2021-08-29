package model

// 管理员所属的部门 --- 一个管理员可以属于多个部门

// ManagerOpen 表模型ManagerOpen表
type ManagerOpen struct {
	ManagerID     uint   `json:"manager_id"`
	OpenName      string `json:"open_name"`
	OpenType      uint8  `json:"open_type"`
	OpenID        string `json:"open_id"`
	AccessToken   string `json:"access_token"`
	TokenExpireIn uint   `json:"token_expire_in"`
	RefreshToken  string `json:"refresh_token"`
	Figure        string `json:"figure"`
	UnionID       string `json:"union_id"`
	Enable        uint8  `json:"enable"`
	BaseField            // 引入基础通用字段--主键ID、创建时间、更新时间、软删除时间(若有需要)
	model                // 引入基础通用方法
}

// TableName 返回表名称方法
func (t ManagerOpen) TableName() string {
	return dbPrefix() + "manager"
}
