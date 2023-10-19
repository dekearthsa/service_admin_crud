package controller

import (
	"context"
	"io"
	"log"
	"net/http"
	"service_admin_crud/model"
	"cloud.google.com/go/datastore"
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func ControllerContent(c *gin.Context) {

	var arrayPath []string
	form, err := c.MultipartForm()
	ctx := context.Background()
	if err != nil {
		log.Println("err before FormFile => ", err)
	}

	files := form.File["images"]
	isText := form.Value["TestID"]
	if files == nil {
		log.Println("err before from.file => ", err)
	}

	clientDatastore, err := datastore.NewClient(ctx, "scg-iat-project-coretech")
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't find projectID."})
	}
	key := datastore.IncompleteKey("test_go_product", nil)

	for _, file := range files {
		size := file.Size
		if size >= 5000000 {
			log.Fatal("err before files[0].Size => ", err)
			c.JSON(http.StatusRequestHeaderFieldsTooLarge, gin.H{"Status": "file must less than 5MB."})
		}

		src, err := file.Open()
		if err != nil {
			log.Println("err files[0].Open() => ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "file"})
		}
		defer src.Close()

		imagePath := isText[0] + "_" + file.Filename

		arrayPath = append(arrayPath, imagePath)
		clientStorage, err := storage.NewClient(ctx)
		if err != nil {
			log.Println("err storage.NewClient(ctx) => ", err)
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "internal error."})
		}

		bucket := clientStorage.Bucket("demostoragebucketearth")
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
	}

	payload := model.Testing{
		TestID:   isText[0],
		ArrayImg: arrayPath,
	}

	if _, err := clientDatastore.Put(ctx, key, &payload); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't insert data."})
	}
	c.JSON(http.StatusOK, gin.H{"payload": payload})
}
