package models

import "time"

type CustLossManagement struct {
	Id              int       `orm:"description(编号);pk;auto" json:"id"`
	CustomerID      int       `orm:"description(客户编号)" json:"customerID"`
	CustomerName    string    `orm:"description(客户姓名)" json:"customerName"`
	CustomerManager string    `orm:"description(客户经理)" json:"customerManage"`
	LastOrderTime   time.Time `orm:"auto_now_add;type(datetime);description(最后下单时间)" json:"lastOrderTime"`
	LossReasons     string    `orm:"description(流失原因)" json:"lossReasons"`
}

func (c *CustLossManagement) TableName() string {
	return "cust_loss_management"
}

func (c *CustLossManagement) CustLossManagementToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":              c.Id,
		"CustomerID":      c.CustomerID,
		"CustomerName":    c.CustomerName,
		"CustomerManager": c.CustomerManager,
		"LastOrderTime":   c.LastOrderTime,
		"LossReasons":     c.LossReasons,
	}
	return respInfo
}
