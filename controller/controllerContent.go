package controller

import (
	"context"
	"log"
	"net/http"
	"service_admin_crud/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func ControllerContent(c *gin.Context) {
	var req model.Content
	if err := c.BindJSON(&req); err != nil {
		log.Println("err BindJSON =>", err)
		// c.JSON(http.StatusBadRequest, gin.H{"message": "payload is empty"})
	}
	ctx := context.Background()
	clientDatastore, err := datastore.NewClient(ctx, "scg-iat-project-coretech")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't find projectID."})
	}

	key := datastore.IncompleteKey("go_product", nil)
	if _, err := clientDatastore.Put(ctx, key, &req); err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't insert data."})
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
