package controllers

/*login相关处理函数*/
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learningProject/models"
	_ "log"
	"net/http"
)

//登出函数
func Logout(c *gin.Context) {
	c.SetCookie("name", "", -1, "/", "localhost", false, true)
	c.SetCookie("pwd", "", -1, "/", "localhost", false, true)

	c.Redirect(302, "/")
}

//处理与get相关的操作，包括渲染html页面，判断并执行账号登出操作
func LoginGet(c *gin.Context) {
	if c.Query("exit") == "true" {
		Logout(c)

	}
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "login",
	})

}

//处理与post相关的操作，
//包括验证输入账号密码是否存在于数据库
//若数据库检测不存在，则跳转到注册页面
func LoginPost(c *gin.Context) {
	name := c.PostForm("user")
	pwd := c.PostForm("passwd")
	ck := models.QueryRow(name)
	fmt.Println(name)
	fmt.Println(ck)
	if ck == pwd {
		fmt.Printf("登录成功！")
		c.SetCookie("name", name, 3600, "/", "localhost", false, true)
		c.SetCookie("pwd", pwd, 3600, "/", "localhost", false, true)
		c.Redirect(302, "/")
	} else {
		fmt.Println("查无此人")
		c.Redirect(302, "/register")
	}

}

//用于判断是否有账户已经登录
func CheckAccount(c *gin.Context) bool {
	var ret bool
	user, err := c.Cookie("name")
	if err != nil || user == "" {
		return false
	}
	var pwd string
	pwd, err = c.Cookie("pwd")
	if models.QueryRow(user) == pwd {
		ret = true
	} else {
		ret = false
	}
	return ret
}
