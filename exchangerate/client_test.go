package exchangerate

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/amazingmarvin/stripe-go"
	_ "github.com/amazingmarvin/stripe-go/testing"
)

func TestExchangeRateGet(t *testing.T) {
	rates, err := Get(string(stripe.CurrencyUSD), nil)
	assert.Nil(t, err)
	assert.NotNil(t, rates)
}

func TestExchangeRateList(t *testing.T) {
	i := List(&stripe.ExchangeRateListParams{})

	// Verify that we can get at least one exchange_rate
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ExchangeRate())
}
