package handler

import (
	"fmt"

	"github.com/gabrielbsx/opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, ErrorResponse{
		Message: message,
	})
}

func sendSuccess(ctx *gin.Context, code int, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(code, SuccessResponse{
		Message: fmt.Sprintf("%s successfully", operation),
		Data:    data.(schemas.OpeningResponse),
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
