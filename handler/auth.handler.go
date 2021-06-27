package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/dto"
	"github.com/ydhnwb/opaku-dummy-backend/service"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthHandler(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_BAD_REQUEST, http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_MESSAGE_UNAUTHORIZED, http.StatusUnauthorized)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := c.userService.FindUserByEmail(loginRequest.Email)

	token := c.jwtService.GenerateToken(fmt.Sprintf("%v", user.ID))
	user.Token = token
	ctx.JSON(http.StatusOK, user)

}

func (c *authHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := common.BuildErrorResponse(common.DEFAULT_ERROR_BAD_REQUEST, http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		splittedError := strings.Split(err.Error(), "\n")
		response := common.BuildErrorResponse(splittedError[0], http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	token := c.jwtService.GenerateToken(fmt.Sprintf("%v", user.ID))
	user.Token = token
	ctx.JSON(http.StatusCreated, user)

}
