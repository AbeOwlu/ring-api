package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AbeOwlu/ring-api/api"
	"github.com/AbeOwlu/ring-api/internal/tools"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	loggInit := zap.NewExample()
	logger := loggInit.Sugar()

	params := api.CoinBalanceParam{}
	decoder := schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		logger.Info(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	var tokenDet *tools.CoinDetails
	tokenDet = (*database).GetUserCoins(params.Username)
	if tokenDet == nil {
		logger.Info(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance:     (*&tokenDet).Coins,
		RespondCode: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Info(err)
		api.InternalErrorHandler(w, err)
		return
	}

}
