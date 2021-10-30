package ethplorer

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
)

var client = resty.New()

func GetLastblock() (uint64, error) {
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": "freekey",
		}).
		SetResult(LastBlock{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getLastBlock")
	if err != nil {
		return 0, err
	}
	result := resp.Result().(*LastBlock)
	return result.LastBlock, nil
}

func GetTokenInfo(address string, apiKey string) (*TokenInfo, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
		}).
		SetResult(TokenInfo{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokenInfo/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenInfo)
	return result, nil
}

func GetAddressInfo(address string, apiKey string) (*AddressInfo, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
			"showETHTotals": "true",
		}).
		SetResult(AddressInfo{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getAddressInfo/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AddressInfo)
	return result, nil
}

func GetAddressTokenInfo(address string, token string, apiKey string) (*AddressInfo, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
			"token": token,
			"showETHTotals": "true",
		}).
		SetResult(AddressInfo{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getAddressInfo/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AddressInfo)
	return result, nil
}

func GetTokensNew(apiKey string) (*[]TokenInfo, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
		}).
		SetResult([]TokenInfo{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokensNew/")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*[]TokenInfo)
	return result, nil
}

func GetTokenHistory(address string, params GetTokenHistoryParams, apiKey string) (*TokenHistory, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}

	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if params.Timestamp != 0 {
		queryParams["timestamp"] = strconv.FormatUint(params.Timestamp, 10)
	}

	if params.Limit != 0 {
		queryParams["limit"] = strconv.FormatUint(params.Limit, 10)
	}

	if params.Type != "" {
		queryParams["type"] = params.Type
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TokenHistory{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokenHistory/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenHistory)
	return result, nil
}

func GetAddressHistory(address string, params GetAddressHistoryParams, apiKey string) (*TokenHistory, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if params.Timestamp != 0 {
		queryParams["timestamp"] = strconv.FormatUint(params.Timestamp, 10)
	}

	if params.Limit != 0 {
		queryParams["limit"] = strconv.FormatUint(params.Limit, 10)
	}

	if params.Type != "" {
		queryParams["type"] = params.Type
	}

	if params.Token != "" {
		queryParams["token"] = params.Token
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TokenHistory{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokenHistory/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenHistory)
	return result, nil
}

func GetAddressTransactions(address string, params GetAddressTransactionsParams, apiKey string) (*[]AddressTransaction, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey
	if params.Timestamp != 0 {
		queryParams["timestamp"] = strconv.FormatUint(params.Timestamp, 10)
	}

	if params.Limit != 0 {
		queryParams["limit"] = strconv.FormatUint(params.Limit, 10)
	}

	if params.Limit != 0 {
		queryParams["limit"] = strconv.FormatUint(params.Limit, 10)
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult([]AddressTransaction{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getAddressTransactions/" + address)
	if err != nil {
		return nil, err
	}

	result := resp.Result().(*[]AddressTransaction)
	return result, nil
}
