package main

import "io/ioutil"

func main() {

	pp, _ := NewApiClinet("https://api.binance.com")

	tt, _ := pp.GetPriceTicker("BTCUSDT")

	body, _ := ioutil.ReadAll(tt.Body)
	println(string(body))

	p, _ := NewApiClinet("https://api.binance.com")

	t, _ := p.GetPrice24hr()

	var test price24hr

	test.ParseJson(t)
}
