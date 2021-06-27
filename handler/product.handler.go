package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/service"
	"gorm.io/gorm"
)

type ProductHandler interface {
	All(ctx *gin.Context)
	FindOneProductByID(ctx *gin.Context)
	FindProductsByName(ctx *gin.Context)
}

type productHandler struct {
	productService service.ProductService
	jwtService     service.JWTService
}

func NewProductHandler(productService service.ProductService, jwtService service.JWTService) ProductHandler {
	return &productHandler{
		productService: productService,
		jwtService:     jwtService,
	}
}

func (c *productHandler) All(ctx *gin.Context) {
	products, err := c.productService.All()
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *productHandler) FindOneProductByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.productService.FindOneProductByID(id)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response := common.BuildErrorResponse(splittedError[0], http.StatusNotFound)
			ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		} else {
			response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		}
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *productHandler) FindProductsByName(ctx *gin.Context) {
	q := ctx.Query("q")
	if q == "" {
		response := common.BuildErrorResponse("Search query cannot be empty", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.productService.FindProductsByName(q)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
