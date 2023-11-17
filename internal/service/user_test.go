package service_test

import (
	"testing"

	"github.com/nobeluc/ecommerce-api/internal/model"
	"github.com/nobeluc/ecommerce-api/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userService := service.NewUserService(nil) // Pass the actual repo here

	user := &model.User{
		// Initialize with valid values
	}

	err := userService.CreateUser(user, "password")
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	userService := service.NewUserService(nil) // Pass the actual repo here

	user := &model.User{
		// Initialize with valid values
	}

	err := userService.UpdateUser(user)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	userService := service.NewUserService(nil) // Pass the actual repo here

	user := &model.User{
		// Initialize with valid values
	}

	err := userService.DeleteUser(user)
	assert.NoError(t, err)
}

func TestFindUserByID(t *testing.T) {
	userService := service.NewUserService(nil) // Pass the actual repo here

	_, err := userService.FindUserByID(1)
	assert.NoError(t, err)
}

func TestListAllUsers(t *testing.T) {
	userService := service.NewUserService(nil) // Pass the actual repo here

	_, err := userService.ListAllUsers()
	assert.NoError(t, err)
}
