// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/eval"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestEvalModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*eval.EvalModel)(nil)
	schema := eval.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
