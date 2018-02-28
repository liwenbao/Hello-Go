//Demo2
//Deego Orm 基本CRUD操作
package main

import (
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dataBaseSource = "root:@tcp(127.0.0.1:5566)/study?charset=utf8&sql_mode=ANSI"
	tablePrefix    = "beego_orm_demo2_"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataBaseSource)
	orm.Debug = true
}

func main() {
	log.Println("hello beego orm demo2")
	orm.RunCommand()
	ormer := orm.NewOrm()

	//Create
	u1 := new(User)
	u1.Name = "zhang san"
	u1.Age = 18
	ormer.Insert(u1) //创建后会自动对 auto 的 field 赋值
	log.Println("u1", u1)

	//Read
	u2 := new(User)
	u2.Id = u1.Id
	ormer.Read(u2)
	log.Println("u2", u2) //SELECT `id`, `name`, `age` FROM `beego_orm_demo2_user` WHERE `id` = ? ] - `4`

	//Read 制定列为查询条件
	u2.Id = 0
	u2.Age = 0
	ormer.Read(u2, "name") //[SELECT `id`, `name`, `age` FROM `beego_orm_demo2_user` WHERE `name` = ? ] - `zhang san`
	log.Println("u2", u2)

	//Update
	u1.Age += 1
	ormer.Update(u1)
	u3 := new(User)
	u3.Id = u1.Id
	ormer.Read(u3)
	log.Println("u3", u3)

	//Delete
	ormer.Delete(u1)
	u4 := new(User)
	u4.Id = u1.Id
	ormer.Read(u4)
	log.Println("u4", u4)

	//ReadOrCreate
	u5 := &User{Name: "li si"}
	create, _, _ := ormer.ReadOrCreate(u5, "name")
	if create {
		log.Println("u5 create", u5)
	} else {
		log.Println("u5 read", u5)
	}
	ormer.Delete(u5)

	//InsertMulti
	mds := []*User{
		&User{Name: "U1", Age: 18},
		&User{Name: "U2", Age: 18},
		&User{Name: "U3", Age: 18}}
	ormer.InsertMulti(2, mds) //参数bulk设置为2，即每次插入2个，最后一次插入剩余1个。

	rawSeter := ormer.Raw("delete from " + tablePrefix + "user")
	rawSeter.Exec()
}
