package coingecko

import (
	"context"
	"fmt"
	"net/url"
)

type CoinsOHLCParams struct {
	ID         string
	VsCurrency string
	Days       string
	Interval   string
	Precision  string
}

type OHLCCandle struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
}

func (c *Client) GetCoinsOHLC(ctx context.Context, params CoinsOHLCParams) ([]OHLCCandle, error) {
	if params.ID == "" {
		return nil, fmt.Errorf("coin ID is required")
	}
	if params.Days == "" {
		return nil, fmt.Errorf("days parameter is required")
	}

	if params.VsCurrency == "" {
		params.VsCurrency = "usd"
	}

	queryParams := url.Values{}
	queryParams.Set("vs_currency", params.VsCurrency)
	queryParams.Set("days", params.Days)

	if params.Interval != "" {
		queryParams.Set("interval", params.Interval)
	}
	if params.Precision != "" {
		queryParams.Set("precision", params.Precision)
	}

	endpoint := fmt.Sprintf("/coins/%s/ohlc", params.ID)
	request := NewRequest(c.Url(endpoint), queryParams)

	var rawResponse [][]float64
	_, err := c.doCall(ctx, request, &rawResponse)
	if err != nil {
		return nil, err
	}

	candles := make([]OHLCCandle, len(rawResponse))
	for i, raw := range rawResponse {
		if len(raw) < 5 {
			continue
		}
		candles[i] = OHLCCandle{
			Timestamp: int64(raw[0]),
			Open:      raw[1],
			High:      raw[2],
			Low:       raw[3],
			Close:     raw[4],
		}
	}

	return candles, nil
}
