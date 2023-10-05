package controller

import (
	"net/http"
	"service_admin_crud/model"

	"github.com/gin-gonic/gin"
)

func ControllerContent(ctx *gin.Context) {
	var setContent model.Content
	ctx.JSON(http.StatusOK, gin.H{"status": setContent})
}
