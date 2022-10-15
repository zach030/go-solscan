package solscan

import (
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
