package models

type Logins struct {
	Id     int64  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Mobile string `gorm:"column:mobile;type:char(11);comment:手机号;not null;" json:"mobile"` // 手机号
}
