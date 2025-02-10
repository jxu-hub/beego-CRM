package models

type User struct {
	Id       int    `orm:"description(用户ID);pk;auto" json:"id"`
	Username string `orm:"description(用户名)" json:"username"`
	Password string `orm:"description(密码)" json:"password"`
}

func (u *User) TableName() string {
	return "user"
}

func (user *User) UserToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"Id":       user.Id,
		"Username": user.Username,
	}
	return respInfo
}
