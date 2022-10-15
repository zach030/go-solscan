package solscan

import (
	"log"

	"github.com/zach030/go-solscan/model"
)

//ChainInfo Blockchain overall information
func (c *Client) ChainInfo() (*model.ChainInfo, error) {
	resp := &model.ChainInfo{}
	_, err := c.client.R().SetResult(resp).Get(model.SolscanChainInfoURL)
	if err != nil {
		log.Println("ChainInfo failed,err=", err)
		return nil, err
	}
	return resp, nil
}
