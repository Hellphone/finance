package jsonapi

import (
	"encoding/json"
	"errors"
)

type AddMoneyToUser struct {
	UserId      int     `json:"user_id"`
	MoneyAmount float64 `json:"money_amount"`
}

func UnmarshalAddMoneyToUser(bs []byte) (*AddMoneyToUser, error) {
	req := &AddMoneyToUser{}
	if err := json.Unmarshal(bs, req); err != nil {
		return nil, errors.New("incorrect data format")
	}

	if req.UserId == 0 {
		return nil, errors.New("incorrect data format. Field `user_id` is required")
	}

	if req.MoneyAmount == 0 {
		return nil, errors.New("incorrect data format. Field `money_amount` is required")
	}

	return req, nil
}
