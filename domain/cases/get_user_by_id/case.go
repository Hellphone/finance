package get_user_by_id

import (
	"github.com/hellphone/finance/domain/model"
)

type Request struct {
	ID int
}

type Response struct {
	User *model.User
}

func GetUserByID() (*model.User, error) {
	// TODO: get user by Id from DB

	return nil, nil
}
