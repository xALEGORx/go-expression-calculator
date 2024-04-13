package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/xALEGORx/go-expression-calculator/internal/orchestrator/repositories"
	"github.com/xALEGORx/go-expression-calculator/pkg/jwt"
	"github.com/xALEGORx/go-expression-calculator/pkg/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := repositories.UserRepository().GetById(jwt.New().JwtUserId(c)); err != nil {
			response.BadRequest(c, "the current user does not exist or has been logged out")
			c.Abort()
			return
		}

		c.Next()
	}
}
