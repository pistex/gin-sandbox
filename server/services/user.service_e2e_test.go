package services

// Basic imports
import (
	"kwanjai/helpers"
	"kwanjai/interfaces"
	"kwanjai/messages"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	suite.SetupAllSuite
	suite.SetupTestSuite
	ctx     interfaces.IContext
	service IUserService
}

func (s *UserServiceTestSuite) SetupSuite() {
	err := helpers.LoadENV("../.env")
	s.NoError(err)
	db, err := helpers.NewTestDatabase()
	s.NotEmpty(db)
	s.NoError(err)
	s.ctx = interfaces.NewContext(nil, nil, db, nil)
	s.service = NewUserService(s.ctx)
	s.NoError(err)
}

func (s *UserServiceTestSuite) SetupTest() {
	_, err := s.ctx.DB().Exec(`DELETE FROM users WHERE 1=1`)
	s.NoError(err)
}

func (suite *UserServiceTestSuite) TestCreateUserSuccess() {
	user, err := suite.service.Create("john@example.com", "johnpassword")
	suite.NotEmpty(user)
	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestCreateUserDuplicatedEmail() {
	user, err := suite.service.Create("john@example.com", "johnpassword")
	suite.NotEmpty(user)
	suite.NoError(err)

	user, err = suite.service.Create("john@example.com", "johnpassword")
	suite.Empty(user)
	suite.ErrorIs(messages.ErrDuplicatedEmail, err)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
