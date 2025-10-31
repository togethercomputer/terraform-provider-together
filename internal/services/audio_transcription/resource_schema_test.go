// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_transcription_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio_transcription"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestAudioTranscriptionModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio_transcription.AudioTranscriptionModel)(nil)
	schema := audio_transcription.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
