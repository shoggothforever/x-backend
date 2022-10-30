package models

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var JwtSecret string

func Init() {
	//Configure the viper
	Config := viper.New()
	Config.SetConfigFile("./config/config.yaml")
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath("./config")
	err := Config.ReadInConfig() // 查找并读取配置文件
	if err != nil {              // 处理读取配置文件的错误
		logrus.Error("Read config file failed: %s \n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logrus.Info("no error in config file")
		} else {
			logrus.Error("found error in config file\n", ok)
		}
	}
	jwtInfo := Config.GetStringMapString("Jwt")
	JwtSecret = jwtInfo["secret"]
	loginInfo := Config.GetStringMapString("mysql")
	Dsn := loginInfo["predsn"] + "todo" + loginInfo["mode"]
	Db, err = gorm.Open(mysql.Open(Dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("gorm OPENS MySQL failed")
	}
	err = Db.AutoMigrate(&Todos{}, &Users{})
	if err != nil {
		logrus.Error("build tables corrupt!\n", err)
	}
}
