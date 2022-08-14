package initialize

import (
	"fmt"
	"log"
	"os"

	"github.com/akazwz/url-shortener-kafka/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {
	if os.Getenv("ENV") != "prod" {
		return InitSqlite()
	} else {
		return InitMysql()
	}
}

func InitSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("初始化 sqlite 失败")
	}
	return db
}

func InitMysql() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("初始化 mysql 失败")
	}
	return db
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.VisitsLog{},
	)
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("数据库表迁移失败")
	}
}
