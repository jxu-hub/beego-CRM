package models

type Services struct {
	ServiceID     int    `orm:"description(服务ID);pk;auto" json:"serviceID"`
	ServiceName   string `orm:"description(服务名称)" json:"serviceName"`
	ServiceStatus string `rom:"description(服务状态);default(未开启服务)" json:"serviceStatus"`
	CustomerID    int    `orm:"description(客户ID外键)" json:"customerID"`
}

func (s *Services) TableName() string {
	return "services"
}
