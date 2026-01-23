package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetMarketChart(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := MarketChartParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
		Days:       "7",
	}

	result, err := client.GetMarketChart(ctx, params)
	if err != nil {
		t.Fatalf("GetMarketChart failed: %v", err)
	}

	if len(result.Prices) == 0 {
		t.Fatal("Expected non-empty prices")
	}

	t.Logf("Total price points: %d", len(result.Prices))
	t.Logf("Total market cap points: %d", len(result.MarketCaps))
	t.Logf("Total volume points: %d", len(result.TotalVolumes))

	if len(result.Prices) > 0 {
		point := result.Prices[0]
		t.Logf("First price point - Timestamp: %d, Value: %.2f", point.Timestamp, point.Value)
	}
}

func TestGetMarketChartMissingID(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := MarketChartParams{
		VsCurrency: "usd",
		Days:       "7",
	}

	_, err := client.GetMarketChart(ctx, params)
	if err == nil {
		t.Fatal("Expected error for missing coin ID")
	}

	t.Logf("Expected error: %v", err)
}

func TestGetMarketChartMissingDays(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := MarketChartParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
	}

	_, err := client.GetMarketChart(ctx, params)
	if err == nil {
		t.Fatal("Expected error for missing days parameter")
	}

	t.Logf("Expected error: %v", err)
}
