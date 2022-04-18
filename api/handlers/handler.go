package handlers

import (
	"github.com/hellphone/finance/api"
	"github.com/hellphone/finance/domain/model"
	"io/ioutil"
	"net/http"

	"github.com/hellphone/finance/domain/cases/add_money_to_user"
	"github.com/hellphone/finance/presentor/jsonapi"
)

func AddMoneyToUser(ctx *api.Context, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	requestBody, err := jsonapi.UnmarshalAddMoneyToUser(body)
	if err != nil {
		return nil, err
	}

	request := &add_money_to_user.Request{
		User: &model.User{
			Id:          requestBody.UserId,
			MoneyAmount: requestBody.MoneyAmount,
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
	// TODO: transferring user Id, receiving user Id and money amount
	return nil, nil
}
