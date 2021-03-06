package model

// 定义所有模型实例的单例，模型变量instance
//   ① 本质是一个结构体零值
//   ② 因模型需要嵌入通用方法需要在init方法中调用不可导出的初始化方法进行模型初始化
//   ③ 模型结构体会在控制器、服务类乃至helper、entry方法中调用，模型定义不可反向应用这些位置的包，否则可能导致循环引用错误
var (
	DeptModel        Dept
	ManagerModel     Manager
	MenuModel        Menu
	RoleModel        Role
	ManagerDeptModel ManagerDept
	ManagerRoleModel ManagerRole
	ManagerOpenModel ManagerOpen
	MenuRouteModel   MenuRoute
	RoleMenuModel    RoleMenu
)

// init 模型初始化，设置通用方法依赖的模型本身
func init() {
	DeptModel.model.construct(&DeptModel)
	ManagerModel.model.construct(&ManagerModel)
	MenuModel.model.construct(&MenuModel)
	RoleModel.model.construct(&RoleModel)
	ManagerDeptModel.model.construct(&ManagerDeptModel)
	ManagerRoleModel.model.construct(&ManagerRoleModel)
	ManagerOpenModel.model.construct(&ManagerOpenModel)
	MenuRouteModel.model.construct(&MenuRouteModel)
	RoleMenuModel.model.construct(&RoleMenuModel)
}
