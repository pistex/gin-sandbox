package controllers

import (
	"errors"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/libraries"
	"kwanjai/messages"
	"kwanjai/models"
	"kwanjai/requests"
	"kwanjai/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	ctx interfaces.IContext
}

func NewAuthController(ctx interfaces.IContext) *authController {
	return &authController{ctx: ctx}
}

func (c *authController) Request() gin.HandlerFunc {
	return func(g *gin.Context) {
		codeChallenge := g.Query("code_challenge")
		if codeChallenge == "" {
			g.Status(http.StatusNotFound)
			return
		}
		codeChallengeMethod := g.Query("code_challenge_method")
		if codeChallengeMethod != c.ctx.Config().CodeChallengeMethod {
			g.Status(http.StatusNotFound)
			return
		}

		authService := services.NewAuthService(c.ctx)
		authorizationCode, err := authService.Create(codeChallenge, codeChallengeMethod)
		if err != nil {
			httpError := messages.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		g.JSON(http.StatusOK, gin.H{"authorizationCode": authorizationCode})
	}
}

// Login endpoint
func (c *authController) Login() gin.HandlerFunc {
	return func(g *gin.Context) {
		request := &requests.AuthLogin{}
		err := g.ShouldBindJSON(request)
		if err != nil {
			httpError := messages.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		authService := services.NewAuthService(c.ctx)
		nonce, err := authService.Login(request.Email, request.Password, request.AuthorizationCode, request.CodeVerifier)
		if errors.Is(err, messages.ErrCredentialMismatch) || errors.Is(err, messages.ErrBadAuthenticationSession) {
			httpError := messages.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		if err != nil {
			httpError := messages.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		// TODO: Set Secure cookie in production
		g.SetCookie("nonce", nonce, int(c.ctx.Config().NonceLifetime.Seconds()), "", "", false, true)
		g.Status(http.StatusNoContent)
	}
}

// Logout endpoint
func (c *authController) Logout() gin.HandlerFunc {
	return func(g *gin.Context) {
		nonce, err := g.Cookie("nonce")
		if err != nil || nonce == "" {
			g.Status(http.StatusUnauthorized)
			return
		}

		user, exist := g.Get("user")
		if !exist {
			g.Status(http.StatusInternalServerError)
			return
		}

		_, ok := user.(*models.User)
		if !ok {
			g.Status(http.StatusInternalServerError)
			return
		}

		authService := services.NewAuthService(c.ctx)
		err = authService.Logout(nonce, user.(*models.User).ID.String(), g.ClientIP())
		if err != nil {
			httpError := messages.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		g.Status(http.StatusOK)
	}
}

// Create token
func (c *authController) CreateToken() gin.HandlerFunc {
	return func(g *gin.Context) {
		nonce, err := g.Cookie("nonce")
		if err != nil || nonce == "" {
			g.Status(http.StatusUnauthorized)
			return
		}

		authorizationCode := g.Query("authorization_code")
		if authorizationCode == "" {
			g.Status(http.StatusNotFound)
			return
		}

		codeVerifier := g.Query("code_verifier")
		if codeVerifier == "" {
			g.Status(http.StatusNotFound)
			return
		}

		authService := services.NewAuthService(c.ctx)
		nonce, token, err := authService.CreateToken(authorizationCode, codeVerifier, nonce, g.ClientIP())
		if errors.Is(err, messages.ErrBadAuthenticationSession) {
			httpError := messages.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		if err != nil {
			httpError := messages.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		// TODO: Set Secure cookie in production
		g.SetCookie("nonce", nonce, int(c.ctx.Config().NonceLifetime.Seconds()), "", "", false, true)
		g.JSON(http.StatusOK, gin.H{"accessToken": token})
	}
}

type passwordUpdate struct {
	OldPassword  string `json:"old_password" binding:"required,min=8"`
	NewPassword1 string `json:"new_password1" binding:"required,min=8"`
	NewPassword2 string `json:"new_password2" binding:"required,min=8"`
}

// PasswordUpdate endpoint
func PasswordUpdate(ginContext *gin.Context) {
	passwordForm := new(passwordUpdate)
	if err := ginContext.ShouldBindJSON(passwordForm); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	if passwordForm.NewPassword1 != passwordForm.NewPassword2 {
		ginContext.JSON(http.StatusBadRequest, gin.H{"message": "Password confrimation failed."})
	}
	username := helpers.GetUsername(ginContext)
	newPassword, _ := libraries.HashPassword(passwordForm.NewPassword1)
	libraries.FirestoreUpdateField("users", username, "HashedPassword", newPassword)
	ginContext.JSON(http.StatusOK, gin.H{"message": "Password updated."})
}
