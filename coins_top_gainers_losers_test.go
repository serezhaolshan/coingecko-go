package coingecko

import (
	"context"
	"os"
	"testing"
)

func deref(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0.0
}

func TestGetTopGainersLosers(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := TopGainersLosersParams{
		VsCurrency:            "usd",
		Duration:              Duration24h,
		PriceChangePercentage: PriceChange24h,
		TopCoins:              TopCoins300,
	}

	result, err := client.GetTopGainersLosers(ctx, params)
	if err != nil {
		t.Fatalf("GetTopGainersLosers failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	t.Logf("Top Gainers count: %d", len(result.TopGainers))
	t.Logf("Top Losers count: %d", len(result.TopLosers))

	if len(result.TopGainers) > 0 {
		top := result.TopGainers[0]
		t.Logf("Biggest gainer: %s (%s) - %.2f%%", top.Name, top.Symbol, deref(top.Usd24hChange))
	}

	if len(result.TopLosers) > 0 {
		top := result.TopLosers[0]
		t.Logf("Biggest loser: %s (%s) - %.2f%%", top.Name, top.Symbol, deref(top.Usd24hChange))
	}
}

func TestGetTopGainersLosersDefaultCurrency(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := TopGainersLosersParams{
		Duration: Duration24h,
	}

	result, err := client.GetTopGainersLosers(ctx, params)
	if err != nil {
		t.Fatalf("GetTopGainersLosers failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	t.Logf("Result: %+v", result)
}
