package model

type GetTxResponse struct {
	BlockTime    int      `json:"blockTime"`
	Slot         int      `json:"slot"`
	TxHash       string   `json:"txHash"`
	Fee          int      `json:"fee"`
	Status       string   `json:"status"`
	Lamport      int      `json:"lamport"`
	Signer       []string `json:"signer"`
	LogMessage   []string `json:"logMessage"`
	InputAccount []struct {
		Account     string `json:"account"`
		Signer      bool   `json:"signer"`
		Writable    bool   `json:"writable"`
		PreBalance  int64  `json:"preBalance"`
		PostBalance int64  `json:"postBalance"`
	} `json:"inputAccount"`
	RecentBlockhash   string `json:"recentBlockhash"`
	InnerInstructions []struct {
		Index              int `json:"index"`
		ParsedInstructions []struct {
			ProgramId string `json:"programId"`
			Program   string `json:"program"`
			Type      string `json:"type"`
			Name      string `json:"name"`
			Params    struct {
				Source        string `json:"source,omitempty"`
				Destination   string `json:"destination,omitempty"`
				Authority     string `json:"authority,omitempty"`
				Amount        string `json:"amount"`
				Account       string `json:"account,omitempty"`
				Mint          string `json:"mint,omitempty"`
				MintAuthority string `json:"mintAuthority,omitempty"`
			} `json:"params"`
			Extra struct {
				Source           string `json:"source"`
				Destination      string `json:"destination"`
				Authority        string `json:"authority"`
				Amount           string `json:"amount"`
				TokenAddress     string `json:"tokenAddress"`
				Decimals         int    `json:"decimals"`
				Symbol           string `json:"symbol"`
				Icon             string `json:"icon"`
				DestinationOwner string `json:"destinationOwner"`
				SourceOwner      string `json:"sourceOwner,omitempty"`
			} `json:"extra,omitempty"`
		} `json:"parsedInstructions"`
	} `json:"innerInstructions"`
	TokenBalanes []struct {
		Account string `json:"account"`
		Amount  struct {
			PostAmount string `json:"postAmount"`
			PreAmount  string `json:"preAmount"`
		} `json:"amount"`
		Token struct {
			Decimals     int    `json:"decimals"`
			TokenAddress string `json:"tokenAddress"`
			Name         string `json:"name,omitempty"`
			Symbol       string `json:"symbol,omitempty"`
			Icon         string `json:"icon,omitempty"`
		} `json:"token"`
	} `json:"tokenBalanes"`
	ParsedInstruction []struct {
		ProgramId  string `json:"programId"`
		Program    string `json:"program,omitempty"`
		Type       string `json:"type"`
		Data       string `json:"data,omitempty"`
		DataEncode string `json:"dataEncode,omitempty"`
		Name       string `json:"name"`
		Params     struct {
			NewAccount            string  `json:"newAccount,omitempty"`
			Source                string  `json:"source,omitempty"`
			TransferAmountSOL     float64 `json:"transferAmount(SOL),omitempty"`
			ProgramOwner          string  `json:"programOwner,omitempty"`
			TokenAddress          string  `json:"tokenAddress,omitempty"`
			InitAcount            string  `json:"initAcount,omitempty"`
			Owner                 string  `json:"owner,omitempty"`
			Delegate              string  `json:"delegate,omitempty"`
			Amount                string  `json:"amount,omitempty"`
			TokenSwap             string  `json:"Token Swap,omitempty"`
			Authority             string  `json:"Authority,omitempty"`
			UserTransferAuthority string  `json:"User Transfer Authority,omitempty"`
			UserSource            string  `json:"User Source,omitempty"`
			PoolSource            string  `json:"Pool Source,omitempty"`
			PoolDestination       string  `json:"Pool Destination,omitempty"`
			UserDestination       string  `json:"User Destination,omitempty"`
			PoolMint              string  `json:"Pool Mint,omitempty"`
			FeeAccount            string  `json:"Fee Account,omitempty"`
			TOKENPROGRAMID        string  `json:"TOKEN_PROGRAM_ID,omitempty"`
			AmountIn              int64   `json:"Amount In,omitempty"`
			MinimumAmountOut      int     `json:"Minimum Amount Out,omitempty"`
			ClosedAccount         string  `json:"closedAccount,omitempty"`
		} `json:"params"`
	} `json:"parsedInstruction"`
	Confirmations       interface{}   `json:"confirmations"`
	Version             string        `json:"version"`
	TokenTransfers      []interface{} `json:"tokenTransfers"`
	SolTransfers        []interface{} `json:"solTransfers"`
	SerumTransactions   []interface{} `json:"serumTransactions"`
	RaydiumTransactions []interface{} `json:"raydiumTransactions"`
	UnknownTransfers    []struct {
		ProgramId string `json:"programId"`
		Event     []struct {
			Source           string `json:"source"`
			Destination      string `json:"destination"`
			Amount           string `json:"amount"`
			Type             string `json:"type"`
			TokenAddress     string `json:"tokenAddress"`
			Decimals         int    `json:"decimals"`
			Symbol           string `json:"symbol"`
			Icon             string `json:"icon"`
			DestinationOwner string `json:"destinationOwner"`
			SourceOwner      string `json:"sourceOwner,omitempty"`
		} `json:"event"`
	} `json:"unknownTransfers"`
}
