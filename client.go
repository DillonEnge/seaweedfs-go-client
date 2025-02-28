package seaweedfs

import "net/http"

type Client struct {
	httpClient *http.Client
	config     Config
}

func NewClient(config Config) *Client {
	return &Client{
		httpClient: &http.Client{},
		config:     config,
	}
}
