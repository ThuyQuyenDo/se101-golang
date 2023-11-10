package middleware

import (
	"go-ecommerce/common"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Avoid case that response result type is text
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// Re-enable panicking mechanism for `Gin lib` cuz `Gin` has its own recovery
					panic(err)
				}

				// `err.(error)` just return a error cuz `err` is of type result `recover()`
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
