package controllers

/*home相关处理函数*/
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//处理与get相关的操作，包括渲染html页面，检测是否已有账户登录
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title":   "home",
		"IsLogin": CheckAccount(c),
	})
}
