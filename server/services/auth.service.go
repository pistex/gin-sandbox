package services

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/messages"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type IAuthService interface {
	Create(codeChallenge string, codeChallengeMethod string) (authorizationCode string, err error)
	Revoke(name string) error
	Login(email string, password string, authorizationCode string, codeVerifier string) (nonce string, err error)
	Logout(nonce string, userID string, userIP string) (err error)
	CreateToken(authorizationCode string, codeVerifier string, nonce string, userIP string) (newNonce string, token string, err error)
	VerifyToken(tokenString string) (token *jwt.Token, err error)
}

type authService struct {
	ctx         interfaces.IContext
	userService IUserService
}

func NewAuthService(ctx interfaces.IContext) IAuthService {
	return &authService{
		ctx:         ctx,
		userService: NewUserService(ctx),
	}
}

func (s *authService) Login(email string, password string, authorizationCode string, codeVerifier string) (string, error) {
	codeChallenge := s.ctx.Cache().Get(s.ctx.Config().Context, authorizationCode).Val()
	if codeChallenge == "" {
		return "", messages.ErrBadAuthenticationSession
	}

	hashedCodeVerifier := helpers.HashStringToBase64(codeVerifier, crypto.SHA256)
	if hashedCodeVerifier != codeChallenge {
		err := s.Revoke(authorizationCode)
		if err != nil {
			return "", err
		}
		return "", messages.ErrBadAuthenticationSession
	}

	err := s.ctx.Cache().Del(s.ctx.Config().Context, authorizationCode).Err()
	if err != nil {
		return "", err
	}

	user, err := s.userService.FindByEmail(email)
	if err != nil {
		return "", messages.ErrCredentialMismatch
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return "", messages.ErrCredentialMismatch
	}

	nonce := helpers.HashString(user.ID.String()+time.Now().String(), crypto.SHA256)
	err = s.ctx.Cache().Set(s.ctx.Config().Context, nonce, user.ID.String(), s.ctx.Config().NonceLifetime).Err()
	if err != nil {
		return "", err
	}
	return nonce, err
}

func (s *authService) Create(codeChallenge string, codeChallengeMethod string) (string, error) {
	authorizationCode := helpers.HashString(time.Now().String(), crypto.SHA256)
	return authorizationCode, s.ctx.Cache().Set(
		s.ctx.Config().Context,
		authorizationCode,
		codeChallenge,
		s.ctx.Config().AuthorizationRequestLifetime,
	).Err()
}

func (s *authService) Revoke(authorizationCode string) error {
	return s.ctx.Cache().Del(s.ctx.Config().Context, authorizationCode).Err()
}

func (s *authService) CreateToken(authorizationCode string, codeVerifier string, nonce string, userIP string) (string, string, error) {
	codeChallenge := s.ctx.Cache().Get(s.ctx.Config().Context, authorizationCode).Val()
	if codeChallenge == "" {
		return "", "", messages.ErrBadAuthenticationSession
	}

	hashedCodeVerifier := helpers.HashStringToBase64(codeVerifier, crypto.SHA256)
	if hashedCodeVerifier != codeChallenge {
		err := s.Revoke(authorizationCode)
		if err != nil {
			return "", "", err
		}
		return "", "", messages.ErrBadAuthenticationSession
	}

	err := s.ctx.Cache().Del(s.ctx.Config().Context, authorizationCode).Err()
	if err != nil {
		return "", "", err
	}

	uid := s.ctx.Cache().Get(s.ctx.Config().Context, nonce).Val()
	if uid == "" {
		return "", "", messages.ErrNonceUsedOrExpired
	}

	err = s.ctx.Cache().Del(s.ctx.Config().Context, nonce).Err()
	if err != nil {
		return "", "", err
	}

	id, err := uuid.Parse(uid)
	if err != nil {
		return "", "", err
	}

	user, err := s.userService.Find(id)
	if err != nil {
		return "", "", err
	}

	nonce = helpers.HashString(user.ID.String()+time.Now().String(), crypto.SHA256)
	err = s.ctx.Cache().Set(s.ctx.Config().Context, nonce, user.ID.String(), s.ctx.Config().NonceLifetime).Err()
	if err != nil {
		return "", "", err
	}

	claims, token, err := s.newJWT(user.ID.String())
	if err != nil {
		return "", "", err
	}

	sessionName := fmt.Sprintf("%s:%s", user.ID.String(), userIP)
	err = s.ctx.Cache().Set(s.ctx.Config().Context, sessionName, claims.Id, s.ctx.Config().JWTTokenLifetime).Err()
	if err != nil {
		return "", "", err
	}

	return nonce, token, err
}

func (s *authService) VerifyToken(tokenString string) (*jwt.Token, error) {
	privateKey, err := s.getPrivateKey()
	if err != nil {
		return nil, messages.ErrLoadPrivateKey
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &privateKey.PublicKey, nil
	}
	return jwt.Parse(tokenString, keyFunc)
}

func (s *authService) Logout(nonce string, userID string, userIP string) error {
	err := s.Revoke(nonce)
	if err != nil {
		return nil
	}

	return s.Revoke(fmt.Sprintf("%s:%s", userID, userIP))
}

func (s *authService) getPrivateKey() (*ecdsa.PrivateKey, error) {
	b, err := base64.StdEncoding.DecodeString(helpers.ENVGetString("ECDSA_PRIVATE_KEY"))
	if err != nil {
		return nil, err
	}

	return x509.ParseECPrivateKey(b)
}

func (s *authService) newJWT(uid string) (*jwt.StandardClaims, string, error) {
	now := time.Now()
	claims := &jwt.StandardClaims{
		Id:        uuid.NewString(),
		Subject:   uid,
		Issuer:    fmt.Sprintf("%s/auth/token", s.ctx.Config().BackendURL),
		IssuedAt:  now.Unix(),
		NotBefore: now.Unix(),
		Audience:  s.ctx.Config().BackendURL,
		ExpiresAt: now.Add(s.ctx.Config().JWTTokenLifetime).Unix(),
	}

	privateKey, err := s.getPrivateKey()
	if err != nil {
		return nil, "", messages.ErrLoadPrivateKey
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return nil, "", err
	}

	return claims, signedToken, nil
}
