package entry

// CreateDeptReq 创建部门请求
// swagger:parameters CreateDeptReq
type CreateDeptReq struct {
	// 上级部门ID(顶级部门给0)
	// in: formData
	// required: true
	Pid *uint `form:"pid" json:"pid" binding:"required" comment:"上级部门ID"`
	// 部门名称（最大191字符）
	// in: formData
	// required: true
	Name string `form:"name" json:"name" binding:"required,max=191" comment:"部门名称"`
	// 自定义排序（大于0的数字，数字越小越靠前）
	// in: formData
	// required: false
	Sort uint `form:"sort" json:"sort" binding:"omitempty" comment:"排序"`
	// 备注描述（可空仅用于内部注释备注）
	// in: formData
	// required: false
	Remark string `form:"remark" json:"remark" binding:"omitempty,max=191" comment:"备注描述"`
}

// EditDeptReq 编辑部门请求
// swagger:parameters EditDeptReq
type EditDeptReq struct {
	// 被编辑的部门ID
	// in: formData
	// required: true
	ID uint `form:"id" json:"id" binding:"required" comment:"部门ID"`
	// 上级部门ID(顶级部门给0)
	// in: formData
	// required: true
	Pid *uint `form:"pid" json:"pid" binding:"required" comment:"上级部门ID"`
	// 部门名称（最大191字符）
	// in: formData
	// required: true
	Name string `form:"name" json:"name" binding:"required,max=191" comment:"部门名称"`
	// 自定义排序（大于0的数字，数字越小越靠前）
	// in: formData
	// required: false
	Sort uint `form:"sort" json:"sort" binding:"omitempty" comment:"排序"`
	// 备注描述（可空仅用于内部注释备注）
	// in: formData
	// required: false
	Remark string `form:"remark" json:"remark" binding:"omitempty,max=191" comment:"备注描述"`
}
