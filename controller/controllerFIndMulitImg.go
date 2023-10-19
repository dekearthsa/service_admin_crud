package controller

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"service_admin_crud/model"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func ControllerFindMulitImg(c *gin.Context) {
	var req model.ArrayImgPath
	if err := c.BindJSON(&req); err != nil {
		log.Println("error BindJSON => ", err)
	}
	log.Println(req)
	var arrayBase64 []string
	for _, element := range req.ArrayImg {

		// imagePath := req.ID + "_" + element
		// log.Println("imagePath => ", imagePath)

		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Println("can't create new client ", err)
		}

		buckets := client.Bucket("demostoragebucketearth")
		rc, err := buckets.Object(element).NewReader(ctx)
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
		// // application/octet-stream // //
		str := base64.StdEncoding.EncodeToString(byteFile)
		arrayBase64 = append(arrayBase64, str)
		// c.Header("Content-Disposition", "attachment; filename=file-name.png")
		// c.Data(http.StatusOK, "image/png", byteFile)
	}

	c.JSON(http.StatusOK, gin.H{"arrayImg": arrayBase64})

}
