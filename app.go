package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hellphone/finance/api"
	v1 "github.com/hellphone/finance/api/v1"
	"github.com/hellphone/finance/repository"
	"github.com/hellphone/finance/repository/postgresql"
)

type App struct {
	Port         string
	Repositories repository.Factory
}

func (a *App) Init(cfg *Config) error {
	a.Port = cfg.Port

	factory, err := postgresql.NewFactory(cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName, cfg.DB.Username, cfg.DB.Password)
	if err != nil {
		return err
	}

	a.Repositories = factory

	return nil
}

func (a *App) Run() {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			c := context.WithValue(req.Context(), "context", &api.Context{
				Repositories: a.Repositories,
			})

			next.ServeHTTP(w, req.WithContext(c))
		})
	})

	v1.Init(r)

	log.Println("server started")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", a.Port), r))
}
