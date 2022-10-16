package model

import "errors"

const (
	TxQueryLimit     = 50
	LastQueryLimit   = 20
	HolderQueryLimit = 100
	// block url
	SolscanBlocksURL      = "https://public-api.solscan.io/block/last"
	SolscanBlockDetailURL = "https://public-api.solscan.io/block"
	SolscanBlockTxsURL    = "https://public-api.solscan.io/block/transactions"
	// tx url
	SolscanLastTxsURL  = "https://public-api.solscan.io/transaction/last"
	SolscanTxDetailURL = "https://public-api.solscan.io/transaction"
	// account url
	SolscanAccountTokensURL             = "https://public-api.solscan.io/account/tokens"
	SolscanAccountTxsURL                = "https://public-api.solscan.io/account/transaction"
	SolscanAccountStakeAccountsURL      = "https://public-api.solscan.io/account/stakeAccounts"
	SolscanAccountSplTransfersURL       = "https://public-api.solscan.io/account/splTransfers"
	SolscanAccountSolTransfersURL       = "https://public-api.solscan.io/account/solTransfers"
	SolscanAccountExportTransactionsURL = "https://public-api.solscan.io/account/exportTransactions"
	SolscanAccountInfoURL               = "https://public-api.solscan.io/account"

	ExportTypeTokenChange = "tokenchange"
	ExportTypeSolTransfer = "soltransfer"
	ExportTypeAll         = "all"

	// token url
	SolscanTokenHoldersURL = "https://public-api.solscan.io/token/holders"
	SolscanTokenMetaURL    = "https://public-api.solscan.io/token/meta"
	SolscanTokenListURL    = "https://public-api.solscan.io/token/list"

	// chain url
	SolscanChainInfoURL = "https://public-api.solscan.io/chaininfo"
	// market url
	SolscanMarketTokenURL = "https://public-api.solscan.io/market/token"
)

var (
	ErrExceedQueryLimit  = errors.New("exceed max query limit")
	ErrInvalidExportType = errors.New("invalid export type")
)
