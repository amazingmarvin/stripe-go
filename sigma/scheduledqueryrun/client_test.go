package scheduledqueryrun

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/amazingmarvin/stripe-go"
	_ "github.com/amazingmarvin/stripe-go/testing"
)

func TestSigmaScheduledQueryRunGet(t *testing.T) {
	run, err := Get("sqr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, run)
	assert.Equal(t, "scheduled_query_run", run.Object)
}

func TestSigmaScheduledQueryRunList(t *testing.T) {
	i := List(&stripe.SigmaScheduledQueryRunListParams{})

	// Verify that we can get at least one scheduled query run
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SigmaScheduledQueryRun())
	assert.Equal(t, "scheduled_query_run", i.SigmaScheduledQueryRun().Object)
}
