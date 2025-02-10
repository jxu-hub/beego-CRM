package models

type ServiceCreate struct {
	Id          int    `orm:"description(编号);pk;auto" json:"id"`
	ServiceName string `orm:"description(服务类型)" json:"serviceName"`
	Outline     string `orm:"description(概要)" json:"outline"`
	CustomerID  int    `orm:"description(客户编号)" json:"customerID"`
	//CustomerID     *Customers `orm:"rel(fk);description(客户编号)" json:"customerID"`
	ServiceStatus  string `orm:"description(服务状态)" json:"serviceStatus"`
	ServiceRequest string `orm:"description(服务请求)" json:"serviceRequest"`
	Username       string `orm:"description(服务创建人)" json:"username"`
	//Username       *User      `orm:"rel(fk);description(服务创建人)" json:"username"`
}

func (s *ServiceCreate) TableName() string {
	return "service_create"
}

func (s *ServiceCreate) ServicesToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":             s.Id,
		"ServiceName":    s.ServiceName,
		"Outline":        s.Outline,
		"CustomerID":     s.CustomerID,
		"ServiceStatus":  s.ServiceStatus,
		"ServiceRequest": s.ServiceRequest,
		"Username":       s.Username,
	}
	return respInfo
}
