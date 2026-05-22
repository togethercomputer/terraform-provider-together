// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_storage_test

import (
	"context"
	"testing"

	"github.com/togethercomputer/terraform-provider-together/internal/services/beta_cluster_storage"
	"github.com/togethercomputer/terraform-provider-together/internal/test_helpers"
)

func TestBetaClusterStorageModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*beta_cluster_storage.BetaClusterStorageModel)(nil)
	schema := beta_cluster_storage.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
