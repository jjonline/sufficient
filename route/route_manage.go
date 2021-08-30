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
		// 部门
		manageRoute.GET("dept", manage.DeptController.List)
		manageRoute.GET("dept/:id", manage.DeptController.Detail)
		manageRoute.POST("dept", manage.DeptController.Create)
		manageRoute.PUT("dept", manage.DeptController.Edit)
		manageRoute.DELETE("dept/:id", manage.DeptController.Delete)

		// 角色
		manageRoute.GET("role", manage.RoleController.List)
		manageRoute.GET("role/:id", manage.RoleController.Detail)
		manageRoute.POST("role", manage.RoleController.Create)
		manageRoute.PUT("role", manage.RoleController.Edit)
		manageRoute.DELETE("role/:id", manage.RoleController.Delete)

		// 管理员
		manageRoute.GET("manager", manage.ManagerController.List)

		// 菜单
		manageRoute.GET("menu", manage.MenuController.List)
	}
}
