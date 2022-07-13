package models

import (
	"fmt"
	"test2/internal/app/common"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func WhereMap(db *gorm.DB, maps map[string]interface{}) (tx *gorm.DB) {
	for i := range maps {
		db = db.Where(i, maps[i])
	}
	return db
}

func Init(config *common.MysqlConfig) {
	dns := GetMysqlDns(config)
	var err error

	var timeout bool = true
	go func() {
		select {
		case <-time.After(3 * time.Second):
			if timeout {
				fmt.Println("连接超时了")
				panic(fmt.Sprintf("mysql comment timeout! link addr: %s:%d", config.Host, config.Port))
			}
			return
		}
	}()
	DB, err = gorm.Open(mysql.Open(dns))
	timeout = false
	if err != nil {
		fmt.Println("gorm open ", err.Error())
		return
	}
	err = DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置连接配置
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100000)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)

	// 开启debug
	DB.Debug()
	fmt.Println("mysql初始化完成")
}

func GetMysqlDns(config *common.MysqlConfig) string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True",
		config.User, config.Password, config.Host, config.Port, config.Db, config.Charset)
}
