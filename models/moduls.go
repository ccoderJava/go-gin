package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin-examples/pkg/setting"
	"log"
)

type Model struct {
	Id         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var db *gorm.DB

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	section, err := setting.Cfg.GetSection("database")
	if err != nil {
		return
	}
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = section.Key("TYPE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("password").String()
	host = section.Key("HOST").String()
	tablePrefix = section.Key("TABLE_PREFIX").String()

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
