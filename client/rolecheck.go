package client

import (
	"bytes"
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
type roleClientBody struct {
	designation string
	apiName     string
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
func (r *RollBaseClient) GetRole(designation, role string) error {
	roleBaseUrl := "http://192.168.0.232:3000/v1/role"
	// roleBaseUrl is the url at which you want to hit the API
	sender :=roleClientBody{designation: "Admin",apiName: "Add"}
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodPost, roleBaseUrl, bytes.NewBufferString(""))
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	requestURL, err := url.Parse(roleBaseUrl)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	resp, err := r.Client.Do(&http.Request{
		Method: http.MethodGet,
		URL:    requestURL,
	})
	fmt.Println("Received Response", resp)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return err
	}
	// use resp to get your desired data
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("failed to get expected response from role base service, got status code: %d", resp.StatusCode)
		return nil
	}
	fmt.Println("Role successfully got from role base service")
	return nil
}
