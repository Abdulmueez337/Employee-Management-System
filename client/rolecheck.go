package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

// RollBaseClient - http client for alert service
type RollBaseClient struct {
	Client *http.Client
}

// NewRollBaseClient - initializes new client
func NewRollBaseClient() *RollBaseClient {
	return &RollBaseClient{
		Client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{Timeout: 30 * time.Second}).DialContext,
			},
		},
	}
}

// GetRole - gets role from the roll base service
func (r *RollBaseClient) GetRole(designation, role string) int {
	fmt.Println("Welcome to role check")
	roleBaseUrl := "http://192.168.100.165:3000/v1/role"
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, roleBaseUrl, nil)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return 500
	}
	req.Header.Set("Content-Type", "application/json")
	roleBaseUrl = roleBaseUrl + "/" + designation + "/" + role
	fmt.Println("querry",roleBaseUrl)
	requestURL, err := url.Parse(roleBaseUrl)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return 500
	}
	resp, err := r.Client.Do(&http.Request{
		Method: http.MethodGet,
		URL:    requestURL,
	})
	if err != nil {
		fmt.Errorf("error: %v", err)
		return 500
	}
	// use resp to get your desired data
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("User is Authorized")
		return 200
	} else if resp.StatusCode == http.StatusBadRequest {
		fmt.Println("User is Not Authorized")
		return 400
	} else if resp.StatusCode == http.StatusForbidden {
		fmt.Println("RoleCheck > Bad Request")
		return 403
	} else {
		fmt.Println("RoleCheck > INTERNAL SERVER ERROR")
		return 500
	}
	return 500
}
