// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio_translation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*AudioTranslationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"file": schema.StringAttribute{
				Description:   "Audio file to translate",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"prompt": schema.StringAttribute{
				Description:   "Optional text to bias decoding.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"language": schema.StringAttribute{
				Description:   "Target output language. Optional ISO 639-1 language code. If omitted, language is set to English.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("en"),
			},
			"model": schema.StringAttribute{
				Description: "Model to use for translation\nAvailable values: \"openai/whisper-large-v3\".",
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
				Description: "The translated text",
				Computed:    true,
			},
			"segments": schema.ListNestedAttribute{
				Description: "Array of translation segments",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioTranslationSegmentsModel](ctx),
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
			"words": schema.ListNestedAttribute{
				Description: "Array of translation words (only when timestamp_granularities includes 'word')",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AudioTranslationWordsModel](ctx),
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

func (r *AudioTranslationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AudioTranslationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
