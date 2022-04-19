package handlers

import (
	"github.com/hellphone/finance/domain/cases/transfer_money"
	"io/ioutil"
	"net/http"

	"github.com/hellphone/finance/api"
	"github.com/hellphone/finance/domain/cases/add_money_to_user"
	"github.com/hellphone/finance/domain/model"
	"github.com/hellphone/finance/presenter/jsonapi"
)

func AddMoneyToUser(ctx *api.Context, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	rb, err := jsonapi.UnmarshalAddMoneyToUser(body)
	if err != nil {
		return nil, err
	}

	request := &add_money_to_user.Request{
		User: &model.User{
			Id:          rb.UserId,
			MoneyAmount: rb.MoneyAmount,
		},
		UserRepository: ctx.Repositories.UserRepository(),
	}

	resp, err := add_money_to_user.Run(request)
	if err != nil {
		return nil, err
	}

	return jsonapi.MarshalUser(resp.User)
}

func TransferMoney(ctx *api.Context, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	rb, err := jsonapi.UnmarshalTransferMoney(body)
	if err != nil {
		return nil, err
	}

	request := &transfer_money.Request{
		UserFromId:     rb.UserFromId,
		UserToId:       rb.UserToId,
		MoneyAmount:    rb.MoneyAmount,
		UserRepository: ctx.Repositories.UserRepository(),
	}

	resp, err := transfer_money.Run(request)
	if err != nil {
		return nil, err
	}

	return jsonapi.MarshalUsers([]*model.User{
		resp.UserTo,
		resp.UserFrom,
	})
}
