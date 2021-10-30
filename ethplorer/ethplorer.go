package ethplorer

import (
	"github.com/go-resty/resty/v2"
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

func GetTokenInfo(address string) (*TokenInfo, error) {
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": "freekey",
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

func GetAddressInfo(address string) (*AddressInfo, error) {
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": "freekey",
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

func GetAddressTokenInfo(address string, token string) (*AddressInfo, error) {
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": "freekey",
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

func GetTokensNew() (*[]TokenInfo, error) {
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"apiKey": "freekey",
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
