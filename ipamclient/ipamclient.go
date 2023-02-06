package xtratusipamclient

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient - Construct a new HTTP Client to interact with the APIM REST API
func NewClient(host, authToken *string) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // <--- Problem
	}
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: tr},
	}

	// set client values, if provided
	if host != nil {
		c.HostURL = *host
	}
	if authToken != nil {
		c.Token = *authToken
	}

	return &c, nil
}

// doRequest -
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	//perform request
	req.Header.Add("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	//read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//write error not StatusOK
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
