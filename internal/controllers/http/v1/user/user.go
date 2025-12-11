package user

import (
	"context"
	"main/internal/services/user"
	use_case "main/internal/usecase/user"
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

func (ca Controller) Update(c *gin.Context) {
	var data user.Update
	authHeader := c.GetHeader("Authorization")
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	avatar, _ := c.FormFile("avatar")

	ctx := context.Background()
	if avatar != nil {
		filePath, err := ca.useCase.Upload(ctx, avatar, "./users/avatars")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		data.Avatar = &filePath
	}

	detail, err := ca.useCase.Update(ctx, data, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"detail": detail, "message": "ok!"})
}

func (ca Controller) Delete(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	ctx := context.Background()

	if err := ca.useCase.Delete(ctx, authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!"})
}
