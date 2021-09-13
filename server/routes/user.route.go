package routes

import (
	"kwanjai/controllers"
	"kwanjai/interfaces"
	"kwanjai/middlewares"
)

func UseUserRouter(ctx interfaces.IContext) {
	userController := controllers.NewUserController(ctx)
	user := ctx.Server().Group("/user")
	user.POST("/", userController.Create())
	user.GET("/me", middlewares.JWT(ctx), userController.Profile())
}
