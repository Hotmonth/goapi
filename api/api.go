package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}


type CoinBalanceRespone struct{
	Code int
	Balance int64
}

type Error struct{
	Code int
	Message string
}


