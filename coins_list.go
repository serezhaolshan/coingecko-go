package coingecko

import (
	"context"
	"net/url"
)

type CoinsListParams struct {
	IncludePlatform bool
	Status          string
}

type CoinsListItem struct {
	ID        string            `json:"id"`
	Symbol    string            `json:"symbol"`
	Name      string            `json:"name"`
	Platforms map[string]string `json:"platforms,omitempty"`
}

func (c *Client) GetCoinsList(ctx context.Context, params *CoinsListParams) ([]CoinsListItem, error) {
	queryParams := url.Values{}

	if params != nil {
		if params.IncludePlatform {
			queryParams.Set("include_platform", "true")
		}
		if params.Status != "" {
			queryParams.Set("status", params.Status)
		}
	}

	request := NewRequest(c.Url("/coins/list"), queryParams)
	var response []CoinsListItem
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
