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

// endregion

// region 定义where查询条件结构

// Where 定义where查询条件
//  切片Where多个条件最终SQL中将使用 AND 符拼接
//  - Field 查询的字段名称字符串
//  - Op    查询的条件[IN、NOT IN、BETWEEN、NOT BETWEEN、LIKE、FIND_IN_SET、=、<>、>=、<、<=、RAW]
//    当使用RAW时将忽略Op值，直接使用gorm提供的原生Where方法构建<注意转义特殊字符特别留意SQL注入风险>
//  - Value 查询的条件值，当Op为LIKE时无需添加前后的百分号<%>，方法体自动添加前后百分号
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
		query = query.Where(fmt.Sprintf("`%s` BETWEEN ? AND ?", w.Field), v[0], v[1])
	case "NOT BETWEEN":
		v := w.Value.([]interface{})
		query = query.Where(fmt.Sprintf("`%s` NOT BETWEEN ? AND ?", w.Field), v[0], v[1])
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
//     使用errors.Is(result.Error, gorm.ErrRecordNotFound)进行判断
//  - pVal：Primary Value 需要查询的主键字段的值<仅支持数值类型主键，字符串类型不支持>
// 	   自定义模型中主键字段通过tag标签 gorm:"primaryKey" 指定的即为主键
// 	   若未定义主键，则取定义模型中的第一个字段作为主键字段使用
//	   绝大多数情况下自增主键可使用；复合主键或字符串主键请使用 FindByWhere
//  - target：查询结果集模型引用，形参为 ModelIFace 接口类型<注意：指针也是接口类型>，需传指针
//  - fields：查询的字段，可选不传表示默认查询出所有字段
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) FindByPrimary(pVal interface{}, target ModelIFace, fields ...string) (err error) {
	return client.DB.Table(target.tableName()).Select(fields).Take(target, pVal).Error
}

// FindByWhere 按where条件查询1条记录<若有多条记录符合要求取按主键升序的第一条记录>
//  - 查询不到记录返回 gorm.ErrRecordNotFound 的error
// 	   使用errors.Is(result.Error, gorm.ErrRecordNotFound)进行判断
//  - where：查询条件 Where 结构体切片
//  - target：查询结果集模型引用，形参为 ModelIFace 接口类型<注意：指针也是接口类型>，需传指针
//  - fields：查询的字段，可选不传表示默认查询出所有字段
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) FindByWhere(where []Where, target ModelIFace, fields ...string) (err error) {
	return parseWhere(client.DB.Table(target.tableName()), where).Select(fields).First(target).Error
}

// CountByWhere 分页查询，按where条件分页查询获取分页列表数和总记录数
//  - 查询不到记录返回返回0和空切片
//  - model：查询的model实例对象
// 	   例如：var a Ad 传参 a
//  - where：查询条件 Where 结构体切片
//  - targetTotal：查询结果集总条数指针引用
func (m model) CountByWhere(model ModelIFace, where []Where, targetTotal *int64) (err error) {
	return parseWhere(client.DB.Table(model.tableName()), where).Count(targetTotal).Error
}

// ListByWhere 按where条件查询多条确认数量有限的列表记录
//	特别注意：该方法用于查询有限数量的多条记录，不设置limit条件
//  - 查询不到记录返回 gorm.ErrRecordNotFound 的error
// 	   使用errors.Is(result.Error, gorm.ErrRecordNotFound)进行判断
//  - model：查询的model实例对象
// 	   例如：var a Ad 传参 a
//  - where：查询条件 Where 结构体切片
//  - target：查询结果集模型的切片引用，形参为 interface，需要传具体model实现<ModelIFace>的切片引用：需传指针
//     例如：var a []Ad 传参 &a
//  - orderBy：排序条件字符串
//     例子1：`name` <表示按name字段升序>
//     例子2：`name` ASC <表示按name字段升序>
//     例子3：`name` ASC, `ID` DESC <表示按name字段升序后按ID降序>
//     例子4：给空字符串表示不设置排序条件
//  - fields：查询的字段，可选不传表示默认查询出所有字段
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) ListByWhere(model ModelIFace, where []Where, target interface{}, orderBy string, fields ...string) (err error) {
	if orderBy == "" {
		return parseWhere(client.DB.Table(model.tableName()), where).Select(fields).Take(target).Error
	}
	return parseWhere(client.DB.Table(model.tableName()), where).Select(fields).Order(orderBy).Take(target).Error
}

