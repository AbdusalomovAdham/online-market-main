package products

import (
	"context"
	"log"
	"main/internal/entity"
	product "main/internal/services/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service product.Service
}

func NewController(service product.Service) Controller {
	return Controller{service: service}
}

func (as Controller) CreateProduct(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var data product.Create
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	images := form.File["images"]
	log.Println("Images uploaded successfully", images)
	ctx := context.Background()
	if len(images) > 0 {
		imgFile, err := as.service.MultipleUpload(ctx, images, "../media/products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		data.Images = &imgFile
	}

	id, err := as.service.CreateProduct(c, data, authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "ok!"})
}

func (as Controller) GetById(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "uz"
	}
	ctx := context.Background()
	product, err := as.service.GetById(ctx, productId, lang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok!", "data": product})
}

func (as Controller) GetProductsList(c *gin.Context) {
	filter := entity.Filter{}
	query := c.Request.URL.Query()

	// category id
	categoryId := c.Query("category_id")
	categoryInt, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Category id must be number!",
		})
		return
	}
	categoryInt64 := int64(categoryInt)
	filter.CategoryId = &categoryInt64

	// limit
	limitQ := query["limit"]
	if len(limitQ) > 0 {
		queryInt, err := strconv.Atoi(limitQ[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Limit must be number!",
			})
			return
		}

		filter.Limit = &queryInt
	}

	// offset
	offsetQ := query["offset"]
	if len(offsetQ) > 0 {
		queryInt, err := strconv.Atoi(offsetQ[0])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Offset must be number!",
			})
			return
		}
		filter.Offset = &queryInt
	}

	// discount
	discountQ := query["discount"]
	if len(discountQ) > 0 && discountQ[0] != "" {
		val := discountQ[0] == "true"
		filter.DiscountOnly = &val
	}

	// max price
	maxPriceQ := query["max_price"]
	if len(maxPriceQ) > 0 && maxPriceQ[0] != "" {
		intVal, err := strconv.Atoi(maxPriceQ[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "max price must be a number!",
			})
			return
		}
		floatVal := float64(intVal)
		filter.PriceMax = &floatVal
	}

	// min price
	minPriceQ := query["min_price"]
	if len(minPriceQ) > 0 && minPriceQ[0] != "" {
		intVal, err := strconv.Atoi(minPriceQ[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "min price must be a number!",
			})
			return
		}
		floatVal := float64(intVal)
		filter.PriceMin = &floatVal
	}

	// rating
	ratingQ := query["rating"]
	if len(ratingQ) > 0 && ratingQ[0] != "" {
		intVal, err := strconv.Atoi(ratingQ[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "rating must be a number!",
			})
			return
		}

		if intVal > 5 || intVal < 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "rating invalid",
			})
			return
		}

		floatVal := float64(intVal)
		filter.PriceMin = &floatVal
	}

	// search
	searchQ := query["search"]
	if len(searchQ) > 0 && searchQ[0] != "" {
		filter.Search = &searchQ[0]
	}

	// order
	orderQ := query["order"]
	if len(orderQ) > 0 {
		filter.Order = &orderQ[0]
	}

	ctx := context.Background()

	authHeader := c.GetHeader("Authorization")
	lang := c.GetHeader("Accept-Language")

	if lang == "" {
		defaultLang := "uz"
		filter.Language = &defaultLang
	} else {
		filter.Language = &lang
	}

	list, count, err := as.service.GetList(ctx, filter, authHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"data": map[string]any{
			"results": list,
			"count":   count,
		},
	})
}
