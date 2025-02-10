package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EditInfoController struct {
	beego.Controller
}

var id int

func (e *EditInfoController) EditUserMessageGet() {
	e.TplName = "editUserMessage.html"
}

// 获取改id数据库中的数据并传给前端
func (e *EditInfoController) EditUserMessagePost() {
	var userInfo models.UserInfo
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	o := orm.NewOrm()
	id, _ = e.GetInt(":id") // 获取到url当中id变量的值
	o.QueryTable("user_info").Filter("id", id).One(&userInfo)
	e.Data["UserInfo"] = userInfo
	e.ServeJSON()
}

// 编辑客户信息
func (e *EditInfoController) EditUserMessagePUT() {
	var userInfo models.UserInfo
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &userInfo)
	if err != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	o := orm.NewOrm()
	o.QueryTable("user_info").Filter("id", userInfo.Id).One(&userInfo)
	e.Data["userInfo"] = userInfo
	var r orm.RawSeter
	r = o.Raw("UPDATE user_info set client_name = ?,area = ?,manager = ?,grade = ?,satisfaction = ?,creditworthiness = ?,address = ?,phone_number = ? WHERE id = ?",
		userInfo.ClientName, userInfo.Area, userInfo.Manager, userInfo.Grade, userInfo.Satisfaction, userInfo.Creditworthiness, userInfo.Area, userInfo.PhoneNumber, id)
	_, error := r.Exec()
	if error != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "更新失败"}
		errResponse := util.NewError(516)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	r = o.Raw("UPDATE customers set customer_name=? WHERE customer_i_d = ?", userInfo.ClientName, id)
	_, error = r.Exec()
	if error != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "更新失败"}
		errResponse := util.NewError(516)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	e.Data["json"] = LoginResponse{Success: true, Message: "更新成功"}
	e.ServeJSON()
}

type GetEditInfo struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	Area             string `json:"area"`
	Manager          string `json:"manager"`
	Grade            string `json:"grade"`
	Satisfaction     string `json:"satisfaction"`
	Creditworthiness string `json:"creditworthiness"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phoneNumber"`
}

var getEditInfo GetEditInfo

// 把从后端查到的数据返回给前端
func (e *EditInfoController) GetEditInfo() {
	var userInfo models.UserInfo
	id, _ = e.GetInt(":id")
	o := orm.NewOrm()
	o.QueryTable("user_info").Filter("id", id).One(&userInfo)
	getEditInfo = GetEditInfo{
		Id:               userInfo.Id,
		Name:             userInfo.ClientName,
		Area:             userInfo.Area,
		Manager:          userInfo.Manager,
		Grade:            userInfo.Grade,
		Satisfaction:     userInfo.Satisfaction,
		Creditworthiness: userInfo.Creditworthiness,
		Address:          userInfo.Address,
		PhoneNumber:      userInfo.PhoneNumber,
	}
	e.Data["getEditInfo"] = getEditInfo
	e.Data["json"] = getEditInfo
	e.ServeJSON()
	e.TplName = "editUserMessage.html"
}
func (e *EditInfoController) PostEditInfo() {
	var userInfo models.UserInfo
	o := orm.NewOrm()
	o.QueryTable("user_info").Filter("id", getEditInfo.Id).One(&userInfo)
	e.Data["getEditInfo"] = getEditInfo
	e.Data["json"] = getEditInfo
	e.ServeJSON()
}

// 编辑服务创建信息
func (e *EditInfoController) EditServiceGet() {
	e.TplName = "editServiceInfo.html"
}

var serviceId int

type OneServiceInfo struct {
	Id             int    `json:"id"`
	ServiceName    string `json:"serviceName"`
	Outline        string `json:"outline"`
	CustomerID     int    `json:"customerID"`
	ServiceStatus  string `json:"serviceStatus"`
	ServiceRequest string `json:"serviceRequest"`
	Username       string `json:"username"`
}

var oneServiceInfo OneServiceInfo

func (e *EditInfoController) GetOneServiceInfo() {
	serviceId, _ = e.GetInt(":id")
	e.ServeJSON()
}

func (e *EditInfoController) PostOneServiceInfo() {
	var serviceCreate models.ServiceCreate
	o := orm.NewOrm()
	o.QueryTable("serviceCreate").Filter("Id", serviceId).One(&serviceCreate)
	fmt.Println("Id = ", serviceId)
	e.Data["json"] = serviceCreate
	e.ServeJSON()
}

// 编辑服务创建信息
func (e *EditInfoController) EditServiceInfoPUT() {
	var serviceCreate models.ServiceCreate
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &serviceCreate)
	fmt.Println("err = ", err)
	if err != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE service_create set service_name = ?,outline = ?,customer_i_d = ?,service_status = ?,service_request = ?,username = ? WHERE id = ?",
		serviceCreate.ServiceName, serviceCreate.Outline, serviceCreate.CustomerID, serviceCreate.ServiceStatus, serviceCreate.ServiceRequest, serviceCreate.Username, serviceId)
	_, error := r.Exec()
	if error != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "编辑失败"}
		errResponse := util.NewError(516)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	e.Data["json"] = LoginResponse{Success: true, Message: "编辑成功"}
	e.ServeJSON()
}

// 编辑客户流失信息
func (e *EditInfoController) GetEditHTML() {
	e.TplName = "editCustLossInfo.html"
}

var editId int

func (e *EditInfoController) GetOneCussLossId() {
	editId, _ = e.GetInt(":id")
	e.ServeJSON()
}

var customerID int

func (e *EditInfoController) GetOneCussLossInfo() {
	var custLossInfo models.CustLossManagement
	o := orm.NewOrm()
	o.QueryTable("cust_loss_management").Filter("Id", editId).One(&custLossInfo)
	customerID = custLossInfo.CustomerID
	e.Data["json"] = custLossInfo
	e.ServeJSON()
}
func (e *EditInfoController) EditCussLossInfoPUT() {
	var custLossInfo models.CustLossManagement
	err := json.Unmarshal(e.Ctx.Input.RequestBody, &custLossInfo)
	fmt.Println("err = ", err)
	if err != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE cust_loss_management set customer_i_d = ?,customer_name=?,customer_manager=?,last_order_time=?,loss_reasons=? WHERE id=?",
		customerID, custLossInfo.CustomerName, custLossInfo.CustomerManager, custLossInfo.LastOrderTime, custLossInfo.LossReasons, editId)
	_, error := r.Exec()
	if error != nil {
		//e.Data["json"] = LoginResponse{Success: false, Message: "编辑失败"}
		errResponse := util.NewError(516)
		e.Data["json"] = errResponse
		e.ServeJSON()
		return
	}
	e.Data["json"] = LoginResponse{Success: true, Message: "编辑成功"}
	e.ServeJSON()
}
