package solscan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zach030/go-solscan/model"
)

//GetLastTxs get last limit transactions
func (c *Client) GetLastTxs(limit int) ([]*model.GetTxResponse, error) {
	if limit > model.LastQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := make([]*model.GetTxResponse, 0)
	_, err := c.client.R().SetResult(resp).SetQueryParam("limit", strconv.Itoa(limit)).Get(model.SolscanLastTxsURL)
	if err != nil {
		log.Printf("GetLastTxs failed, param=%+v,err=%v\n", limit, err)
		return nil, err
	}
	return resp, nil
}

//GetTxWithSignature get transaction detail by signature
func (c *Client) GetTxWithSignature(sig string) (*model.GetTxResponse, error) {
	resp := &model.GetTxResponse{}
	_, err := c.client.R().SetResult(resp).Get(fmt.Sprintf("%s/%s", model.SolscanTxDetailURL, sig))
	if err != nil {
		log.Printf("GetTxWithSignature failed, param=%+v,err=%v\n", sig, err)
		return nil, err
	}
	return resp, nil
}
