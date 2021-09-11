package services

import (
	"database/sql"
	"errors"
	"kwanjai/consts"
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type IUserService interface {
	Create(email string, password string) (*models.User, error)
	Find(id uuid.UUID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user models.User) (*models.User, error)
	ChangePassword(id uuid.UUID, password string, newPassword string) error
}

type userService struct {
	ctx interfaces.IContext
}

func NewUserService(ctx interfaces.IContext) IUserService {
	return &userService{
		ctx: ctx,
	}
}

func (s *userService) Create(email string, password string) (*models.User, error) {
	_, err := s.FindByEmail(email)
	if !errors.Is(sql.ErrNoRows, err) {
		return nil, consts.DuplicatedEmail
	}

	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		return nil, err
	}

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.NewUser(email, hashedPassword)
	query := `INSERT INTO users (id, email, password, is_active) VALUES (:id, :email, :password, :is_active)`
	_, err = s.ctx.DB().NamedExec(query, user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (s *userService) Find(id uuid.UUID) (*models.User, error) {
	panic("implement me")
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := s.ctx.DB().Get(user, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(user models.User) (*models.User, error) {
	panic("implement me")
}

func (s *userService) ChangePassword(id uuid.UUID, password string, newPassword string) error {
	panic("implement me")
}
