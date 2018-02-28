package main

type User struct {
	Id   int64  `orm:"aito"`
	Name string `orm:"size(20)"`
	Age  int8
}
