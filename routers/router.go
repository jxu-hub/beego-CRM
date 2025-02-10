package routers

import (
	"crmManaSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}, "Post:Login")
	beego.Router("/login", &controllers.LoginController{}, "GET:Get")
	beego.Router("/home", &controllers.HomeController{})

	// 需要登录才可以进入
	beego.Router("/test/index", &controllers.IndexController{})
	beego.Router("/test/login", &controllers.LoginController{}, "POST:Login")
	beego.Router("/users", &controllers.UserListController{})
	beego.Router("/users", &controllers.UserListController{}, "Get:GetAll")
	beego.Router("/searchUsername", &controllers.UserListController{}, "Post:SearchUserName")
	beego.Router("/dologin", &controllers.LoginController{}, "POST:Login")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/register", &controllers.RegisterController{}, "POST:Post")
	// 添加客户信息
	beego.Router("/addUserMessage", &controllers.AddInfoController{}, "Get:UserMessageAddGet")
	beego.Router("/addUserMessagePost", &controllers.AddInfoController{}, "Post:AddUserMessagePost")
	// 删除客户信息
	beego.Router("/deleteUserInfo/:id", &controllers.DeleteInfo{}, "DELETE:DeleteUserInfo")
	// 修改客户信息
	beego.Router("/editUserMessage", &controllers.EditInfoController{}, "Get:EditUserMessageGet")
	beego.Router("/doeditUserMessage/:id", &controllers.EditInfoController{}, "Post:EditUserMessagePost")
	beego.Router("/doeditUserMessage", &controllers.EditInfoController{}, "PUT:EditUserMessagePUT")
	beego.Router("/getEditInfo/:id", &controllers.EditInfoController{}, "Get:GetEditInfo")
	beego.Router("/getEditInfo", &controllers.EditInfoController{}, "Post:PostEditInfo")

	beego.Router("/getall", &controllers.UserListController{}, "GET:GetAll")
	beego.Router("/deleteuser/:id", &controllers.UserListController{}, "DELETE:DeleteUser")
	beego.Router("/updateUser", &controllers.UserListController{}, "PUT:UpdateUser")
	beego.Router("/page", &controllers.UserListController{}, "POST:ShowUserByPage")
	// 基础信息
	beego.Router("basicInfo", &controllers.BasicInfoController{}, "Get:BasicInfo")
	// 搜索客户信息
	beego.Router("/searchClient", &controllers.BasicInfoController{}, "POST:SearchUser")
	// 展示基本信息
	beego.Router("/showBasicInfo", &controllers.BasicInfoController{}, "Get:ShowBasicInfo")
	beego.Router("/basicInfoPage", &controllers.BasicInfoController{}, "Post:ShowBasicByPage")
	// 展示联系人信息
	beego.Router("/showContacts", &controllers.ShowContactsController{}, "Get:ShowContactsGet")
	beego.Router("/showContactsPost/:id", &controllers.ShowContactsController{}, "Post:ShowContactsPost")
	beego.Router("/getContactsInfo/:id", &controllers.ShowContactsController{}, "Get:GetContactsInfo")
	beego.Router("/getAllContactsInfo", &controllers.ShowContactsController{}, "Get:GetAllContactsInfo")
	// 订单查看
	beego.Router("orderView", &controllers.OrderViewController{}, "Get:OrderView")
	// 搜索订单号
	beego.Router("/searchOrderNumber", &controllers.OrderViewController{}, "Post:SearchOrderNumber")
	// 订单查看->展示所有的订单
	beego.Router("/showAllOrder", &controllers.OrderViewController{}, "Post:ShowAllOrderView")
	beego.Router("/orderInfoPage", &controllers.OrderViewController{}, "Post:ShowOrderByPage")
	// 订单查看->生成订单
	beego.Router("/addOrderInfo", &controllers.OrderViewController{}, "Get:AddOrderInfoGet")
	beego.Router("/generateOrder", &controllers.AddInfoController{}, "Post:GenerateOrder")
	beego.Router("/generateOrderAdd", &controllers.AddInfoController{}, "Post:AddOrderInfo")
	// 客户流失页面
	beego.Router("/showCustLossInfoGet", &controllers.CustLossManageController{})
	beego.Router("/showCustLossInfoPost", &controllers.CustLossManageController{}, "Post:ShowCustLossInfoPost")
	beego.Router("/custLossInfoPage", &controllers.CustLossManageController{}, "Post:ShowCustLossByPage")
	// 客户流失->编辑
	beego.Router("/editCustLossInfo", &controllers.EditInfoController{}, "Get:GetEditHTML")
	beego.Router("/getOneCustLossId/:id", &controllers.EditInfoController{}, "Get:GetOneCussLossId")
	beego.Router("/getOneCustLossInfo", &controllers.EditInfoController{}, "Get:GetOneCussLossInfo")
	beego.Router("/editCustLossInfoPut", &controllers.EditInfoController{}, "PUT:EditCussLossInfoPUT")
	// 搜索流失客户姓名
	beego.Router("/searchCustLossName", &controllers.CustLossManageController{}, "Post:SearchCustLossName")
	// 客户流失
	beego.Router("/getCustLossInfoGet/:id", &controllers.CustLossManagementController{}, "Get:CustLossManagementGet")
	beego.Router("/custLossInfo", &controllers.CustLossManagementController{})
	beego.Router("/addCustLossInfoPost", &controllers.AddInfoController{}, "Post:AddCustLossInfo")
	beego.Router("/deleteCustLossInfo", &controllers.DeleteInfo{}, "DELETE:DeleteCustLossInfo")
	beego.Router("/deleteCustLossHtmlInfo/:id", &controllers.DeleteInfo{}, "DELETE:DeleteCustLossHTMLInfo")
	// 服务创建
	beego.Router("serviceCreate", &controllers.ServiceCreateController{})
	beego.Router("/showAllServiceInfo", &controllers.ServiceCreateController{}, "Post:ShowAllServiceInfo")
	beego.Router("/serviceCreatePage", &controllers.ServiceCreateController{}, "Post:ShowServiceInfoByPage")
	// 所有服务创建客户
	beego.Router("/searchServiceName", &controllers.ServiceCreateController{}, "Post:SearchServiceName")
	// 服务创建->添加
	beego.Router("addServiceInfo", &controllers.ServiceCreateController{}, "Get:AddServiceInfoGet")
	beego.Router("serviceCreateAdd", &controllers.ServiceCreateController{}, "Post:ServiceCreateAdd")
	// 服务创建->删除
	beego.Router("/deleteServiceInfo/:id", &controllers.DeleteInfo{}, "DELETE:DeleteServiceInfo")
	// 服务创建->编辑
	beego.Router("editServiceInfo", &controllers.ServiceCreateController{}, "Get:EditServiceInfoGet")
	beego.Router("/getOneServiceInfo/:id", &controllers.EditInfoController{}, "Get:GetOneServiceInfo")
	beego.Router("/getOneServiceInfo", &controllers.EditInfoController{}, "Post:PostOneServiceInfo")
	beego.Router("/serviceCreateEdit", &controllers.EditInfoController{}, "PUT:EditServiceInfoPUT")
	// 删除选中客户信息
	beego.Router("/getDeleteUserId/:id", &controllers.DeleteInfo{}, "Get:GetDeleteUserId")
	beego.Router("/deleteAllUserInfo", &controllers.DeleteInfo{}, "DELETE:DeleteAllUserInfo")
	// 删除选中客户流失信息
	beego.Router("/getDeleteCustLossId/:id", &controllers.DeleteInfo{}, "Get:GetDeleteCustLossId")
	beego.Router("/deleteAllCustLossInfo", &controllers.DeleteInfo{}, "DELETE:DeleteAllCustInfo")
}
