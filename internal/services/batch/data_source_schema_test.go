// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package batch_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/batch"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestBatchDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*batch.BatchDataSourceModel)(nil)
	schema := batch.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
