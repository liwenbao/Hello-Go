package main

import (
	//"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//注册ORM驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//注册ORM数据库，必须注册default数据库
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:5566)/study?charset=utf8&sql_mode=ANSI")

	//使用前缀注册模型
	orm.RegisterModelWithPrefix("beego_orm_demo_", new(User))

	//开启ORM日志
	orm.Debug = true

	//监听ORM命令行。可通过./demo1 orm syncdb -force=false -v=true 命令来实现自动同步数据库的操作。
	orm.RunCommand()

	//	//自动同步数据库
	//	orm.RunSyncdb("default", false, true)
}

func main() {
	log.Println("hello beego orm")
	o := orm.NewOrm()
	u1 := &User{Name: "zhangsan", Age: 18}
	ret, err := o.Insert(u1)
	if err != nil {
		log.Println("insert user fail, error is", err)
	} else {
		log.Println("insert user successfully, ret is", ret)
	}

	//fmt.Scanln(new(string))
}
