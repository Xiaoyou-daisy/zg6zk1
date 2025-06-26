package inits

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zg6zk1/apiway/basic/global"
)

var err error

func MysqlInit() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "zg6zk:2003225zyh@tcp(14.103.243.149:3306)/zg6zk?charset=utf8mb4&parseTime=True&loc=Local"
	global.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("mysql connect success")
	}

}
