package route

import "github.com/jjonline/sufficient/app/controller/manage"

// manageRoute 管理后台api路由定义
func manageRoute() {
	manageRoute := router.Group("manage")

	// 无需鉴权的路由
	manageRoute.Use()
	{
		manageRoute.GET("login", manage.SignController.Login)
	}

	// 需鉴权的路由
	manageRoute.Use()
	{
		manageRoute.GET("dept", manage.DeptController.List)
		manageRoute.GET("role", manage.RoleController.List)
		manageRoute.GET("manager", manage.ManagerController.List)
		manageRoute.GET("menu", manage.MenuController.List)
	}
}
