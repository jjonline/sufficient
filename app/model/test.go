package model

type Test struct {
	Name      string
	Type      uint
	BaseField // 引入基础通用字段
	model     // 引入基础通用方法
}

// init 模型初始化，设置通用方法依赖的模型本身
func init() {
	TestModel.model.construct(&TestModel)
}

// TableName 返回表名称方法
func (t Test) TableName() string  {
	return dbPrefix() + "test"
}
