package models

//struct
type Todos struct { //根据需求补充标签
	Id     uint   `gorm:"primary_key" json:"id" binding:"required"`
	Name   string `json:"name"`
	Passwd string `json:"psw"`
}
type Users struct {
	Id      uint   `gorm:"primary key" json:"id"`
	User_id uint   `json:"userid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
