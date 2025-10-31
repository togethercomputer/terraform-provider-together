// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package job_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/job"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestJobDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*job.JobDataSourceModel)(nil)
	schema := job.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
