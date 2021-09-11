package routes

import (
	"kwanjai/controllers"
	"kwanjai/interfaces"
)

func UseUserRouter(ctx interfaces.IContext) {
	userController := controllers.NewUserController(ctx)
	user := ctx.Server().Group("/user")
	user.POST("/", userController.Create())
	user.POST("/:user_id/change-password/:email_token", userController.ChangePassword(ctx))
}
