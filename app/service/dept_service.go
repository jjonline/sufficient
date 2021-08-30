package service

import (
	"context"
	"github.com/jjonline/sufficient/app/entry"
	"github.com/jjonline/sufficient/app/model"
	"github.com/jjonline/sufficient/render"
)

// 部门服务类
type deptService struct{}

// Create 新建部门
func (s *deptService) Create(ctx context.Context, req entry.CreateDeptReq) error {
	var level uint8 = 1 // 部门所处级别，默认顶级部门级别为1上级部门id为0
	if *req.Pid != 0 {
		var pDept model.Dept
		if err := model.DeptModel.DeptByPid(ctx, *req.Pid, &pDept); err != nil {
			return render.ErrDefineWithMsg.Wrap(err, "上级部门不存在")
		}

		level = pDept.Level + 1
	}

	// insert one
	newDept := model.Dept{
		Name:   req.Name,
		Pid:    *req.Pid,
		Level:  level,
		Sort:   req.Sort,
		Remark: req.Remark,
	}

	fields := []string{
		"name",
		"pid",
		"level",
		"sort",
		"remark",
	}

	return model.DeptModel.InsertOne(ctx, &newDept, fields...)
}

// Edit 编辑部门
func (s *deptService) Edit(ctx context.Context, req entry.EditDeptReq) error {
	// 待编辑数据
	var eDept model.Dept
	if err := model.DeptModel.FindByPrimary(ctx, req.ID, &eDept); err != nil {
		return render.ErrDefineWithMsg.Wrap(err, "拟编辑部门数据不存在")
	}

	var level uint8 = 1 // 部门所处级别，默认顶级部门级别为1上级部门id为0
	if *req.Pid != 0 {
		var pDept model.Dept
		if err := model.DeptModel.DeptByPid(ctx, *req.Pid, &pDept); err != nil {
			return render.ErrDefineWithMsg.Wrap(err, "上级部门不存在")
		}
		level = pDept.Level + 1
	}

	// insert one
	upDept := model.Dept{
		Name:   req.Name,
		Pid:    *req.Pid,
		Level:  level,
		Sort:   req.Sort,
		Remark: req.Remark,
		BaseField: model.BaseField{
			ID: req.ID, // set update where condition
		},
	}

	// 指定字段，更新零值
	fields := []string{
		"name",
		"pid",
		"level",
		"sort",
		"remark",
	}

	_, err := model.DeptModel.UpdateOne(ctx, &upDept, fields...)
	return err
}

// Delete 删除部门
func (s *deptService) Delete(ctx context.Context, id uint) error {
	// 待删除数据
	var eDept model.Dept
	if err := model.DeptModel.FindByPrimary(ctx, id, &eDept); err != nil {
		return render.ErrDefineWithMsg.Wrap(err, "拟删除部门数据不存在")
	}

	_, err := model.DeptModel.DeleteByPrimary(ctx, id)
	return err
}
