package models

type User struct {
	BaseModel
	NickName string `json:"nick_name" gorm:"type:varchar(50);comment:'昵称'"`
	Name     string `json:"name" gorm:"type:varchar(50);comment:'姓名'"`
	Password string `json:"password" gorm:"type:varchar(50);comment:'密码'"`
	Role     int    `gorm:"size:255;default:1" json:"role"`
	Email    string `json:"email" gorm:"type:varchar(250);comment:'邮箱'"`
}

func (u User) Comment() string {
	return "用户表"
}
