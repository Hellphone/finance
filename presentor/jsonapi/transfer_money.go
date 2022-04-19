package jsonapi

import (
	"encoding/json"
	"errors"
)

type TransferMoney struct {
	UserFromId  int     `json:"user_from_id"`
	UserToId    int     `json:"user_to_id"`
	MoneyAmount float64 `json:"money_amount"`
}

func UnmarshalTransferMoney(bs []byte) (*TransferMoney, error) {
	req := &TransferMoney{}
	if err := json.Unmarshal(bs, req); err != nil {
		return nil, errors.New("incorrect data format")
	}

	if req.UserFromId == 0 {
		return nil, errors.New("incorrect data format. Field `user_from_id` is required")
	}

	if req.UserToId == 0 {
		return nil, errors.New("incorrect data format. Field `user_to_id` is required")
	}

	if req.MoneyAmount == 0 {
		return nil, errors.New("incorrect data format. Field `money_amount` is required")
	}

	return req, nil
}
