package session

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/amazingmarvin/stripe-go"
	_ "github.com/amazingmarvin/stripe-go/testing"
)

func TestCheckoutSessionGet(t *testing.T) {
	session, err := Get("cs_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, session)
}

func TestCheckoutSessionNew(t *testing.T) {
	session, err := New(&stripe.CheckoutSessionParams{
		CancelURL:         stripe.String("https://stripe.com/cancel"),
		ClientReferenceID: stripe.String("1234"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Amount:      stripe.Int64(1234),
				Currency:    stripe.String(string(stripe.CurrencyUSD)),
				Description: stripe.String("description"),
				Images: stripe.StringSlice([]string{
					"https://stripe.com/image1",
				}),
				Name:     stripe.String("name"),
				Quantity: stripe.Int64(2),
			},
		},
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			Description: stripe.String("description"),
			Shipping: &stripe.ShippingDetailsParams{
				Address: &stripe.AddressParams{
					Line1: stripe.String("line1"),
					City:  stripe.String("city"),
				},
				Carrier: stripe.String("carrier"),
				Name:    stripe.String("name"),
			},
		},
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
				{
					Plan:     stripe.String("plan"),
					Quantity: stripe.Int64(2),
				},
			},
		},
		SuccessURL: stripe.String("https://stripe.com/success"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, session)
}
