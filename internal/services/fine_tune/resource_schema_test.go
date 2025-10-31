// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tune_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/fine_tune"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestFineTuneModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*fine_tune.FineTuneModel)(nil)
	schema := fine_tune.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
