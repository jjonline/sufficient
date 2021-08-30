package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/jjonline/sufficient/app/entry"
	"github.com/jjonline/sufficient/render"
)

// 管理后台部门（组织、公司）相关控制器
type deptController struct{}

// Create swagger:route POST /manage/dept Manage CreateDeptReq
// 创建部门
//
// 提交创建部门数据
//     Responses:
// 		 default: ErrorRes
//       200: BoolRes
func (m *deptController) Create(ctx *gin.Context) {
	var req entry.CreateDeptReq
	if err := ctx.ShouldBind(&req); err != nil {
		render.F(ctx, err)
	}
}

func (m *deptController) Edit(ctx *gin.Context) {
	render.S(ctx, "Edit")
}

func (m *deptController) Delete(ctx *gin.Context) {
	render.S(ctx, "Delete")
}

func (m *deptController) List(ctx *gin.Context) {
	render.S(ctx, "List")
}

func (m *deptController) Detail(ctx *gin.Context) {
	render.S(ctx, "Detail")
}

