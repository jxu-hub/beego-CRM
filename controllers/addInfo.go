package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type AddInfoController struct {
	beego.Controller
}

func (a *AddInfoController) UserMessageAddGet() {
	//o := orm.NewOrm()
	//userInfo := []models.UserInfo{}
	//o.QueryTable(new(models.UserInfo)).RelatedSel().All(&userInfo)
	a.TplName = "addUserMessage.html"
}

func (a *AddInfoController) AddUserMessagePost() {
	var userInfo models.UserInfo
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var contacts models.Contacts
	err = json.Unmarshal(a.Ctx.Input.RequestBody, &contacts)
	_, err = o.Insert(&userInfo)
	_, err1 := o.Insert(&contacts)
	err = o.QueryTable("user_info").Filter("client_name", userInfo.ClientName).One(&userInfo)
	customers := models.Customers{
		CustomerName: userInfo.ClientName,
		ContactID:    &contacts,
	}
	_, err2 := o.Insert(&customers)
	if err != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "客户信息插入失败"}
		errResponse := util.NewError(507)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	if err1 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "联系人信息插入失败"}
		errResponse := util.NewError(508)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	fmt.Println("err2 = ", err2)
	if err2 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "客户姓名插入失败"}
		errResponse := util.NewError(509)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	a.Data["json"] = LoginResponse{Success: true, Message: "添加成功"}
	a.ServeJSON()
}

func (a *AddInfoController) AddCustLossInfo() {
	var custLossManagement models.CustLossManagement
	err2 := json.Unmarshal(a.Ctx.Input.RequestBody, &custLossManagement)
	if err2 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	var userInfo models.UserInfo
	o := orm.NewOrm()
	o.QueryTable("user_info").Filter("Id", getId).One(&userInfo)
	var orderView models.OrderView
	// 查询该客户最后下单的时间
	err1 := o.Raw("SELECT * FROM order_view ORDER BY order_time DESC LIMIT 1").QueryRow(&orderView)
	fmt.Println("err1 = ", err1)
	if err1 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "查询失败"}
		errResponse := util.NewError(511)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	custLossManagement = models.CustLossManagement{
		CustomerID:      userInfo.Id,
		CustomerName:    userInfo.ClientName,
		CustomerManager: userInfo.Manager,
		LastOrderTime:   orderView.OrderTime,
		LossReasons:     custLossManagement.LossReasons,
	}
	_, err := o.Insert(&custLossManagement)
	if err != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "添加失败"}
		errResponse := util.NewError(512)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	a.Data["json"] = LoginResponse{Success: true, Message: "添加成功"}
	a.ServeJSON()
}

var (
	customerId int
)

func (a *AddInfoController) GenerateOrder() {
	var customer models.Customers
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &customer)
	if err != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	customerId = customer.CustomerID
	fmt.Println("customer = ", customer.CustomerID)
	a.TplName = "orderView.html"
	a.ServeJSON()
}
func (a *AddInfoController) AddOrderInfo() {
	var orderView models.OrderView
	err1 := json.Unmarshal(a.Ctx.Input.RequestBody, &orderView)
	fmt.Println("err1 = ", err1)
	if err1 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	fmt.Println("Price = ", orderView.Price)
	var salesOrder models.SalesOrders
	err2 := json.Unmarshal(a.Ctx.Input.RequestBody, &salesOrder)
	fmt.Println("err2 = ", err2)
	if err2 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	//// 生成订单号
	//orderNumber := GenerateOrderNumber()
	currentTime := time.Now()
	fmt.Println("currentTime = ", currentTime)

	orderView = models.OrderView{
		CustomerID:  customerId,
		OrderNumber: orderNumber,
		Address:     orderView.Address,
		OrderTime:   currentTime,
		Price:       orderView.Price,
		CreateTime:  currentTime,
	}
	o := orm.NewOrm()
	_, err3 := o.Insert(&orderView)
	if err3 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "创建订单失败"}
		errResponse := util.NewError(513)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	salesOrder = models.SalesOrders{
		OrderName:   orderNumber,
		OrderStatus: salesOrder.OrderStatus,
		CustomerID:  customerId,
	}
	_, err4 := o.Insert(&salesOrder)
	if err4 != nil {
		//a.Data["json"] = LoginResponse{Success: false, Message: "创建订单失败"}
		errResponse := util.NewError(513)
		a.Data["json"] = errResponse
		a.ServeJSON()
		return
	}
	a.Data["json"] = LoginResponse{Success: true, Message: "订单创建成功"}
	a.ServeJSON()
}
