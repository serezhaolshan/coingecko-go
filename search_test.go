package coingecko

import (
	"context"
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	apiKey := os.Getenv("COINGECKO_API_KEY")
	client := NewClient(DefaultURL, apiKey)

	result, err := client.Search(context.Background(), SearchParams{Query: "bitcoin"})
	if err != nil {
		t.Fatal(err)
	}

	if len(result.Coins) == 0 {
		t.Fatal("no coins returned")
	}

	t.Logf("Found %d coins", len(result.Coins))
}