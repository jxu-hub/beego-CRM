package models

//type ContactInfo struct {
//	ContactPhoneNumber string `orm:"description(联系人姓名)" json:"contactPhoneNumber"`
//	ContractEmail      string `orm:"description(联系人邮箱)" json:"contractEmail"`
//}

type Contacts struct {
	ContactID   int    `orm:"description(联系人ID);pk;auto" json:"contactID"`
	ContactName string `orm:"description(联系人姓名)" json:"contactName"`
	ContactInfo string `orm:"description(联系人信息)" json:"contactInfo"`
}

func (c *Contacts) TableName() string {
	return "contacts"
}

func (contacts *Contacts) ContactsToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"ContactID":   contacts.ContactID,
		"ContactName": contacts.ContactName,
		"ContactInfo": contacts.ContactInfo,
	}
	return respInfo
}
