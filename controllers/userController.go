package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
)

type PageParam struct {
	Pagesize int `json:"pagesize"` // 每页显示多少条
	Pagenum  int `json:"pagenum"`  // 第几页
}

type UserController struct {
	beego.Controller
}

type MainController struct {
	beego.Controller
}

type BasicInfoController struct {
	beego.Controller
}

func (b *BasicInfoController) BasicInfo() {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user_info").Count()
	b.Data["count"] = count
	b.TplName = "basicInfo.html"
}

func (b *BasicInfoController) ShowBasicByPage() {
	var pageInfo PageParam
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &pageInfo)
	if err != nil {
		//b.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		b.Data["json"] = errResponse
		b.ServeJSON()
		return
	}
	page := pageInfo.Pagenum
	pageSize := pageInfo.Pagesize
	userInfos, err := models.GetBasicInfo(page, pageSize)
	if err != nil {
		//b.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		errResponse := util.NewError(506)
		b.Data["json"] = errResponse
		b.ServeJSON()
		return
	}
	if len(userInfos) > 0 {
		var respList []interface{}
		for _, userInfos := range userInfos {
			respList = append(respList, userInfos.UserInfoReSpDesc())
		}
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取客户信息成功",
			Data:    respList,
		}
		b.Data["json"] = response
	}
	b.ServeJSON()
}

type SearchClient struct {
	SearchClient string `json:"searchClient"`
}

// 搜索客户
func (b *BasicInfoController) SearchUser() {
	var searchClient SearchClient
	var userInfos []models.UserInfo
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &searchClient)
	fmt.Println("err = ", err)
	if err != nil {
		//b.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		b.Data["json"] = errResponse
		b.ServeJSON()
		return
	}
	o := orm.NewOrm()
	//_, err1 := o.QueryTable("user_info").Filter("ClientName", searchClient.SearchClient).All(&userInfos)
	// 模糊查询
	sql := "SELECT * FROM user_info WHERE client_name LIKE ?"
	_, err1 := o.Raw(sql, "%"+searchClient.SearchClient+"%").QueryRows(&userInfos)
	if err1 != nil {
		//b.Data["json"] = LoginResponse{Success: false, Message: "没有找到该订单"}
		errResponse := util.NewError(511)
		b.Data["json"] = errResponse
		b.ServeJSON()
		return
	}
	var respList []interface{}
	for _, client := range userInfos {
		respList = append(respList, client)
	}
	b.Data["json"] = respList
	b.ServeJSON()
}

// 展示基本信息
func (b *BasicInfoController) ShowBasicInfo() {
	var userInfos []models.UserInfo
	o := orm.NewOrm()
	o.QueryTable("userInfo").All(&userInfos)
	var respList []interface{}
	for _, userInfo := range userInfos {
		respList = append(respList, userInfo.UserInfoReSpDesc())
	}
	b.Data["json"] = respList
	b.ServeJSON()
}

// 查看联系人
type ShowContactsController struct {
	beego.Controller
}

func (s *ShowContactsController) ShowContactsGet() {
	s.TplName = "showContact.html"
}

var getInt int

// 获取被点击客户的id
func (s *ShowContactsController) GetContactsInfo() {
	getInt, _ = s.GetInt(":id")
	s.ServeJSON()
	s.TplName = "showContact.html"
}

// 展示客户信息
func (s *ShowContactsController) GetAllContactsInfo() {
	var contact models.Contacts
	o := orm.NewOrm()
	o.QueryTable("contacts").Filter("ContactID", getInt).One(&contact)
	var respList []interface{}
	respList = append(respList, contact)
	s.Data["json"] = respList
	s.ServeJSON()
	s.TplName = "showContact.html"
}

// 把一条联系人信息传到前端
func (s *ShowContactsController) ShowContactsPost() {
	var contact models.Contacts
	o := orm.NewOrm()
	o.QueryTable("contacts").Filter("ContactID", getInt).One(&contact)
	s.Data["json"] = contact
	s.ServeJSON()
}

// 订单查看
type OrderViewController struct {
	beego.Controller
}

func (o *OrderViewController) OrderView() {
	or := orm.NewOrm()
	count, _ := or.QueryTable("order_view").Count()
	o.Data["count"] = count
	o.TplName = "orderView.html"
}

var orderNumber string

func (o *OrderViewController) AddOrderInfoGet() {
	orderNumber = GenerateOrderNumber()
	o.Data["orderNumber"] = orderNumber
	o.TplName = "addOrderInfo.html"
	//o.ServeJSON()
}

// 展示所有订单
func (o *OrderViewController) ShowAllOrderView() {
	var orderViews []models.OrderView
	or := orm.NewOrm()
	or.QueryTable("order_view").All(&orderViews)
	var respList []interface{}
	for _, orderView := range orderViews {
		respList = append(respList, orderView)
	}
	o.Data["json"] = respList
	o.ServeJSON()
}

// 订单号结构体
type SearchOrderNumber struct {
	SearchOrderNumber string `json:"searchOrderNumber"`
}

