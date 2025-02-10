package controllers

import (
	"crmManaSystem/models"
	"crmManaSystem/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DeleteInfo struct {
	beego.Controller
}

// 删除用户信息
func (d *DeleteInfo) DeleteUserInfo() {
	id, err := d.GetInt(":id")
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"}
		errResponse := util.NewError(506)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	o := orm.NewOrm()
	_, err = o.Delete(&models.UserInfo{Id: id})
	_, err = o.Delete(&models.Contacts{ContactID: id})
	_, err = o.Delete(&models.Customers{CustomerID: id})
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"}
		errResponse := util.NewError(514)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	d.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	d.ServeJSON()
}

// 删除客户流失信息
func (d *DeleteInfo) DeleteCustLossInfo() {
	o := orm.NewOrm()
	_, err := o.Delete(&models.UserInfo{Id: getId})
	fmt.Println()
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "流失操作失败"}
		errResponse := util.NewError(515)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	d.Data["json"] = LoginResponse{Success: true, Message: "流失完成"}
	d.ServeJSON()
}

// 删除客户流失页面信息
func (d *DeleteInfo) DeleteCustLossHTMLInfo() {
	id, _ := d.GetInt(":id")
	o := orm.NewOrm()
	_, err := o.Delete(&models.CustLossManagement{Id: id})
	fmt.Println()
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "流失操作失败"}
		errResponse := util.NewError(515)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	d.Data["json"] = LoginResponse{Success: true, Message: "流失完成"}
	d.ServeJSON()
}

// 删除服务创建信息
func (d *DeleteInfo) DeleteServiceInfo() {
	id, err := d.GetInt(":id")
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "获取参数失败"}
		errResponse := util.NewError(506)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	o := orm.NewOrm()
	_, err = o.Delete(&models.ServiceCreate{Id: id})
	_, err = o.Delete(&models.Services{ServiceID: id})
	if err != nil {
		//d.Data["json"] = LoginResponse{Success: false, Message: "删除数据失败"}
		errResponse := util.NewError(514)
		d.Data["json"] = errResponse
		d.ServeJSON()
		return
	}
	d.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	d.ServeJSON()
}

// 删除选中用户信息
var getDeleteUserId int
var deleteUserIds []int

func (d *DeleteInfo) GetDeleteUserId() {
	getDeleteUserId, _ = d.GetInt(":id")
	deleteUserIds = append(deleteUserIds, getDeleteUserId)
	fmt.Println("id = ", getDeleteUserId)
	d.TplName = "basicInfo.html"
}
func (d *DeleteInfo) DeleteAllUserInfo() {
	o := orm.NewOrm()
	for _, deleteUserId := range deleteUserIds {
		_, err := o.Delete(&models.UserInfo{Id: deleteUserId})
		_, err = o.Delete(&models.Contacts{ContactID: deleteUserId})
		_, err = o.Delete(&models.Customers{CustomerID: deleteUserId})
		if err != nil {
			errResponse := util.NewError(514)
			d.Data["json"] = errResponse
			d.ServeJSON()
			return
		}
	}

	d.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	d.ServeJSON()
}

// 删除选中客户流失信息
var getDeleteCustLossId int
var deleteCustLossIds []int

func (d *DeleteInfo) GetDeleteCustLossId() {
	getDeleteCustLossId, _ = d.GetInt(":id")
	deleteCustLossIds = append(deleteCustLossIds, getDeleteCustLossId)
	fmt.Println("id = ", getDeleteCustLossId)
	d.TplName = "custLossManagement.html"
}
func (d *DeleteInfo) DeleteAllCustInfo() {
	o := orm.NewOrm()
	for _, deleteCustLossId := range deleteCustLossIds {
		_, err := o.Delete(&models.CustLossManagement{Id: deleteCustLossId})
		if err != nil {
			errResponse := util.NewError(514)
			d.Data["json"] = errResponse
			d.ServeJSON()
			return
		}
	}
	d.Data["json"] = LoginResponse{Success: true, Message: "删除成功"}
	d.ServeJSON()
}
