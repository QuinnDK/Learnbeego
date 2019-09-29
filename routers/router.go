package routers

import (
	"github.com/astaxie/beego"
	"word/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("hello/", &controllers.HelloControllers{})
	beego.Router("/login", &controllers.Logincontroller{}, "Get:Get;Post:Post")
}
