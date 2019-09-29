package routers

import (
	"word/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("hello/",&controllers.HelloControllers{})
}
