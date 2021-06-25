package model

import (
	"github.com/jjonline/golang-backend/client"
)

// Model 定义公共模型接口嵌入类提供统一标准方法
type Model interface {
	tableName() string
}

// region 定义model`基类`

type model struct{}

func (m model) tableName() string {
	return ""
}

// endregion

// FindByPrimary 单主键查询1条记录<仅支持单字段主键，多字段主键自助在子类model实现>
//  - 参数ID：需要查询的主键字段的值
// 	   自定义模型中主键字段通过tag标签 gorm:"primaryKey" 指定的即为主键
// 	   若未定义主键，则取定义模型中的第一个字段作为主键字段使用
//	   绝大多数情况下自增主键，其他情况下
//  - 参数target：查询结果集模型引用，形参为Model接口类型<注意：指针也是接口类型>，实际调用时需要传指针
//  - 参数target：查询结果集模型引用，形参为Model接口类型<注意：指针也是接口类型>，实际调用时需要传指针
func (m model) FindByPrimary(ID interface{}, target Model, fields ...map[string]string) (err error) {
	return client.DB.Table((&m).tableName()).Find(target, ID).Error
}

func (m model) List(ID uint32, target interface{}) (err error) {
	return client.DB.Model(&m).Find(target, ID).Error
}
