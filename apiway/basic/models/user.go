package models

import (
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Id          int64          `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Username    string         `gorm:"column:username;type:varchar(50);comment:用户名称;not null;" json:"username"`      // 用户名称
	HeaderImage string         `gorm:"column:header_image;type:varchar(150);comment:头像;not null;" json:"header_image"` // 头像
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间;" json:"deleted_at"` // 删除时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"`
	Mobile      string         `gorm:"column:mobile;type:char(11);not null;" json:"mobile"`
}
