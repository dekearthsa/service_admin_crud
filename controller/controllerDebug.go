package controller

import "github.com/gin-gonic/gin"

func ControllerDebug(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status": "OK"})
}
