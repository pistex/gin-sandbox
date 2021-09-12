package controllers

import (
	"errors"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/messages"
	"kwanjai/requests"
	"kwanjai/services"
	"kwanjai/views"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController(ctx interfaces.IContext) *userController {
	return &userController{ctx: ctx}
}

type userController struct {
	ctx interfaces.IContext
}

func (c *userController) Create() gin.HandlerFunc {
	return func(g *gin.Context) {
		request := &requests.UserCreate{}
		err := g.ShouldBindJSON(request)
		if err != nil {
			httpError := helpers.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		userSerice := services.NewUserService(c.ctx)
		user, err := userSerice.Create(request.Email, request.Password)
		if errors.Is(err, messages.ErrDuplicatedEmail) {
			httpError := helpers.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		if err != nil {
			httpError := helpers.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		g.JSON(http.StatusCreated, views.NewUserView(user))
	}
}

func (c *userController) ChangePassword(ctx interfaces.IContext) gin.HandlerFunc {
	return func(g *gin.Context) {
		userID, err := uuid.Parse(g.Param("user_id"))
		if err != nil {
			g.Status(http.StatusNotFound)
			return
		}

		request := &requests.UserChangePassword{}
		err = g.ShouldBindJSON(request)
		if err != nil {
			httpError := helpers.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		userSerice := services.NewUserService(ctx)
		err = userSerice.ChangePassword(userID, request.Password, request.NewPassword)
		if err == messages.ErrCredentialMismatch {
			httpError := helpers.NewBadRequestError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		if err != nil {
			httpError := helpers.NewInternalServerError(err)
			g.JSON(httpError.GetStatus(), httpError.GetJSON())
			return
		}

		g.Status(http.StatusOK)
	}
}
