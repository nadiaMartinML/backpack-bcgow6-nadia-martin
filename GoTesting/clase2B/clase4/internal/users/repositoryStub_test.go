package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s StubStore) ReadFile(data interface{}) error {
	users := []User{
		{
			Id:       1,
			Name:     "Martina",
			LastName: "Guillen",
		},
		{
			Id:       2,
			Name:     "Yamila",
			LastName: "Sacchi",
		},
	}
	// como recibe una interfaz vacía, necesitamos declarar una para poder incluirle los datos que queremos, esto se hace con el
	// data. >> el data es el nombre que viene del parametro y luego como en el repository "la data" estan con un puntero, aca necesitamos
	// referenciarla tambien como puntero de un puntero, por eso el *
	usersData := data.(*[]User)
	// cuando ya esta declarada la variable utilizamos el append para incluir los datos, siempre utilizando los punteros
	*usersData = append(*usersData, users...)
	// retornamos nil porque no estamos testeando si ocurre un error
	return nil
}

// esta funcion la dejamos con return nil porque no la necesitamos, pero como estamos llamando al repository necesitamos que todos
// los metodos esten
func (s StubStore) WriteFile(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange

	myStubStore := StubStore{}
	motor := NewRepository(myStubStore)
	usersExpect := []User{
		{
			Id:       1,
			Name:     "Martina",
			LastName: "Guillen",
		},
		{
			Id:       2,
			Name:     "Yamila",
			LastName: "Sacchi",
		},
	}

	//act
	usersResult, _ := motor.GetAllUsers()

	//assert
	assert.Equal(t, usersExpect, usersResult)
}
