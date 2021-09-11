package types

import (
	"context"
	"time"
)

type Config struct {
	BaseDirectory            string
	Port                     string
	FrontendURL              string
	BackendURL               string
	EmailServicePassword     string
	EmailVerficationLifetime *time.Duration
	JWTAccessTokenSecretKey  string
	JWTRefreshTokenSecretKey string
	JWTAccessTokenLifetime   *time.Duration
	JWTRefreshTokenLifetime  *time.Duration
	Context                  *context.Context
	OmisePublicKey           string
	OmiseSecretKey           string
}
