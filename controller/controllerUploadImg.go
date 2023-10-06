package controller

import (
	"context"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

// Buckets := "scg-iat-project-coretech.appspot.com"

func ControllerUploadImage(c *gin.Context) {
	uid := c.Request.Header.Get("userID")
	setFileName, err := c.FormFile("images")
	if err != nil {
		log.Println("err c.FormFile(images)", err)
		c.JSON(http.StatusNoContent, gin.H{"Status": "not found file name."})
	}

	form, err := c.MultipartForm()
	ctx := context.Background()
	if err != nil {
		// panic(err)
		log.Println("err before FormFile => ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"Status": "Not found file."})
	}
	files := form.File["images"]
	if files != nil {
		log.Println("err before from.file => ", err)
		// c.JSON(http.StatusNoContent, gin.H{"Status": "file is empty."})
	}

	size := files[0].Size
	if size >= 5000000 {
		log.Println("err before files[0].Size => ", err)
		c.JSON(http.StatusRequestHeaderFieldsTooLarge, gin.H{"Status": "file must less than 5MB."})
	}

	src, err := files[0].Open()
	if err != nil {
		log.Println("err files[0].Open() => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "file"})
	}
	defer src.Close()

	imagePath := uid + "_" + setFileName.Filename
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Println("err storage.NewClient(ctx) => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Status": "internal error."})
	}

	bucket := client.Bucket("demostoragebucketearth")
	wc := bucket.Object(imagePath).NewWriter(ctx)
	_, err = io.Copy(wc, src)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "StatusServiceUnavailable"})
	}

	err = wc.Close()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "StatusServiceUnavailable"})
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Upload image sucess."})

}
