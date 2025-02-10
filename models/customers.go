package models

type Customers struct {
	CustomerID int `orm:"description(客户ID);pk;auto" json:"customerID"`
	//CustomerName *UserInfo `orm:"rel(fk);description(客户名称)" json:"customerName"`
	CustomerName string    `orm:"description(客户姓名)" json:"customerName"`
	ContactID    *Contacts `orm:"rel(fk);description(联系人ID)" json:"contactID"`
}

func (c *Customers) TableName() string {
	return "customers"
}

func (c *Customers) CustomersToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"CustomerID":   c.CustomerID,
		"CustomerName": c.CustomerName,
		"ContactID":    c.ContactID,
	}
	return respInfo
}
