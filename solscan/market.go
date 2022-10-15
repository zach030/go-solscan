package solscan

import (
	"fmt"
	"log"

	"github.com/zach030/go-solscan/model"
)

func (c *Client) GetTokenMarket(address string) (*model.TokenMarket, error) {
	resp := &model.TokenMarket{}
	_, err := c.client.R().SetResult(resp).Get(fmt.Sprintf("%s/%s", model.SolscanMarketTokenURL, address))
	if err != nil {
		log.Printf("GetTokenMarket failed, param=%+v,err=%v\n", address, err)
		return nil, err
	}
	return resp, nil
}
