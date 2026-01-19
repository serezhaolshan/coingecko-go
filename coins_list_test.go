package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetCoinsList(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	result, err := client.GetCoinsList(ctx, nil)
	if err != nil {
		t.Fatalf("GetCoinsList failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	found := false
	for _, coin := range result {
		if coin.ID == "bitcoin" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected bitcoin in results")
	}

	t.Logf("Total coins: %d", len(result))
}

func TestGetCoinsListWithPlatforms(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := &CoinsListParams{
		IncludePlatform: true,
	}

	result, err := client.GetCoinsList(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsList failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	for _, coin := range result {
		if len(coin.Platforms) > 0 {
			t.Logf("Coin with platforms: %s (%s), platforms: %v", coin.Name, coin.ID, coin.Platforms)
			break
		}
	}

	t.Logf("Total coins: %d", len(result))
}

func TestGetCoinsListInactive(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := &CoinsListParams{
		Status: "inactive",
	}

	result, err := client.GetCoinsList(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsList failed: %v", err)
	}

	t.Logf("Total inactive coins: %d", len(result))
}
