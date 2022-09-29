package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/TovarischSuhov/go-template/internal/domain"
	"github.com/TovarischSuhov/log"
)

type HTTPClient struct {
	host        string
	path        string
	accessToken string
	consumerKey string
	client      *http.Client
}

func NewHTTPClient(consumerKey, accessToken, host, path string) Client {
	client := HTTPClient{
		host:        host,
		path:        path,
		consumerKey: consumerKey,
		accessToken: accessToken,
		client:      &http.Client{},
	}
	return &client
}

func (c *HTTPClient) GetTopicsList() (*domain.Response, error) {
	requestURL := url.URL{
		Scheme:   "https",
		Host:     c.host,
		Path:     c.path,
		RawQuery: fmt.Sprintf("access_token=%s&consumer_key=%s&state=all", c.accessToken, c.consumerKey),
	}
	rawURL := requestURL.String()
	log.Debug(rawURL)
	request, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respStruct domain.Response
	err = json.Unmarshal(body, &respStruct)
	return &respStruct, err
}
