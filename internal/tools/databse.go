package tools

import (
	"go.uber.org/zap"
	// "github.com/AbeOwlu/ring-api/internal/tools"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	loggInit := zap.NewExample()
	logger := loggInit.Sugar()
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		logger.Info(err)
		return nil, err
	}

	return &database, err
}
