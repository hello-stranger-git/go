package mysqlutils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	toolBoxDB *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	username := viper.GetString("db.username")
	password := viper.GetString("db.password")
	addr := viper.GetString("db.addr")
	name := viper.GetString("db.name")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	database, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic("Database connection failed. Database name,err: " + err.Error())
	}
	sqlDB, _ := database.DB()
	sqlDB.SetMaxOpenConns(1000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxIdleConns(100)  // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	DB = &Database{
		toolBoxDB: database,
	}
}
