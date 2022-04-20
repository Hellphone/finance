package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	domainRepository "github.com/hellphone/finance/domain/repository"
	"github.com/hellphone/finance/repository"
	_ "github.com/lib/pq"
)

type factory struct {
	db *sql.DB
}

func NewFactory(host, port, dbname, username, password string) (repository.Factory, error) {
	if username != "" && dbname != "" && password != "" && host != "" && port != "" {
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, 5432, username, password, dbname)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}
		if err = db.Ping(); err != nil {
			return nil, err
		}

		return &factory{
			db: db,
		}, nil
	}
	return nil, errors.New("Postgresql credentials are empty")
}

func (f *factory) UserRepository() domainRepository.User {
	return &userRepository{
		db: f.db,
	}
}
