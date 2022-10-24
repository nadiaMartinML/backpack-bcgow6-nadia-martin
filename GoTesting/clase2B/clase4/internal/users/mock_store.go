package users

type MockStore struct {
	users     []User
	readCall  bool
	writeCall bool
}

func (m *MockStore) ReadFile(data interface{}) error {
	m.readCall = true
	usersData := data.(*[]User)
	*usersData = m.users
	return nil
}

func (m *MockStore) WriteFile(data interface{}) error {
	m.writeCall = true
	usersData := data.([]User)
	m.users = usersData
	return nil
}
