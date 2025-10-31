// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chat_completion

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ChatCompletionResource)(nil)

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
			"messages": schema.ListNestedAttribute{
				Description: "A list of messages comprising the conversation so far.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"content": schema.StringAttribute{
							Optional: true,
						},
						"role": schema.StringAttribute{
							Description: `Available values: "system", "user", "assistant", "tool", "function".`,
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"system",
									"user",
									"assistant",
									"tool",
									"function",
								),
							},
						},
						"name": schema.StringAttribute{
							Optional: true,
						},
						"function_call": schema.SingleNestedAttribute{
							Optional:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Attributes: map[string]schema.Attribute{
								"arguments": schema.StringAttribute{
									Required: true,
								},
								"name": schema.StringAttribute{
									Required: true,
								},
							},
						},
						"tool_calls": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Required: true,
									},
									"function": schema.SingleNestedAttribute{
										Required: true,
										Attributes: map[string]schema.Attribute{
											"arguments": schema.StringAttribute{
												Required: true,
											},
											"name": schema.StringAttribute{
												Required: true,
											},
										},
									},
									"index": schema.Float64Attribute{
										Required: true,
									},
									"type": schema.StringAttribute{
										Description: `Available values: "function".`,
										Required:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("function"),
										},
									},
								},
							},
						},
						"tool_call_id": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
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
			"function_call": schema.StringAttribute{
				Description: `Available values: "none", "auto".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("none", "auto"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
				Description:   "A number between 0 and 1 that can be used as an alternative to top_p and top-k.",
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
			"reasoning_effort": schema.StringAttribute{
				Description: "Controls the level of reasoning effort the model should apply when generating responses. Higher values may result in more thoughtful and detailed responses but may take longer to generate.\nAvailable values: \"low\", \"medium\", \"high\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"low",
						"medium",
						"high",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"tool_choice": schema.StringAttribute{
				Description:   "Controls which (if any) function is called by the model. By default uses `auto`, which lets the model pick between generating a message or calling a function.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
			"response_format": schema.SingleNestedAttribute{
				Description: "An object specifying the format that the model must output.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"schema": schema.MapAttribute{
						Description: "The schema of the response format.",
						Optional:    true,
						ElementType: jsontypes.NormalizedType{},
					},
					"type": schema.StringAttribute{
						Description: "The type of the response format.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"tools": schema.ListNestedAttribute{
				Description: "A list of tools the model may call. Currently, only functions are supported as a tool. Use this to provide a list of functions the model may generate JSON inputs for.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"function": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Optional: true,
								},
								"name": schema.StringAttribute{
									Optional: true,
								},
								"parameters": schema.MapAttribute{
									Description: "A map of parameter names to their values.",
									Optional:    true,
									ElementType: jsontypes.NormalizedType{},
								},
							},
						},
						"type": schema.StringAttribute{
							Optional: true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"context_length_exceeded_behavior": schema.StringAttribute{
				Description: "Defined the behavior of the API when max_tokens exceed the maximum context length of the model. When set to 'error', API will return 400 with appropriate error message. When set to 'truncate', override the max_tokens with maximum context length of the model.\nAvailable values: \"truncate\", \"error\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("truncate", "error"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("error"),
			},
			"created": schema.Int64Attribute{
				Computed: true,
			},
			"object": schema.StringAttribute{
				Description: `Available values: "chat.completion".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("chat.completion"),
				},
			},
			"choices": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ChatCompletionChoicesModel](ctx),
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
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"logprobs": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ChatCompletionChoicesLogprobsModel](ctx),
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
						"message": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ChatCompletionChoicesMessageModel](ctx),
							Attributes: map[string]schema.Attribute{
								"content": schema.StringAttribute{
									Computed: true,
								},
								"role": schema.StringAttribute{
									Description: `Available values: "assistant".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("assistant"),
									},
								},
								"function_call": schema.SingleNestedAttribute{
									Computed:           true,
									DeprecationMessage: "This attribute is deprecated.",
									CustomType:         customfield.NewNestedObjectType[ChatCompletionChoicesMessageFunctionCallModel](ctx),
									Attributes: map[string]schema.Attribute{
										"arguments": schema.StringAttribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"reasoning": schema.StringAttribute{
									Computed: true,
								},
								"tool_calls": schema.ListNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectListType[ChatCompletionChoicesMessageToolCallsModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
											},
											"function": schema.SingleNestedAttribute{
												Computed:   true,
												CustomType: customfield.NewNestedObjectType[ChatCompletionChoicesMessageToolCallsFunctionModel](ctx),
												Attributes: map[string]schema.Attribute{
													"arguments": schema.StringAttribute{
														Computed: true,
													},
													"name": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"index": schema.Float64Attribute{
												Computed: true,
											},
											"type": schema.StringAttribute{
												Description: `Available values: "function".`,
												Computed:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("function"),
												},
											},
										},
									},
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
				CustomType: customfield.NewNestedObjectType[ChatCompletionUsageModel](ctx),
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
			"warnings": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ChatCompletionWarningsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"message": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *ChatCompletionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ChatCompletionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
