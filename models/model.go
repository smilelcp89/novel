package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"novel/global"
	"novel/global/config"
)

var orm *gorm.DB

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.DbConfig.User,
		config.DbConfig.Password,
		config.DbConfig.Host,
		config.DbConfig.Port,
		config.DbConfig.Database,
		config.DbConfig.Charset)
	var err error
	orm, err = gorm.Open(config.DbConfig.Type, dsn)
	if err != nil {
		global.Logger.Info(dsn)
		global.Logger.Errorf("连接数据库失败： %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.DbConfig.TablePrefix + defaultTableName
	}
}
