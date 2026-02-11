package coingecko

import (
	"context"
	"net/url"
)

type SearchParams struct {
	Query string
}

type SearchResponse struct {
	Coins      []SearchCoin     `json:"coins"`
	Exchanges  []SearchExchange `json:"exchanges"`
	Categories []SearchCategory `json:"categories"`
	NFTs       []SearchNFT      `json:"nfts"`
}

type SearchCoin struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	APISymbol     string `json:"api_symbol"`
	Symbol        string `json:"symbol"`
	MarketCapRank *int   `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

type SearchExchange struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MarketType string `json:"market_type"`
	Thumb      string `json:"thumb"`
	Large      string `json:"large"`
}

type SearchCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SearchNFT struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Thumb  string `json:"thumb"`
}

func (c *Client) Search(ctx context.Context, params SearchParams) (*SearchResponse, error) {
	q := url.Values{}
	if params.Query != "" {
		q.Set("query", params.Query)
	}

	var resp SearchResponse
	_, err := c.doCall(ctx, NewRequest(c.Url("/search"), q), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}