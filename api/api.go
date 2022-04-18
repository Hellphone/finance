package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/hellphone/finance/repository"

	"github.com/gorilla/mux"
)

type Context struct {
	Repositories repository.Factory
}

type RouteHandler func(*Context, *http.Request) ([]byte, error)

func HandleRoute(r *mux.Router, path string, handler RouteHandler) *mux.Route {
	return r.HandleFunc(path, func(resp http.ResponseWriter, req *http.Request) {
		ctx, err := extractContext(req)
		if err != nil {
			RespondWithErrorJSON(resp, err)
			return
		}

		log.Println(req.Method+" "+req.RequestURI)

		// Run handler
		res, err := handler(ctx, req)

		// Handle error
		if err != nil {
			log.Fatalf("%s %s: %s", req.Method, req.RequestURI, err)
			RespondWithErrorJSON(resp, err)
			return
		}

		RespondWithJSON(resp, http.StatusOK, res)
	})
}

func RespondWithErrorJSON(w http.ResponseWriter, err error) {
	//e := domain.ToDomainError(err)
	//
	//code := http.StatusInternalServerError
	//switch {
	//case e.IsNotFoundError():
	//	code = http.StatusNotFound
	//case e.IsValidateError():
	//	code = http.StatusBadRequest
	//}

	//RespondWithJSON(w, code, jsonapi.NewErrorResponse(e).ToBytes())

	log.Fatalf("Handle request: %s", err.Error())
}

func RespondWithJSON(w http.ResponseWriter, code int, payload []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payload)
}

func extractContext(req *http.Request) (*Context, error) {
	ctx, ok := req.Context().Value("context").(*Context)
	if !ok {
		return nil, errors.New("Invalid request context")
	}
	return ctx, nil
}
