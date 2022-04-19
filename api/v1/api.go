package v1

import (
	"github.com/gorilla/mux"
	"github.com/hellphone/finance/api"
	"github.com/hellphone/finance/api/handlers"
	"net/http"
)

func Init(r *mux.Router) {
	// 2 values: user ID and money amount
	api.HandleRoute(r, "/api/v1/add_money_to_user", handlers.AddMoneyToUser).Methods(http.MethodPost)
	// 3 values: user IDs and money amount
	api.HandleRoute(r, "/api/v1/transfer_money", handlers.TransferMoney).Methods(http.MethodPost)
}
