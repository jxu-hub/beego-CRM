package controllers

import (
	"crmManaSystem/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	var user models.User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Id", LoginId).One(&user)
	i.Data["loginUser"] = user.Username
	i.TplName = "index.html"
}
