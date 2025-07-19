package mexcsdk

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryCurrencyInfo_RealAPI(t *testing.T) {
	// Skip if not running integration tests
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	// Get credentials from environment variables
	apiKey := os.Getenv("MEXC_API_KEY")
	apiSecret := os.Getenv("MEXC_API_SECRET")
	if apiKey == "" || apiSecret == "" {
		t.Fatal("API credentials not set in environment variables")
	}

	// Initialize client with real MEXC API endpoint
	client := NewSpot(apiKey, apiSecret, "https://api.mexc.com/api/v3", nil)

	// Test actual API call
	currencies, err := client.QueryCurrencyInfo(nil)

	// Basic assertions
	assert.NoError(t, err, "API call should succeed")
	assert.NotEmpty(t, currencies, "Should return at least one currency")

	// Verify structure of the first currency
	first := currencies[0]
	assert.NotEmpty(t, first.Coin, "Currency should have a coin code")
	assert.NotEmpty(t, first.NetworkList, "Currency should have network info")

	t.Logf("Successfully fetched %d currencies", len(currencies))
	t.Logf("First currency: %s with %d networks", first.Coin, len(first.NetworkList))
}
