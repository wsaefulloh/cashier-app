package repos

import (
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
)

type UserRepMock struct {
	Mocks mock.Mock
}

var dataUser = models.User{
	Name:      "testing",
	Id:        1,
	UserName:  "usertest",
	Email:     "eemail@gmail.com",
	Role:      "admin",
	Password:  "abcd1234",
	CreatedAt: time.Now(),
	UpdateAt:  time.Now(),
}

var dataUsers = models.Users{
	dataUser,
}

func (pr *UserRepMock) FindAll() (*models.Users, error) {
	args := pr.Mocks.Called()

	if args == nil {
		return nil, nil
	}

	return &dataUsers, nil
}

func (pr *UserRepMock) Save(prod *models.User) error {
	args := pr.Mocks.Called(prod)

	if args == nil {
		return nil
	}

	return nil
}

func (pr *UserRepMock) GetPass(username string) (string, error) {
	args := pr.Mocks.Called(username)

	if args == nil {
		return "", nil
	}

	return dataUser.Password, nil
}

func (pr *UserRepMock) GetEmail(username string) (string, error) {
	args := pr.Mocks.Called(username)

	if args == nil {
		return "", nil
	}

	return dataUser.Email, nil
}

func (pr *UserRepMock) GetUsername(username string) (string, error) {
	args := pr.Mocks.Called(username)

	if args == nil {
		return "", nil
	}

	return dataUser.UserName, nil
}

func (pr *UserRepMock) GetRole(username string) (string, error) {
	args := pr.Mocks.Called(username)

	if args == nil {
		return "", nil
	}

	return dataUser.Role, nil
}
