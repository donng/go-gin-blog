package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-blog/conf"
)

var db *gorm.DB

// 初始化数据库服务
func init() {
	databaseConf := conf.Setting.Database

	fmt.Println(databaseConf)

	var err error
	db, err = gorm.Open(databaseConf.Type, getDatabaseSource(databaseConf))
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return databaseConf.TablePrefix + defaultTableName
	}
}

func getDatabaseSource(conf conf.Database) string {
	switch conf.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.DBName)
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", conf.Host, conf.Port, conf.User, conf.DBName, conf.Password)
	case "sqlite3":
		return "/tmp/gorm.db"
	case "mssql":
		return fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", conf.User, conf.Password, conf.Host, conf.Port, conf.DBName)
	default:
		panic(fmt.Sprintf("error database type: %s", conf.Type))
	}
}
