package model

type ChainInfo struct {
	BlockHeight      int   `json:"blockHeight"`
	CurrentEpoch     int   `json:"currentEpoch"`
	AbsoluteSlot     int   `json:"absoluteSlot"`
	TransactionCount int64 `json:"transactionCount"`
}
