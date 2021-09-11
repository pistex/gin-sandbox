package controllers

import (
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/libraries"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authenticationController struct {
	ctx interfaces.IContext
}

func NewAuthenticationController(ctx interfaces.IContext) *authenticationController {
	return &authenticationController{ctx: ctx}
}

// Login endpoint
func (c *authenticationController) Login() gin.HandlerFunc {
	panic("implement me")
}

// Logout endpoint
func (c *authenticationController) Logout() gin.HandlerFunc {
	panic("implement me")
}

// RefreshToken endpiont
func RefreshToken(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		token := new(libraries.Token)
		ginContext.ShouldBind(token)
		extractedToken := strings.Split(ginContext.Request.Header.Get("Authorization"), "Bearer ")
		if len(extractedToken) != 2 {
			token.AccessToken = ""
		} else {
			token.AccessToken = extractedToken[1]
		}
		if token.RefreshToken == "" {
			ginContext.JSON(http.StatusBadRequest, gin.H{"message": "No refresh token provied."})
			return
		}
		_, refreshUsername, _, err := libraries.VerifyToken(token.RefreshToken, "refresh")
		if err != nil {
			if refreshUsername == "anonymous" {
				ginContext.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
				return
			}
			log.Panicln(err)
		}
		_, accessUsername, tokenID, err := libraries.VerifyToken(token.AccessToken, "access") // if token is expried here, it's got delete.
		if accessUsername != "anonymous" && err == nil {                                      // user != "anonymous" means token is still valid.
			_, err = libraries.FirestoreDelete("tokens", tokenID)
			if err != nil {
				log.Panicln(err)
			}
		}
		newToken, err := libraries.CreateToken("access", refreshUsername)
		token.AccessToken = newToken
		ginContext.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// TokenVerification endpiont
func TokenVerification(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.Status(http.StatusOK)
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
