package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type price24hr struct {
	Symbol             string  `json:"symbol"`
	PriceChange        float64 `json:"priceChange"`
	PriceChangePercent float64 `json:"priceChangePercent"`
	WeightedAvgPrice   float64 `json:"weightedAvgPrice"`
	PrevClosePrice     float64 `json:"prevClosePrice"`
	LastPrice          float64 `json:"lastPrice"`
	LastQty            float64 `json:"lastQty"`
	BidPrice           float64 `json:"bidPrice"`
	BidQty             float64 `json:"bidQty"`
	AskPrice           float64 `json:"askPrice"`
	AskQty             float64 `json:"askQty"`
	OpenPrice          float64 `json:"openPrice"`
	HighPrice          float64 `json:"highPrice"`
	LowPrice           float64 `json:"lowPrice"`
	Volume             float64 `json:"volume"`
	QuoteVolume        float64 `json:"quoteVolume"`
	OpenTime           float64 `json:"openTime"`
	CloseTime          float64 `json:"closeTime"`
	FirstId            int     `json:"firstId"`
	LastId             int     `json:"LastId"`
	Count              int     `json:"Count"`
}

func (jsonPrice24hr *price24hr) ParseJson(resp *http.Response) {

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, jsonPrice24hr)
	if err != nil {
		panic(err)
	}
}
