package controllers

import (
	"fmt"
	"myapp/app/models"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open("mysql", getConnectionString())

	if err != nil {
		// revel.ERROR.Println("FATAL", err)
		panic(err)
	}

	db.DB()
	db.AutoMigrate(&models.Post{})
	DB = db
}

func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			// revel.ERROR.Fatal("Cound not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectionString() string {
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")
	timezone := getParamString("db.timezone", "parseTime=true&loc=Asia%2FTokyo")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s?%s", user, pass, protocol, host, port, dbname, dbargs, timezone)
}
