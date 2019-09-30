package routers

import (
	"github.com/astaxie/beego"
	"word/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("hello/", &controllers.HelloControllers{})
	beego.Router("/login", &controllers.Logincontroller{}, "Get:Get;Post:Post")
	beego.Router("/cookie", &controllers.Cookiecontroller{})
	beego.Router("/Register", &controllers.Registercontrollers{})
	beego.Router("/Upload", &controllers.Uploadcontroller{})
}
