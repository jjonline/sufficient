package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/jjonline/sufficient/app/entry"
	"github.com/jjonline/sufficient/render"
)

// 管理后台角色相关控制器
type roleController struct{}

// Create swagger:route POST /manage/role Manage CreateRoleReq
// 创建角色
//
// 提交创建角色数据
//     Responses:
//       default: ErrorRes
//       200: BoolRes
func (m *roleController) Create(ctx *gin.Context) {
	var req entry.CreateDeptReq
	if err := ctx.ShouldBind(&req); err != nil {
		render.F(ctx, err)
	}
}

func (m *roleController) Edit(ctx *gin.Context) {
	render.S(ctx, "Edit")
}

func (m *roleController) Delete(ctx *gin.Context) {
	render.S(ctx, "Delete")
}

func (m *roleController) List(ctx *gin.Context) {
	render.S(ctx, "List")
}

func (m *roleController) Detail(ctx *gin.Context) {
	render.S(ctx, "Detail")
}
