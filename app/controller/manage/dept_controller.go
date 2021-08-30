package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/jjonline/sufficient/app/entry"
	"github.com/jjonline/sufficient/app/service"
	"github.com/jjonline/sufficient/render"
	"github.com/jjonline/sufficient/utils/convert"
)

// 管理后台部门（组织、公司）相关控制器
type deptController struct{}

// Create swagger:route POST /manage/dept Manage CreateDeptReq
// 创建部门
//
// 提交创建部门
//     Responses:
//       default: ErrorRes
//       200: BoolRes
func (m *deptController) Create(ctx *gin.Context) {
	var req entry.CreateDeptReq
	if err := ctx.ShouldBind(&req); err != nil {
		render.F(ctx, err)
		return
	}

	if err := service.DeptService.Create(ctx, req); err != nil {
		render.F(ctx, err)
		return
	}

	render.S(ctx, true)
}

// Edit swagger:route PUT /manage/dept Manage EditDeptReq
// 编辑部门
//
// 提交编辑部门
//     Responses:
//       default: ErrorRes
//       200: BoolRes
func (m *deptController) Edit(ctx *gin.Context) {
	var req entry.EditDeptReq
	if err := ctx.ShouldBind(&req); err != nil {
		render.F(ctx, err)
		return
	}

	if err := service.DeptService.Edit(ctx, req); err != nil {
		render.F(ctx, err)
		return
	}

	render.S(ctx, true)
}

// Delete swagger:route DELETE /manage/dept/:id Manage DelDeptNoReq
// 删除部门
//
// 提交删除部门
//     Responses:
//       default: ErrorRes
//       200: BoolRes
func (m *deptController) Delete(ctx *gin.Context) {
	deptID := convert.String(ctx.Param("id")).UInt()
	if deptID == 0 {
		render.F(ctx, render.InvalidParams)
		return
	}

	if err := service.DeptService.Delete(ctx, deptID); err != nil {
		render.F(ctx, err)
		return
	}

	render.S(ctx, true)
}

func (m *deptController) List(ctx *gin.Context) {
	render.S(ctx, "List")
}

func (m *deptController) Detail(ctx *gin.Context) {
	render.S(ctx, "Detail")
}
