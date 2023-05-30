package main

import (
	"net/http"
	"net/url"
)

/*
		реализуем  структуру
	которая будет хранить в себе все, что связано с
	get запросом к ресурсу
*/

type ApiCLient struct {
	baseUrl    *url.URL     //Структура пакета url. Хранит всю информацию о URL
	httpClient *http.Client //интерфейс для отправки запросов на сервер
}

/*
	функция возвращающая готовую структуру ApiClient
*/
func NewApiClinet(baseURl string) (*ApiCLient, error) {

	parsedURl, err := url.Parse(baseURl) // парсим ссылку на ресурс

	if err != nil {
		return nil, err
	}

	client := &ApiCLient{ //записываем информацию в структуру
		baseUrl:    parsedURl,
		httpClient: &http.Client{},
	}
	return client, nil //возвращаем наполненую структуру типа ApiClient
}

func (apiclientSTR *ApiCLient) GetPriceTicker(resourceID string) (*http.Response, error) {
	endPoint := apiclientSTR.baseUrl.ResolveReference(&url.URL{Path: "api/v3/ticker/price", RawQuery: "symbol=" + resourceID})
	return GetResponse(endPoint, apiclientSTR)
}

func (apiclientSTR *ApiCLient) GetExchangeTime() (*http.Response, error) {
	endPoint := apiclientSTR.baseUrl.ResolveReference(&url.URL{Path: "/api/v3/time"})
	return GetResponse(endPoint, apiclientSTR)
}

func (apiclientSTR *ApiCLient) GetPrice24hr() (*http.Response, error) {
	endPoint := apiclientSTR.baseUrl.ResolveReference(&url.URL{Path: "/api/v3/ticker/24hr"})
	return GetResponse(endPoint, apiclientSTR)
}

func GetResponse(endPoint *url.URL, apiclientSTR *ApiCLient) (*http.Response, error) {

	request, err := http.NewRequest("GET", endPoint.String(), nil)
	if err != nil {
		return nil, err
	}

	//отправка запроса к серверу
	response, err := apiclientSTR.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
