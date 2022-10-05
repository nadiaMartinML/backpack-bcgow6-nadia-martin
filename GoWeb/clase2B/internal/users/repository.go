package users

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
