package controllers

import (
	"class5/models"
	//"class5/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
type Student struct{
	Sid string
	Name string
	Age int64
	Sex string
}
func init()  {
	models.RegisterDB()
}
/* 获取全部学生信息 op == getAllStudents；
 */
func StudentGet(c *gin.Context)  {
	//c.JSON(http.StatusOK,"test 1")
	//c.JSON(http.StatusOK,"test 2")
	if c.Query("op") == "getAllStudents"{
		c.JSON(http.StatusOK, models.QueryAll())
	}

}
/* 已知学号获取信息: op == getOneStudent, 输入：学号：sId；
	输入学生信息: op == sInsert, 输入：json数组 输出：重复插入的Student数据；
	删除学生信息(一个，多个)：op == sDelete, 输入：Sid数组；
	删除学生信息（全部）：op == sDeleteAll；
	修改学生信息: op == sModify, 输入：修改了数据的Student信息（完整信息）的json数组；
 */
func StudentPost(c *gin.Context)  {
	if c.Query("op") == "getOneStudent"{
		s := models.QueryRow(c.PostForm("sId"))
		if s == nil {
			c.JSON(http.StatusOK, "查无此人")
		}else{
			c.JSON(http.StatusOK, s)
		}

	}
	if c.Query("op") == "sInsert"{
		s := make([]models.Student,0)
		err:=c.BindJSON(&s)
		//err :=c.ShouldBindJSON(&s)
		if err != nil{
			fmt.Println(err)
		}
		sIn := models.Insert(s)
		c.JSON(http.StatusOK, sIn)

	}
	if c.Query("op") == "sDelete" {
		s := c.PostFormArray("sId")
		models.DeleteFew(s)
	}
	if c.Query("op") == "sDeleteAll"{
		models.DeleteAll()
	}
	if c.Query("op") == "sModify"{
		s := make([]models.Student,0)
		c.BindJSON(&s)
		models.Modify(s)
	}
}