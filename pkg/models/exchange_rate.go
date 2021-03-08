package models

import "strings"

// LatestExchangeRateResponse returns a view-object which contains latest exchange rate
type LatestExchangeRateResponse struct {
	DataSource    string                  `json:"dataSource"`
	ReferenceUrl  string                  `json:"referenceUrl"`
	UpdateTime    int64                   `json:"updateTime"`
	BaseCurrency  string                  `json:"baseCurrency"`
	ExchangeRates LatestExchangeRateSlice `json:"exchangeRates"`
}

// LatestExchangeRate represents a data pair of currency and exchange rate
type LatestExchangeRate struct {
	Currency string `json:"currency"`
	Rate     string `json:"rate"`
}

// LatestExchangeRateSlice represents the slice data structure of LatestExchangeRate
type LatestExchangeRateSlice []*LatestExchangeRate

// Len returns the count of items
func (c LatestExchangeRateSlice) Len() int {
	return len(c)
}

// Swap swaps two items
func (c LatestExchangeRateSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// Less reports whether the first item is less than the second one
func (c LatestExchangeRateSlice) Less(i, j int) bool {
	return strings.Compare(c[i].Currency, c[j].Currency) < 0
}
