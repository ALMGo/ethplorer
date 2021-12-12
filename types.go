package ethplorer

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

// UnmarshalJSON decodes an int64 timestamp into a time.Time object
func (p *Timestamp) UnmarshalJSON(bytes []byte) error {
	// 1. Decode the bytes into an int64
	var raw int64
	err := json.Unmarshal(bytes, &raw)

	if err != nil {
		fmt.Printf("error decoding timestamp: %s\n", err)
		return err
	}

	// 2 - Parse the unix timestamp
	*&p.Time = time.Unix(raw, 0)
	return nil
}

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
	Address            string     `json:"address"`
	Name               string     `json:"name"`
	Decimals           uint64     `json:"decimals,string"`
	Symbol             string     `json:"symbol"`
	TotalSupply        string     `json:"totalSupply"`
	Owner              string     `json:"owner"`
	TxsCount           uint64     `json:"txsCount"`
	TransfersCount     uint64     `json:"transfersCount"`
	LastUpdated        int64      `json:"lastUpdated"`
	Slot               uint64     `json:"slot"`
	StorageTotalSupply uint64     `json:"StorageTotalSupply"`
	IssuancesCount     uint64     `json:"issuancesCount"`
	HoldersCount       uint64     `json:"holdersCount"`
	Image              string     `json:"image"`
	Description        string     `json:"description"`
	Website            string     `json:"website"`
	Telegram           string     `json:"telegram"`
	Twitter            string     `json:"twitter"`
	Reddit             string     `json:"reddit"`
	Facebook           string     `json:"facebook"`
	Coingecko          string     `json:"coingecko"`
	EthTransferCount   uint64     `json:"ethTransfersCount"`
	Price              TokenPrice `json:"price"`
	CountOps           uint64     `json:"countOps"`
	PublicTags         []string   `json:"publicTags"`
	OpCount            uint64     `json:"opCount"`
	Added              uint64     `json:"added"`
}

type TokenFinancials struct {
	Balance    float64 `json:"balance"`
	RawBalance string  `json:"rawBalance"`
	TotalIn    float64 `json:"totalIn"`
	TotalOut   float64 `json:"totalOut"`
}

type Holders struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
	Share   float64 `json:"share"`
}

type TopTokenHolders struct {
	Holders []Holders `json:"holders"`
}

type ETH struct {
	TokenFinancials
	Price TokenPrice `json:"price"`
}

type ContractInfo struct {
	CreatorHash     string    `json:"creatorHash"`
	TransactionHash string    `json:"transactionHash"`
	Timestamp       Timestamp `json:"timestamp"`
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

type Operations struct {
	Timestamp       Timestamp `json:"timestamp"`
	TransactionHash string    `json:"transactionHash"`
	TokenInfo       TokenInfo `json:"tokenInfo"`
	Type            string    `json:"type"` // transfer, approve, issuance, mint, burn
	Address         string    `json:"address"`
	From            string    `json:"from"`
	To              string    `json:"to"`
	Value           uint64    `json:"value,string"`
}

type TokenHistory struct {
	Operations []Operations `json:"operations"`
}

type AddressTransaction struct {
	Timestamp Timestamp `json:"timestamp"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Hash      string    `json:"hash"`
	Value     float64   `json:"value"`
	Input     string    `json:"input"`
	Success   bool      `json:"success"`
}

type TopTokens struct {
	Tokens  []TokenInfo `json:"tokens"`
	OpCount uint64      `json:"opCount"`
}

type GetTokenHistoryParams struct {
	Type      string
	Limit     uint64
	Timestamp Timestamp
}

type TransactionDate struct {
	Year  uint64 `json:"year"`
	Month uint64 `json:"month"`
	Day   uint64 `json:"day"`
}

type CountTxs struct {
	Id  TransactionDate `json:"_id"`
	Ts  uint64          `json:"ts"`
	Cnt uint64          `json:"cnt"`
}

type TokenDailyTransactionCounts struct {
	CountTxs []CountTxs `json:"countTxs"`
}

type Price struct {
	Ts        uint64  `json:"ts"`
	Date      string  `json:"date"`
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    float64 `json:"volume"`
	VolumeUSD float64 `json:"volumeConverted"`
	Average   float64 `json:"average"`
}

type History struct {
	TokenDailyTransactionCounts
	Prices []Price `json:"prices"`
}

type TokenDailyPriceHistory struct {
	History History `json:"history"`
}

type GetAddressHistoryParams struct {
	GetTokenHistoryParams
	Token string `json:"token"`
}

type GetAddressTransactionsParams struct {
	Limit          uint64
	Timestamp      Timestamp
	ShowZeroValues bool
}

type GetTopParams struct {
	Limit    uint64
	Criteria string
}

type GetAddressTokenInfoParams struct {
	Token         string
	ShowEthTotals bool
}
