package types

import (
	"context"
	"time"
)

type Config struct {
	BaseDirectory                string
	Port                         string
	FrontendURL                  string
	BackendURL                   string
	EmailServicePassword         string
	EmailVerficationLifetime     time.Duration
	Context                      context.Context
	OmisePublicKey               string
	OmiseSecretKey               string
	CodeChallengeMethod          string
	AuthorizationRequestLifetime time.Duration
	JWTTokenLifetime             time.Duration
	NonceLifetime                time.Duration
}
