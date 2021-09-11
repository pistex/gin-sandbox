package views

import (
	"kwanjai/models"

	"github.com/jinzhu/copier"
)

type user struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func NewUserView(model *models.User) *user {
	view := &user{}
	_ = copier.Copy(view, model)
	return view
}
