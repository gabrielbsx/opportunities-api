package handler

import (
	"net/http"

	"github.com/gabrielbsx/opportunities-api/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Failed to create opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "Failed to create opening")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "create-opening", opening)
}