// 搜索订单号
func (o *OrderViewController) SearchOrderNumber() {
	var searchOrderNumber SearchOrderNumber
	var orderViews []models.OrderView
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &searchOrderNumber)
	if err != nil {
		//o.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		o.Data["json"] = errResponse
		o.ServeJSON()
		return
	}
	or := orm.NewOrm()
	//_, err1 := or.QueryTable("order_view").Filter("OrderNumber", searchOrderNumber.SearchOrderNumber).All(&orderViews)
	// 模糊查询
	sql := "SELECT * FROM order_view WHERE order_number LIKE ?"
	_, err1 := or.Raw(sql, "%"+searchOrderNumber.SearchOrderNumber+"%").QueryRows(&orderViews)
	if err1 != nil {
		//o.Data["json"] = LoginResponse{Success: false, Message: "没有找到该订单"}
		errResponse := util.NewError(511)
		o.Data["json"] = errResponse
		o.ServeJSON()
		return
	}
	var respList []interface{}
	for _, orderView := range orderViews {
		respList = append(respList, orderView.OrderViewToRespDesc())
	}
	o.Data["json"] = respList
	o.ServeJSON()
}

// 订单分页
func (o *OrderViewController) ShowOrderByPage() {
	var pageInfo PageParam
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &pageInfo)
	if err != nil {
		//o.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		o.Data["json"] = errResponse
		o.ServeJSON()
		return
	}
	page := pageInfo.Pagenum
	pageSize := pageInfo.Pagesize
	orderViews, err := models.GetOrderInfo(page, pageSize)
	if err != nil {
		//o.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		errResponse := util.NewError(506)
		o.Data["json"] = errResponse
		o.ServeJSON()
		return
	}
	if len(orderViews) > 0 {
		var respList []interface{}
		for _, orderView := range orderViews {
			respList = append(respList, orderView.OrderViewToRespDesc())
		}
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取客户信息成功",
			Data:    respList,
		}
		o.Data["json"] = response
	}
	o.ServeJSON()
}

// 随机生成订单号
func GenerateOrderNumber() string {
	// 获取当前时间
	currentTime := time.Now()
	year := currentTime.Year()
	// 生成8位随机数字和大小写字母组合的字符串
	str := "0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	var temp string
	var orderNumber = fmt.Sprintf("%v", year)
	for i := 0; i < 8; i++ {
		temp = fmt.Sprintf("%c", str[rand.Intn(len(str))])
		orderNumber += temp
	}
	// 将年份转换为字符串并拼接在订单号前面
	return orderNumber
}

// 流失管理
type CustLossManagementController struct {
	beego.Controller
}

var getId int

func (c *CustLossManagementController) Get() {
	c.TplName = "custLossInfo.html"
}

func (c *CustLossManagementController) CustLossManagementGet() {
	getId, _ = c.GetInt(":id")
	var custLossManagement models.CustLossManagement
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &custLossManagement)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		c.Data["json"] = errResponse
		c.ServeJSON()
		return
	}
	c.TplName = "custLossInfo.html"
}

type CustLossManageController struct {
	beego.Controller
}

func (c *CustLossManageController) Get() {
	c.TplName = "custLossManagement.html"
	o := orm.NewOrm()
	count, _ := o.QueryTable("cust_loss_management").Count()
	c.Data["count"] = count
}

// 展示客户流失信息
func (c *CustLossManageController) ShowCustLossInfoPost() {
	var custLossManagements []models.CustLossManagement
	o := orm.NewOrm()
	o.QueryTable("cust_loss_management").All(&custLossManagements)
	var respList []interface{}
	for _, custLossManagement := range custLossManagements {
		respList = append(respList, custLossManagement.CustLossManagementToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

// 客户流失分页
func (c *CustLossManageController) ShowCustLossByPage() {
	var pageInfo PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageInfo)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		c.Data["json"] = errResponse
		c.ServeJSON()
		return
	}
	page := pageInfo.Pagenum
	pageSize := pageInfo.Pagesize
	custLosses, err := models.GetCustLossInfo(page, pageSize)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		errResponse := util.NewError(506)
		c.Data["json"] = errResponse
		c.ServeJSON()
		return
	}
	if len(custLosses) > 0 {
		var respList []interface{}
		for _, custLoss := range custLosses {
			respList = append(respList, custLoss.CustLossManagementToRespDesc())
		}
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取客户信息成功",
			Data:    respList,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}

// 临时结构体
type SearchUsername struct {
	SearchUsername string `json:"searchUsername"`
}

// 搜索流失客户姓名
func (c *CustLossManageController) SearchCustLossName() {
	var searchUsername SearchUsername
	var custLossManages []models.CustLossManagement
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &searchUsername)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		c.Data["json"] = errResponse
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	//_, err1 := o.QueryTable("cust_loss_management").Filter("CustomerName", searchUsername.SearchUsername).All(&custLossManages)
	// 模糊查询
	sql := "SELECT * FROM cust_loss_management WHERE customer_name LIKE ?"
	_, err1 := o.Raw(sql, "%"+searchUsername.SearchUsername+"%").QueryRows(&custLossManages)
	if err1 != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "没有找到该客户"}
		errResponse := util.NewError(511)
		c.Data["json"] = errResponse
		c.ServeJSON()
		return
	}
	var respList []interface{}
	for _, custLossManage := range custLossManages {
		respList = append(respList, custLossManage.CustLossManagementToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

func (u *UserController) Get() {
	u.TplName = "users.html"
}

func (m *MainController) Get() {
	m.TplName = "home.html"
}
