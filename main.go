package main

import (
	"ginVue/common"
	"ginVue/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
)

type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null;"`
	Telephone string `gorm:"varchar(20);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

// @title 2508yyds队后端接口文档
// @version 1.0
// @description 基于区块链的煤炭大数据物流平台的后端接口文档

// @contact.name yzy
// @contact.url https://www.yucoding.club
// @contact.email 1014007642@qq.com

// @host 127.0.0.1:8081
// @BasePath /api/v1

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port") //启动端口
	if port != "" {
		panic(r.Run(":" + port))
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()                 //获取工作路径
	viper.SetConfigName("application")       //配置文件名
	viper.SetConfigType("yml")               //配置文件类型
	viper.AddConfigPath(workDir + "/config") //文件路径
	err := viper.ReadInConfig()
	utils.ChainInit()
	if err != nil {
		panic(err)
	}
}
