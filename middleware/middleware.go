package middleware

import (
	"context"
	"net/http"
	"strings"
	"github.com/1Nelsonel/Savannah/config"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if token == authorizationHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
		c.Abort()
		return
	}

	idToken, err := config.Verifier.Verify(context.Background(), token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	var claims struct {
		Email string `json:"email"`
	}
	if err := idToken.Claims(&claims); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse claims"})
		c.Abort()
		return
	}

	c.Set("email", claims.Email)
	c.Next()
}