package order

import (
	"context"
	"encoding/json"
	"main/internal/services/order"
	use_case "main/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *use_case.UseCase
}

func NewController(useCase *use_case.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

func (uc Controller) CreateOrder(c *gin.Context) {
	var order order.Create

	authHeader := c.GetHeader("Authorization")
	if err := c.ShouldBind(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	err := uc.useCase.CreateOrder(ctx, order, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order created"})
}

func (uc Controller) GetList(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	list, err := uc.useCase.GetList(context.Background(), authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": list})
}

func (uc Controller) OrderAccept(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	orderUUID := c.Param("uuid")

	var data order.Update

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	detail, err := uc.useCase.OrderAccept(context.Background(), authHeader, orderUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(detail), &jsonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid JSON from detail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order accepted", "data": jsonData})
}
