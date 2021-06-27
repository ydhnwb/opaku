package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/opaku-dummy-backend/common"
	"github.com/ydhnwb/opaku-dummy-backend/service"
	// "github.com/ydhnwb/golang_heroku/common/response"
	// "github.com/ydhnwb/golang_heroku/service"
)

//AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res := common.BuildErrorResponse(common.DEFAULT_ERROR_BAD_REQUEST, http.StatusBadRequest)
			c.AbortWithStatusJSON(http.StatusBadRequest, res)
			return
		}

		token := jwtService.ValidateToken(authHeader, c)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			res := common.BuildErrorResponse(common.DEFAULT_ERROR_TOKEN_NOT_VALID, http.StatusBadRequest)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		}
	}
}
