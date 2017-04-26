package a3rt

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
)

const (
	DefaultBaseUri = "https://api.a3rt.recruit-tech.co.jp"
)

var (
	ErrNoApiKey = errors.New("no API key")
)

type Config struct {
	BaseUri *url.URL
	ApiKey  string
}

type Client struct {
	*http.Client
	config *Config
}

func NewClient() *Client {
	u, _ := url.Parse(DefaultBaseUri)

	return &Client{
		Client: http.DefaultClient,
		config: &Config{
			BaseUri: u,
		},
	}
}

func (client *Client) SetBaseUri(baseUri string) error {
	u, err := url.Parse(baseUri)
	if err != nil {
		return err
	}

	client.config.BaseUri = u

	return nil
}

func (client *Client) SetApiKey(apiKey string) {
	client.config.ApiKey = apiKey
}

func (client *Client) doApi(ctx context.Context, method string, uri string, params url.Values, res interface{}) error {
	if len(client.config.ApiKey) == 0 {
		return ErrNoApiKey
	}

	u := client.config.BaseUri
	u.Path = path.Join(u.Path, uri)

	params.Add("apikey", client.config.ApiKey)

	var body io.Reader
	switch method {
	case http.MethodGet:
		u.RawQuery = params.Encode()
	default:
		body = strings.NewReader(params.Encode())
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if res == nil {
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(&res)
}
