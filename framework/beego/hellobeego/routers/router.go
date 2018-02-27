package routers

import (
	"Hello-Go/framework/beego/hellobeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello/", &controllers.Hello{})
}
