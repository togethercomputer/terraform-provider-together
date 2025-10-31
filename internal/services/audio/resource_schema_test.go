// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestAudioModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio.AudioModel)(nil)
	schema := audio.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
