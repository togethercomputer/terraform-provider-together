// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/eval"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestEvalDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*eval.EvalDataSourceModel)(nil)
	schema := eval.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
