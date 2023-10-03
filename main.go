package main

import (
	"fmt"
	"project_shopping_tour/service_admin_crud/controller"

	"github.com/gin-gonic/gin"
)

const port string = ":9292"

func main() {
	router := gin.New()

	router.GET("/debug", controller.ControllerDebug)
	router.POST("/api/upload", controller.ControllerUploadFile)
	router.POST("/api/update/product", controller.ControllerProduct)
	router.POST("api/update/content", controller.ControllerContent)

	err := router.Run(port)
	if err != nil {
		fmt.Print("Service admin CRUD fail to start" + err.Error())
	}
	fmt.Println("Service admin CRUD start at port" + port + "debug => " + "http://localhost:" + port + "/debug")
}
