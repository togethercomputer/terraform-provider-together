// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package completion_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/completion"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestCompletionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*completion.CompletionModel)(nil)
	schema := completion.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
