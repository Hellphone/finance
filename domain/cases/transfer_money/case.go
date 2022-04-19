package transfer_money

import (
	"errors"
	"github.com/hellphone/finance/domain/model"
	"github.com/hellphone/finance/domain/repository"
)

type Request struct {
	UserFromId     int
	UserToId       int
	MoneyAmount    float64
	UserRepository repository.User
}

type Response struct {
	UserFrom *model.User
	UserTo   *model.User
}

func (r *Request) validate() error {
	if r == nil {
		return errors.New("transfer_money Request should not be nil")
	}

	if r.UserFromId == 0 {
		return errors.New("UserFromId should not be zero")
	}

	if r.UserToId == 0 {
		return errors.New("UserToId should not be zero")
	}

	if r.UserToId == r.UserFromId {
		return errors.New("user IDs should not be equal")
	}

	if r.MoneyAmount <= 0 {
		return errors.New("money_amount should be greater than zero")
	}

	return nil
}

func Run(r *Request) (*Response, error) {
	if err := r.validate(); err != nil {
		return nil, err
	}

	userFrom, err := r.UserRepository.GetOneById(r.UserFromId)
	if err != nil {
		return nil, err
	}

	userTo, err := r.UserRepository.GetOneById(r.UserToId)
	if err != nil {
		return nil, err
	}

	moneyFromAmount := userFrom.MoneyAmount - r.MoneyAmount
	if moneyFromAmount < 0 {
		return nil, errors.New("not enough money to transfer")
	}

	moneyToAmount := userTo.MoneyAmount + r.MoneyAmount

	resp := &Response{
		UserFrom: &model.User{
			Id:          userFrom.Id,
			Name:        userFrom.Name,
			MoneyAmount: moneyFromAmount,
		},
		UserTo: &model.User{
			Id:          userTo.Id,
			Name:        userTo.Name,
			MoneyAmount: moneyToAmount,
		},
	}

	_, err = r.UserRepository.Store(resp.UserFrom)
	if err != nil {
		return nil, err
	}

	_, err = r.UserRepository.Store(resp.UserTo)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
