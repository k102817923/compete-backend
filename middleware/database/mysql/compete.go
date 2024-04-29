package mysql

import (
	"fmt"
	"log"
	"paracraft-compete-backend/common/setting"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var CompeteDB *gorm.DB

func initDB(dbName string) *gorm.DB {
	dbType := setting.LoadStringParam(dbName, "TYPE")
	user := setting.LoadStringParam(dbName, "USER")
	password := setting.LoadStringParam(dbName, "PASSWORD")
	host := setting.LoadStringParam(dbName, "HOST")
	port := setting.LoadIntParam(dbName, "PORT")

	connect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True", user, password, host, port, dbName)
	db, err := gorm.Open(dbType, connect)
	if err != nil {
		log.Fatalf("mysql connect error: %v", err)
	}
	return db
}

func InitCompeteDB() {
	dbName := "compete_test"
	maxIdle := setting.LoadIntParam("mysql_setting", "MAX_IDLE_CONNECTION")
	maxOpen := setting.LoadIntParam("mysql_setting", "MAX_OPEN_CONNECTION")
	CompeteDB = initDB(dbName)
	CompeteDB.DB().SetMaxIdleConns(maxIdle)
	CompeteDB.DB().SetMaxOpenConns(maxOpen)
	CompeteDB.DB().SetConnMaxLifetime(10 * time.Second)
	// 如果设置为true, `User`的默认表名为`user`, 使用`TableName`设置的表名不受影响
	CompeteDB.SingularTable(false)
}
