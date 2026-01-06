// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/beta_cluster"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestBetaClusterDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*beta_cluster.BetaClusterDataSourceModel)(nil)
	schema := beta_cluster.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
