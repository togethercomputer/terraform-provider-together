// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rerank_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/rerank"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestRerankModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*rerank.RerankModel)(nil)
	schema := rerank.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
