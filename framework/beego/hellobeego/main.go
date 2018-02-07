package main

import (
	_ "Hello-Go/framework/beego/hellobeego/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
