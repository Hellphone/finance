package jsonapi

import (
	"encoding/json"
	"errors"
	"github.com/hellphone/finance/domain/model"
)

type User struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	MoneyAmount float64 `json:"money_amount"`
}

type userResponse struct {
	Data *User `json:"data"`
}

type usersResponse struct {
	Data []*User `json:"data"`
}

func MarshalUser(user *model.User) ([]byte, error) {
	data := userResponse{
		Data: newUserFromModel(user),
	}

	return json.Marshal(data)
}

func MarshalUsers(users []*model.User) ([]byte, error) {
	usersData := make([]*User, len(users))
	for i, user := range users {
		usersData[i] = newUserFromModel(user)
	}

	data := usersResponse{Data: usersData}

	return json.Marshal(data)
}

func newUserFromModel(u *model.User) *User {
	return &User{
		Id:          u.Id,
		Name:        u.Name,
		MoneyAmount: u.MoneyAmount,
	}
}

func UnmarshalUser(data []byte) (*model.User, error) {
	userResponse := userResponse{}

	err := json.Unmarshal(data, &userResponse)
	if err != nil {
		return nil, err
	}

	if userResponse.Data == nil {
		return nil, errors.New("Data should not be empty")
	}

	return userResponse.Data.toModel(), nil
}

func (u *User) toModel() *model.User {
	return &model.User{
		Id:          u.Id,
		Name:        u.Name,
		MoneyAmount: u.MoneyAmount,
	}
}
