// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_test

import (
	"context"
	"testing"

	"github.com/togethercomputer/terraform-provider-together/internal/services/beta_cluster"
	"github.com/togethercomputer/terraform-provider-together/internal/test_helpers"
)

func TestBetaClusterModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*beta_cluster.BetaClusterModel)(nil)
	schema := beta_cluster.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
