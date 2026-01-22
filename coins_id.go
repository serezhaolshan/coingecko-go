package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

type TVLValue struct {
	USD float64
}

func (t *TVLValue) UnmarshalJSON(data []byte) error {
	var obj struct {
		USD float64 `json:"usd"`
	}
	if err := json.Unmarshal(data, &obj); err == nil {
		t.USD = obj.USD
		return nil
	}

	var val float64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	t.USD = val
	return nil
}

func (t TVLValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.USD)
}

type CoinParams struct {
	ID            string
	DexPairFormat string
}

type CoinImage struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type CoinLinks struct {
	Homepage                    []string            `json:"homepage"`
	WhitePaper                  string              `json:"whitepaper"`
	BlockchainSite              []string            `json:"blockchain_site"`
	OfficialForumURL            []string            `json:"official_forum_url"`
	ChatURL                     []string            `json:"chat_url"`
	AnnouncementURL             []string            `json:"announcement_url"`
	TwitterScreenName           string              `json:"twitter_screen_name"`
	FacebookUsername            string              `json:"facebook_username"`
	BitcoinTalkThreadIdentifier string              `json:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   string              `json:"telegram_channel_identifier"`
	SubredditURL                string              `json:"subreddit_url"`
	ReposURL                    map[string][]string `json:"repos_url"`
}

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

type CoinMarketData struct {
	CurrentPrice                           map[string]float64 `json:"current_price"`
	TotalValueLocked                       *TVLValue          `json:"total_value_locked,omitempty"`
	McapToTVLRatio                         *float64           `json:"mcap_to_tvl_ratio,omitempty"`
	FdvToTVLRatio                          *float64           `json:"fdv_to_tvl_ratio,omitempty"`
	ROI                                    *ROI               `json:"roi,omitempty"`
	ATH                                    map[string]float64 `json:"ath"`
	ATHChangePercentage                    map[string]float64 `json:"ath_change_percentage"`
	ATHDate                                map[string]string  `json:"ath_date"`
	ATL                                    map[string]float64 `json:"atl"`
	ATLChangePercentage                    map[string]float64 `json:"atl_change_percentage"`
	ATLDate                                map[string]string  `json:"atl_date"`
	MarketCap                              map[string]float64 `json:"market_cap"`
	MarketCapRank                          int                `json:"market_cap_rank"`
	FullyDilutedValuation                  map[string]float64 `json:"fully_diluted_valuation"`
	MarketCapFDVRatio                      float64            `json:"market_cap_fdv_ratio"`
	TotalVolume                            map[string]float64 `json:"total_volume"`
	High24h                                map[string]float64 `json:"high_24h"`
	Low24h                                 map[string]float64 `json:"low_24h"`
	PriceChange24h                         float64            `json:"price_change_24h"`
	PriceChangePercentage24h               float64            `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64            `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64            `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64            `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64            `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64            `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64            `json:"price_change_percentage_1y"`
	MarketCapChange24h                     float64            `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64            `json:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               map[string]float64 `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      map[string]float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     map[string]float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      map[string]float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     map[string]float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     map[string]float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     map[string]float64 `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    map[string]float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      map[string]float64 `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           map[string]float64 `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency map[string]float64 `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            float64            `json:"total_supply"`
	MaxSupply                              float64            `json:"max_supply"`
	CirculatingSupply                      float64            `json:"circulating_supply"`
	LastUpdated                            string             `json:"last_updated"`
}

type CoinCommunityData struct {
	FacebookLikes            any     `json:"facebook_likes"`
	TwitterFollowers         int     `json:"twitter_followers"`
	RedditAveragePosts48h    float64 `json:"reddit_average_posts_48h"`
	RedditAverageComments48h float64 `json:"reddit_average_comments_48h"`
	RedditSubscribers        int     `json:"reddit_subscribers"`
	RedditAccountsActive48h  int     `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount int     `json:"telegram_channel_user_count"`
}

type CoinDeveloperData struct {
	Forks                          int            `json:"forks"`
	Stars                          int            `json:"stars"`
	Subscribers                    int            `json:"subscribers"`
	TotalIssues                    int            `json:"total_issues"`
	ClosedIssues                   int            `json:"closed_issues"`
	PullRequestsMerged             int            `json:"pull_requests_merged"`
	PullRequestContributors        int            `json:"pull_request_contributors"`
	CodeAdditionsDeletions4Weeks   map[string]int `json:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks              int            `json:"commit_count_4_weeks"`
	Last4WeeksCommitActivitySeries []int          `json:"last_4_weeks_commit_activity_series"`
}

type CoinPublicInterestStats struct {
	AlexaRank   int `json:"alexa_rank"`
	BingMatches int `json:"bing_matches"`
}

type CoinResponse struct {
	ID                           string                  `json:"id"`
	Symbol                       string                  `json:"symbol"`
	Name                         string                  `json:"name"`
	WebSlug                      string                  `json:"web_slug"`
	AssetPlatformID              string                  `json:"asset_platform_id"`
	Platforms                    map[string]string       `json:"platforms"`
	DetailPlatforms              map[string]any          `json:"detail_platforms"`
	BlockTimeInMinutes           int                     `json:"block_time_in_minutes"`
	HashingAlgorithm             string                  `json:"hashing_algorithm"`
	Categories                   []string                `json:"categories"`
	PreviewListing               bool                    `json:"preview_listing"`
	PublicNotice                 any                     `json:"public_notice"`
	AdditionalNotices            []any                   `json:"additional_notices"`
	Description                  map[string]string       `json:"description"`
	Links                        CoinLinks               `json:"links"`
	Image                        CoinImage               `json:"image"`
	CountryOrigin                string                  `json:"country_origin"`
	GenesisDate                  string                  `json:"genesis_date"`
	SentimentVotesUpPercentage   float64                 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                 `json:"sentiment_votes_down_percentage"`
	WatchlistPortfolioUsers      int                     `json:"watchlist_portfolio_users"`
	MarketCapRank                int                     `json:"market_cap_rank"`
	MarketData                   CoinMarketData          `json:"market_data"`
	CommunityData                CoinCommunityData       `json:"community_data"`
	DeveloperData                CoinDeveloperData       `json:"developer_data"`
	PublicInterestStats          CoinPublicInterestStats `json:"public_interest_stats"`
	StatusUpdates                []any                   `json:"status_updates"`
	LastUpdated                  string                  `json:"last_updated"`
}

func (c *Client) GetCoin(ctx context.Context, params CoinParams) (*CoinResponse, error) {
	if params.ID == "" {
		return nil, fmt.Errorf("coin ID is required")
	}

	queryParams := url.Values{}
	if params.DexPairFormat != "" {
		queryParams.Set("dex_pair_format", params.DexPairFormat)
	}

	endpoint := fmt.Sprintf("/coins/%s", params.ID)
	request := NewRequest(c.Url(endpoint), queryParams)
	var response CoinResponse
	_, err := c.doCall(ctx, request, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
