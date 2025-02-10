package models

import "time"

type OrderView struct {
	Id          int       `orm:"description(编号)" json:"id"`
	CustomerID  int       `orm:"description(客户编号)" json:"customerID"`
	OrderNumber string    `orm:"description(订单号)" json:"orderNumber"`
	Address     string    `orm:"description(地址)" json:"address"`
	OrderTime   time.Time `orm:"auto_now_add;type(datetime);description(下单时间)" json:"orderTime"`
	Price       float64   `orm:"description(价格)" json:"price"`
	CreateTime  time.Time `orm:"auto_now_add;type(datetime);description(创建时间)" json:"createTime"`
}

func (o *OrderView) TableName() string {
	return "order_view"
}

func (o *OrderView) OrderViewToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":          o.Id,
		"CustomerID":  o.CustomerID,
		"OrderNumber": o.OrderNumber,
		"Address":     o.Address,
		"OrderTime":   o.OrderTime,
		"Price":       o.Price,
		"CreateTime":  o.CreateTime,
	}
	return respInfo
}
