package controllers

import "github.com/astaxie/beego"

type Uploadcontroller struct {
	beego.Controller
}

func (this *Uploadcontroller) Get() {
	this.TplName = "Upload.html"
}
func (this *Uploadcontroller) Post() {
	file, header, err := this.GetFile("file")
	if err != nil {
		this.Ctx.WriteString("获取文件失败")
		return
	}
	defer file.Close()

	filename := header.Filename
	err = this.SaveToFile("file", "static/"+filename)
	if err != nil {
		this.Ctx.WriteString("上传失败")
	} else {
		this.Ctx.WriteString("上传成功")
	}
}
