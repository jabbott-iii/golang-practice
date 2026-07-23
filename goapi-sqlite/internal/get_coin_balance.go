package internal

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}

	var database *DatabaseInterface
	database, err = NewDatabase()
	if err != nil {
		InternalErrorHandler(w)
		return
	}

	var tokenDetails *CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}
	
	var response = CoinBalanceResponse {
		Balance: (*tokenDetails).Coins,
		Code: http.StatusOK,
	}

	w.Header().Set("Constant-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		InternalErrorHandler(w)
		return
	}
}