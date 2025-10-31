// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/video"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestVideoDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*video.VideoDataSourceModel)(nil)
	schema := video.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
