package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/service"
)

type CartHandler interface {
	FindAllCart(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
}

type cartHandler struct {
	cartService service.CartService
	jwtService  service.JWTService
}

func NewCartHandler(cartService service.CartService, jwtService service.JWTService) CartHandler {
	return &cartHandler{
		cartService: cartService,
		jwtService:  jwtService,
	}
}

func (c *cartHandler) FindAllCart(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.cartService.FindAllCart(userID)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (c *cartHandler) AddToCart(ctx *gin.Context) {
	var addToCartRequest dto.AddToCartRequest
	err := ctx.ShouldBind(&addToCartRequest)

	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	id, _ := strconv.ParseUint(userID, 10, 64)

	res, err := c.cartService.AddToCart(addToCartRequest, id)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}
	ctx.JSON(http.StatusOK, res)

}

func (c *cartHandler) DeleteCart(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	print(userID)

	//consider to check if cart.user.id is equal to current logged in user
	id := ctx.Param("id")

	err := c.cartService.DeleteCart(id)

	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	res := common.BuildErrorResponse("OK!", http.StatusOK)
	ctx.JSON(http.StatusOK, res)

}
