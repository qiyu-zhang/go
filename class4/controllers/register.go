package controllers

/*register相关处理函数*/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learningProject/models"
	"net/http"
)

//处理与get相关的操作，包括渲染html页面
func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "register",
	})

}

//处理与post相关的操作，包括将输入账号给存入数据库
func RegisterPost(c *gin.Context) {
	name := c.PostForm("user")
	pwd := c.PostForm("passwd")
	if models.InsertRow(name, pwd) {
		fmt.Println("注册成功")
		c.SetCookie("name", name, 3600, "/", "localhost", false, true)
		c.SetCookie("pwd", pwd, 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		fmt.Println("注册失败")
		c.Redirect(302, "/register")
	}
}
