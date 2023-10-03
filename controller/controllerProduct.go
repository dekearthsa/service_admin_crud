package controller

import (
	"net/http"
	"project_shopping_tour/service_admin_crud/model"

	"github.com/gin-gonic/gin"
)

func ControllerProduct(ctx *gin.Context) {
	var setProduct model.Product
	ctx.JSON(http.StatusOK, gin.H{"status": setProduct})
}
