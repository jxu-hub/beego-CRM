package models

type SalesOrders struct {
	OrderID     int    `orm:"description(销售订单ID);pk;auto" json:"orderID"`
	OrderName   string `orm:"description(订单名称)" json:"orderName"`
	OrderStatus string `orm:"description(订单状态)" json:"orderStatus"`
	//CustomerID  *Customers `orm:"rel(fk);description(客户ID外键)" json:"customerID"`
	CustomerID int `orm:"description(客户ID外键)" json:"customerID"`
}

func (s *SalesOrders) TableName() string {
	return "sales_orders"
}
