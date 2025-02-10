package controllers

import (
	"crmManaSystem/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func MyFilter(ctx *context.Context) {
	username := ctx.Input.Session("username")
	fmt.Println()
	redisUtil := util.NewRedisCache()
	username1, _ := redisUtil.Get("login_username")
	if username == nil && username1 == nil {
		//ctx.WriteString("没有登录,请前往登陆页面")
		ctx.Redirect(302, beego.URLFor("LoginController.Get"))
	}
}
