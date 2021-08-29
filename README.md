# sufficient

日常开发过程中提炼封装的go语言开发框架，自带通用管理后台实现

## 一、项目描述

本项目规划了一套自定义项目结构，使用了 `gin`、`gorm`、`go-redis`等开源组件构造了一套自定义后端api开发框架，并自带了一个通用管理后台控制面板，便于新项目快速构建。

## 二、项目结构

项目目录树和介绍如下：

````
├── app                App应用目录
│   ├── console        终端控制台应用汇总
│   │   ├── command    终端命令应用
│   │   └── job        队列任务类job
│   ├── controller     控制器
│   ├── entry          实体类
│   ├── model          model模型
│   └── service        service服务类
├── client             全局句柄客户端
├── conf               全局配置
├── define             全局常量变量
├── migrations         数据库迁移文件
├── render             渲染器
├── route              路由
├── utils              utils帮助类
├── conf.toml.example  配置文件样例
└──  main.go           main包入口
````
## 三、项目约定

> 现代项目一般都强调：约定大于配置

1. 主键字段约定使用 `uint` 类型，因为现代部署机器绝大多数是64位，实际部署时`uint`将指向`uint64`

## 四、Model模型定义

1. 第一步，在`./app/model/`定义模型，示例见下方代码
2. 第二步，在`./app/model/instance.go`定义模型实例

````
package model

// Test 定义Test模型结构体，实现 schema.Tabler 接口即可
type Test struct {
	Name      string
	Type      uint
	BaseField // 引入基础通用字段
	model     // 引入基础通用方法
}

// TableName 返回表名称方法
func (t Test) TableName() string  {
	return dbPrefix() + "test"
}
````