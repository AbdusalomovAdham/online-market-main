package auth

import (
	"context"
	auth_service "main/internal/services/auth"
	"main/internal/usecase/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	useCase *auth.UseCase
}

func NewController(useCase *auth.UseCase) *Controller {
	return &Controller{useCase: useCase}
}

// func (as Controller) SignUp(c *gin.Context) {
// 	var email string
// 	if err := c.BindJSON(&email); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}

// 	ctx := context.Background()
// 	token, detail, err := as.useCase.SignUp(ctx, email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"token":   token,
// 		"data":    detail,
// 		"message": "ok!",
// 	})
// }

func (ac Controller) SendEmailCode(c *gin.Context) {
	var data auth_service.SendEmailCode

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()
	token, email, err := ac.useCase.SendEmailCode(ctx, data.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "email": email})
}

func (as Controller) CheckCode(c *gin.Context) {
	var data auth_service.CheckCode

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()

	detail, token, err := as.useCase.CheckCode(ctx, data.Code, data.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!", "data": detail, "token": token})
}

func (as Controller) ResendCode(c *gin.Context) {
	var data auth_service.ResendCode

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx := context.Background()
	if err := as.useCase.ResendCode(ctx, data.Token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "code resend"})
}
