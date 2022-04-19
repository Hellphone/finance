package add_money_to_user

import (
	"errors"

	"github.com/hellphone/finance/domain/model"
	"github.com/hellphone/finance/domain/repository"
)

type Request struct {
	User           *model.User
	UserRepository repository.User
}

type Response struct {
	User *model.User
}

func (r *Request) validate() error {
	if r == nil {
		return errors.New("add_money_to_user Request should not be nil")
	}

	if r.User == nil {
		return errors.New("User should not be nil")
	}

	return nil
}

func Run(r *Request) (*Response, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}

	user, err := r.UserRepository.GetOneById(r.User.Id)
	if err != nil {
		return nil, err
	}

	moneyAmount := user.MoneyAmount + r.User.MoneyAmount
	if moneyAmount < 0 {
		return nil, errors.New("money amount cannot be negative")
	}

	resp := &Response{
		&model.User{
			Id:          user.Id,
			Name:        user.Name,
			MoneyAmount: moneyAmount,
		},
	}

	user, err = r.UserRepository.Store(resp.User)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
