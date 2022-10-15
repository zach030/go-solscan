package model

type AccountTokens struct {
	TokenAddress string `json:"tokenAddress"`
	TokenAmount  struct {
		Amount         string `json:"amount"`
		Decimals       int    `json:"decimals"`
		UiAmount       int    `json:"uiAmount"`
		UiAmountString string `json:"uiAmountString"`
	} `json:"tokenAmount"`
	TokenAccount string `json:"tokenAccount"`
	TokenName    string `json:"tokenName"`
	TokenIcon    string `json:"tokenIcon"`
	RentEpoch    int    `json:"rentEpoch"`
	Lamports     int    `json:"lamports"`
}

type AccountTx struct {
	BlockTime          int      `json:"blockTime"`
	Slot               int      `json:"slot"`
	TxHash             string   `json:"txHash"`
	Fee                int      `json:"fee"`
	Status             string   `json:"status"`
	Lamport            int      `json:"lamport"`
	Signer             []string `json:"signer"`
	IncludeSPLTransfer bool     `json:"includeSPLTransfer"`
	ParsedInstruction  []struct {
		ProgramId string `json:"programId"`
		Type      string `json:"type"`
		Program   string `json:"program,omitempty"`
	} `json:"parsedInstruction"`
}
