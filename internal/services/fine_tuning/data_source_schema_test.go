// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tuning_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/fine_tuning"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestFineTuningDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*fine_tuning.FineTuningDataSourceModel)(nil)
	schema := fine_tuning.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
