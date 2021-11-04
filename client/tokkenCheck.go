package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

// TokenBaseClient - http client for alert service
type TokenBaseClient struct {
	Client *http.Client
}

// NewTokenBaseClient - initializes new client
func NewTokenBaseClient() *TokenBaseClient {
	return &TokenBaseClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{Timeout: 30 * time.Second}).DialContext,
			},
		},
	}
}

// GetValidate - validate the token
func (r *TokenBaseClient) GetValidate(token string) int {
	authBaseUrl := "http://192.168.0.110:3002/validateToken"
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, authBaseUrl, nil)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return 500
	}
	req.Header.Add("Authorization", token)
	resp, err := r.Client.Do(req)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return 500
	}
	// use resp to get your desired data
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return 200
	} else if resp.StatusCode == http.StatusUnauthorized {
		return 401
	} else {
		return 500
	}
	return 500
}
