package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"kwanjai/controllers"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/middlewares"
	"kwanjai/routes"
	"kwanjai/types"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Server(ctx interfaces.IContext) *gin.Engine {

	ginEngine := gin.Default()

	user := ginEngine.Group("/user")
	user.Use(middlewares.AuthenticatedOnly(ctx))
	user.GET("/all", controllers.AllUsernames(ctx))
	user.GET("/my_profile", controllers.MyProfile())
	user.POST("/update_password", controllers.PasswordUpdate)
	user.PATCH("/profile_picture", controllers.ProfilePicture(ctx))
	user.PATCH("/update_profile", controllers.UpdateProfile())
	user.POST("/pay", controllers.UpgradePlan(ctx))
	user.POST("/unsubscribe", controllers.Unsubscribe(ctx))
	project := ginEngine.Group("/project")
	project.Use(middlewares.AuthenticatedOnly(ctx))
	{
		project.GET("/all", controllers.AllProject())
		project.POST("/new", controllers.NewProject())
		project.POST("/find", controllers.FindProject())
		project.PATCH("/update", controllers.UpdateProject())
		project.DELETE("/delete", controllers.DeleteProject(ctx))
	}
	board := ginEngine.Group("/board")
	board.Use(middlewares.AuthenticatedOnly(ctx))
	{
		board.POST("/all", controllers.AllBoard(ctx))
		board.POST("/new", controllers.NewBoard(ctx))
		board.POST("/find", controllers.FindBoard(ctx))
		board.PATCH("/update", controllers.UpdateBoard(ctx))
		board.DELETE("/delete", controllers.DeleteBoard(ctx))
	}
	post := ginEngine.Group("/post")
	post.Use(middlewares.AuthenticatedOnly(ctx))
	{
		post.POST("/all", controllers.AllPost())
		post.POST("/new", controllers.NewPost())
		post.PATCH("/update", controllers.UpdatePost())
		post.DELETE("/delete", controllers.DeletePost())
		post.POST("/comment/new", controllers.NewComment())
		post.PATCH("/comment/update", controllers.UpdateComment())
		post.DELETE("/comment/delete", controllers.DeleteComment())
	}
	return ginEngine
}

func main() {

	err := helpers.LoadENV()
	helpers.CheckErrorAndPanic(err)

	db, err := helpers.NewDatabase()
	helpers.CheckErrorAndPanic(err)
	defer db.Close()

	ctx := interfaces.NewContext(&types.Config{Port: viper.GetString("PORT")}, gin.Default(), db)

	routes.UseUserRouter(ctx)

	helpers.CheckErrorAndPanic(ctx.Server().Run(fmt.Sprintf(":%s", ctx.Config().Port)))
}
