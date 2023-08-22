package service

import "github.com/hfs1988/jagaad_test/entity"

type UserService interface {
	FetchUsers() ([]entity.User, error)
	WriteUsers(users []entity.User) error
	GetUsers(tags []string) ([]entity.User, error)
}
