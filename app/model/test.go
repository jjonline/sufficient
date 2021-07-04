package model

// Test 表模型test定义
type Test struct {
	Name      string
	Type      uint
	BaseField // 引入基础通用字段
	model     // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Test) TableName() string  {
	return dbPrefix() + "test"
}
