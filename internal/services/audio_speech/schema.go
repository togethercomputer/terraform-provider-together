// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_speech

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*AudioSpeechResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				Description:   "Input text to generate the audio for",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The name of the model to query.<br> <br> [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#audio-models) The current supported tts models are: - cartesia/sonic - hexgrad/Kokoro-82M - canopylabs/orpheus-3b-0.1-ft",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"voice": schema.StringAttribute{
				Description:   `The voice to use for generating the audio. The voices supported are different for each model. For eg - for canopylabs/orpheus-3b-0.1-ft, one of the voices supported is tara, for hexgrad/Kokoro-82M, one of the voices supported is af_alloy and for cartesia/sonic, one of the voices supported is "friendly sidekick". <br> <br> You can view the voices supported for each model using the /v1/voices endpoint sending the model name as the query parameter. [View all supported voices here](https://docs.together.ai/docs/text-to-speech#voices-available).`,
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"stream": schema.BoolAttribute{
				Description:   "If true, output is streamed for several characters at a time instead of waiting for the full response. The stream terminates with `data: [DONE]`. If false, return the encoded audio as octet stream",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"language": schema.StringAttribute{
				Description: "Language of input text.\nAvailable values: \"en\", \"de\", \"fr\", \"es\", \"hi\", \"it\", \"ja\", \"ko\", \"nl\", \"pl\", \"pt\", \"ru\", \"sv\", \"tr\", \"zh\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"en",
						"de",
						"fr",
						"es",
						"hi",
						"it",
						"ja",
						"ko",
						"nl",
						"pl",
						"pt",
						"ru",
						"sv",
						"tr",
						"zh",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("en"),
			},
			"response_encoding": schema.StringAttribute{
				Description: "Audio encoding of response\nAvailable values: \"pcm_f32le\", \"pcm_s16le\", \"pcm_mulaw\", \"pcm_alaw\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pcm_f32le",
						"pcm_s16le",
						"pcm_mulaw",
						"pcm_alaw",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("pcm_f32le"),
			},
			"response_format": schema.StringAttribute{
				Description: "The format of audio output. Supported formats are mp3, wav, raw if streaming is false. If streaming is true, the only supported format is raw.\nAvailable values: \"mp3\", \"wav\", \"raw\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"mp3",
						"wav",
						"raw",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("wav"),
			},
			"sample_rate": schema.Int64Attribute{
				Description:   "Sampling rate to use for the output audio. The default sampling rate for canopylabs/orpheus-3b-0.1-ft and hexgrad/Kokoro-82M is 24000 and for cartesia/sonic is 44100.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(44100),
			},
		},
	}
}

func (r *AudioSpeechResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AudioSpeechResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
