package controller

import (
	"context"
	"log"
	"service_admin_crud/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func ControllerProduct(c *gin.Context) {
	var res model.Product
	ctx := context.Background()

	// log.Println("1. res => ", res)
	if err := c.BindJSON(&res); err != nil {
		log.Fatal("err BindJSON =>", err)
	}

	client, err := datastore.NewClient(ctx, "scg-iat-project-coretech")
	if err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"message": "can't find projectID."})
	}

	// log.Println("2. res => ", res)

	// log.Println("3. e1 => ", e1)
	key := datastore.IncompleteKey("go_product", nil)
	if _, err := client.Put(ctx, key, &res); err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"message": "can't insert data."})
	}
	c.JSON(200, gin.H{"status": "Inserted."})
}
