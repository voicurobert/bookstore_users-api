package services

import (
	"github.com/voicurobert/bookstore_users-api/domain/users"
	"github.com/voicurobert/bookstore_users-api/utils/crypto_utils"
	"github.com/voicurobert/bookstore_users-api/utils/date_utils"
	"github.com/voicurobert/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestError)
	CreateUser(users.User) (*users.User, *errors.RestError)
	UpdateUser(users.User, bool) (*users.User, *errors.RestError)
	DeleteUser(int64) *errors.RestError
	Search(string) (users.Users, *errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestError)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{ID: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = crypto_utils.GetMd5(user.Password)
	user.DateCreated = date_utils.GetNowAsDBFormat()
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {
	current, err := s.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) *errors.RestError {
	current := &users.User{ID: userId}
	return current.Delete()
}

func (s *usersService) Search(status string) (users.Users, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	user := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
