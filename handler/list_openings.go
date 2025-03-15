package handler

import (
	"net/http"

	"github.com/gabrielbsx/opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "could not list openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
