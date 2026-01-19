package coingecko

import (
	"context"
	"net/url"
	"strings"
)

type SimplePriceParams struct {
	IDs                  []string
	Names                []string
	Symbols              []string
	VsCurrencies         []string
	IncludeTokens        string
	IncludeMarketCap     bool
	Include24hrVol       bool
	Include24hrChange    bool
	IncludeLastUpdatedAt bool
}

type SimplePriceResponse map[string]map[string]any

func (c *Client) GetSimplePrice(ctx context.Context, params SimplePriceParams) (SimplePriceResponse, error) {
	queryParams := url.Values{}

	if len(params.VsCurrencies) == 0 {
		params.VsCurrencies = []string{"usd"}
	}
	queryParams.Set("vs_currencies", strings.Join(params.VsCurrencies, ","))

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

	if params.IncludeMarketCap {
		queryParams.Set("include_market_cap", "true")
	}

	if params.Include24hrVol {
		queryParams.Set("include_24hr_vol", "true")
	}

	if params.Include24hrChange {
		queryParams.Set("include_24hr_change", "true")
	}

	if params.IncludeLastUpdatedAt {
		queryParams.Set("include_last_updated_at", "true")
	}

	request := NewRequest(c.Url("/simple/price"), queryParams)
	var response SimplePriceResponse
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
