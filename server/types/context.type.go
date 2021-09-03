package types

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	BaseDirectory                   string
	Port                            string
	FirebaseProjectID               string
	FrontendURL                     string
	BackendURL                      string
	EmailServicePassword            string
	EmailVerficationLifetime        *time.Duration
	JWTAccessTokenSecretKey         string
	JWTRefreshTokenSecretKey        string
	JWTAccessTokenLifetime          *time.Duration
	JWTRefreshTokenLifetime         *time.Duration
	Context                         *context.Context
	DefaultAuthenticationMiddleware gin.HandlerFunc
	OmisePublicKey                  string
	OmiseSecretKey                  string
}
