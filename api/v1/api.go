package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hellphone/finance/api"
	"github.com/hellphone/finance/api/handlers"
)

func Init(r *mux.Router) {
	api.HandleRoute(r, "/api/v1/add_money_to_user", handlers.AddMoneyToUser).Methods(http.MethodPost)
	api.HandleRoute(r, "/api/v1/transfer_money", handlers.TransferMoney).Methods(http.MethodPost)
}
