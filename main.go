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
	// router.POST("/api/test/user", func(c *gin.Context) {

	// 	var res model.Product
	// 	ctx := context.Background()

	// 	// log.Println("1. res => ", res)
	// 	if err := c.BindJSON(&res); err != nil {
	// 		log.Fatal("err BindJSON =>", err)
	// 	}

	// 	client, err := datastore.NewClient(ctx, "scg-iat-project-coretech")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		c.JSON(200, gin.H{"message": "can't find projectID."})
	// 	}

	// 	// log.Println("2. res => ", res)

	// 	// log.Println("3. e1 => ", e1)
	// 	key := datastore.IncompleteKey("go_product", nil)
	// 	if _, err := client.Put(ctx, key, &res); err != nil {
	// 		log.Fatal(err)
	// 		c.JSON(200, gin.H{"message": "can't insert data."})
	// 	}
	// 	c.JSON(200, gin.H{"status": "Inserted."})
	// })

	err := router.Run(port)
	if err != nil {
		fmt.Print("Service admin CRUD fail to start" + err.Error())
	}
	fmt.Println("Service admin CRUD start at port" + port + "debug => " + "http://localhost:" + port + "/debug")
}
