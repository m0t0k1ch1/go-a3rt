package a3rt

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type SmallTalkResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Results []*SmallTalkResult `json:"results"`
}

type SmallTalkResult struct {
	Perplexity float64 `json:"perplexity"`
	Reply      string  `json:"reply"`
}

func (client *Client) SmallTalk(ctx context.Context, query string) (*SmallTalkResult, error) {
	v := url.Values{}
	v.Add("query", query)

	var resp SmallTalkResponse
	if err := client.doApi(ctx, http.MethodPost, "talk/v1/smalltalk", v, &resp); err != nil {
		return nil, err
	}

	if resp.Status != 0 {
		return nil, errors.New(fmt.Sprintf("%d: %s", resp.Status, resp.Message))
	}

	return resp.Results[0], nil
}
