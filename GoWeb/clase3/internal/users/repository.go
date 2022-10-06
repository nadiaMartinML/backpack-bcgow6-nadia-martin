package users

import (
	"fmt"

	"github.com/nadiaMartinML/backpack-bcgow6-nadia-martin/GoWeb/clase3/pkg/store"
)

type User struct {
	Id         int     `json:"Id"`
	Name       string  `form:"Name" json:"Name"`
	LastName   string  `form:"LastName" json:"LastName"`
	Email      string  `form:"Email" json:"Email"`
	Age        int     `form:"Age" json:"Age"`
	Height     float64 `form:"Height" json:"Height"`
	Active     bool    `form:"Active" json:"Active"`
	DateCreate string  `form:"DateCreate" json:"DateCreate"`
}

var users []User

type Repository interface {
	GetAllUsers() ([]User, error)
	GetOneUser()
	SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	LastId() (int, error)
	UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	DeleteUser(Id int) error
	UpdatePatch(Id int, LastName string, Age int) (User, error)
}

type repository struct {
	dataBase store.Store
}

func NewRepository(dataBase store.Store) Repository {
	return &repository{
		dataBase: dataBase,
	}
}

// LastId implements Repository
func (r *repository) LastId() (int, error) {
	var usersData []User
	if err := r.dataBase.ReadFile(&usersData); err != nil {
		return 0, err
	}
	if len(usersData) == 0 {
		return 0, nil
	}

	return usersData[len(usersData)-1].Id, nil
}

// getAllUsers implements Repository
func (r *repository) GetAllUsers() ([]User, error) {
	var usersData []User
	r.dataBase.ReadFile(&usersData)
	return usersData, nil
}

// getOneUser implements Repository
func (r *repository) GetOneUser() {
}

// saveUser implements Repository
func (r *repository) SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	var usersData []User
	r.dataBase.ReadFile(&usersData)

	u := User{Id, Name, LastName, Email, Age, Height, Active, DateCreate}
	usersData = append(usersData, u)

	if err := r.dataBase.WriteFile(usersData); err != nil {
		return User{}, err
	}

	return u, nil
}

// UpdateUser implements Repository
func (r *repository) UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	u := User{Name: Name, LastName: LastName, Email: Email, Age: Age, Height: Height, Active: Active, DateCreate: DateCreate}
	updated := false
	for i := range users {
		if users[i].Id == Id {
			u.Id = Id
			users[i] = u
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", Id)
	}
	return u, nil
}

// DeleteUser implements Repository
func (r *repository) DeleteUser(Id int) error {
	deleted := false
	var index int

	for i := range users {
		if users[i].Id == Id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Usuario %d no encontrado", Id)
	}

	users = append(users[:index], users[index+1:]...)
	return nil
}

// Update implements Repository
func (r *repository) UpdatePatch(Id int, LastName string, Age int) (User, error) {
	var u User
	updated := false

	for i := range users {
		if users[i].Id == Id {
			users[i].LastName = LastName
			users[i].Age = Age
			updated = true
			u = users[i]
		}
	}

	if !updated {
		return User{}, fmt.Errorf("Usuario %d no encontrado", Id)
	}

	return u, nil
}
