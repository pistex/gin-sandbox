package middlewares

import (
	"errors"
	"fmt"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/messages"
	"kwanjai/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// JWT middleware
func JWT(ctx interfaces.IContext) gin.HandlerFunc {
	return func(g *gin.Context) {
		// Get header from Authorization header
		authHeader := g.GetHeader("Authorization")
		if authHeader == "" {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Split authHeader which should be in from of `TokenPrefix Token` by space
		split := strings.Split(authHeader, " ")
		if len(split) != 2 {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if token prefix is `Bearer`
		if split[0] != "Bearer" {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Verify token
		auhtService := services.NewAuthService(ctx)
		token, err := auhtService.VerifyToken(split[1])
		if errors.Is(err, messages.ErrLoadPrivateKey) {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if err != nil {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claim := &jwt.StandardClaims{}
		err = helpers.JsonMapper(token.Claims, claim)
		if err != nil {
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if token is still has not been revoked or expired
		sessionName := fmt.Sprintf("%s:%s", claim.Subject, g.ClientIP())
		tokenID := ctx.Cache().Get(ctx.Config().Context, sessionName).Val()
		if tokenID != claim.Id {
			// This means token is not issued by authService
			g.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Store user in context
		// If an error occurs from this step should be from internal
		userService := services.NewUserService(ctx)
		id, err := uuid.Parse(claim.Subject)
		if err != nil {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		user, err := userService.Find(id)
		if err != nil {
			g.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		g.Set("user", user)
		g.Set("tokenID", tokenID)
	}
}

// AuthenticatedOnly disallows "anonymous" user.
func AuthenticatedOnly(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		username := helpers.GetUsername(ginContext)
		if username == "anonymous" {
			ginContext.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
