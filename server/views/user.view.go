package views

import (
	"kwanjai/models"
	"time"

	"github.com/jinzhu/copier"
)

type user struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}

func NewUserView(model *models.User) *user {
	view := &user{}
	_ = copier.Copy(view, model)
	return view
}
