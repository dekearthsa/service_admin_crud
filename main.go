package main

import (
	"fmt"
	"service_admin_crud/controller"

	"github.com/gin-gonic/gin"
)

const port string = ":7475"

func main() {
	router := gin.New()

	router.GET("/api/debug", controller.ControllerDebug)
	router.POST("/api/upload", controller.ControllerUploadFile)
	router.POST("/api/update/product", controller.ControllerProduct)
	router.POST("/api/update/content", controller.ControllerContent)

	err := router.Run(port)
	if err != nil {
		fmt.Print("Service admin CRUD fail to start" + err.Error())
	}
	fmt.Println("Service admin CRUD start at port" + port + "debug => " + "http://localhost:" + port + "/debug")
}
