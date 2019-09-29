package controllers

import "github.com/astaxie/beego"

type Cookiecontroller struct {
	beego.Controller
}

func (this *Cookiecontroller) Get() {
	if this.Ctx.GetCookie("user") == "user" {
		this.Ctx.SetCookie("user", "admin")
		this.Ctx.WriteString("Cookie设置成功")
	} else {
		user := this.Ctx.GetCookie("user")
		this.Ctx.WriteString("user:" + user)
	}
}
