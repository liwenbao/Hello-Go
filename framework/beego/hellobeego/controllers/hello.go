package controllers

import (
	"github.com/astaxie/beego"
)

type Hello struct {
	beego.Controller
}

func (c *Hello) Get() {
	c.Ctx.WriteString("<h1>Hellow Beego</h1>")
}
