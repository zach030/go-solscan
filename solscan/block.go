package solscan

import (
	"fmt"
	"log"
	"strconv"

	"github.com/zach030/go-solscan/model"
)

//GetBlockDetail get detail information of given block
func (c *Client) GetBlockDetail(blockNumber int) (*model.GetBlockResponse, error) {
	resp := &model.GetBlockResponse{}
	_, err := c.client.R().SetResult(resp).Get(fmt.Sprintf("%s/%d", model.SolscanBlockDetailURL, blockNumber))
	if err != nil {
		log.Printf("GetBlockDetail failed, param=%+v,err=%v\n", blockNumber, err)
		return nil, err
	}
	return resp, err
}

//GetBlockTxs get transactions in one block, max limit 50 records per request
func (c *Client) GetBlockTxs(blockNumber string, offset, limit int) (*model.GetBlockTxsResponse, error) {
	if limit > model.BlockTxQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := &model.GetBlockTxsResponse{}
	_, err := c.client.R().SetResult(resp).
		SetQueryParam("block", blockNumber).
		SetQueryParam("offset", strconv.Itoa(offset)).
		SetQueryParam("limit", strconv.Itoa(limit)).
		Get(model.SolscanBlockTxsURL)
	if err != nil {
		log.Printf("GetBlockTxs failed, param=%+v,err=%v\n", blockNumber, err)
		return nil, err
	}
	return resp, nil
}

//GetLastBlocks get last limit blocks
func (c *Client) GetLastBlocks(limit int) ([]*model.GetBlockResponse, error) {
	if limit > model.LastQueryLimit {
		return nil, model.ErrExceedQueryLimit
	}
	resp := make([]*model.GetBlockResponse, 0)
	_, err := c.client.R().SetQueryParam("limit", strconv.Itoa(limit)).SetResult(resp).Get(model.SolscanBlocksURL)
	if err != nil {
		log.Printf("GetLastBlocks failed, param=%+v,err=%v\n", limit, err)
		return nil, err
	}
	return resp, nil
}
