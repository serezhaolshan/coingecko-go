package coingecko

import (
	"context"
	"net/url"
)

type GlobalData struct {
	ActiveCryptocurrencies          int                `json:"active_cryptocurrencies"`
	UpcomingICOs                    int                `json:"upcoming_icos"`
	OngoingICOs                     int                `json:"ongoing_icos"`
	EndedICOs                       int                `json:"ended_icos"`
	Markets                         int                `json:"markets"`
	TotalMarketCap                  map[string]float64 `json:"total_market_cap"`
	TotalVolume                     map[string]float64 `json:"total_volume"`
	MarketCapPercentage             map[string]float64 `json:"market_cap_percentage"`
	MarketCapChangePercentage24hUSD float64            `json:"market_cap_change_percentage_24h_usd"`
	UpdatedAt                       int64              `json:"updated_at"`
}

type GlobalResponse struct {
	Data GlobalData `json:"data"`
}

func (c *Client) GetGlobal(ctx context.Context) (*GlobalData, error) {
	request := NewRequest(c.Url("/global"), url.Values{})
	var response GlobalResponse
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}