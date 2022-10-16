package solscan

import (
	"fmt"
	"log"

	"github.com/zach030/go-solscan/model"
)

//GetTokenMarket Get market information of the given token
func (c *Client) GetTokenMarket(tokenAddr string) (*model.TokenMarket, error) {
	resp := &model.TokenMarket{}
	_, err := c.client.R().SetResult(resp).Get(fmt.Sprintf("%s/%s", model.SolscanMarketTokenURL, tokenAddr))
	if err != nil {
		log.Printf("GetTokenMarket failed, param=%+v,err=%v\n", tokenAddr, err)
		return nil, err
	}
	return resp, nil
}
