package main

import "C"
import (
	_ "github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"trygorm/app/controllers"
	"trygorm/models"
	"trygorm/utils"
)

func init() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetReportCaller(true)
	controllers.Init()
	models.Init()
}
func main() {
	utils.Router()
}
