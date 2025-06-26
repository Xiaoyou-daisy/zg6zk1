package models

type Wxs struct {
	Id     int64  `gorm:"column:id;type:int;primaryKey;" json:"id"`
	Wsnum  string `gorm:"column:wsnum;type:varchar(60);not null;" json:"wsnum"`
	Mobile string `gorm:"column:mobile;type:char(11);not null;" json:"mobile"`
}
