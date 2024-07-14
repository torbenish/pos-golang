package database

import "github.com/torbenish/pos-golang/07-APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
