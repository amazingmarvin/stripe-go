package stripe

import (
	"encoding/json"
)

// CheckoutSessionSubmitType is the list of allowed values for the `submit_type`
// of a Session.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take.
const (
	CheckoutSessionSubmitTypeAuto   CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook   CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay    CheckoutSessionSubmitType = "pay"
)

// CheckoutSessionDisplayItemType is the list of allowed values for the display item type.
type CheckoutSessionDisplayItemType string

// List of values that CheckoutSessionDisplayItemType can take.
const (
	CheckoutSessionDisplayItemTypeCustom CheckoutSessionDisplayItemType = "custom"
	CheckoutSessionDisplayItemTypePlan   CheckoutSessionDisplayItemType = "plan"
	CheckoutSessionDisplayItemTypeSKU    CheckoutSessionDisplayItemType = "sku"
)

// CheckoutSessionMode is the list of allowed values for the mode on a Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take.
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
)

// CheckoutSessionLineItemParams is the set of parameters allowed for a line item
// on a checkout session.
type CheckoutSessionLineItemParams struct {
	// When set, provides configuration for this item's quantity to be adjusted by the customer during Checkout.
	AdjustableQuantity *CheckoutSessionLineItemAdjustableQuantityParams `form:"adjustable_quantity"`
	// [Deprecated] The amount to be collected per unit of the line item. If specified, must also pass `currency` and `name`.
	Amount *int64 `form:"amount"`
	// [Deprecated] Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies). Required if `amount` is passed.
	Currency *string `form:"currency"`
	// [Deprecated] The description for the line item, to be displayed on the Checkout page.
	Description *string `form:"description"`
	// The [tax rates](https://stripe.com/docs/api/tax_rates) that will be applied to this line item depending on the customer's billing/shipping address. We currently support the following countries: US, GB, AU, and all countries in the EU.
	DynamicTaxRates []*string `form:"dynamic_tax_rates"`
	// [Deprecated] A list of image URLs representing this line item. Each image can be up to 5 MB in size. If passing `price` or `price_data`, specify images on the associated product instead.
	Images []*string `form:"images"`
	// [Deprecated] The name for the item to be displayed on the Checkout page. Required if `amount` is passed.
	Name *string `form:"name"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) or [Plan](https://stripe.com/docs/api/plans) object. One of `price` or `price_data` is required.
	Price *string `form:"price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline. One of `price` or `price_data` is required.
	PriceData *CheckoutSessionLineItemPriceDataParams `form:"price_data"`
	// The quantity of the line item being purchased. Quantity should not be defined when `recurring.usage_type=metered`.
	Quantity *int64 `form:"quantity"`
	// The [tax rates](https://stripe.com/docs/api/tax_rates) which apply to this line item.
	TaxRates []*string `form:"tax_rates"`
}

// CheckoutSessionPaymentIntentDataTransferDataParams is the set of parameters allowed for the
// transfer_data hash.
type CheckoutSessionPaymentIntentDataTransferDataParams struct {
	Destination *string `form:"destination"`
}

// CheckoutSessionPaymentIntentDataParams is the set of parameters allowed for the
// payment intent creation on a checkout session.
type CheckoutSessionPaymentIntentDataParams struct {
	Params               `form:"*"`
	ApplicationFeeAmount *int64                                              `form:"application_fee_amount"`
	CaptureMethod        *string                                             `form:"capture_method"`
	Description          *string                                             `form:"description"`
	OnBehalfOf           *string                                             `form:"on_behalf_of"`
	ReceiptEmail         *string                                             `form:"receipt_email"`
	SetupFutureUsage     *string                                             `form:"setup_future_usage"`
	Shipping             *ShippingDetailsParams                              `form:"shipping"`
	StatementDescriptor  *string                                             `form:"statement_descriptor"`
	TransferData         *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data"`
}

// CheckoutSessionSetupIntentDataParams is the set of parameters allowed for the setup intent
// creation on a checkout session.
type CheckoutSessionSetupIntentDataParams struct {
	Params      `form:"*"`
	Description *string `form:"description"`
	OnBehalfOf  *string `form:"on_behalf_of"`
}

