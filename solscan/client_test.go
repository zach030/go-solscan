package solscan

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	client := New(SetMaxRetryCount(3), SetRetryWaitTime(3*time.Second))
	assert.Equal(t, 3, client.maxRetryCount)
	assert.Equal(t, 3*time.Second, client.retryWaitTime)
}

func TestQueryBlockDetail(t *testing.T) {
	var blockNumber = 155457187
	var blockHeight = 140436931
	client := New(SetMaxRetryCount(3), SetRetryWaitTime(3*time.Second))
	block, err := client.GetBlockDetail(blockNumber)
	assert.Nil(t, err)
	assert.Equal(t, block.Result.BlockHeight, blockHeight)
}

func TestChainInfo(t *testing.T) {
	client := New(SetMaxRetryCount(3), SetRetryWaitTime(3*time.Second))
	chain, err := client.ChainInfo()
	assert.Nil(t, err)
	fmt.Printf("%+v", chain)
}
