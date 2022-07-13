package models

import (
	"time"
)

type BaseModel struct {
	Id           int       `gorm:"primary_key" json:"id"`
	CreatedDate  time.Time `gorm:"type:timestamp not null;default:current_timestamp" json:"created_date"`
	ModifiedDate time.Time `gorm:"type:timestamp on update current_timestamp;default:current_timestamp;" json:"modified_date"`
	Yn           int8      `gorm:"default:1;comment:'是否删除: 1存在 2删除'" json:"yn"`
}
