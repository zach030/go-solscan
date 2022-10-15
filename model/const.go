package model

import "errors"

const (
	TxQueryLimit   = 50
	LastQueryLimit = 20
	SolscanBaseURL = "public-api.solscan.io"
	// block url
	SolscanBlocksURL      = "https://public-api.solscan.io/block/last"
	SolscanBlockDetailURL = "https://public-api.solscan.io/block"
	SolscanBlockTxsURL    = "https://public-api.solscan.io/block/transactions"
	// tx url
	SolscanLastTxsURL  = "https://public-api.solscan.io/transaction/last"
	SolscanTxDetailURL = "https://public-api.solscan.io/transaction"
	// account url
	SolscanAccountTokensURL = "https://public-api.solscan.io/account/tokens"
	SolscanAccountTxsURL    = "https://public-api.solscan.io/account/transaction"
	// chain url
	SolscanChainInfoURL = "https://public-api.solscan.io/chaininfo"
	// market url
	SolscanMarketTokenURL = "https://public-api.solscan.io/market/token"
)

var (
	ErrExceedQueryLimit = errors.New("exceed max query limit")
)
