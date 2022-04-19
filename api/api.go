package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/hellphone/finance/domain"
	"github.com/hellphone/finance/presenter/jsonapi"
	"github.com/hellphone/finance/repository"

	"github.com/gorilla/mux"
)

type Context struct {
	Repositories repository.Factory
	Body         []byte
}

type RouteHandler func(*Context, *http.Request) (int, []byte, error)

func HandleRoute(r *mux.Router, path string, handler RouteHandler) *mux.Route {
	return r.HandleFunc(path, func(resp http.ResponseWriter, req *http.Request) {
		ctx, err := extractContext(req)
		if err != nil {
			RespondWithErrorJSON(resp, http.StatusInternalServerError, err)
			return
		}

		log.Println(req.Method + " " + req.RequestURI)

		code, res, err := handler(ctx, req)
		if err != nil {
			RespondWithErrorJSON(resp, code, err)
			return
		}

		if code == 0 {
			RespondWithErrorJSON(resp, code, errors.New("API handler return unexpected response status: 0"))
			return
		}

		RespondWithJSON(resp, http.StatusOK, res)
	})
}

func RespondWithErrorJSON(w http.ResponseWriter, code int, err error) {
	if code == 0 {
		code = http.StatusInternalServerError
	}

	e := domain.ToDomainError(err)

	//switch {
	//case e.IsNotFoundError():
	//	code = http.StatusNotFound
	//case e.IsValidateError():
	//	code = http.StatusBadRequest
	//}

	RespondWithJSON(w, code, jsonapi.NewErrorResponse(e).ToBytes())

	log.Printf("error handling request: %s", err.Error())
}

func RespondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

func DomainError(err error) (int, []byte, error) {
	code := http.StatusInternalServerError

	if domain.ToDomainError(err).IsNotFoundError() {
		code = http.StatusNotFound
	}

	return code, nil, err
}

func BadRequest(err error) (int, []byte, error) {
	return http.StatusBadRequest, nil, err
}

func InternalServerError(err error) (int, []byte, error) {
	return http.StatusInternalServerError, nil, err
}

func OK(body []byte) (int, []byte, error) {
	return http.StatusOK, body, nil
}

func extractContext(req *http.Request) (*Context, error) {
	ctx, ok := req.Context().Value("context").(*Context)
	if !ok {
		return nil, errors.New("Invalid request context")
	}
	return ctx, nil
}
