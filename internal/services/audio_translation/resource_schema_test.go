// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_translation_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio_translation"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestAudioTranslationModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio_translation.AudioTranslationModel)(nil)
	schema := audio_translation.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
