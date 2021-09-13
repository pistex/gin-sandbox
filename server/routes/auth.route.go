package routes

import (
	"kwanjai/controllers"
	"kwanjai/interfaces"
	"kwanjai/middlewares"
)

func UseAuthRouter(ctx interfaces.IContext) {
	authController := controllers.NewAuthController(ctx)
	auth := ctx.Server().Group("/auth")
	auth.GET("", authController.Request())
	auth.POST("/login", authController.Login())
	auth.POST("/logout", middlewares.JWT(ctx), authController.Logout())
	auth.GET("/token", authController.CreateToken())
}
