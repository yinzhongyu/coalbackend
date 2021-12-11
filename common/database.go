package common

import (
	"fmt"
	"ginVue/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	//driverName := "mysql"
	//host := "localhost"
	//port := "3306"
	//databse := "gin"
	//username := "root"
	//password := "123456"
	//charset := "utf8"

	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	databse := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, databse, charset)
	fmt.Println(args)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect mysql")
	}
	db.SingularTable(true)
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
