package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ControllerUploadFile(c *gin.Context) {
	pwd, _ := os.Getwd()
	file, err := c.FormFile("file")
	var dst = pwd + "/" + "files/" + file.Filename

	fmt.Println(file.Size)
	if file.Size > 5000000 {
		c.JSON(http.StatusRequestHeaderFieldsTooLarge, gin.H{"status": "file must less than 5MB"})
	} else {
		if err != nil {
			panic(err)
		}
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		c.JSON(http.StatusOK, gin.H{"status": "uploaded"})
	}

}
