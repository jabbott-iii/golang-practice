package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {

	"tom": {
		AuthToken: "589AGC",
		Username: "tom",
	},
	"jack": {
		AuthToken: "134KFG",
		Username: "jack",
	},
	"todd": {
		AuthToken: "689TDV",
		Username: "todd",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"tom": {
		Coins: 1000,
		Username: "tom",
	},
	"jack": {
		Coins: 99999,
		Username: "jack",
	},
	"todd": {
		Coins: 22211,
		Username: "todd",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//simulate database call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// ismulate database call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}