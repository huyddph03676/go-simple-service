package middleware

import (
	"go-simple-service/common"
	"go-simple-service/component"

	"github.com/gin-gonic/gin"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Context-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)

					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)

				return
			}
		}()

		c.Next()
	}
}
