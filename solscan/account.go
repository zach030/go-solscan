package solscan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zach030/go-solscan/model"
)

//GetAccountTokens get token balances in given account
func (c *Client) GetAccountTokens(account string) ([]*model.AccountTokens, error) {
	resp := make([]*model.AccountTokens, 0)
	_, err := c.client.R().SetQueryParam("account", account).SetResult(resp).Get(model.SolscanAccountTokensURL)
	if err != nil {
		log.Printf("GetAccountTokens failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}

//GetAccountTxs get transactions in given account with query parameters
func (c *Client) GetAccountTxs(account, beforeHash string, limit int) ([]*model.AccountTx, error) {
	if limit > model.TxQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := make([]*model.AccountTx, 0)
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("account", account).
		SetQueryParam("beforeHash", beforeHash).
		SetQueryParam("limit", strconv.Itoa(limit)).
		Get(model.SolscanAccountTxsURL)
	if err != nil {
		log.Printf("GetAccountTxs failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetAccountStakeAccounts(account string) (*model.AccountStakeInfo, error) {
	// todo fix return model
	resp := &model.AccountStakeInfo{}
	_, err := c.client.R().SetResult(resp).SetQueryParam("account", account).Get(model.SolscanAccountStakeAccountsURL)
	if err != nil {
		log.Printf("GetAccountStakeAccounts failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}

// GetAccountSplTransfers Get list of transactions make tokenBalance changes. MaxLimit 50 records per request
func (c *Client) GetAccountSplTransfers(account string, limit, offset int, fromTime, toTime int) (*model.GetAccountSplTransfersResponse, error) {
	if limit > model.TxQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := &model.GetAccountSplTransfersResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("account", account).
		SetQueryParam("fromTime", strconv.Itoa(fromTime)).
		SetQueryParam("toTime", strconv.Itoa(toTime)).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("offset", strconv.Itoa(offset)).
		Get(model.SolscanAccountSplTransfersURL)
	if err != nil {
		log.Printf("GetAccountSplTransfers failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}

//GetAccountSolTransfers Get list of SOL transfers. MaxLimit 50 records per request
func (c *Client) GetAccountSolTransfers(account string, limit, offset int, fromTime, toTime int) (*model.GetAccountSolTransfersResponse, error) {
	if limit > model.TxQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := &model.GetAccountSolTransfersResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("account", account).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("offset", strconv.Itoa(offset)).
		SetQueryParam("fromTime", strconv.Itoa(fromTime)).
		SetQueryParam("toTime", strconv.Itoa(toTime)).
		Get(model.SolscanAccountSolTransfersURL)
	if err != nil {
		log.Printf("GetAccountSolTransfers failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}

//GetAccountExportTransactions Export transactions to CSV. MaxLimit 5000 records per request
func (c *Client) GetAccountExportTransactions(account, typ string, fromTime, toTime int) error {
	if typ != model.ExportTypeSolTransfer && typ != model.ExportTypeAll && typ != model.ExportTypeTokenChange {
		return model.ErrInvalidExportType
	}
	_, err := c.client.R().
		SetQueryParam("account", account).
		SetQueryParam("type", typ).
		SetQueryParam("fromTime", strconv.Itoa(fromTime)).
		SetQueryParam("toTime", strconv.Itoa(toTime)).
		Get(model.SolscanAccountExportTransactionsURL)
	if err != nil {
		log.Printf("GetAccountExportTransactions failed, param=%+v,err=%v\n", account, err)
		return err
	}
	return nil
}

//GetAccountInfo Get overall account information, including program account, NFT metadata information
func (c *Client) GetAccountInfo(account string) (*model.GetAccountInfoResponse, error) {
	resp := &model.GetAccountInfoResponse{}
	_, err := c.client.R().SetResult(resp).Get(fmt.Sprintf("%s/%s", model.SolscanAccountInfoURL, account))
	if err != nil {
		log.Printf("GetAccountInfo failed, param=%+v,err=%v\n", account, err)
		return nil, err
	}
	return resp, nil
}
