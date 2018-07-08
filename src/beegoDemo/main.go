package main

import (
	_ "beegoDemo/routers"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController)Get(){
	this.Ctx.WriteString("hello world")
}


func main() {
	beego.Router("/main", &MainController{})
	beego.Run()
}

