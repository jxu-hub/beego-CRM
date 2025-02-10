package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	orm.RegisterModel(new(User), new(OrderView), new(Customers), new(Contacts), new(Services), new(SalesOrders), new(UserInfo), new(CustLossManagement), new(ServiceCreate))
	driverName := beego.AppConfig.String("driverName")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	err := orm.RegisterDataBase("default", "mysql", datasource)
	if err != nil {
		fmt.Printf("连接数据库出错 %s", err)
	}
	err1 := orm.RunSyncdb("default", false, true)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
	}
	fmt.Println("连接数据库成功")
	orm.RegisterDriver(driverName, orm.DRMySQL)
	// 启用ORM日志
	orm.Debug = true
}
