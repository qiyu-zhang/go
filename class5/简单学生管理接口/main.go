package main

import (
	"class5/controllers"
	"github.com/gin-gonic/gin"
)


func main()  {
	r := gin.Default()
	r.GET("/",controllers.StudentGet)
	r.POST("/", controllers.StudentPost)
	r.Run(":8080")
}