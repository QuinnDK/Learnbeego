package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type Registercontrollers struct {
	beego.Controller
}

func (this *Registercontrollers) Get() {
	this.TplName = "Register.html"
}

func (this *Registercontrollers) Post() {
	Name := this.GetString("name")
	Password := this.GetString("pwd")
	Sex := this.GetString("sex")
	Age, err := this.GetInt("age")
	if err == nil {
		this.Ctx.WriteString("Name=" + Name + "\nPwd=" + Password + "\nSex=" + Sex + "\nAge=" + strconv.Itoa(Age))
	} else {
		this.Ctx.WriteString("Name=" + Name + "\nPwd=" + Password + "\nSex=" + Sex)
	}
}
