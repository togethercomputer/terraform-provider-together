// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package completion

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*CompletionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The name of the model to query.<br> <br> [See all of Together AI's chat models](https://docs.together.ai/docs/serverless-models#chat-models)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"prompt": schema.StringAttribute{
				Description:   "A string providing context for the model to complete.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"echo": schema.BoolAttribute{
				Description:   "If true, the response will contain the prompt. Can be used with `logprobs` to return prompt logprobs.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"frequency_penalty": schema.Float64Attribute{
				Description:   "A number between -2.0 and 2.0 where a positive value decreases the likelihood of repeating tokens that have already been mentioned.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"logprobs": schema.Int64Attribute{
				Description: "An integer between 0 and 20 of the top k tokens to return log probabilities for at each generation step, instead of just the sampled token. Log probabilities help assess model confidence in token predictions.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(0, 20),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"max_tokens": schema.Int64Attribute{
				Description:   "The maximum number of tokens to generate.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"min_p": schema.Float64Attribute{
				Description:   "A number between 0 and 1 that can be used as an alternative to top-p and top-k.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"n": schema.Int64Attribute{
				Description: "The number of completions to generate for each prompt.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(1, 128),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"presence_penalty": schema.Float64Attribute{
				Description:   "A number between -2.0 and 2.0 where a positive value increases the likelihood of a model talking about new topics.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"repetition_penalty": schema.Float64Attribute{
				Description:   "A number that controls the diversity of generated text by reducing the likelihood of repeated sequences. Higher values decrease repetition.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"safety_model": schema.StringAttribute{
				Description:   "The name of the moderation model used to validate tokens. Choose from the available moderation models found [here](https://docs.together.ai/docs/inference-models#moderation-models).",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"seed": schema.Int64Attribute{
				Description:   "Seed value for reproducibility.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"stream": schema.BoolAttribute{
				Description:   "If true, stream tokens as Server-Sent Events as the model generates them instead of waiting for the full model response. The stream terminates with `data: [DONE]`. If false, return a single JSON object containing the results.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"temperature": schema.Float64Attribute{
				Description:   "A decimal number from 0-1 that determines the degree of randomness in the response. A temperature less than 1 favors more correctness and is appropriate for question answering or summarization. A value closer to 1 introduces more randomness in the output.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"top_k": schema.Int64Attribute{
				Description:   "An integer that's used to limit the number of choices for the next predicted word or token. It specifies the maximum number of tokens to consider at each step, based on their probability of occurrence. This technique helps to speed up the generation process and can improve the quality of the generated text by focusing on the most likely options.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"top_p": schema.Float64Attribute{
				Description:   "A percentage (also called the nucleus parameter) that's used to dynamically adjust the number of choices for each predicted token based on the cumulative probabilities. It specifies a probability threshold below which all less likely tokens are filtered out. This technique helps maintain diversity and generate more fluent and natural-sounding text.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"logit_bias": schema.MapAttribute{
				Description:   "Adjusts the likelihood of specific tokens appearing in the generated output.",
				Optional:      true,
				ElementType:   types.Float64Type,
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
			},
			"stop": schema.ListAttribute{
				Description:   `A list of string sequences that will truncate (stop) inference text output. For example, "</s>" will stop generation as soon as the model generates the given token.`,
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"created": schema.Int64Attribute{
				Computed: true,
			},
			"object": schema.StringAttribute{
				Description: `Available values: "text.completion".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("text.completion"),
				},
			},
			"choices": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[CompletionChoicesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"finish_reason": schema.StringAttribute{
							Description: `Available values: "stop", "eos", "length", "tool_calls", "function_call".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"stop",
									"eos",
									"length",
									"tool_calls",
									"function_call",
								),
							},
						},
						"logprobs": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[CompletionChoicesLogprobsModel](ctx),
							Attributes: map[string]schema.Attribute{
								"token_ids": schema.ListAttribute{
									Description: "List of token IDs corresponding to the logprobs",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Float64](ctx),
									ElementType: types.Float64Type,
								},
								"token_logprobs": schema.ListAttribute{
									Description: "List of token log probabilities",
									Computed:    true,
									CustomType:  customfield.NewListType[types.Float64](ctx),
									ElementType: types.Float64Type,
								},
								"tokens": schema.ListAttribute{
									Description: "List of token strings",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
						"seed": schema.Int64Attribute{
							Computed: true,
						},
						"text": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"usage": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[CompletionUsageModel](ctx),
				Attributes: map[string]schema.Attribute{
					"completion_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"prompt_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"total_tokens": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *CompletionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CompletionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
