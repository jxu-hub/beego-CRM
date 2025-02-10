package models

type UserInfo struct {
	Id               int    `orm:"description(用户ID);pk;auto" json:"id"`
	ClientName       string `orm:"description(客户姓名)" json:"clientName"`
	Area             string `orm:"description(地区)" json:"area"`
	Manager          string `orm:"description(经理)" json:"manager"`
	Grade            string `orm:"description(级别)" json:"grade"`
	Satisfaction     string `orm:"description(满意度)" json:"satisfaction"`
	Creditworthiness string `orm:"description(信用度)" json:"creditworthiness"`
	Address          string `orm:"description(地址)" json:"address"`
	PhoneNumber      string `orm:"description(手机号)" json:"phoneNumber"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func (userInfo *UserInfo) UserInfoReSpDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":               userInfo.Id,
		"ClientName":       userInfo.ClientName,
		"Area":             userInfo.Area,
		"Manager":          userInfo.Manager,
		"Grade":            userInfo.Grade,
		"Satisfaction":     userInfo.Satisfaction,
		"Creditworthiness": userInfo.Creditworthiness,
		"Address":          userInfo.Address,
		"PhoneNumber":      userInfo.PhoneNumber,
	}
	return respInfo
}
