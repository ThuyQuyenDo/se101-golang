package middleware

import (
	"go-ecommerce/common"
	"go-ecommerce/config"
	"go-ecommerce/provider/tokenprovider/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong authen header",
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearer {token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func JWT(cfg *config.AppConfig) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(cfg.JWTSecretKey)

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user := common.JWTUserData{
			ID:   *payload.UserID,
			Role: payload.Role,
		}

		if err != nil {
			panic(err)
		}

		c.Set(common.CurrentRequester, user)
		c.Next()
	}
}
