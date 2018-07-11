package controllers

import "github.com/astaxie/beego"

type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.Data["content"] = "value"
	this.Layout = "admin/layout.html"
	this.TplName = "admin/add.tpl"
}


