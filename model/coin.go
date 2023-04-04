package model

import "time"

type ResponseData struct {
	Data   []Data `json:"data"`
	Status Status `json:"status"`
}
type Usd struct {
	Price                 float64   `json:"price"`
	Volume24H             int64     `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    int       `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
type Btc struct {
	Price                 int       `json:"price"`
	Volume24H             int       `json:"volume_24h"`
	VolumeChange24H       int       `json:"volume_change_24h"`
	PercentChange1H       int       `json:"percent_change_1h"`
	PercentChange24H      int       `json:"percent_change_24h"`
	PercentChange7D       int       `json:"percent_change_7d"`
	MarketCap             int       `json:"market_cap"`
	MarketCapDominance    int       `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}
type Quote struct {
	Usd Usd `json:"USD"`
	Btc Btc `json:"BTC"`
}
type Eth struct {
	Price                 int       `json:"price"`
	Volume24H             int       `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       int       `json:"percent_change_1h"`
	PercentChange24H      int       `json:"percent_change_24h"`
	PercentChange7D       int       `json:"percent_change_7d"`
	MarketCap             int       `json:"market_cap"`
	MarketCapDominance    int       `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}

type Data struct {
	ID                            int         `json:"id"`
	Name                          string      `json:"name"`
	Symbol                        string      `json:"symbol"`
	Slug                          string      `json:"slug"`
	CmcRank                       int         `json:"cmc_rank,omitempty"`
	NumMarketPairs                int         `json:"num_market_pairs"`
	CirculatingSupply             int         `json:"circulating_supply"`
	TotalSupply                   int         `json:"total_supply"`
	MaxSupply                     int         `json:"max_supply"`
	LastUpdated                   time.Time   `json:"last_updated"`
	DateAdded                     time.Time   `json:"date_added"`
	Tags                          []string    `json:"tags"`
	Platform                      interface{} `json:"platform"`
	SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply,omitempty"`
	SelfReportedMarketCap         interface{} `json:"self_reported_market_cap,omitempty"`
	Quote                         Quote       `json:"quote,omitempty"`
}
type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}
