package routes

import (
	"kwanjai/controllers"
	"kwanjai/interfaces"
	"kwanjai/middlewares"
)

func UseAuthentionRouter(ctx interfaces.IContext) {
	authentication := ctx.GetServer().Group("/authentication")
	authentication.POST("/login", controllers.Login(ctx))
	authentication.POST("/register", controllers.Register(ctx))
	authentication.POST("/logout", middlewares.AuthenticatedOnly(ctx), controllers.Logout(ctx))
	authentication.POST("/verify_email/:ID", controllers.VerifyEmail(ctx))
	authentication.POST("/resend_verification_email", controllers.ResendVerifyEmail(ctx))
	authentication.POST("/token/refresh", controllers.RefreshToken(ctx))
	authentication.GET("/token/verify", middlewares.AuthenticatedOnly(ctx), controllers.TokenVerification(ctx))
}
