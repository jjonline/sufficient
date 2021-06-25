package model

import (
	"fmt"
	"github.com/jjonline/golang-backend/client"
	"gorm.io/gorm"
)

// ModelIFace 定义公共模型接口嵌入类提供统一标准方法
type ModelIFace interface {
	tableName() string
}

// region 定义model`基类`

type model struct{}

func (m model) tableName() string {
	return ""
}

// endregion

// region 定义where查询条件结构

// Where 定义where查询条件
type Where struct {
	Field string
	Op    string
	Value interface{}
}

// toWhere 将where传给db查询构造器
func (w Where) toWhere(query *gorm.DB) *gorm.DB {
	switch w.Op {
	case "IN":
		query = query.Where(fmt.Sprintf("`%s` IN (?)", w.Field), w.Value)
	case "NOT IN":
		query = query.Where(fmt.Sprintf("`%s` NOT IN (?)", w.Field), w.Value)
	case "BETWEEN":
		v := w.Value.([]interface{})
		query = query.Where(fmt.Sprintf("`%s` BETWEEN ? and ?", w.Field), v[0], v[1])
	case "NOT BETWEEN":
		v := w.Value.([]interface{})
		query = query.Where(fmt.Sprintf("`%s` NOT BETWEEN ? and ?", w.Field), v[0], v[1])
	case "LIKE":
		query = query.Where(fmt.Sprintf("`%s` LIKE ?", w.Field), "%"+w.Value.(string)+"%")
	case "FIND_IN_SET":
		query = query.Where(fmt.Sprintf("FIND_IN_SET(?, %s)", w.Field), w.Value)
	case "RAW":
		query = query.Where(w.Field, w.Value.([]interface{})...)
	default:
		// "<>", "!=", ">", ">=", "=", "<", "<="
		query = query.Where(fmt.Sprintf("`%s` %s ?", w.Field, w.Op), w.Value)
	}
	return query
}

// parseWhere 解析where条件
func parseWhere(query *gorm.DB, wheres []Where) *gorm.DB {
	if wheres == nil {
		return query
	}
	for _, w := range wheres {
		query = w.toWhere(query)
	}
	return query
}

// endregion

// FindByPrimary 单主键查询1条记录<仅支持单字段主键不支持符合主键>
//  - 查询不到记录返回 gorm.ErrRecordNotFound 的error
// 	   使用errors.Is(result.Error, gorm.ErrRecordNotFound)进行判断
//  - pVal：Primary Value 需要查询的主键字段的值<仅支持数值类型主键，字符串类型不支持>
// 	   自定义模型中主键字段通过tag标签 gorm:"primaryKey" 指定的即为主键
// 	   若未定义主键，则取定义模型中的第一个字段作为主键字段使用
//	   绝大多数情况下自增主键可使用；复合主键或字符串主键请使用 FindByWhere
//  - target：查询结果集模型引用，形参为 ModelIFace 接口类型<注意：指针也是接口类型>，需传指针
//  - fields：查询的字段，可选不传表示默认查询所有
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) FindByPrimary(pVal interface{}, target ModelIFace, fields ...string) (err error) {
	if len(fields) > 0 {
		return client.DB.Table((&m).tableName()).Select(fields).Take(target, pVal).Error
	}
	return client.DB.Table((&m).tableName()).Take(target, pVal).Error
}

// FindByWhere 按where条件查询1条记录
//  - 查询不到记录返回 gorm.ErrRecordNotFound 的error
// 	   使用errors.Is(result.Error, gorm.ErrRecordNotFound)进行判断
//  -
func (m model) FindByWhere(where []Where, target ModelIFace, fields ...string) (err error) {
	if len(fields) > 0 {
		return parseWhere(client.DB.Table((&m).tableName()), where).Select(fields).Take(target).Error
	}
	return parseWhere(client.DB.Table((&m).tableName()), where).Take(target).Error
}

func (m model) List(ID uint32, target interface{}) (err error) {
	return client.DB.Model(&m).Find(target, ID).Error
}
