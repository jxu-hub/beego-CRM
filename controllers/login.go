package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
)

type LoginController struct {
	beego.Controller
}

var getRand string
var LoginId int

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	USERNAME = "username"
	LOGIN    = "login"
)

func (l *LoginController) Get() {
	getRandom := GetRandCaptcha()
	l.Data["captchaslice"] = getRandom
	l.TplName = "login.html"
}

func GetRandCaptcha() string {
	str := "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	var temp string
	var str1 = ""
	for i := 0; i < 4; i++ {
		temp = fmt.Sprintf("%c", str[rand.Intn(len(str))])
		str1 += temp
	}
	return str1
}

func (l *LoginController) Login() {
	username := l.GetString("username")
	password := l.GetString("password")
	captcha := l.GetString("captcha")

	if username == "" || password == "" {
		l.Data["json"] = map[string]interface{}{"code": 400, "msg": "用户名或密码不能为空"}
		l.ServeJSON()
		return
	}

	getRand = l.GetString("captchaslice")

	if captcha == "" {
		l.Data["json"] = map[string]interface{}{"code": 400, "msg": "验证码不能为空"}
		l.ServeJSON()
		return
	}
	if !strings.EqualFold(captcha, getRand) {
		l.Data["json"] = map[string]interface{}{"code": 400, "msg": "验证码错误"}
		l.ServeJSON()
	}

	o := orm.NewOrm()

	// 生成密码哈希
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Username: username,
		Password: string(passwordHash),
	}
	o.QueryTable("user").Filter("Username", username).One(&user)
	LoginId = user.Id
	// 验证密码是否匹配
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		l.Data["json"] = map[string]string{"error": "Invalid username or password."}
		l.ServeJSON()
		return
	}
	if err := o.Read(&user, "Username", "Password"); err == nil {
		redisCache := util.NewRedisCache()
		redisCache.Set("login_username", user.Username)
		l.SetSession(USERNAME, username)
		l.Redirect("/test/index", 302)
		l.Ctx.WriteString("登录成功")
		return
	} else {
		fmt.Println("登录失败")
		l.Redirect("/login", 302)
		l.Ctx.WriteString("登录失败")
	}
}
