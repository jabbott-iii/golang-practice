package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params CoinBalanceParams
	var decoder = schema.NewDecoder()

	if err := decoder.Decode(&params, r.URL.Query()); err != nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}

	database, err := NewDatabase()
	if err != nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}

	tokenDetails := database.GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error("no coin details found for user:", params.Username)
		InternalErrorHandler(w)
		return
	}

	var response = CoinBalanceResponse{
		Balance: tokenDetails.Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}
}
