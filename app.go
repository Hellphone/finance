package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", a.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			c := context.WithValue(req.Context(), "context", &api.Context{
				Repositories: a.Repositories,
			})

			next.ServeHTTP(w, req.WithContext(c))
		})
	})

	v1.Init(r)

	go func() {
		log.Println("starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
