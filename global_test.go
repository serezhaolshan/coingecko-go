package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetGlobal(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	result, err := client.GetGlobal(ctx)
	if err != nil {
		t.Fatalf("GetGlobal failed: %v", err)
	}

	if result.ActiveCryptocurrencies == 0 {
		t.Error("Expected ActiveCryptocurrencies to be non-zero")
	}

	if result.Markets == 0 {
		t.Error("Expected Markets to be non-zero")
	}

	if len(result.TotalMarketCap) == 0 {
		t.Error("Expected TotalMarketCap to have entries")
	}

	if len(result.TotalVolume) == 0 {
		t.Error("Expected TotalVolume to have entries")
	}

	if len(result.MarketCapPercentage) == 0 {
		t.Error("Expected MarketCapPercentage to have entries")
	}

	if result.UpdatedAt == 0 {
		t.Error("Expected UpdatedAt to be non-zero")
	}

	t.Logf("Active cryptocurrencies: %d", result.ActiveCryptocurrencies)
	t.Logf("Markets: %d", result.Markets)
	t.Logf("BTC dominance: %.2f%%", result.MarketCapPercentage["btc"])
	t.Logf("ETH dominance: %.2f%%", result.MarketCapPercentage["eth"])
	t.Logf("Total market cap (USD): %.0f", result.TotalMarketCap["usd"])
}