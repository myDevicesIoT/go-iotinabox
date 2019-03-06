package iotinabox

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	apiUrl    = "https://api.iotinabox.com/"
	authUrl   = "https://accounts.iotinabox.io"
	userAgent = "go-tina-client"
)

type Client struct {
	client *http.Client

	ApiUrl *url.URL

	Devices *DeviceService
	//Company *CompanyService
	Locations *LocationService
	//Account *AccountService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(apiUrl)

	c := &Client{client: httpClient, ApiUrl: baseUrl}
	c.Devices = &DeviceService{client: c}
	c.Locations = &LocationService{client: c}

	return c
}

func (c *Client) NewRequest(method, pathString string, body interface{}) (*http.Request, error) {
	u, err := c.ApiUrl.Parse(pathString)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}
	log.Println(u.String())
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	log.Println(resp.Body)

	decErr := json.NewDecoder(resp.Body).Decode(v)

	if decErr == io.EOF {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		err = decErr
	}

	return resp, err
}

type Response struct {
	*http.Response
}
