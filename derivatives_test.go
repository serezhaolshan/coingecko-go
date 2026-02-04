package coingecko

import (
	"context"
	"testing"
)

func TestGetDerivativesTickers(t *testing.T) {
	client := NewClient(DefaultURL, "")

	tickers, err := client.GetDerivativesTickers(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tickers) == 0 {
		t.Fatal("expected at least one ticker")
	}

	ticker := tickers[0]
	if ticker.Market == "" {
		t.Error("expected market to be set")
	}
	if ticker.Symbol == "" {
		t.Error("expected symbol to be set")
	}
}

func TestGetDerivativesTickersByIndexID(t *testing.T) {
	client := NewClient(DefaultURL, "")

	tickers, err := client.GetDerivativesTickersByIndexID(context.Background(), "BTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(tickers) == 0 {
		t.Fatal("expected at least one BTC ticker")
	}

	for _, ticker := range tickers {
		if ticker.IndexID != "BTC" {
			t.Errorf("expected index_id to be BTC, got %s", ticker.IndexID)
		}
	}
}

func TestGetAggregatedDerivativesData(t *testing.T) {
	client := NewClient(DefaultURL, "")

	data, err := client.GetAggregatedDerivativesData(context.Background(), "BTC")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data == nil {
		t.Fatal("expected data to be returned")
	}

	if data.IndexID != "BTC" {
		t.Errorf("expected index_id to be BTC, got %s", data.IndexID)
	}

	if data.TotalOpenInterest <= 0 {
		t.Error("expected positive total open interest for BTC")
	}

	if data.TickerCount <= 0 {
		t.Error("expected positive ticker count for BTC")
	}
}
