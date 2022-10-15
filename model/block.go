package model

type GetBlockResponse struct {
	CurrentSlot int   `json:"currentSlot"`
	Result      Block `json:"result"`
}

type Block struct {
	BlockHeight       int    `json:"blockHeight"`
	BlockTime         int    `json:"blockTime"`
	BlockHash         string `json:"blockhash"`
	ParentSlot        int    `json:"parentSlot"`
	PreviousBlockHash string `json:"previousBlockhash"`
	FeeRewards        int    `json:"feeRewards"`
	Validator         string `json:"validator"`
	TransactionCount  int    `json:"transactionCount"`
}

type GetBlockTxsResponse struct {
	Meta struct {
		Err *struct {
			InstructionError []interface{} `json:"InstructionError"`
		} `json:"err"`
		Fee               int `json:"fee"`
		InnerInstructions []struct {
			Index        int `json:"index"`
			Instructions []struct {
				Accounts  []string `json:"accounts,omitempty"`
				Data      string   `json:"data,omitempty"`
				ProgramId string   `json:"programId"`
				Parsed    struct {
					Info struct {
						Amount      string `json:"amount"`
						Authority   string `json:"authority"`
						Destination string `json:"destination"`
						Source      string `json:"source"`
					} `json:"info"`
					Type string `json:"type"`
				} `json:"parsed,omitempty"`
				Program string `json:"program,omitempty"`
			} `json:"instructions"`
		} `json:"innerInstructions"`
		LogMessages       []string `json:"logMessages"`
		PostBalances      []int64  `json:"postBalances"`
		PostTokenBalances []struct {
			AccountIndex  int    `json:"accountIndex"`
			Mint          string `json:"mint"`
			Owner         string `json:"owner"`
			ProgramId     string `json:"programId"`
			UiTokenAmount struct {
				Amount         string   `json:"amount"`
				Decimals       int      `json:"decimals"`
				UiAmount       *float64 `json:"uiAmount"`
				UiAmountString string   `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"postTokenBalances"`
		PreBalances      []int64 `json:"preBalances"`
		PreTokenBalances []struct {
			AccountIndex  int    `json:"accountIndex"`
			Mint          string `json:"mint"`
			Owner         string `json:"owner"`
			ProgramId     string `json:"programId"`
			UiTokenAmount struct {
				Amount         string   `json:"amount"`
				Decimals       int      `json:"decimals"`
				UiAmount       *float64 `json:"uiAmount"`
				UiAmountString string   `json:"uiAmountString"`
			} `json:"uiTokenAmount"`
		} `json:"preTokenBalances"`
		Rewards interface{} `json:"rewards"`
		Status  struct {
			Ok  interface{} `json:"Ok"`
			Err struct {
				InstructionError []interface{} `json:"InstructionError"`
			} `json:"Err,omitempty"`
		} `json:"status"`
	} `json:"meta"`
	Transaction struct {
		Message struct {
			AccountKeys []struct {
				Pubkey   string `json:"pubkey"`
				Signer   bool   `json:"signer"`
				Source   string `json:"source"`
				Writable bool   `json:"writable"`
			} `json:"accountKeys"`
			AddressTableLookups interface{} `json:"addressTableLookups"`
			Instructions        []struct {
				Parsed struct {
					Info struct {
						ClockSysvar      string `json:"clockSysvar"`
						SlotHashesSysvar string `json:"slotHashesSysvar"`
						Vote             struct {
							Hash      string `json:"hash"`
							Slots     []int  `json:"slots"`
							Timestamp int    `json:"timestamp"`
						} `json:"vote"`
						VoteAccount   string `json:"voteAccount"`
						VoteAuthority string `json:"voteAuthority"`
					} `json:"info"`
					Type string `json:"type"`
				} `json:"parsed,omitempty"`
				Program   string   `json:"program,omitempty"`
				ProgramId string   `json:"programId"`
				Accounts  []string `json:"accounts,omitempty"`
				Data      string   `json:"data,omitempty"`
			} `json:"instructions"`
			RecentBlockhash string `json:"recentBlockhash"`
		} `json:"message"`
		Signatures []string `json:"signatures"`
	} `json:"transaction"`
	Version string `json:"version"`
}

type GetBlockDetailResponse struct {
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
				Source      string `json:"source"`
				Destination string `json:"destination"`
				Amount      int    `json:"amount"`
			} `json:"params"`
		} `json:"parsedInstructions"`
	} `json:"innerInstructions"`
	TokenBalanes      []interface{} `json:"tokenBalanes"`
	ParsedInstruction []struct {
		ProgramId  string `json:"programId"`
		Type       string `json:"type"`
		Data       string `json:"data"`
		DataEncode string `json:"dataEncode"`
		Name       string `json:"name"`
		Params     struct {
			Account0 string `json:"Account0"`
			Account1 string `json:"Account1"`
			Account2 string `json:"Account2"`
			Account3 string `json:"Account3"`
			Account4 string `json:"Account4"`
			Account5 string `json:"Account5"`
			Account6 string `json:"Account6"`
		} `json:"params"`
	} `json:"parsedInstruction"`
	Confirmations       int           `json:"confirmations"`
	TokenTransfers      []interface{} `json:"tokenTransfers"`
	SolTransfers        []interface{} `json:"solTransfers"`
	SerumTransactions   []interface{} `json:"serumTransactions"`
	RaydiumTransactions []interface{} `json:"raydiumTransactions"`
	UnknownTransfers    []struct {
		ProgramId string `json:"programId"`
		Event     []struct {
			Source       string `json:"source"`
			Destination  string `json:"destination"`
			Amount       int    `json:"amount"`
			Type         string `json:"type"`
			Decimals     int    `json:"decimals"`
			Symbol       string `json:"symbol"`
			TokenAddress string `json:"tokenAddress"`
		} `json:"event"`
	} `json:"unknownTransfers"`
}
