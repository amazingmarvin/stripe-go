package orderreturn

import (
	"net/http"

	stripe "github.com/amazingmarvin/stripe-go"
	"github.com/amazingmarvin/stripe-go/form"
)

// Client is used to invoke /orders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an order return
func Get(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	return getC().Get(id, params)
}

// Get returns the details of an order return
func (c Client) Get(id string, params *stripe.OrderReturnParams) (*stripe.OrderReturn, error) {
	path := stripe.FormatURLPath("/v1/order_returns/%s", id)
	orderReturn := &stripe.OrderReturn{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, orderReturn)
	return orderReturn, err
}

// List returns a list of order returns.
func List(params *stripe.OrderReturnListParams) *Iter {
	return getC().List(params)
}

// List returns a list of order returns.
func (c Client) List(listParams *stripe.OrderReturnListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.OrderReturnList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/order_returns", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for order returns.
type Iter struct {
	*stripe.Iter
}

// OrderReturn returns the order return which the iterator is currently pointing to.
func (i *Iter) OrderReturn() *stripe.OrderReturn {
	return i.Current().(*stripe.OrderReturn)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
