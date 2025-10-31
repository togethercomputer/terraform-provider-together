// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package batch_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/batch"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestBatchModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*batch.BatchModel)(nil)
	schema := batch.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
