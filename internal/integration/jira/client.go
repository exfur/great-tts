package jira

import "net/http"

type Client struct {
	baseURL    string
	httpClient *http.Client
	authHeader string
}

func NewClient(baseURL, user, token string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
		// authHeader generation logic here...
	}
}
