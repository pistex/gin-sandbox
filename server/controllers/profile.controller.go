package controllers

import (
	"kwanjai/interfaces"
	"kwanjai/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProfilePicture endpoint
func ProfilePicture(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		panic("implement me")
	}
}

// MyProfile endpoint
func MyProfile() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		user, _ := ginContext.Get("user") // user always exists
		userObject := user.(*models.User)

		ginContext.JSON(http.StatusOK, gin.H{
			"message": "Get profile successfully",
			"profile": userObject,
		})
	}
}

// UpdateProfile endpoint
func UpdateProfile() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		panic("implement me")
	}
}

type paymentData struct {
	Token string `json:"token"`
	Price int    `json:"price"`
}

// UpgradePlan endpoint
func UpgradePlan(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		panic("implement me")
	}
}

// Unsubscribe endpoint
func Unsubscribe(ctx interfaces.IContext) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		panic("implement me")
	}
}
