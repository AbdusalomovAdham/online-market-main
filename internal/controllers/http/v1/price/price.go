package price

import (
	"context"
	"main/internal/services/price"
	use_case "main/internal/usecase/price"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *use_case.UseCase
}

func NewController(useCase *use_case.UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

func (cu Controller) GetPriceByLocation(c *gin.Context) {
	var price price.GetPrice
	if err := c.ShouldBindJSON(&price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	prices, err := cu.useCase.GetPrice(ctx, price.FromLocationID, price.ToLocationID, price.TariffID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "ok!", "data": prices})
}
