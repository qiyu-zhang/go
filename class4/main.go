package main

/*主程序，
1、绑定html页面；
2、绑定各个html的处理函数
3、启动服务
*/
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"learningProject/controllers"
)

func main() {
	r := gin.Default()
	//html加载模板
	r.LoadHTMLGlob("templates/*")
	//绑定路由
	r.GET("/", controllers.Home)
	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)
	r.POST("/register", controllers.RegisterPost)
	r.GET("/register", controllers.RegisterGet)
	//启动服务
	r.Run(":8080")
}
