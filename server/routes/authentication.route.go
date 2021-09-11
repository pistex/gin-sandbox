package routes

import (
	"kwanjai/controllers"
	"kwanjai/interfaces"
	"kwanjai/middlewares"
)

func UseAuthentionRouter(ctx interfaces.IContext) {
	authenticationController := controllers.NewAuthenticationController(ctx)
	authentication := ctx.Server().Group("/authentication")
	authentication.POST("/login", authenticationController.Login())
	authentication.POST("/logout", middlewares.AuthenticatedOnly(ctx), authenticationController.Logout())
	authentication.POST("/verify_email/:ID", controllers.VerifyEmail(ctx))
	authentication.POST("/resend_verification_email", controllers.ResendVerifyEmail(ctx))
	authentication.POST("/token/refresh", controllers.RefreshToken(ctx))
	authentication.GET("/token/verify", middlewares.AuthenticatedOnly(ctx), controllers.TokenVerification(ctx))
}
