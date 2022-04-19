package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/hellphone/finance/domain"
	"github.com/hellphone/finance/domain/model"
	"github.com/hellphone/finance/domain/repository"

	_ "github.com/lib/pq"
)

const userTableName = "users"

type userRepository struct {
	db *sql.DB
}

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	MoneyAmount float64 `json:"money_amount"`
}

func (u *User) toModel() *model.User {
	return &model.User{
		Name:        u.Name,
		MoneyAmount: u.MoneyAmount,
	}
}

func (r *userRepository) GetList(f repository.User) ([]*model.User, error) {
	return r.getList(f, 0, 0)
}

func (r *userRepository) getList(f repository.User, limit, offset int) ([]*model.User, error) {

	return nil, nil
}

// TODO: probably save several users at once
func (r *userRepository) Store(u *model.User) (*model.User, error) {
	if u.Id != 0 {
		// update
		query := fmt.Sprintf(
			`UPDATE %s SET name = '%s', money_amount = %.2f WHERE id=%d`,
			userTableName,
			u.Name,
			u.MoneyAmount,
			u.Id,
		)

		_, err := r.db.Exec(query)
		if err != nil {
			return nil, err
		}

		return u, nil
	}
	// store
	query := fmt.Sprintf(
		`INSERT INTO %s ("name", "money_amount") VALUES ("name"=%s, "money_amount"=%f)`,
		userTableName,
		u.Name,
		u.MoneyAmount,
	)

	_, err := r.db.Exec(query)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetOneById(id int) (*model.User, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = %d;`, userTableName, id)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	user := &model.User{}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.MoneyAmount)
		if err != nil {
			return nil, err
		}
	}

	if user.Id == 0 {
		return nil, domain.NewNotFoundError(fmt.Sprintf("user with ID [%d] not found", id))
	}

	return user, nil
}
