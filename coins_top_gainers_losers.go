package coingecko

import (
	"context"
	"net/url"
)

type TopGainersLosersParams struct {
	VsCurrency            string
	Duration              Duration
	PriceChangePercentage PriceChangePercentage
	TopCoins              TopCoins
}

type TopGainerLoserCoin struct {
	ID            string   `json:"id"`
	Symbol        string   `json:"symbol"`
	Name          string   `json:"name"`
	Image         string   `json:"image"`
	MarketCapRank int      `json:"market_cap_rank"`
	Usd           float64  `json:"usd"`
	Usd24hVol     float64  `json:"usd_24h_vol"`
	Usd24hChange  *float64 `json:"usd_24h_change,omitempty"`
}

type TopGainersLosersResponse struct {
	TopGainers []TopGainerLoserCoin `json:"top_gainers"`
	TopLosers  []TopGainerLoserCoin `json:"top_losers"`
}

func (c *Client) GetTopGainersLosers(ctx context.Context, params TopGainersLosersParams) (*TopGainersLosersResponse, error) {
	queryParams := url.Values{}

	vsCurrency := params.VsCurrency
	if vsCurrency == "" {
		vsCurrency = "usd"
	}
	queryParams.Set("vs_currency", vsCurrency)

	if params.Duration != "" {
		queryParams.Set("duration", string(params.Duration))
	}

	if params.PriceChangePercentage != "" {
		queryParams.Set("price_change_percentage", string(params.PriceChangePercentage))
	}

	if params.TopCoins != "" {
		queryParams.Set("top_coins", string(params.TopCoins))
	}

	request := NewRequest(c.Url("/coins/top_gainers_losers"), queryParams)
	var response TopGainersLosersResponse
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
