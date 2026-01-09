package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetTrending(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	result, err := client.GetTrending(ctx, &TrendingParams{
		Include: "coins",
	})
	if err != nil {
		t.Fatalf("GetTrending failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	t.Logf("Found %d trending coins", len(result.Coins))
	t.Logf("Found %d trending NFTs", len(result.NFTs))
	t.Logf("Found %d trending categories", len(result.Categories))

	if len(result.Coins) > 0 {
		coin := result.Coins[0].Item
		t.Logf("Top trending coin: %s (%s)", coin.Name, coin.Symbol)
	}
}

func TestGetTrendingAll(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	result, err := client.GetTrending(ctx, &TrendingParams{
		Include: "coins,categories",
	})
	if err != nil {
		t.Fatalf("GetTrending failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	t.Logf("Result: %+v", result)
}
