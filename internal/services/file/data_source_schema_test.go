// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package file_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/file"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestFileDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*file.FileDataSourceModel)(nil)
	schema := file.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
