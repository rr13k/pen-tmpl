package models

import (
	"time"
)

// gorm doc: https://gorm.io/
type BaseModel struct {
	Id           int       `gorm:"primary_key;comment:'id'" json:"id"`
	CreatedDate  time.Time `gorm:"comment:'创建日期';type:timestamp not null;default:current_timestamp" json:"created_date"`
	ModifiedDate time.Time `gorm:"comment:'修改日期';type:timestamp on update current_timestamp;" json:"modified_date"`
	Yn           int8      `gorm:"default:1;comment:'是否删除: 1存在 2删除'" json:"yn"`
}
