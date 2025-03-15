package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, gin.H{
		"error": message,
	})
}

func sendSuccess(ctx *gin.Context, code int, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, gin.H{
		"message": fmt.Sprintf("%s successfully", operation),
		"data":    data,
	})
}
