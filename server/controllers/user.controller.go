package controllers

import (
	"errors"
	"kwanjai/consts"
	"kwanjai/helpers"
	"kwanjai/interfaces"
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
			g.JSON(helpers.NewBadRequestError(err).GetStatus(), helpers.NewBadRequestError(err).GetJSON())
			return
		}

		userSerice := services.NewUserService(c.ctx)
		user, err := userSerice.Create(request.Email, request.Password)
		if errors.Is(err, consts.DuplicatedEmail) {
			g.JSON(helpers.NewBadRequestError(err).GetStatus(), helpers.NewBadRequestError(err).GetJSON())
			return
		}

		if err != nil {
			g.JSON(helpers.NewInternalServerError(err).GetStatus(), helpers.NewInternalServerError(err).GetJSON())
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
			g.JSON(helpers.NewBadRequestError(err).GetStatus(), helpers.NewBadRequestError(err).GetJSON())
			return
		}

		userSerice := services.NewUserService(ctx)
		err = userSerice.ChangePassword(userID, request.Password, request.NewPassword)
		if err == consts.CredentialMismatch {
			g.JSON(helpers.NewBadRequestError(err).GetStatus(), helpers.NewBadRequestError(err).GetJSON())
			return
		}
		if err != nil {
			g.JSON(helpers.NewInternalServerError(err).GetStatus(), helpers.NewInternalServerError(err).GetJSON())
			return
		}

		g.Status(http.StatusOK)
	}
}
