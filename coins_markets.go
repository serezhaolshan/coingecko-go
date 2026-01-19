package coingecko

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type CoinsMarketsParams struct {
	VsCurrency            string
	IDs                   []string
	Names                 []string
	Symbols               []string
	IncludeTokens         string
	Category              string
	Order                 string
	PerPage               int
	Page                  int
	Sparkline             bool
	PriceChangePercentage string
	Locale                string
	Precision             string
}

type CoinsMarketsItem struct {
	ID                           string         `json:"id"`
	Symbol                       string         `json:"symbol"`
	Name                         string         `json:"name"`
	Image                        string         `json:"image"`
	CurrentPrice                 *float64       `json:"current_price"`
	MarketCap                    *float64       `json:"market_cap"`
	MarketCapRank                *int           `json:"market_cap_rank"`
	FullyDilutedValuation        *float64       `json:"fully_diluted_valuation"`
	TotalVolume                  *float64       `json:"total_volume"`
	High24h                      *float64       `json:"high_24h"`
	Low24h                       *float64       `json:"low_24h"`
	PriceChange24h               *float64       `json:"price_change_24h"`
	PriceChangePercentage24h     *float64       `json:"price_change_percentage_24h"`
	MarketCapChange24h           *float64       `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h *float64       `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            *float64       `json:"circulating_supply"`
	TotalSupply                  *float64       `json:"total_supply"`
	MaxSupply                    *float64       `json:"max_supply"`
	ATH                          *float64       `json:"ath"`
	ATHChangePercentage          *float64       `json:"ath_change_percentage"`
	ATHDate                      string         `json:"ath_date"`
	ATL                          *float64       `json:"atl"`
	ATLChangePercentage          *float64       `json:"atl_change_percentage"`
	ATLDate                      string         `json:"atl_date"`
	ROI                          *MarketROI     `json:"roi"`
	LastUpdated                  string         `json:"last_updated"`
	SparklineIn7d                *SparklineIn7d `json:"sparkline_in_7d,omitempty"`
	PriceChangePercentage1h      *float64       `json:"price_change_percentage_1h_in_currency,omitempty"`
	PriceChangePercentage7d      *float64       `json:"price_change_percentage_7d_in_currency,omitempty"`
	PriceChangePercentage14d     *float64       `json:"price_change_percentage_14d_in_currency,omitempty"`
	PriceChangePercentage30d     *float64       `json:"price_change_percentage_30d_in_currency,omitempty"`
	PriceChangePercentage200d    *float64       `json:"price_change_percentage_200d_in_currency,omitempty"`
	PriceChangePercentage1y      *float64       `json:"price_change_percentage_1y_in_currency,omitempty"`
}

type MarketROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type SparklineIn7d struct {
	Price []float64 `json:"price"`
}

func (c *Client) GetCoinsMarkets(ctx context.Context, params CoinsMarketsParams) ([]CoinsMarketsItem, error) {
	queryParams := url.Values{}

	if params.VsCurrency == "" {
		params.VsCurrency = "usd"
	}
	queryParams.Set("vs_currency", params.VsCurrency)

	if len(params.IDs) > 0 {
		queryParams.Set("ids", strings.Join(params.IDs, ","))
	}
	if len(params.Names) > 0 {
		queryParams.Set("names", strings.Join(params.Names, ","))
	}
	if len(params.Symbols) > 0 {
		queryParams.Set("symbols", strings.Join(params.Symbols, ","))
	}
	if params.IncludeTokens != "" {
		queryParams.Set("include_tokens", params.IncludeTokens)
	}
	if params.Category != "" {
		queryParams.Set("category", params.Category)
	}
	if params.Order != "" {
		queryParams.Set("order", params.Order)
	}
	if params.PerPage > 0 {
		queryParams.Set("per_page", fmt.Sprintf("%d", params.PerPage))
	}
	if params.Page > 0 {
		queryParams.Set("page", fmt.Sprintf("%d", params.Page))
	}
	if params.Sparkline {
		queryParams.Set("sparkline", "true")
	}
	if params.PriceChangePercentage != "" {
		queryParams.Set("price_change_percentage", params.PriceChangePercentage)
	}
	if params.Locale != "" {
		queryParams.Set("locale", params.Locale)
	}
	if params.Precision != "" {
		queryParams.Set("precision", params.Precision)
	}

	request := NewRequest(c.Url("/coins/markets"), queryParams)
	var response []CoinsMarketsItem
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
