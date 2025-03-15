package handler

import (
	"net/http"

	"github.com/gabrielbsx/opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "could not delete opening")
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-opening", opening)
}
