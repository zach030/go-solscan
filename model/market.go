package model

type TokenMarket struct {
	PriceUsdt      int     `json:"priceUsdt"`
	VolumeUsdt     int64   `json:"volumeUsdt"`
	MarketCapFD    int64   `json:"marketCapFD"`
	MarketCapRank  int     `json:"marketCapRank"`
	PriceChange24H float64 `json:"priceChange24h"`
}
