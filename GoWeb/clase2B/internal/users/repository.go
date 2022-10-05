package users

import "fmt"

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
var lastId int

type Repository interface {
	GetAllUsers() ([]User, error)
	GetOneUser()
	SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	LastId() (int, error)
	UpdateUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error)
	DeleteUser(Id int) error
	UpdatePatch(Id int, LastName string, Age int) (User, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// LastId implements Repository
func (r *repository) LastId() (int, error) {
	return lastId, nil
}

// getAllUsers implements Repository
func (r *repository) GetAllUsers() ([]User, error) {
	return users, nil
}

// getOneUser implements Repository
func (r *repository) GetOneUser() {
}

// saveUser implements Repository
func (r *repository) SaveUser(Id int, Name, LastName, Email string, Age int, Height float64, Active bool, DateCreate string) (User, error) {
	u := User{Id, Name, LastName, Email, Age, Height, Active, DateCreate}
	users = append(users, u)
	lastId = u.Id
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
