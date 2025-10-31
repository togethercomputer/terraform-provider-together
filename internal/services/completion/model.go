// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package completion

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type CompletionModel struct {
	ID                types.String                                         `tfsdk:"id" json:"id,computed"`
	Model             types.String                                         `tfsdk:"model" json:"model,required"`
	Prompt            types.String                                         `tfsdk:"prompt" json:"prompt,required"`
	Echo              types.Bool                                           `tfsdk:"echo" json:"echo,optional"`
	FrequencyPenalty  types.Float64                                        `tfsdk:"frequency_penalty" json:"frequency_penalty,optional"`
	Logprobs          types.Int64                                          `tfsdk:"logprobs" json:"logprobs,optional"`
	MaxTokens         types.Int64                                          `tfsdk:"max_tokens" json:"max_tokens,optional"`
	MinP              types.Float64                                        `tfsdk:"min_p" json:"min_p,optional"`
	N                 types.Int64                                          `tfsdk:"n" json:"n,optional"`
	PresencePenalty   types.Float64                                        `tfsdk:"presence_penalty" json:"presence_penalty,optional"`
	RepetitionPenalty types.Float64                                        `tfsdk:"repetition_penalty" json:"repetition_penalty,optional"`
	SafetyModel       types.String                                         `tfsdk:"safety_model" json:"safety_model,optional"`
	Seed              types.Int64                                          `tfsdk:"seed" json:"seed,optional"`
	Stream            types.Bool                                           `tfsdk:"stream" json:"stream,optional"`
	Temperature       types.Float64                                        `tfsdk:"temperature" json:"temperature,optional"`
	TopK              types.Int64                                          `tfsdk:"top_k" json:"top_k,optional"`
	TopP              types.Float64                                        `tfsdk:"top_p" json:"top_p,optional"`
	LogitBias         *map[string]types.Float64                            `tfsdk:"logit_bias" json:"logit_bias,optional"`
	Stop              *[]types.String                                      `tfsdk:"stop" json:"stop,optional"`
	Created           types.Int64                                          `tfsdk:"created" json:"created,computed"`
	Object            types.String                                         `tfsdk:"object" json:"object,computed"`
	Choices           customfield.NestedObjectList[CompletionChoicesModel] `tfsdk:"choices" json:"choices,computed"`
	Usage             customfield.NestedObject[CompletionUsageModel]       `tfsdk:"usage" json:"usage,computed"`
}

func (m CompletionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CompletionModel) MarshalJSONForUpdate(state CompletionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CompletionChoicesModel struct {
	FinishReason types.String                                             `tfsdk:"finish_reason" json:"finish_reason,computed"`
	Logprobs     customfield.NestedObject[CompletionChoicesLogprobsModel] `tfsdk:"logprobs" json:"logprobs,computed"`
	Seed         types.Int64                                              `tfsdk:"seed" json:"seed,computed"`
	Text         types.String                                             `tfsdk:"text" json:"text,computed"`
}

type CompletionChoicesLogprobsModel struct {
	TokenIDs      customfield.List[types.Float64] `tfsdk:"token_ids" json:"token_ids,computed"`
	TokenLogprobs customfield.List[types.Float64] `tfsdk:"token_logprobs" json:"token_logprobs,computed"`
	Tokens        customfield.List[types.String]  `tfsdk:"tokens" json:"tokens,computed"`
}

type CompletionUsageModel struct {
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" json:"completion_tokens,computed"`
	PromptTokens     types.Int64 `tfsdk:"prompt_tokens" json:"prompt_tokens,computed"`
	TotalTokens      types.Int64 `tfsdk:"total_tokens" json:"total_tokens,computed"`
}
