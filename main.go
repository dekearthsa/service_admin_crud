package main

import (
	"fmt"
	"service_admin_crud/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const port string = ":7475"

func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers"},
	}))

	router.GET("/api/debug", controller.ControllerDebug)
	router.POST("/api/upload", controller.ControllerUploadFile)
	router.POST("/api/update/product", controller.ControllerProduct)
	router.POST("/api/update/content", controller.ControllerContent)
	router.POST("/api/send/img", controller.ControllerSendImg)
	router.PUT("/api/upload/img", controller.ControllerUploadImage)

	err := router.Run(port)
	if err != nil {
		fmt.Print("Service admin CRUD fail to start" + err.Error())
	}
	fmt.Println("Service admin CRUD start at port" + port + "debug => " + "http://localhost:" + port + "/debug")
}
