package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type UserListController struct {
	beego.Controller
}

func (u *UserListController) Get() {
	if !IsUserLoggedIn(&u.Controller) {
		// 用户未登录，返回错误信息或者跳转到登陆页面
		u.Redirect("/login", 302)
		return
	}
	o := orm.NewOrm()
	count, _ := o.QueryTable("user").Count()
	u.Data["count"] = count
	u.TplName = "users.html"
}

func (u *UserListController) GetAll() {
	var users []models.User
	o := orm.NewOrm()
	o.QueryTable("user").All(&users)
	var respList []interface{}
	for _, user := range users {
		respList = append(respList, user.UserToRespDesc())
	}
	u.Data["json"] = respList
	u.ServeJSON()
}

func IsUserLoggedIn(u *beego.Controller) bool {
	username := u.GetSession(USERNAME)
	if username == nil {
		return false
	}
	return true
}

func (u *UserListController) DeleteUser() {
	id, err := u.GetInt(":id") // 获取到url当中id变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		//u.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"}
		errResponse := util.NewError(506)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的Delete方法，&models.User{Id:id}表示删除是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.User{Id: id})
	if err != nil { //如果添加错误就返回信息：删除数据失败
		//u.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"}
		errResponse := util.NewError(514)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	u.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	u.ServeJSON()
}

func (u *UserListController) UpdateUser() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		//u.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	fmt.Println("username = ", user.Username, "password = ", user.Password, "id = ", user.Id)
	// 生成密码哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	fmt.Println("0 = > ", passwordHash)
	fmt.Println("err = ", err)
	if err != nil {
		//u.Data["json"] = models.UnifiedResponse{Code: 501, Message: "密码加密失败"}
		errResponse := util.NewError(505)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE user set Username = ?,Password = ? WHERE Id= ?", user.Username, string(passwordHash), user.Id)
	_, error := r.Exec()
	if error != nil {
		//u.Data["json"] = LoginResponse{Success: false, Message: "更新失败"}
		errResponse := util.NewError(516)
		u.Data["json"] = errResponse
	} else {
		u.Data["json"] = LoginResponse{Success: true, Message: "更新成功"}
	}
	u.ServeJSON()
}

func (u *UserListController) ShowUserByPage() {
	var pageinfo PageParam
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		//u.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	page := pageinfo.Pagenum
	pagesize := pageinfo.Pagesize
	users, err := models.GetUsers(page, pagesize)
	if err != nil {
		//u.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		errResponse := util.NewError(506)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	if len(users) > 0 {
		var repList []interface{}
		for _, user := range users {
			repList = append(repList, user.UserToRespDesc())
		}
		// 构建统一响应结构体
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取用户成功",
			Data:    repList,
		}
		u.Data["json"] = response
	}
	u.ServeJSON()
}

// 搜索用户名
func (u *UserListController) SearchUserName() {
	var username models.User
	var users []models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &username)
	fmt.Println("err = ", err)
	fmt.Println("username = ", username.Username)
	if err != nil {
		//u.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		u.Data["json"] = errResponse
		u.ServeJSON()
		return
	}
	o := orm.NewOrm()
	//o.QueryTable("user").Filter("Username", username.Username).All(&users)
	// 模糊查询
	sql := "SELECT * FROM user WHERE username LIKE ?"
	o.Raw(sql, "%"+username.Username+"%").QueryRows(&users)
	var respList []interface{}
	for _, user := range users {
		respList = append(respList, user.UserToRespDesc())
	}
	u.Data["json"] = respList
	u.ServeJSON()
}
