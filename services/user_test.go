package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wsaefulloh/go-solid-principle/models"
	"github.com/wsaefulloh/go-solid-principle/repos"
)

var userRepo = repos.UserRepMock{Mocks: mock.Mock{}}
var userService = user_service{rep: &userRepo}

func TestUser(t *testing.T) {
	t.Run("Get All User", func(t *testing.T) {
		userRepo.Mocks.On("FindAll")
		data, err := userService.FindAll()
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Save User", func(t *testing.T) {
		prod := models.User{
			Name:      "testing",
			Id:        1,
			UserName:  "usertest",
			Email:     "eemail@gmail.com",
			Role:      "admin",
			Password:  "abcd1234",
			CreatedAt: time.Now(),
			UpdateAt:  time.Now(),
		}
		userRepo.Mocks.On("Save", &prod)
		data, err := userService.AddData(&prod)
		res := data.Status
		assert.Equal(t, 200, res, "Status must be 200")
		assert.Nil(t, err)
	})

	t.Run("Get Pass User by Username", func(t *testing.T) {
		userRepo.Mocks.On("GetPass", "usertest")
		data, err := userService.GetPass("usertest")
		assert.Equal(t, "abcd1234", data, "Password must be 'abcd1234'")
		assert.Nil(t, err)
	})

	t.Run("Get Email by Username", func(t *testing.T) {
		userRepo.Mocks.On("GetEmail", "usertest")
		data, err := userService.GetEmail("usertest")
		assert.Equal(t, "eemail@gmail.com", data, "Email must be 'eemail@gmail.com'")
		assert.Nil(t, err)
	})

	t.Run("Get Role by Username", func(t *testing.T) {
		userRepo.Mocks.On("GetRole", "usertest")
		data, err := userService.GetRole("usertest")
		assert.Equal(t, "admin", data, "Result must be 'admin'")
		assert.Nil(t, err)
	})

	t.Run("Check Already Username", func(t *testing.T) {
		userRepo.Mocks.On("GetUsername", "usertest")
		data, err := userService.GetUsername("usertest")
		assert.Equal(t, "usertest", data, "Username must be 'usertest'")
		assert.Nil(t, err)
	})
}

//TEST TABLE

type testTable struct {
	name    string
	params  string
	returns string
}

func TestGetUsernameTable(t *testing.T) {
	var test = []testTable{
		{
			name:    "test1",
			params:  "wahyu",
			returns: "usertest wahyu",
		},
		{
			name:    "test2",
			params:  "wahyu2",
			returns: "usertest wahyu2",
		},
	}

	for _, val := range test {
		t.Run(val.name, func(t *testing.T) {
			userRepo.Mocks.On("GetUsername", "usertest")
			data, err := userService.GetUsername("usertest")
			result := data + " " + val.params
			assert.Equal(t, val.returns, result)
			assert.Nil(t, err)
		})
	}
}
