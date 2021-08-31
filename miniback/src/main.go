package main

import (
	"gominiback/backend/controller"
	"gominiback/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.POST("/data", controller.GetStudentData)
	r.GET("/data", controller.MySelectedData)
	r.GET("/updated", controller.UpdateStudent)
	r.GET("/delete", controller.DeleteStudent)
	r.Run(":3000")
}
