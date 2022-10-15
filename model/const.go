package model

import "errors"

const (
	BlockTxQueryLimit     = 50
	LastQueryLimit        = 20
	SolscanBaseURL        = "public-api.solscan.io"
	SolscanBlocksURL      = "https://public-api.solscan.io/block/last"
	SolscanBlockDetailURL = "https://public-api.solscan.io/block"
	SolscanBlockTxsURL    = "https://public-api.solscan.io/block/transactions"

	SolscanLastTxsURL  = "https://public-api.solscan.io/transaction/last"
	SolscanTxDetailURL = "https://public-api.solscan.io/transaction"
)

var (
	ErrExceedQueryLimit = errors.New("exceed max query limit")
)
