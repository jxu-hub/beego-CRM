package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

type UnifiedResponse struct {
	Code    int         `json:"code"`    // 响应状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 用户
func GetUsers(page, pageSize int) ([]User, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码合法
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var user []User
	_, err := o.QueryTable("user").Limit(pageSize, (page-1)*pageSize).All(&user)
	return user, err
}

// 客户基本信息
func GetBasicInfo(page, pageSize int) ([]UserInfo, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("user_info").Count()
	fmt.Println("总数量", count)
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var userInfo []UserInfo
	_, err := o.QueryTable("user_info").Limit(pageSize, (page-1)*pageSize).All(&userInfo)
	return userInfo, err
}

// 订单信息
func GetOrderInfo(page, pageSize int) ([]OrderView, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("order_view").Count()
	fmt.Println("总数量", count)
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var orderView []OrderView
	_, err := o.QueryTable("order_view").Limit(pageSize, (page-1)*pageSize).All(&orderView)
	return orderView, err
}

// 客户流失
func GetCustLossInfo(page, pageSize int) ([]CustLossManagement, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("cust_loss_management").Count()
	fmt.Println("总数量", count)
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var custLosses []CustLossManagement
	_, err := o.QueryTable("cust_loss_management").Limit(pageSize, (page-1)*pageSize).All(&custLosses)
	return custLosses, err
}

// 服务创建
func GetServiceInfo(page, pageSize int) ([]ServiceCreate, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("service_create").Count()
	fmt.Println("总数量", count)
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var serviceInfos []ServiceCreate
	_, err := o.QueryTable("service_create").Limit(pageSize, (page-1)*pageSize).All(&serviceInfos)
	return serviceInfos, err
}
