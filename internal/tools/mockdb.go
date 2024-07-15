package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789EFG",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    456,
		Username: "jason",
	},
	"marie": {
		Coins:    789,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second)
	client, ok := mockLoginDetails[username]
	if ok != true {
		return nil
	}

	return &client
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second)

	clientData, ok := mockCoinDetails[username]
	if ok != true {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
