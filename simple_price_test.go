package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetSimplePrice(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := SimplePriceParams{
		IDs:                  []string{"bitcoin", "ethereum"},
		VsCurrencies:         []string{"usd", "eur"},
		IncludeMarketCap:     true,
		Include24hrVol:       true,
		Include24hrChange:    true,
		IncludeLastUpdatedAt: true,
	}

	result, err := client.GetSimplePrice(ctx, params)
	if err != nil {
		t.Fatalf("GetSimplePrice failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	if _, ok := result["bitcoin"]; !ok {
		t.Error("Expected bitcoin in results")
	}

	if _, ok := result["ethereum"]; !ok {
		t.Error("Expected ethereum in results")
	}

	t.Logf("Result: %+v", result)
}

func TestGetSimplePriceWithSymbols(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := SimplePriceParams{
		Symbols:          []string{"btc", "eth"},
		VsCurrencies:     []string{"usd"},
		IncludeMarketCap: true,
	}

	result, err := client.GetSimplePrice(ctx, params)
	if err != nil {
		t.Fatalf("GetSimplePrice failed: %v", err)
	}

	t.Logf("Result: %+v", result)
}
