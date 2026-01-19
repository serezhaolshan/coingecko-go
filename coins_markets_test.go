package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetCoinsMarkets(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsMarketsParams{
		VsCurrency: "usd",
		PerPage:    10,
		Page:       1,
	}

	result, err := client.GetCoinsMarkets(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsMarkets failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	if result[0].ID != "bitcoin" {
		t.Logf("First coin is %s (expected bitcoin)", result[0].ID)
	}

	t.Logf("Total coins returned: %d", len(result))
	for _, coin := range result[:3] {
		t.Logf("Coin: %s, Price: %v, Market Cap Rank: %v", coin.Name, coin.CurrentPrice, coin.MarketCapRank)
	}
}

func TestGetCoinsMarketsWithIDs(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsMarketsParams{
		VsCurrency: "usd",
		IDs:        []string{"bitcoin", "ethereum"},
	}

	result, err := client.GetCoinsMarkets(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsMarkets failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 coins, got %d", len(result))
	}

	for _, coin := range result {
		t.Logf("Coin: %s (%s), Price: $%.2f", coin.Name, coin.Symbol, *coin.CurrentPrice)
	}
}

func TestGetCoinsMarketsWithSparkline(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsMarketsParams{
		VsCurrency: "usd",
		IDs:        []string{"bitcoin"},
		Sparkline:  true,
	}

	result, err := client.GetCoinsMarkets(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsMarkets failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	if result[0].SparklineIn7d == nil {
		t.Error("Expected sparkline data")
	} else {
		t.Logf("Sparkline data points: %d", len(result[0].SparklineIn7d.Price))
	}
}

func TestGetCoinsMarketsWithPriceChange(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsMarketsParams{
		VsCurrency:            "usd",
		IDs:                   []string{"bitcoin"},
		PriceChangePercentage: "1h,24h,7d",
	}

	result, err := client.GetCoinsMarkets(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsMarkets failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	coin := result[0]
	t.Logf("Bitcoin price changes - 1h: %v, 24h: %v, 7d: %v",
		coin.PriceChangePercentage1h,
		coin.PriceChangePercentage24h,
		coin.PriceChangePercentage7d)
}

func TestGetCoinsMarketsByCategory(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsMarketsParams{
		VsCurrency: "usd",
		Category:   "layer-1",
		PerPage:    5,
	}

	result, err := client.GetCoinsMarkets(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsMarkets failed: %v", err)
	}

	t.Logf("Layer-1 coins returned: %d", len(result))
	for _, coin := range result {
		t.Logf("Coin: %s (%s)", coin.Name, coin.Symbol)
	}
}
