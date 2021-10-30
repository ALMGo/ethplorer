package ethplorer

import (
	"encoding/json"
	"strconv"
)

type LastBlock struct {
	LastBlock uint64 `json:"lastBlock"`
}

type TokenPrice struct {
	Rate            float64 `json:"rate"`
	Currency        string  `json:"currency"`
	Diff            float64 `json:"diff"`
	Diff1           float64 `json:"diff1"`
	Diff7d          float64 `json:"diff7d"`
	Diff30d         float64 `json:"diff30d"`
	MarketCapUsd    float64 `json:"marketCapUsd"`
	AvailableSupply float64 `json:"availableSupply"`
	Volume24h       float64 `json:"volume24h"`
	Ts              int64   `json:"ts"`
}

func (w *TokenPrice) UnmarshalJSON(data []byte) error {
	type TokenPrice2 TokenPrice
	var result TokenPrice2
	if _, err := strconv.ParseBool(string(data)); err == nil {
		*w = TokenPrice{}
		return nil
	}
	err := json.Unmarshal(data, &result)
	*w = TokenPrice(result)
	return err
}

type TokenInfo struct {
	Address          string     `json:"address"`
	TotalSupply      string     `json:"totalSupply"`
	Name             string     `json:"name"`
	Decimals         string     `json:"decimals"'`
	Symbol           string     `json:"symbol"`
	Price            TokenPrice `json:"price"`
	Owner            string     `json:"owner"`
	CountOps         int64      `json:"countOps"`
	TransfersCount   int64      `json:"transfersCount"`
	HoldersCount     int64      `json:"holdersCount"`
	IssuancesCount   int64      `json:"issuancesCount"`
	LastUpdated      int64      `json:"lastUpdated"`
	Image            string     `json:"image"`
	Description      string     `json:"description"`
	Website          string     `json:"website"`
	Facebook         string     `json:"facebook"`
	Telegram         string     `json:"talagram"`
	Twitter          string     `json:"twitter"`
	Reddit           string     `json:"reddit"`
	Coingecko        string     `json:"coingecko"`
	EthTransferCount uint64     `json:"ethTransfersCount"`
	PublicTags       []string   `json:"publicTags"`
	Added            uint64     `json:"added"`
}

type TokenFinancials struct {
	Balance    float64 `json:"balance"`
	RawBalance string  `json:"rawBalance"`
	TotalIn    float64 `json:"totalIn"`
	TotalOut   float64 `json:"totalOut"`
}

type ETH struct {
	TokenFinancials
	Price TokenPrice `json:"price"`
}

type ContractInfo struct {
	CreatorHash     string `json:"creatorHash"`
	TransactionHash string `json:"transactionHash"`
	Timestamp       uint64 `json:"timestamp"`
}

type Token struct {
	TokenInfo TokenInfo `json:"tokenInfo"`
	TokenFinancials
}

type AddressInfo struct {
	Address      string       `json:"address"`
	ETH          ETH          `json:"eth"`
	ContractInfo ContractInfo `json:"contractInfo"`
	TokenInfo    TokenInfo    `json:"tokenInfo"`
	Tokens       []Token      `json:"tokens"`
}
