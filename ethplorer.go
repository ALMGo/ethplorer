package ethplorer

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

var client = resty.New()

// TODO: delete todo
func GetLastBlock(apiKey string) (uint64, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
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

// TODO: delete todo
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

// TODO: delete todo
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

// TODO: delete todo
func GetAddressTokenInfo(address string, params GetAddressTokenInfoParams, apiKey string) (*AddressInfo, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}

	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey
	queryParams["showETHTotals"] = strconv.FormatBool(params.ShowEthTotals)

	if params.Token != "" {
		queryParams["token"] = params.Token
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(AddressInfo{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getAddressInfo/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*AddressInfo)
	return result, nil
}

// TODO: delete todo
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

// TODO: delete todo
func GetTokenHistory(address string, params GetTokenHistoryParams, apiKey string) (*TokenHistory, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}

	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if  !params.Timestamp.IsZero() {
		queryParams["timestamp"] = strconv.FormatInt(params.Timestamp.UnixNano(), 10)
	}

	if params.Limit != 0 {
		limit := params.Limit
		if limit > 1000 {
			limit = 1000
		}
		queryParams["limit"] = strconv.FormatUint(limit, 10)
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

// TODO: delete todo
func GetAddressHistory(address string, params GetAddressHistoryParams, apiKey string) (*TokenHistory, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if !params.Timestamp.IsZero() {
		queryParams["timestamp"] = strconv.FormatInt(params.Timestamp.UnixNano(), 10)
	}

	if params.Limit != 0 {
		limit := params.Limit
		if limit > 1000 {
			limit = 1000
		}
		queryParams["limit"] = strconv.FormatUint(limit, 10)
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
		Get("https://api.ethplorer.io/getAddressHistory/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenHistory)
	return result, nil
}

// TODO: delete todo
func GetAddressTransactions(address string, params GetAddressTransactionsParams, apiKey string) (*[]AddressTransaction, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey
	if !params.Timestamp.IsZero() {
		queryParams["timestamp"] = strconv.FormatInt(params.Timestamp.UnixNano(), 10)
	}

	if params.Limit != 0 {
		limit := params.Limit
		if limit > 1000 {
			limit = 1000
		}
		queryParams["limit"] = strconv.FormatUint(limit, 10)
	}

	queryParams["showZeroValues"] = strconv.FormatBool(params.ShowZeroValues)

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

// TODO: delete todo
func GetTopTokens(apiKey string) (*TopTokens, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": apiKey,
		}).
		SetResult(TopTokens{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTopTokens/")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TopTokens)
	return result, nil
}

func GetTop(params GetTopParams, apiKey string) (*TopTokens, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey
	if params.Criteria != "" {
		queryParams["criteria"] = params.Criteria
	}

	if params.Limit != 0 {
		limit := params.Limit
		if limit > 1000 {
			limit = 1000
		}
		queryParams["limit"] = strconv.FormatUint(limit, 10)
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TopTokens{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTop/")
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TopTokens)
	return result, nil
}

func GetTopTokenHolders(address string, limit uint64, apiKey string) (*TopTokenHolders, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if limit != 0 {
		if limit > 1000 {
			limit = 1000
		}
		queryParams["limit"] = strconv.FormatUint(limit, 10)
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TopTokenHolders{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTopTokenHolders/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TopTokenHolders)
	return result, nil
}

func GetTokenDailyTransactionCounts(address string, period uint64, apiKey string) (*TokenDailyTransactionCounts, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if period != 0 {
		if period > 90 {
			period = 90
		}
		queryParams["period"] = strconv.FormatUint(period, 10)
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TokenDailyTransactionCounts{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokenHistoryGrouped/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenDailyTransactionCounts)
	return result, nil
}

func GetTokenDailyPriceHistory(address string, period uint64, apiKey string) (*TokenDailyPriceHistory, error) {
	if apiKey == "" {
		apiKey = "freekey"
	}
	queryParams := make(map[string]string)
	queryParams["apiKey"] = apiKey

	if period != 0 {
		if period > 90 {
			period = 90
		}
		queryParams["period"] = strconv.FormatUint(period, 10)
	}

	resp, err := client.R().
		SetQueryParams(queryParams).
		SetResult(TokenDailyPriceHistory{}).
		ForceContentType("application/json").
		Get("https://api.ethplorer.io/getTokenPriceHistoryGrouped/" + address)
	if err != nil {
		return nil, err
	}
	result := resp.Result().(*TokenDailyPriceHistory)
	return result, nil
}
