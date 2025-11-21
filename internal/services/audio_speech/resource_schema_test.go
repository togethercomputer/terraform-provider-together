// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_speech_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio_speech"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestAudioSpeechModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*audio_speech.AudioSpeechModel)(nil)
	schema := audio_speech.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
