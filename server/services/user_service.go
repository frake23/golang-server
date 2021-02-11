package services

type IUserService interface {
	GetUsers() (*Users, error)
	GetUser(id string) (*User, error)
	AddUser(id string, name string) error
	DeleteUser(id string) error
}

type UserService struct {
	db *JsonDb
}

func (service UserService) GetUsers() (*Users, error) {
	return nil, nil
}

func (service UserService) GetUser(id string) (*User, error) {
	return nil, nil
}

func (service UserService) AddUser(id string, name string) error {
	return nil
}

func (service UserService) DeleteUser(id string) error {
	return nil
}
