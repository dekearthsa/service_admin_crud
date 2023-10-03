package controller

import (
	"net/http"
	"project_shopping_tour/service_admin_crud/model"

	"github.com/gin-gonic/gin"
)

func ControllerContent(ctx *gin.Context) {
	var setContent model.Content
	ctx.JSON(http.StatusOK, gin.H{"status": setContent})
}
