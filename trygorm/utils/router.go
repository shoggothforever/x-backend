package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"keys"
	"net/http"
	"trygorm/app/controllers"
	"trygorm/app/middleware"
	"trygorm/models"
	"trygorm/views"
)

//MySQL create/retrieve/update/deletej
func CRUD(db *gorm.DB) {
	//build connect
	dsn := keys.MySQLdsn + "todo" + keys.DefaultMode
	db, err := gorm.Open("mysql", dsn)
	//Checkerr(err)
	if err != nil {
		logrus.WithFields(logrus.Fields{"error": err}).Error("OPEN MySQL failed")
	}
	defer db.Close()
	user := models.Users{1, 1, "dog", "huhu"}
	todo := models.Todos{1, "cat", "fish"}
	//Create Tables
	db.AutoMigrate(&models.Users{}, &models.Todos{})
	var u = new(models.Users)
	db.First(&u)
	var t = new(models.Todos)
	db.First(t)
	db.Delete(&u)
	db.Delete(&t)
	//Create row
	db.Create(user)
	db.Create(todo)
	//Retrieve row
	var p = new(models.Users)
	db.Where("Title=?", "dog").Find(&p)
	if p != nil {
		logrus.WithFields(logrus.Fields{"result": nil}).Info("查询不到数据")
		fmt.Println(*p)
	} else {
		fmt.Println("can't find any data")
	}
	//Update row
	db.Model(&p).Update("Content", "wangwang")
	//Delete row
	//db.Delete(&p)
	db.Where("Titile=?", "dog").Find(&p)
	fmt.Println(*p)
}

//Task 1.2
func Router() {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})
	router.GET("/ping", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Pong!")
	}) //output Pong!
	/**setting router group**/
	r := router.Group("/print")
	{
		//print query
		r.GET("/query", views.Query)
		//read body
		r.POST("/body", views.PrintBody)

		//request infomation
		r.GET("/bodyinfo", views.PrintBodyInfo)
	}
	//login api
	a := router.Group("/api")
	{
		a.POST("/login", controllers.Login)
	}
	/**router for Todo Api**/
	m := a.Group("/todo")
	m.Use(middleware.Authorize())
	{
		/**find data**/
		m.GET("/get/:id/:token", views.Search) //url form:localhost:9090/api/todo/get/id/token
		/**add data**/
		m.POST("/add/:id/:name/:psw/:token", views.Add) //url form:localhost:9090/api/todo/add/id/name/psw/token
		/**delete data**/
		m.DELETE("/delete/:id/:token", views.Delete)
		/**update data**/
		m.PUT("/update/:id/:name/:psw/:token", views.Update)
	}
	router.Run(":9090")
}
