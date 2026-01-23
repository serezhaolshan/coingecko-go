package coingecko

import (
	"context"
	"fmt"
	"net/url"
)

type MarketChartParams struct {
	ID         string
	VsCurrency string
	Days       string
	Interval   string
	Precision  string
}

type MarketChartDataPoint struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}

type MarketChartResponse struct {
	Prices      []MarketChartDataPoint `json:"prices"`
	MarketCaps  []MarketChartDataPoint `json:"market_caps"`
	TotalVolumes []MarketChartDataPoint `json:"total_volumes"`
}

func (c *Client) GetMarketChart(ctx context.Context, params MarketChartParams) (*MarketChartResponse, error) {
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

	endpoint := fmt.Sprintf("/coins/%s/market_chart", params.ID)
	request := NewRequest(c.Url(endpoint), queryParams)

	var rawResponse struct {
		Prices       [][]float64 `json:"prices"`
		MarketCaps   [][]float64 `json:"market_caps"`
		TotalVolumes [][]float64 `json:"total_volumes"`
	}
	_, err := c.doCall(ctx, request, &rawResponse)
	if err != nil {
		return nil, err
	}

	response := &MarketChartResponse{
		Prices:       make([]MarketChartDataPoint, len(rawResponse.Prices)),
		MarketCaps:   make([]MarketChartDataPoint, len(rawResponse.MarketCaps)),
		TotalVolumes: make([]MarketChartDataPoint, len(rawResponse.TotalVolumes)),
	}

	for i, raw := range rawResponse.Prices {
		if len(raw) < 2 {
			continue
		}
		response.Prices[i] = MarketChartDataPoint{
			Timestamp: int64(raw[0]),
			Value:     raw[1],
		}
	}

	for i, raw := range rawResponse.MarketCaps {
		if len(raw) < 2 {
			continue
		}
		response.MarketCaps[i] = MarketChartDataPoint{
			Timestamp: int64(raw[0]),
			Value:     raw[1],
		}
	}

	for i, raw := range rawResponse.TotalVolumes {
		if len(raw) < 2 {
			continue
		}
		response.TotalVolumes[i] = MarketChartDataPoint{
			Timestamp: int64(raw[0]),
			Value:     raw[1],
		}
	}

	return response, nil
}