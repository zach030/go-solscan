package solscan

import (
	"log"
	"strconv"

	"github.com/zach030/go-solscan/model"
)

//GetTokenHolders Get token holders
func (c *Client) GetTokenHolders(tokenAddr string, limit, offset int) (*model.GetTokenHoldersResponse, error) {
	if limit > model.HolderQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := &model.GetTokenHoldersResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("tokenAddress", tokenAddr).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("offset", strconv.Itoa(offset)).
		Get(model.SolscanTokenHoldersURL)
	if err != nil {
		log.Printf("GetTokenHolders failed, param=%+v,err=%v\n", tokenAddr, err)
		return nil, err
	}
	return resp, nil
}

//GetTokenMeta Get metadata of given token
func (c *Client) GetTokenMeta(tokenAddr string) (*model.GetTokenMetaResponse, error) {
	resp := &model.GetTokenMetaResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("tokenAddress", tokenAddr).
		Get(model.SolscanTokenMetaURL)
	if err != nil {
		log.Printf("GetTokenMeta failed, param=%+v,err=%v\n", tokenAddr, err)
		return nil, err
	}
	return resp, nil
}

//GetTokenList Get list of tokens. MaxLimit 50 records per request
// sortBy : arket_cap, volume, holder, price, price_change_24h, price_change_7d, price_change_14d, price_change_30d, price_change_60d, price_change_200d, price_change_1y
// direction: asc/desc
func (c *Client) GetTokenList(sortBy, direction string, limit, offset int) (*model.GetTokenListResponse, error) {
	if limit > model.TxQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := &model.GetTokenListResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("sortBy", sortBy).
		SetQueryParam("direction", direction).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetQueryParam("offset", strconv.Itoa(offset)).
		Get(model.SolscanTokenListURL)
	if err != nil {
		log.Printf("GetTokenList failed, param=%+v,err=%v\n", sortBy, err)
		return nil, err
	}
	return resp, nil
}
