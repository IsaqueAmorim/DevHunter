package handler

import (
	"github.com/IsaqueAmorim/DevHunter/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) *schemas.Opening {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		return nil
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Salary:   request.Salary,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening %v", err.Error())
		return nil
	}

	return &opening
}
