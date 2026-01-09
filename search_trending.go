package coingecko

import (
	"context"
	"net/url"
)

type TrendingParams struct {
	Include string
}

type TrendingCoinItemData struct {
	Price                    float64            `json:"price"`
	PriceBTC                 string             `json:"price_btc"`
	PriceChangePercentage24h map[string]float64 `json:"price_change_percentage_24h"`
	MarketCap                string             `json:"market_cap"`
	MarketCapBTC             string             `json:"market_cap_btc"`
	TotalVolume              string             `json:"total_volume"`
	TotalVolumeBTC           string             `json:"total_volume_btc"`
	Sparkline                string             `json:"sparkline"`
	Content                  struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"content"`
}

type TrendingCoinItem struct {
	ID            string               `json:"id"`
	CoinID        int                  `json:"coin_id"`
	Name          string               `json:"name"`
	Symbol        string               `json:"symbol"`
	MarketCapRank int                  `json:"market_cap_rank"`
	Thumb         string               `json:"thumb"`
	Small         string               `json:"small"`
	Large         string               `json:"large"`
	Slug          string               `json:"slug"`
	PriceBTC      float64              `json:"price_btc"`
	Score         int                  `json:"score"`
	Data          TrendingCoinItemData `json:"data"`
}

type TrendingNFTItemData struct {
	FloorPrice                         string `json:"floor_price"`
	FloorPriceInUSD24hPercentageChange string `json:"floor_price_in_usd_24h_percentage_change"`
	H24Volume                          string `json:"h24_volume"`
	H24AverageSalePrice                string `json:"h24_average_sale_price"`
	Sparkline                          string `json:"sparkline"`
	Content                            struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"content"`
}

type TrendingNFTItem struct {
	ID                            string               `json:"id"`
	Name                          string               `json:"name"`
	Symbol                        string               `json:"symbol"`
	Thumb                         string               `json:"thumb"`
	NFTContractID                 int                  `json:"nft_contract_id"`
	NativeCurrencySymbol          string               `json:"native_currency_symbol"`
	FloorPriceInNativeCurrency    float64              `json:"floor_price_in_native_currency"`
	FloorPrice24hPercentageChange float64              `json:"floor_price_24h_percentage_change"`
	Data                          *TrendingNFTItemData `json:"data"`
}

type TrendingCategoryItem struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	MarketCap1hChange float64 `json:"market_cap_1h_change"`
	Slug              string  `json:"slug"`
	CoinsCount        string  `json:"coins_count"`
	Data              struct {
		MarketCap                    float64            `json:"market_cap"`
		MarketCapBTC                 float64            `json:"market_cap_btc"`
		TotalVolume                  float64            `json:"total_volume"`
		TotalVolumeBTC               float64            `json:"total_volume_btc"`
		MarketCapChangePercentage24h map[string]float64 `json:"market_cap_change_percentage_24h"`
		Sparkline                    string             `json:"sparkline"`
	} `json:"data"`
}

type TrendingResponse struct {
	Coins []struct {
		Item TrendingCoinItem `json:"item"`
	} `json:"coins"`
	NFTs []struct {
		Item TrendingNFTItem `json:"item"`
	} `json:"nfts"`
	Categories []TrendingCategoryItem `json:"categories"`
}

func (c *Client) GetTrending(ctx context.Context, params *TrendingParams) (*TrendingResponse, error) {
	queryParams := url.Values{}
	if params != nil && params.Include != "" {
		queryParams.Set("include", params.Include)
	}

	request := NewRequest(c.Url("/search/trending"), queryParams)
	var response TrendingResponse
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