// Paginate 分页查询，按where条件分页查询获取分页列表数和总记录数
//  - 查询不到记录返回返回0和空切片
//  - model：查询的model实例对象
// 	   例如：var a Ad 传参 a
//  - where：查询条件 Where 结构体切片
//  - target：查询结果集模型的切片引用，形参为 interface，需要传具体model实现<ModelIFace>的切片引用：需传指针
//     例如：var a []Ad 传参 &a
//  - targetTotal：查询结果集总条数指针引用
//  - page：查询分页的当前页码数，从1开始
//  - limit：查询分页的1页多少条限制
//  - orderBy：排序条件字符串
//     例子1：`name` <表示按name字段升序>
//     例子2：`name` ASC <表示按name字段升序>
//     例子3：`name` ASC, `ID` DESC <表示按name字段升序后按ID降序>
//     例子4：给空字符串表示不设置排序条件
//  - fields：查询的字段，可选不传表示默认查询出所有字段
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) Paginate(model ModelIFace, where []Where, target interface{}, targetTotal *int64, page, limit int, orderBy string, fields ...string) (err error) {
	offset := 0
	if page > 0 {
		offset = (page - 1) * limit
	}

	// calc total count
	_ = parseWhere(client.DB.Table(model.tableName()), where).Count(targetTotal).Error

	if orderBy == "" {
		return parseWhere(client.DB.Table(model.tableName()), where).Offset(offset).Limit(limit).Select(fields).Find(target).Error
	}
	return parseWhere(client.DB.Table(model.tableName()), where).Offset(offset).Limit(limit).Select(fields).Order(orderBy).Find(target).Error
}

// SimplePaginate 简单分页查询，按where条件分页简要查询获取分页列表数据<不返回总条数减少1次查询>
//  - 查询不到记录返回返回空切片
//  - model：查询的model实例对象
// 	   例如：var a Ad 传参 a
//  - where：查询条件 Where 结构体切片
//  - target：查询结果集模型的切片引用，形参为 interface，需要传具体model实现<ModelIFace>的切片引用：需传指针
//     例如：var a []Ad 传参 &a
//  - page：查询分页的当前页码数，从1开始
//  - limit：查询分页的1页多少条限制
//  - orderBy：排序条件字符串
//     例子1：`name` <表示按name字段升序>
//     例子2：`name` ASC <表示按name字段升序>
//     例子3：`name` ASC, `ID` DESC <表示按name字段升序后按ID降序>
//     例子4：给空字符串表示不设置排序条件
//  - fields：查询的字段，可选不传表示默认查询出所有字段
// 	   可变参数查询多个字段，使用字符串切片可变参数展开模式可使用字符串切片数组
//	   []string{"name", "sex"}...
func (m model) SimplePaginate(model ModelIFace, where []Where, target interface{}, page, limit int, orderBy string, fields ...string) (err error) {
	offset := 0
	if page > 0 {
		offset = (page - 1) * limit
	}

	if orderBy == "" {
		return parseWhere(client.DB.Table(model.tableName()), where).Offset(offset).Limit(limit).Select(fields).Find(target).Error
	}
	return parseWhere(client.DB.Table(model.tableName()), where).Offset(offset).Limit(limit).Select(fields).Order(orderBy).Find(target).Error
}

// InsertOne 通过赋值模型创建1条记录
//  - data：model实例对象，请注意不要给主键字段赋值，创建成功后主键字段将填充创建记录的主键值
// 	   例如：var a Ad; a.XX="yy"; 传参 a
//  - fields：限定创建语句写入的字段名，不传留空则取结构体非零字段创建1条新记录
//	   []string{"name", "sex"}...
func (m model) InsertOne(data ModelIFace, fields ...string) error {
	return client.DB.Table(data.tableName()).Select(fields).Create(&data).Error
}

// InsertOneUseMap 通过map键值对创建1条记录
//  - 注意：创建成功后不会回填主键
//  - model：model实例对象
// 	   例如：var a Ad; a.XX="yy"; 传参 a
//  - data：新增1条记录的map数据
//  - where：查询条件 Where 结构体切片
func (m model) InsertOneUseMap(model ModelIFace, data map[string]interface{}) error {
	return client.DB.Table(model.tableName()).Create(data).Error
}

// MultiInsert 通过赋值模型批量创建
//  - model：model实例对象
//  - data：model实例对象切片，请注意不要给主键字段赋值，创建成功后主键字段将填充创建记录的主键值
// 	   例如：var a Ad; a.XX="yy"; 传参 a
//  - fields：限定创建语句写入的字段名，不传留空则取结构体非零字段创建1条新记录
//	   []string{"name", "sex"}...
func (m model) MultiInsert(model ModelIFace, data interface{}, fields ...string) error {
	return client.DB.Table(model.tableName()).Select(fields).Create(data).Error
}

// MultiInsertUseMap 通过map键值对切片批量创建
//  - 注意：创建成功后不会回填主键
//  - model：model实例对象
// 	   例如：var a Ad; a.XX="yy"; 传参 a
//  - data：新增1条记录的map数据
//  - where：查询条件 Where 结构体切片
func (m model) MultiInsertUseMap(model ModelIFace, data []map[string]interface{}) error {
	return client.DB.Table(model.tableName()).Create(data).Error
}