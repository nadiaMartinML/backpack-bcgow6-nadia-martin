package users

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	// arrage
	users := []User{
		{
			Id:         1,
			Name:       "Martina",
			LastName:   "Guillen",
			Email:      "mg@mail.com",
			Age:        30,
			Height:     1.65,
			Active:     true,
			DateCreate: "2022-09-19",
		},
	}

	myMockStore := MockStore{
		users: users,
	}

	repository := NewRepository(&myMockStore)
	service := NewService(repository)

	// act
	userUpdated := User{
		Id:         1,
		Name:       "Norma",
		LastName:   "Vazquez",
		Email:      "nm@vazquez.com",
		Age:        60,
		Height:     1.70,
		Active:     true,
		DateCreate: "2022-10-21",
	}
	userResult, err := service.UpdateUser(1, "Norma", "Vazquez", "nm@vazquez.com", 60, 1.70, true, "2022-10-21")

	// assert
	assert.Nil(t, err)
	assert.Equal(t, userResult, userUpdated)
	assert.True(t, myMockStore.readCall)
}

func TestUpdateUserFail(t *testing.T) {
	// arrage
	users := []User{
		{
			Id:         1,
			Name:       "Martina",
			LastName:   "Guillen",
			Email:      "mg@mail.com",
			Age:        30,
			Height:     1.65,
			Active:     true,
			DateCreate: "2022-09-19",
		},
	}

	myMockStore := MockStore{
		users: users,
	}

	repository := NewRepository(&myMockStore)
	service := NewService(repository)

	// act
	_, err := service.UpdateUser(4, "Norma", "Vazquez", "nm@vazquez.com", 60, 1.70, true, "2022-10-21")

	// assert
	assert.NotNil(t, err)
	assert.True(t, myMockStore.readCall)
}

func TestDeleteUser(t *testing.T) {
	//arrage
	users := []User{
		{
			Id:         1,
			Name:       "Martina",
			LastName:   "Guillen",
			Email:      "mg@mail.com",
			Age:        30,
			Height:     1.65,
			Active:     true,
			DateCreate: "2022-09-19",
		},
		{
			Id:         2,
			Name:       "Martina",
			LastName:   "Guillen",
			Email:      "mg@mail.com",
			Age:        30,
			Height:     1.65,
			Active:     true,
			DateCreate: "2022-09-19",
		},
	}

	myMockStore := MockStore{
		users: users,
	}

	repository := NewRepository(&myMockStore)
	service := NewService(repository)

	//act
	errUserExist := service.DeleteUser(1)
	usersAfterDelete, _ := service.GetAllUsers()
	errUserNotExist := service.DeleteUser(5)
	expectError := fmt.Errorf("Usuario %d no fue encontrado", 5)

	//assert
	assert.Nil(t, errUserExist)
	assert.True(t, myMockStore.readCall)
	assert.Equal(t, len(usersAfterDelete), len(users)-1)
	assert.EqualErrorf(t, errUserNotExist, expectError.Error(), "")
}
