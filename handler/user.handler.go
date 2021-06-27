package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/service"
)

type UserHandler interface {
	Profile(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserHandler(
	userService service.UserService,
	jwtService service.JWTService,
) UserHandler {
	return &userHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userHandler) getUserIDByHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_MESSAGE_UNAUTHORIZED, http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}

func (c *userHandler) Update(ctx *gin.Context) {
	var updateUserRequest dto.UpdateUserRequest

	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_BAD_REQUEST, http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

		return
	}

	id := c.getUserIDByHeader(ctx)

	if id == "" {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_MESSAGE_UNAUTHORIZED, http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_id, _ := strconv.ParseInt(id, 0, 64)
	updateUserRequest.ID = _id
	res, err := c.userService.UpdateUser(updateUserRequest)

	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_MESSAGE_UNAUTHORIZED, http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	ctx.JSON(http.StatusOK, res)

}

func (c *userHandler) Profile(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)

	if token == nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_MESSAGE_UNAUTHORIZED, http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
	}

	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user, err := c.userService.FindUserByID(id)

	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_BAD_REQUEST, http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	ctx.JSON(http.StatusOK, user)
}
