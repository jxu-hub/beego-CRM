package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {

	username := r.GetString("username")
	password := r.GetString("password")
	again_password := r.GetString("again_password")

	// 生成密码哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println("0 = > ", passwordHash)

	if err != nil {
		//r.Data["json"] = models.UnifiedResponse{Code: 501, Message: "密码加密失败"}
		//errResponse := util.NewError(505)
		r.Ctx.WriteString("密码加密失败")
		r.ServeJSON()
		return
	}

	if username == "" || password == "" {
		r.Ctx.WriteString("用户名或密码不能为空")
		return
	}

	if password != again_password {
		r.Ctx.WriteString("两次密码不一致")
		return
	}

	o := orm.NewOrm()
	user := models.User{
		Username: username,
		Password: string(passwordHash),
	}

	// 判断用户名是否存在
	err1 := o.Read(&user, "UserName")
	if err1 == nil {
		r.Ctx.WriteString("用户名已存在")
		return
	}
	var response util.APIResponse

	// 插入数据
	_, err1 = o.Insert(&user)
	if err1 != nil {
		response = util.JSONResponse(502, "注册失败")
		//r.Data["json"] = LoginResponse{Success: false, Message: "用户注册失败"}
		r.ServeJSON()
		return
	}
	response = util.JSONResponse(200, "注册成功")
	//r.Data["json"] = LoginResponse{Success: true, Message: "用户注册成功"}
	r.Redirect("/login", 302)
	r.Data["json"] = response
	r.ServeJSON()
}
