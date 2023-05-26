package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Price struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Time struct {
	Time float64 `json:"time"`
}

func (price *Price) TikerPrice(tiker string) {
	var url = "https://api.binance.com/api/v3/ticker/price?symbol=" + tiker
	var result = UnmarshalBinance(url, &price).(map[string]interface{})
	println(1)
	price = &Price{
		Symbol: result["symbol"].(string),
		Price:  result["price"].(string),
	}

}

func (time Time) BinanceTime() Time {
	var url = "https://api.binance.com/api/v3/time"
	var result = UnmarshalBinance(url, time).(map[string]interface{})
	time = Time{
		Time: result["serverTime"].(float64),
	}
	return time
}

func UnmarshalBinance(url string, allStruct interface{}) interface{} {

	var hhtpRequest = MakeRequest(url)
	if err := json.Unmarshal(hhtpRequest, &allStruct); err != nil {
		panic(err)
	}
	return allStruct
}

func MakeRequest(url string) (result []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	result = []byte(body)
	return result
}
