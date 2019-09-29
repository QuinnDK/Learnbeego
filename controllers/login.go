package controllers

import "github.com/astaxie/beego"

type Logincontroller struct {
	beego.Controller
}

//使用GET方法
func (this *Logincontroller) Get() {
	this.TplName = "login.html" //解析login模板
}

//使用PUT方法
func (this *Logincontroller) Post() {
	user := this.GetString("user")
	pwd := this.GetString("pwd")
	this.Ctx.WriteString("user=" + user + "password=" + pwd)
}
