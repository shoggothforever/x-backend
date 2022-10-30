package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
	"trygorm/models"
)

func Query(c *gin.Context) {
	id := c.DefaultQuery("id", "233")
	fmt.Fprintf(c.Writer, id)
}
func PrintBody(c *gin.Context) {
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

}
func PrintBodyInfo(c *gin.Context) {
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
}

func Search(c *gin.Context) {
	db := models.Db
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error("transform data corrupt")
	}
	uid := uint(id)
	var data []models.Todos
	db.Where("id=?", uid).Find(&data)
	if len(data) == 0 {
		c.JSON(200, gin.H{
			"code":    404,
			"message": "没有查询到数据",
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "查询到数据",
			"id":      data[0].Id,
			"title":   data[0].Name,
			"content": data[0].Passwd,
		})
	}
}
func Add(c *gin.Context) {
	db := models.Db
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
}
func Delete(c *gin.Context) {
	db := models.Db
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
}
func Update(c *gin.Context) {
	db := models.Db
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
}
