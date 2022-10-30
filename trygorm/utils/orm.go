package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"trygorm/models"
)

func Tra1(tx *gorm.DB) error {
	if err := tx.Create(&models.Todos{42, "dump", "end"}).Error; err != nil {
		logrus.WithFields(logrus.Fields{"Error": err}).Error("创建Todo记录失败")
		return err
	}
	if err := tx.Create(&models.Users{42, 233, "travel", "space"}).Error; err != nil {
		logrus.WithFields(logrus.Fields{"Error": err}).Error("创建User记录失败")
		return err
	}
	return nil
}
func Sess(db *gorm.DB) {
	tx := db.Session(&gorm.Session{SkipDefaultTransaction: true}) /*.Begin()*/
	Tra1(tx)
}
