package model

import "time"

type GetTokenHoldersResponse struct {
	Data []struct {
		Address  string `json:"address"`
		Amount   int    `json:"amount"`
		Decimals int    `json:"decimals"`
		Owner    string `json:"owner"`
		Rank     int    `json:"rank"`
	} `json:"data"`
	Total int `json:"total"`
}

type GetTokenMetaResponse struct {
	Symbol      string        `json:"symbol"`
	Name        string        `json:"name"`
	Icon        string        `json:"icon"`
	Website     string        `json:"website"`
	Twitter     string        `json:"twitter"`
	Tag         []interface{} `json:"tag"`
	Decimals    int           `json:"decimals"`
	CoingeckoId string        `json:"coingeckoId"`
	Holder      int           `json:"holder"`
}

type GetTokenListResponse struct {
	Data []struct {
		MintAddress string `json:"mintAddress"`
		TokenSymbol string `json:"tokenSymbol"`
		TokenName   string `json:"tokenName"`
		Decimals    int    `json:"decimals"`
		Icon        string `json:"icon"`
		Website     string `json:"website"`
		Twitter     string `json:"twitter"`
		Extensions  struct {
			Website     string `json:"website"`
			CoingeckoId string `json:"coingeckoId"`
		} `json:"extensions"`
		TokenHolder   bool `json:"tokenHolder"`
		MarketCapRank int  `json:"marketCapRank"`
		Supply        struct {
			Amount         string  `json:"amount"`
			Decimals       int     `json:"decimals"`
			UiAmount       float64 `json:"uiAmount"`
			UiAmountString string  `json:"uiAmountString"`
		} `json:"supply"`
		Holder      int     `json:"holder"`
		PriceUst    float64 `json:"priceUst"`
		MarketCapFD float64 `json:"marketCapFD"`
		Volume      struct {
			VolumeUsd int `json:"volumeUsd"`
			Volume    int `json:"volume"`
		} `json:"volume"`
		CoingeckoInfo struct {
			MarketCapRank int `json:"marketCapRank"`
			CoingeckoRank int `json:"coingeckoRank"`
			MarketData    struct {
				CurrentPrice                 float64   `json:"currentPrice"`
				Ath                          float64   `json:"ath"`
				AthChangePercentage          float64   `json:"athChangePercentage"`
				AthDate                      time.Time `json:"athDate"`
				Atl                          float64   `json:"atl"`
				AtlChangePercentage          float64   `json:"atlChangePercentage"`
				AtlDate                      time.Time `json:"atlDate"`
				MarketCap                    int       `json:"marketCap"`
				MarketCapRank                int       `json:"marketCapRank"`
				FullyDilutedValuation        int64     `json:"fullyDilutedValuation"`
				TotalVolume                  int       `json:"totalVolume"`
				PriceHigh24H                 float64   `json:"priceHigh24h"`
				PriceLow24H                  float64   `json:"priceLow24h"`
				PriceChange24H               float64   `json:"priceChange24h"`
				PriceChangePercentage24H     float64   `json:"priceChangePercentage24h"`
				PriceChangePercentage7D      float64   `json:"priceChangePercentage7d"`
				PriceChangePercentage14D     float64   `json:"priceChangePercentage14d"`
				PriceChangePercentage30D     float64   `json:"priceChangePercentage30d"`
				PriceChangePercentage60D     float64   `json:"priceChangePercentage60d"`
				PriceChangePercentage200D    float64   `json:"priceChangePercentage200d"`
				PriceChangePercentage1Y      float64   `json:"priceChangePercentage1y"`
				MarketCapChange24H           float64   `json:"marketCapChange24h"`
				MarketCapChangePercentage24H float64   `json:"marketCapChangePercentage24h"`
				TotalSupply                  float64   `json:"totalSupply"`
				MaxSupply                    int64     `json:"maxSupply"`
				CirculatingSupply            float64   `json:"circulatingSupply"`
				LastUpdated                  time.Time `json:"lastUpdated"`
			} `json:"marketData"`
		} `json:"coingeckoInfo"`
		Tag []interface{} `json:"tag"`
	} `json:"data"`
	Total int `json:"total"`
}
