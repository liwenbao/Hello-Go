package demo

import (
	"git.kkcoding.com/kk-golang/kk-lib/micro"
)

/*
	定义一个app类型，继承micro.App
	app的实例必须调用micro.Load，以加载app所能处理的任务
	micro.LoadPath方法可以通过配置文件来加载app。
	app中必须包含service的定义。否则无法处理请求。
	必须对app的Name字段正确赋值，其值必须以“/”开头。
*/
type DemoApp struct {
	micro.App
	DB *micro.Database

	DemoObject *DemoObject //DemoObject必须实现接口db.IObject

	DemoService1 *DemoService1
	DemoService2 *DemoService2
}
