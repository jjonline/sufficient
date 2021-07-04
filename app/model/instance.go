package model

// 定义所有模型实例的单例，模型变量
var (
	TestModel Test
)

// init 模型初始化，设置通用方法依赖的模型本身
func init() {
	TestModel.model.construct(&TestModel)
}
