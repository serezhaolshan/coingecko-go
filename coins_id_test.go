package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetCoin(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinParams{
		ID:            "bitcoin",
		DexPairFormat: "symbol",
	}

	result, err := client.GetCoin(ctx, params)
	if err != nil {
		t.Fatalf("GetCoin failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if result.ID != "bitcoin" {
		t.Errorf("Expected ID to be bitcoin, got %s", result.ID)
	}

	t.Logf("Coin: %s (%s)", result.Name, result.Symbol)
	t.Logf("Market Cap Rank: %d", result.MarketCapRank)

	if usdPrice, ok := result.MarketData.CurrentPrice["usd"]; ok {
		t.Logf("Current Price (USD): $%.2f", usdPrice)
	}
}

func TestGetCoinEthereum(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinParams{
		ID: "ethereum",
	}

	result, err := client.GetCoin(ctx, params)
	if err != nil {
		t.Fatalf("GetCoin failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if result.ID != "ethereum" {
		t.Errorf("Expected ID to be ethereum, got %s", result.ID)
	}

	t.Logf("Result: %s", result.Name)
}

func TestGetCoinInvalidID(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinParams{
		ID: "",
	}

	_, err := client.GetCoin(ctx, params)
	if err == nil {
		t.Fatal("Expected error for empty coin ID")
	}

	t.Logf("Expected error: %v", err)
}
