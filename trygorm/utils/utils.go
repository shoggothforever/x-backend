package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"keys"
	"log"
	"strconv"
	"trygorm/models"
)

func Checkerr(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}

//MySQL create/retrieve/update/delete
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
	db := models.Db
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Pong!")
	}) //output Pong!
	/**setting router group**/
	r := router.Group("/print")
	{
		//print query
		r.GET("/query", func(c *gin.Context) {
			id := c.DefaultQuery("id", "233")
			fmt.Fprintf(c.Writer, id)
		})
		//read body
		r.POST("/body", func(c *gin.Context) {
			var to models.Todos

			if err := c.ShouldBind(&to); err != nil {
				logrus.WithFields(logrus.Fields{"id": to.Id, "name": to.Name, "psw": to.Passwd}).Error("infomation unmatched")
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "bind corrupt",
					"data": to,
				})
			} else {
				logrus.WithFields(logrus.Fields{"id": to.Id, "name": to.Name, "psw": to.Passwd}).Info("infomation matched")
				//db.Create(&to)
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "bind Successfully",
					"data": to,
				})
			}

		})

		//request infomation
		r.GET("/bodyinfo", func(c *gin.Context) {
			fmt.Fprintln(c.Writer, c.Request.RequestURI)
			fmt.Fprintln(c.Writer, c.Request.Host)
			fmt.Fprintln(c.Writer, c.Request.Method)
			fmt.Fprintln(c.Writer, c.Request.Proto)
			fmt.Fprintln(c.Writer, c.Request.RemoteAddr)
			body := c.Request.Body
			var buffer []byte
			_, err := body.Read(buffer)
			if err != nil {
				logrus.Info("读取数据失败")
			}
			fmt.Fprintln(c.Writer, buffer)
		})
	}
	/**router for Todo Api**/
	m := router.Group("api/todo")
	{

		/**find data**/
		m.GET("/get/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				logrus.Error("transform data corrupt")
			}
			uid := uint(id)
			var datas []models.Todos
			db.Where("id=?", uid).Find(&datas)
			if len(datas) == 0 {
				c.JSON(200, gin.H{
					"code":    404,
					"message": "没有查询到数据",
				})
			} else {
				for i := 0; i <= len(datas); i++ {
					c.JSON(200, gin.H{
						"code":    200,
						"message": "查询到数据",
						"id":      datas[i].Id,
						"title":   datas[i].Name,
						"content": datas[i].Passwd,
					})
				}
			}
		})
		/**add data**/
		m.POST("/add/:id/:name/:psw", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				logrus.Error("transform data corrupt")
			}
			uid := uint(id)
			name := c.Param("name")
			psw := c.Param("psw")
			var to models.Todos = models.Todos{
				uid, name, psw,
			}
			if tx := db.Create(&to); tx != nil {
				c.JSON(200, gin.H{
					"code":    200,
					"message": "添加成功",
					"id":      to.Id,
				})
			} else {
				c.JSON(200, gin.H{
					"code":    400,
					"message": "添加失败",
					"data":    to,
				})
			}
		})
		/**delete data**/
		m.DELETE("/delete/:id", func(c *gin.Context) {
			var data []models.Todos
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				logrus.Error("transform data corrupt")
			}
			uid := uint(id)
			db.Where("id=?", uid).Find(&data)
			if len(data) == 0 {
				c.JSON(200, gin.H{
					"message": "表中无此数据",
					"code":    400,
				})
			} else {
				db.Where("id=?", uid).Delete(&data)
				c.JSON(200, gin.H{
					"message": "删除成功",
				})
			}
		})
		/**update data**/
		m.PUT("/update/:id/:name/:psw", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {
				logrus.Error("transform data corrupt")
			}
			uid := uint(id)
			name := c.Param("name")
			psw := c.Param("psw")
			var data []models.Todos
			db.Where("id=?", uid).Find(&data)
			if len(data) == 0 {
				c.JSON(200, gin.H{
					"message": "表中无此数据",
					"code":    400,
				})
			} else {
				if tx := db.Where("id=?", uid).Updates(models.Todos{uid, name, psw}); tx != nil {
					c.JSON(200, gin.H{
						"code":    200,
						"message": "修改成功",
						"data":    data,
					})
				} else {
					c.JSON(200, gin.H{
						"code":    400,
						"message": "修改失败",
						"data":    data,
					})
				}
			}
		})

	}
	router.Run(":9090")
}
