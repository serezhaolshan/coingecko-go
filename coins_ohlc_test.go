package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestGetCoinsOHLC(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
		Days:       "7",
	}

	result, err := client.GetCoinsOHLC(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsOHLC failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	t.Logf("Total candles: %d", len(result))
	if len(result) > 0 {
		candle := result[0]
		t.Logf("First candle - Timestamp: %d, Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f",
			candle.Timestamp, candle.Open, candle.High, candle.Low, candle.Close)
	}
}

func TestGetCoinsOHLC30Days(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		ID:         "ethereum",
		VsCurrency: "usd",
		Days:       "30",
	}

	result, err := client.GetCoinsOHLC(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsOHLC failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	t.Logf("Total candles for 30 days: %d", len(result))
}

func TestGetCoinsOHLCDaily(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
		Days:       "14",
		Interval:   "daily",
	}

	result, err := client.GetCoinsOHLC(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsOHLC failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	t.Logf("Total daily candles: %d", len(result))
}

func TestGetCoinsOHLCMissingID(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		VsCurrency: "usd",
		Days:       "7",
	}

	_, err := client.GetCoinsOHLC(ctx, params)
	if err == nil {
		t.Fatal("Expected error for missing coin ID")
	}

	t.Logf("Expected error: %v", err)
}

func TestGetCoinsOHLCMissingDays(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
	}

	_, err := client.GetCoinsOHLC(ctx, params)
	if err == nil {
		t.Fatal("Expected error for missing days parameter")
	}

	t.Logf("Expected error: %v", err)
}

func TestGetCoinsOHLCWithPrecision(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)
	ctx := context.Background()

	params := CoinsOHLCParams{
		ID:         "bitcoin",
		VsCurrency: "usd",
		Days:       "1",
		Precision:  "2",
	}

	result, err := client.GetCoinsOHLC(ctx, params)
	if err != nil {
		t.Fatalf("GetCoinsOHLC failed: %v", err)
	}

	if len(result) == 0 {
		t.Fatal("Expected non-empty result")
	}

	t.Logf("Total candles with precision 2: %d", len(result))
	if len(result) > 0 {
		t.Logf("Sample price: %.2f", result[0].Close)
	}
}
