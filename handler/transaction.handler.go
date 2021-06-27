package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/service"
)

type TransactionHandler interface {
	FindAllMyTransaction(ctx *gin.Context)
	CreateTransaction(ctx *gin.Context)
}

type transactionHandler struct {
	transactionService service.TransactionService
	jwtService         service.JWTService
}

func NewTransactionHandler(transactionService service.TransactionService, jwtService service.JWTService) TransactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
		jwtService:         jwtService,
	}
}

func (c *transactionHandler) FindAllMyTransaction(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	transactions, err := c.transactionService.FindAllMyTransaction(userID)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		res := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (c *transactionHandler) CreateTransaction(ctx *gin.Context) {
	var createTransactionRequest dto.CreateTransactionRequest
	err := ctx.ShouldBind(&createTransactionRequest)

	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		res := common.BuildErrorResponse(splittedError[0], http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(authHeader, ctx)
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])

	res, err := c.transactionService.CreateTransaction(createTransactionRequest, userID)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		res := common.BuildErrorResponse(splittedError[0], http.StatusUnprocessableEntity)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, res)
		return

	}

	ctx.JSON(http.StatusOK, res)

}
