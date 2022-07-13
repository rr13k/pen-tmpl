package models

type User struct {
	BaseModel
	NickName string `json:"nick_name"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     int    `gorm:"size:255;default:1" json:"role"`
	Email    string `json:"email"`
}

func ChenkToken(name, token string) (*User, error) {
	var u User
	err := DB.Model(&User{}).Where("yn", 1).Where("name", name).Where("password", token).Find(&u).Error
	return &u, err
}

func GetUser(maps interface{}) (*User, error) {
	var _user *User
	err := DB.Model(&User{}).Where(maps).Find(&_user).Error
	return _user, err
}

func GetAllUser(pageNum int, pageSize int, maps map[string]interface{}) ([]*User, error) {
	var _users []*User
	err := WhereMap(DB, maps).Model(&User{}).Offset(pageNum).Limit(pageSize).Find(&_users).Error
	return _users, err
}

func GetUserTotal(maps map[string]interface{}) (int, error) {
	var count int64
	if err := WhereMap(DB, maps).Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
