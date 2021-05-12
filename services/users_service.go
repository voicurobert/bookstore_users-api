package services

import (
	"github.com/voicurobert/bookstore_users-api/domain/users"
	"github.com/voicurobert/bookstore_users-api/utils/date_utils"
	"github.com/voicurobert/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{ID: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date_utils.GetNowAsDBFormat()
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(user users.User, isPartial bool) (*users.User, *errors.RestError) {
	current, err := GetUser(user.ID)
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

func DeleteUser(userId int64) *errors.RestError {
	current := &users.User{ID: userId}
	return current.Delete()
}

func Search(status string) ([]users.User, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
