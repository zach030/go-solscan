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

type AccountStakeInfo struct {
	Voter        string   `json:"voter"`
	Amount       string   `json:"amount"`
	Type         string   `json:"type"`
	StakeAccount string   `json:"stakeAccount"`
	Staker       string   `json:"staker"`
	Role         []string `json:"role"`
}

type GetAccountSplTransfersResponse struct {
	Data []struct {
		Id           string   `json:"_id"`
		Address      string   `json:"address"`
		Signature    []string `json:"signature"`
		ChangeType   string   `json:"changeType"`
		ChangeAmount int64    `json:"changeAmount"`
		Decimals     int      `json:"decimals"`
		PostBalance  string   `json:"postBalance"`
		PreBalance   string   `json:"preBalance"`
		TokenAddress string   `json:"tokenAddress"`
		Symbol       string   `json:"symbol"`
		BlockTime    int      `json:"blockTime"`
		Slot         int      `json:"slot"`
		Fee          int      `json:"fee"`
		Owner        string   `json:"owner"`
		Balance      struct {
			Amount   string `json:"amount"`
			Decimals int    `json:"decimals"`
		} `json:"balance"`
	} `json:"data"`
	Total int `json:"total"`
}

type GetAccountSolTransfersResponse struct {
	Data []struct {
		Id                  string `json:"_id"`
		Src                 string `json:"src"`
		Dst                 string `json:"dst"`
		Lamport             int    `json:"lamport"`
		BlockTime           int    `json:"blockTime"`
		Slot                int    `json:"slot"`
		TxHash              string `json:"txHash"`
		Fee                 int    `json:"fee"`
		Status              string `json:"status"`
		Decimals            int    `json:"decimals"`
		TxNumberSolTransfer int    `json:"txNumberSolTransfer"`
	} `json:"data"`
}

type GetAccountInfoResponse struct {
	Lamports     int    `json:"lamports"`
	OwnerProgram string `json:"ownerProgram"`
	Type         string `json:"type"`
	RentEpoch    int    `json:"rentEpoch"`
	Account      string `json:"account"`
}
