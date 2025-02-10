package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ServiceCreateController struct {
	beego.Controller
}

func (s *ServiceCreateController) Get() {
	o := orm.NewOrm()
	count, _ := o.QueryTable("service_create").Count()
	s.Data["count"] = count
	s.TplName = "serviceCreate.html"
}

func (s *ServiceCreateController) AddServiceInfoGet() {
	s.TplName = "addServiceInfo.html"
}

func (s *ServiceCreateController) EditServiceInfoGet() {
	s.TplName = "editServiceInfo.html"
}

// 展示所有服务信息
func (s *ServiceCreateController) ShowAllServiceInfo() {
	var serviceCreates []models.ServiceCreate
	o := orm.NewOrm()
	o.QueryTable("service_create").All(&serviceCreates)
	var respList []interface{}
	for _, serviceCreate := range serviceCreates {
		respList = append(respList, serviceCreate.ServicesToRespDesc())
	}
	s.Data["json"] = respList
	s.ServeJSON()
}

// 添加服务创建
func (s *ServiceCreateController) ServiceCreateAdd() {
	var serviceCreate models.ServiceCreate
	var user models.User
	var service models.Services
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &serviceCreate)
	if err != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Id", LoginId).One(&user)
	err = json.Unmarshal(s.Ctx.Input.RequestBody, &service)
	serviceCreate = models.ServiceCreate{
		ServiceName:    serviceCreate.ServiceName,
		Outline:        serviceCreate.Outline,
		CustomerID:     serviceCreate.CustomerID,
		ServiceStatus:  serviceCreate.ServiceStatus,
		ServiceRequest: serviceCreate.ServiceRequest,
		Username:       user.Username,
	}
	service = models.Services{
		ServiceName:   serviceCreate.ServiceName,
		ServiceStatus: serviceCreate.ServiceStatus,
		CustomerID:    serviceCreate.CustomerID,
	}
	fmt.Println("username = ", serviceCreate.Username)
	_, err = o.Insert(&serviceCreate)
	_, err = o.Insert(&service)
	fmt.Println("err = ", err)
	if err != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "服务创建失败"}
		errResponse := util.NewError(517)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	s.Data["json"] = LoginResponse{Success: true, Message: "服务创建成功"}
	s.ServeJSON()
}

/**

 */
// 服务创建分页
func (s *ServiceCreateController) ShowServiceInfoByPage() {
	var pageInfo PageParam
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &pageInfo)
	if err != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	page := pageInfo.Pagenum
	pageSize := pageInfo.Pagesize
	serviceCreates, err := models.GetServiceInfo(page, pageSize)
	if err != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "获取数据失败"}
		errResponse := util.NewError(506)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	if len(serviceCreates) > 0 {
		var respList []interface{}
		for _, serviceCreate := range serviceCreates {
			respList = append(respList, serviceCreate.ServicesToRespDesc())
		}
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取服务创建信息成功",
			Data:    respList,
		}
		s.Data["json"] = response
	}
	s.ServeJSON()
}

// 搜索服务创建客户
func (s *ServiceCreateController) SearchServiceName() {
	var searchClient SearchClient
	var serviceCreates []models.ServiceCreate
	var customers []models.Customers
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &searchClient)
	if err != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errResponse := util.NewError(504)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	o := orm.NewOrm()
	// 根据客户姓名查询客户编号
	//_, err = o.QueryTable("customers").Filter("CustomerName", searchClient.SearchClient).All(&customers)
	// 模糊查询
	sql := "SELECT * FROM customers WHERE customer_name LIKE ?"
	_, err1 := o.Raw(sql, "%"+searchClient.SearchClient+"%").QueryRows(&customers)
	if err1 != nil {
		//s.Data["json"] = LoginResponse{Success: false, Message: "没有该用户信息"}
		errResponse := util.NewError(511)
		s.Data["json"] = errResponse
		s.ServeJSON()
		return
	}
	var respList []interface{}
	for _, customer := range customers {
		o.QueryTable("service_create").Filter("CustomerID", customer.CustomerID).All(&serviceCreates)
		for _, serviceCreate := range serviceCreates {
			respList = append(respList, serviceCreate.ServicesToRespDesc())
		}
		s.Data["json"] = respList
		s.ServeJSON()
	}
	s.ServeJSON()
}
