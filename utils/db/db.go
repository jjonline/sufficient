package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jjonline/golang-backend/client"
	"gorm.io/gorm"
)

// Model 封装db操作
// 使用的gorm v2版本，文档：https://gorm.io/zh_CN/docs/index.html
type Model struct {
	Table interface{}
	DB    *gorm.DB
	ctx   context.Context
}

// W 定义where查询条件
type W struct {
	Field string
	Op    string
	Value interface{}
}

// DB 创建db对象
func DB(table interface{}, ctx context.Context, dbArr ...*gorm.DB) *Model {
	db := client.DB
	if len(dbArr) > 0 {
		db = dbArr[0]
	}
	return &Model{
		Table: table,
		DB:    db,
		ctx:   ctx,
	}
}

// model 获取当前操作的model对象
func (m *Model) model() *gorm.DB {
	return m.DB.Model(m.Table).WithContext(m.ctx)
}

// One 根据ID获取一条数据
func (m *Model) One(id interface{}, target interface{}) error {
	return m.model().Where("id = ?", id).First(target).Error
}

// First 根据Where获取一条数据
func (m *Model) First(wheres []W, target interface{}) error {
	return ToWhere(m.model(), wheres).First(target).Error
}

// Exist 判断是否存在一条数据
func (m *Model) Exist(wheres []W, field string) (bo bool, err error) {
	row := ToWhere(m.model(), wheres).Select(field).Limit(1).Row()
	var tmp interface{}
	err = row.Scan(&tmp)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}
		return
	}
	return true, nil
}

// List 获取分页列表
func (m *Model) List(wheres []W, offset, limit int, order string, target interface{}) error {
	return ToWhere(m.model(), wheres).Offset(offset).Limit(limit).Order(order).Find(target).Error
}

// Count 获取数量
func (m *Model) Count(wheres []W) (n int64, err error) {
	err = ToWhere(m.model(), wheres).Count(&n).Error
	return
}

// Columns 获取一列
func (m *Model) Columns(wheres []W, field string, order string, target interface{}) error {
	if order == "" {
		order = "id desc"
	}
	return ToWhere(m.model(), wheres).Order(order).Pluck(field, target).Error
}

// Insert 插入数据
// 支持使用数组形式批量插入，还支持map[string]interface{} 和 []map[string]interface{}{}方式
func (m *Model) Insert(data interface{}) error {
	return m.model().Create(data).Error
}

// Save 保存一条数据(插入或更新)
// todo 使用上还有问题，在使用时结构体不能含有时间字段！！目前还找到很好的解决办法，先延缓
func (m *Model) Save() error {
	return m.model().Save(m.Table).Error
}

// UpdateOne 根据ID更新数据
func (m *Model) UpdateOne(id interface{}, data map[string]interface{}) (int64, error) {
	db := m.model().Where("id=?", id).Limit(1).UpdateColumns(data)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Update 更新数据
func (m *Model) Update(wheres []W, data map[string]interface{}) (int64, error) {
	db := ToWhere(m.model(), wheres).UpdateColumns(data)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// DeleteOne 删除一条数据
func (m *Model) DeleteOne(id interface{}) (int64, error) {
	db := m.model().Where("id=?", id).Delete(m.Table)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// Delete 删除数据
func (m *Model) Delete(wheres []W) (int64, error) {
	db := ToWhere(m.model(), wheres).Delete(m.Table)
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// 将where传给db查询构造器
func (w W) toWhere(query *gorm.DB) *gorm.DB {
	switch w.Op {
	case "in":
		query = query.Where(fmt.Sprintf("%s in (?)", w.Field), w.Value)
	case "like":
		query = query.Where(fmt.Sprintf("%s like ?", w.Field), "%"+w.Value.(string)+"%")
	case "between":
		v := w.Value.([]interface{})
		query = query.Where(fmt.Sprintf("%s between ? and ?", w.Field), v[0], v[1])
	case "find_in_set":
		query = query.Where(fmt.Sprintf("find_in_set(?, %s)", w.Field), w.Value)
	case "raw":
		query = query.Where(w.Field, w.Value.([]interface{})...)
	default:
		query = query.Where(fmt.Sprintf("%s %s ?", w.Field, w.Op), w.Value)
	}
	return query
}

// ToWhere 轉換where條件
func ToWhere(query *gorm.DB, wheres []W) *gorm.DB {
	if wheres == nil {
		return query
	}
	for _, w := range wheres {
		query = w.toWhere(query)
	}
	return query
}
