package repository

import "github.com/hellphone/finance/domain/repository"

type Factory interface {
	UserRepository() repository.User
}
