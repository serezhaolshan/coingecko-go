package coingecko

import (
	"context"
	"net/url"
)

type DerivativesTicker struct {
	Market                   string   `json:"market"`
	Symbol                   string   `json:"symbol"`
	IndexID                  string   `json:"index_id"`
	Price                    string   `json:"price"`
	PricePercentageChange24h float64  `json:"price_percentage_change_24h"`
	ContractType             string   `json:"contract_type"`
	Index                    *float64 `json:"index"`
	Basis                    float64  `json:"basis"`
	Spread                   *float64 `json:"spread"`
	FundingRate              float64  `json:"funding_rate"`
	OpenInterest             *float64 `json:"open_interest"`
	Volume24h                *float64 `json:"volume_24h"`
	LastTradedAt             int64    `json:"last_traded_at"`
	ExpiredAt                *int64   `json:"expired_at"`
}

func (c *Client) GetDerivativesTickers(ctx context.Context) ([]DerivativesTicker, error) {
	request := NewRequest(c.Url("/derivatives"), url.Values{})
	var response []DerivativesTicker
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetDerivativesTickersByIndexID(ctx context.Context, indexID string) ([]DerivativesTicker, error) {
	tickers, err := c.GetDerivativesTickers(ctx)
	if err != nil {
		return nil, err
	}

	var filtered []DerivativesTicker
	for _, ticker := range tickers {
		if ticker.IndexID == indexID {
			filtered = append(filtered, ticker)
		}
	}
	return filtered, nil
}

type AggregatedDerivativesData struct {
	IndexID            string  `json:"index_id"`
	TotalOpenInterest  float64 `json:"total_open_interest"`
	TotalVolume24h     float64 `json:"total_volume_24h"`
	AvgFundingRate     float64 `json:"avg_funding_rate"`
	TickerCount        int     `json:"ticker_count"`
	PerpetualCount     int     `json:"perpetual_count"`
	FuturesCount       int     `json:"futures_count"`
}

func (c *Client) GetAggregatedDerivativesData(ctx context.Context, indexID string) (*AggregatedDerivativesData, error) {
	tickers, err := c.GetDerivativesTickersByIndexID(ctx, indexID)
	if err != nil {
		return nil, err
	}

	if len(tickers) == 0 {
		return nil, nil
	}

	result := &AggregatedDerivativesData{
		IndexID: indexID,
	}

	var fundingRateSum float64
	var fundingRateCount int

	for _, ticker := range tickers {
		if ticker.OpenInterest != nil {
			result.TotalOpenInterest += *ticker.OpenInterest
		}
		if ticker.Volume24h != nil {
			result.TotalVolume24h += *ticker.Volume24h
		}
		if ticker.FundingRate != 0 {
			fundingRateSum += ticker.FundingRate
			fundingRateCount++
		}
		result.TickerCount++

		if ticker.ContractType == "perpetual" {
			result.PerpetualCount++
		} else {
			result.FuturesCount++
		}
	}

	if fundingRateCount > 0 {
		result.AvgFundingRate = fundingRateSum / float64(fundingRateCount)
	}

	return result, nil
}
