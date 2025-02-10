package main

import (
	"crmManaSystem/controllers"
	"crmManaSystem/models"
	_ "crmManaSystem/models"
	_ "crmManaSystem/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RunCommand()
	models.Init()
	beego.InsertFilter("/test/*", beego.BeforeRouter, controllers.MyFilter)
	beego.SetViewsPath("views")
	beego.Run()
}
