package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	users    []User
	readCall bool
}

func (m *MockStore) ReadFile(data interface{}) error {
	m.readCall = true
	usersData := data.(*[]User)
	*usersData = m.users
	return nil
}

func (m MockStore) WriteFile(data interface{}) error {
	return nil
}

func TestUpdatePatch(t *testing.T) {
	// arrage
	users := []User{
		{
			Id:       1,
			Name:     "Martina",
			LastName: "Guillen",
			Age:      30,
		},
		{
			Id:       2,
			Name:     "Yamila",
			LastName: "Sacchi",
			Age:      32,
		},
	}

	myMockStore := MockStore{
		users: users,
	}

	motor := NewRepository(&myMockStore)

	// act
	userResult, _ := motor.UpdatePatch(myMockStore.users[0].Id, "Vazquez", 28)

	// assert
	assert.Equal(t, userResult.LastName, "Vazquez")
	assert.Equal(t, userResult.Age, 28)
	assert.True(t, myMockStore.readCall)
}
