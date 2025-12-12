package auth

import (
	"context"
	"main/internal/services/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service auth.Service
}

func NewController(service *auth.Service) Controller {
	return Controller{
		service: *service,
	}
}

func (as Controller) SendOtp(c *gin.Context) {
	var phoneNumer auth.PhoneRequest

	if err := c.ShouldBindJSON(&phoneNumer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	id, token, err := as.service.SendEmailCode(ctx, phoneNumer.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!", "id": id, "token": token})
}
