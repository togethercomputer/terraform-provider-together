// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/video"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestVideoModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*video.VideoModel)(nil)
	schema := video.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
