package middlewares

import (
	"auth-service/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err
		var appErr *utils.AppError
		if errors.As(err, &appErr) {
			c.JSON(appErr.Status, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    "INTERNAL_SERVER_ERROR",
				"message": "internal server error",
			})
		}
	}
}
