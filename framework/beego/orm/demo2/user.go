package main

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int64  `orm:"aito"`
	Name string `orm:"size(20)"`
	Age  int8
}

func init() {
	orm.RegisterModelWithPrefix(tablePrefix, new(User))
}
