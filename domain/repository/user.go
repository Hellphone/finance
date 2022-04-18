package repository

import (
	"github.com/hellphone/finance/domain/model"
)

type User interface {
	Store(user *model.User) (*model.User, error)
	GetOneById(id int) (*model.User, error)
}
