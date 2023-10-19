package controller

import (
	"context"
	"log"
	"net/http"
	"service_admin_crud/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func ControllerGetAllPath(c *gin.Context) {
	ctx := context.Background()

	clientDatastore, err := datastore.NewClient(ctx, "scg-iat-project-coretech")
	if err != nil {
		log.Println(err)
	}
	var itemData []model.StrucImgSelect
	keys, err := clientDatastore.GetAll(ctx, datastore.NewQuery("test_go_product"), &itemData)
	if err != nil {
		log.Println("err in ControllerGetAllPath => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": err})
	}
	for i, key := range keys {
		log.Println(key)
		log.Println(itemData[i])
	}

	c.JSON(http.StatusOK, gin.H{"body": itemData})
}
