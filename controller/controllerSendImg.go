package controller

import (
	"context"
	"io"
	"log"
	"net/http"
	"service_admin_crud/model"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func ControllerSendImg(c *gin.Context) {
	var res model.ImgPath
	ctx := context.Background()
	BUCKET := "demostoragebucketearth"
	imagePath := res.UserID + "_" + res.ImageName
	if err := c.BindJSON(&res); err != nil {
		log.Println("error BindJSON => ", err)
	}
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("can't create new client ", err)
	}

	buckets := client.Bucket(BUCKET)
	rc, err := buckets.Object(imagePath).NewReader(ctx)
	if err != nil {
		log.Println("err when fetch image from bucket", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"Status": "err when fetch image from bucket."})
	}
	byteFile, err := io.ReadAll(rc)
	defer rc.Close()
	if err != nil {
		log.Println("err read file from bucket")
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "err read file from bucket."})
	}

	c.Data(http.StatusOK, "application/octet-stream", byteFile)
	// c.JSON(http.StatusOK, gin.H{"Status": "Sent"})

}
