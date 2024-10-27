// middlewares/auth.go

package middlewares

import (
	"ekeberg.com/go-api-sql-gcp-products/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, humanOrService, err := utils.VerifyToken(token) // Get human_or_service from VerifyToken

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)
	context.Set("human_or_service", humanOrService) // Set human_or_service in context
	context.Next()
}
