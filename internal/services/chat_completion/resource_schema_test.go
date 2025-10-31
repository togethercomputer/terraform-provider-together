// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chat_completion_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/chat_completion"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestChatCompletionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*chat_completion.ChatCompletionModel)(nil)
	schema := chat_completion.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
