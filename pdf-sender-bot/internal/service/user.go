package service

import (
	"reflect"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/internal/repository"
)

type UserService interface {
	GetUserProfile(id int64) (*dto.User, error)
	SaveUserProfile(id int64, data *dto.User) error
	DidUserRegistrated(id int64) (bool, error)
}

type userService struct {
	dao           repository.DAO
	UsersStatuses map[int64]int
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{
		dao:           dao,
		UsersStatuses: map[int64]int{},
	}
}

func (u *userService) GetUserProfile(id int64) (*dto.User, error) {
	user, err := u.dao.NewUserQuery().GetUserData(id)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (u *userService) SaveUserProfile(id int64, data *dto.User) error {
	hasRegistration, err := u.DidUserRegistrated(id)
	if err != nil {
		return err
	}

	if hasRegistration {
		err := u.dao.NewUserQuery().UpdateUserData(data)
		if err != nil {
			return err
		}
	} else {
		err := u.dao.NewUserQuery().InsertUserData(data)
		if err != nil {
			return err
		}
	}
}
func (u *userService) DidUserRegistrated(id int64) (bool, error) {
	data, err := u.GetUserProfile(id)
	if err != nil {
		return false, err
	}

	if *data == (dto.User{}) {
		return false, err
	}

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return false, nil
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			return false, nil
		}
	}

	return true, nil

}