// CheckoutSessionSubscriptionDataItemsParams is the set of parameters allowed for one item on a
// checkout session associated with a subscription.
type CheckoutSessionSubscriptionDataItemsParams struct {
	Plan     *string `form:"plan"`
	Quantity *int64  `form:"quantity"`
}

// CheckoutSessionSubscriptionDataParams is the set of parameters allowed for the subscription
// creation on a checkout session.
type CheckoutSessionSubscriptionDataParams struct {
	Params                `form:"*"`
	ApplicationFeePercent *float64                                      `form:"application_fee_percent"`
	Items                 []*CheckoutSessionSubscriptionDataItemsParams `form:"items"`
	TrialEnd              *int64                                        `form:"trial_end"`
	TrialFromPlan         *bool                                         `form:"trial_from_plan"`
	TrialPeriodDays       *int64                                        `form:"trial_period_days"`
}

// CheckoutSessionParams is the set of parameters that can be used when creating
// a checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/create
type CheckoutSessionParams struct {
	Params                   `form:"*"`
	BillingAddressCollection *string                                 `form:"billing_address_collection"`
	CancelURL                *string                                 `form:"cancel_url"`
	ClientReferenceID        *string                                 `form:"client_reference_id"`
	Customer                 *string                                 `form:"customer"`
	CustomerEmail            *string                                 `form:"customer_email"`
	LineItems                []*CheckoutSessionLineItemParams        `form:"line_items"`
	Locale                   *string                                 `form:"locale"`
	Mode                     *string                                 `form:"mode"`
	PaymentIntentData        *CheckoutSessionPaymentIntentDataParams `form:"payment_intent_data"`
	PaymentMethodTypes       []*string                               `form:"payment_method_types"`
	SetupIntentData          *CheckoutSessionSetupIntentDataParams   `form:"setup_intent_data"`
	SubscriptionData         *CheckoutSessionSubscriptionDataParams  `form:"subscription_data"`
	SubmitType               *string                                 `form:"submit_type"`
	SuccessURL               *string                                 `form:"success_url"`
}

// CheckoutSessionDisplayItemCustom represents an item of type custom in a checkout session
type CheckoutSessionDisplayItemCustom struct {
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Name        string   `json:"name"`
}

// CheckoutSessionDisplayItem represents one of the items in a checkout session.
type CheckoutSessionDisplayItem struct {
	Amount   int64                             `json:"amount"`
	Currency Currency                          `json:"currency"`
	Custom   *CheckoutSessionDisplayItemCustom `json:"custom"`
	Quantity int64                             `json:"quantity"`
	Plan     *Plan                             `json:"plan"`
	SKU      *SKU                              `json:"sku"`
	Type     CheckoutSessionDisplayItemType    `json:"type"`
}

// CheckoutSession is the resource representing a Stripe checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/object
type CheckoutSession struct {
	CancelURL          string                        `json:"cancel_url"`
	ClientReferenceID  string                        `json:"client_reference_id"`
	Customer           *Customer                     `json:"customer"`
	CustomerEmail      string                        `json:"customer_email"`
	Deleted            bool                          `json:"deleted"`
	DisplayItems       []*CheckoutSessionDisplayItem `json:"display_items"`
	ID                 string                        `json:"id"`
	Livemode           bool                          `json:"livemode"`
	Locale             string                        `json:"locale"`
	Mode               CheckoutSessionMode           `json:"mode"`
	Object             string                        `json:"object"`
	PaymentIntent      *PaymentIntent                `json:"payment_intent"`
	PaymentMethodTypes []string                      `json:"payment_method_types"`
	SetupIntent        *SetupIntent                  `json:"setup_intent"`
	Subscription       *Subscription                 `json:"subscription"`
	SubmitType         CheckoutSessionSubmitType     `json:"submit_type"`
	SuccessURL         string                        `json:"success_url"`
}

// UnmarshalJSON handles deserialization of a checkout session.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type session CheckoutSession
	var v session
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = CheckoutSession(v)
	return nil
}
