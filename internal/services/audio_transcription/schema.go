// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_transcription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*AudioTranscriptionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"file": schema.StringAttribute{
				Description:   "Audio file to transcribe",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"max_speakers": schema.Int64Attribute{
				Description:   "Maximum number of speakers expected in the audio. Used to improve diarization accuracy when the approximate number of speakers is known.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"min_speakers": schema.Int64Attribute{
				Description:   "Minimum number of speakers expected in the audio. Used to improve diarization accuracy when the approximate number of speakers is known.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"prompt": schema.StringAttribute{
				Description:   "Optional text to bias decoding.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"diarize": schema.BoolAttribute{
				Description:   "Whether to enable speaker diarization. When enabled, you will get the speaker id for each word in the transcription. In the response, in the words array, you will get the speaker id for each word. In addition, we also return the speaker_segments array which contains the speaker id for each speaker segment along with the start and end time of the segment along with all the words in the segment. <br> <br> For eg - ... \"speaker_segments\": [\n  \"speaker_id\": \"SPEAKER_00\",\n  \"start\": 0,\n  \"end\": 30.02,\n  \"words\": [\n    {\n      \"id\": 0,\n      \"word\": \"Tijana\",\n      \"start\": 0,\n      \"end\": 11.475,\n      \"speaker_id\": \"SPEAKER_00\"\n    },\n    ...",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"language": schema.StringAttribute{
				Description:   "Optional ISO 639-1 language code. If `auto` is provided, language is auto-detected.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("en"),
			},
			"model": schema.StringAttribute{
				Description: "Model to use for transcription\nAvailable values: \"openai/whisper-large-v3\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("openai/whisper-large-v3"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("openai/whisper-large-v3"),
			},
			"response_format": schema.StringAttribute{
				Description: "The format of the response\nAvailable values: \"json\", \"verbose_json\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("json", "verbose_json"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("json"),
			},
			"temperature": schema.Float64Attribute{
				Description: "Sampling temperature between 0.0 and 1.0",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1),
				},
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(0),
			},
			"timestamp_granularities": schema.StringAttribute{
				Description: "Controls level of timestamp detail in verbose_json. Only used when response_format is verbose_json. Can be a single granularity or an array to get multiple levels.\nAvailable values: \"segment\", \"word\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("segment", "word"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
			},
			"duration": schema.Float64Attribute{
				Description: "The duration of the audio in seconds",
				Computed:    true,
			},
			"task": schema.StringAttribute{
				Description: "The task performed\nAvailable values: \"transcribe\", \"translate\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("transcribe", "translate"),
				},
			},
			"text": schema.StringAttribute{
				Description: "The transcribed text",
				Computed:    true,
			},
			"segments": schema.ListNestedAttribute{
				Description: "Array of transcription segments",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioTranscriptionSegmentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Unique identifier for the segment",
							Computed:    true,
						},
						"end": schema.Float64Attribute{
							Description: "End time of the segment in seconds",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "Start time of the segment in seconds",
							Computed:    true,
						},
						"text": schema.StringAttribute{
							Description: "The text content of the segment",
							Computed:    true,
						},
					},
				},
			},
			"speaker_segments": schema.ListNestedAttribute{
				Description: "Array of transcription speaker segments (only when diarize is enabled)",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioTranscriptionSpeakerSegmentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Description: "Unique identifier for the speaker segment",
							Computed:    true,
						},
						"end": schema.Float64Attribute{
							Description: "End time of the speaker segment in seconds",
							Computed:    true,
						},
						"speaker_id": schema.StringAttribute{
							Description: "The speaker identifier",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "Start time of the speaker segment in seconds",
							Computed:    true,
						},
						"text": schema.StringAttribute{
							Description: "The full text spoken by this speaker in this segment",
							Computed:    true,
						},
						"words": schema.ListNestedAttribute{
							Description: "Array of words spoken by this speaker in this segment",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[AudioTranscriptionSpeakerSegmentsWordsModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"end": schema.Float64Attribute{
										Description: "End time of the word in seconds",
										Computed:    true,
									},
									"start": schema.Float64Attribute{
										Description: "Start time of the word in seconds",
										Computed:    true,
									},
									"word": schema.StringAttribute{
										Description: "The word",
										Computed:    true,
									},
									"speaker_id": schema.StringAttribute{
										Description: "The speaker id for the word (only when diarize is enabled)",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"words": schema.ListNestedAttribute{
				Description: "Array of transcription words (only when timestamp_granularities includes 'word')",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioTranscriptionWordsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"end": schema.Float64Attribute{
							Description: "End time of the word in seconds",
							Computed:    true,
						},
						"start": schema.Float64Attribute{
							Description: "Start time of the word in seconds",
							Computed:    true,
						},
						"word": schema.StringAttribute{
							Description: "The word",
							Computed:    true,
						},
						"speaker_id": schema.StringAttribute{
							Description: "The speaker id for the word (only when diarize is enabled)",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *AudioTranscriptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AudioTranscriptionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
