// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package embedding_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/embedding"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestEmbeddingModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*embedding.EmbeddingModel)(nil)
	schema := embedding.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
