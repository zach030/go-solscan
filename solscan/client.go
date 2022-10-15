package solscan

import (
	"time"

	"github.com/go-resty/resty/v2"
)

type (
	Client struct {
		client        *resty.Client
		maxRetryCount int
		retryWaitTime time.Duration
	}
	ClientOption func(*Client)
)

func New(options ...ClientOption) *Client {
	c := &Client{
		client: resty.New(),
	}
	for _, option := range options {
		option(c)
	}
	return c
}

func SetMaxRetryCount(maxRetryCount int) ClientOption {
	return func(client *Client) {
		client.maxRetryCount = maxRetryCount
		client.client.SetRetryCount(maxRetryCount)
	}
}

func SetRetryWaitTime(retryWaitTime time.Duration) ClientOption {
	return func(client *Client) {
		client.retryWaitTime = retryWaitTime
		client.client.SetRetryWaitTime(retryWaitTime)
	}
}

func SetRetryCondition(retryCond resty.RetryConditionFunc) ClientOption {
	return func(client *Client) {
		client.client.AddRetryCondition(retryCond)
	}
}
