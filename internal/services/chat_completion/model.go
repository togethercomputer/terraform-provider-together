// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package chat_completion

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type ChatCompletionModel struct {
	ID                            types.String                                                `tfsdk:"id" json:"id,computed"`
	Model                         types.String                                                `tfsdk:"model" json:"model,required"`
	Messages                      *[]*ChatCompletionMessagesModel                             `tfsdk:"messages" json:"messages,required"`
	Echo                          types.Bool                                                  `tfsdk:"echo" json:"echo,optional"`
	FrequencyPenalty              types.Float64                                               `tfsdk:"frequency_penalty" json:"frequency_penalty,optional"`
	FunctionCall                  types.String                                                `tfsdk:"function_call" json:"function_call,optional"`
	Logprobs                      types.Int64                                                 `tfsdk:"logprobs" json:"logprobs,optional"`
	MaxTokens                     types.Int64                                                 `tfsdk:"max_tokens" json:"max_tokens,optional"`
	MinP                          types.Float64                                               `tfsdk:"min_p" json:"min_p,optional"`
	N                             types.Int64                                                 `tfsdk:"n" json:"n,optional"`
	PresencePenalty               types.Float64                                               `tfsdk:"presence_penalty" json:"presence_penalty,optional"`
	ReasoningEffort               types.String                                                `tfsdk:"reasoning_effort" json:"reasoning_effort,optional"`
	RepetitionPenalty             types.Float64                                               `tfsdk:"repetition_penalty" json:"repetition_penalty,optional"`
	SafetyModel                   types.String                                                `tfsdk:"safety_model" json:"safety_model,optional"`
	Seed                          types.Int64                                                 `tfsdk:"seed" json:"seed,optional"`
	Stream                        types.Bool                                                  `tfsdk:"stream" json:"stream,optional"`
	Temperature                   types.Float64                                               `tfsdk:"temperature" json:"temperature,optional"`
	ToolChoice                    types.String                                                `tfsdk:"tool_choice" json:"tool_choice,optional"`
	TopK                          types.Int64                                                 `tfsdk:"top_k" json:"top_k,optional"`
	TopP                          types.Float64                                               `tfsdk:"top_p" json:"top_p,optional"`
	LogitBias                     *map[string]types.Float64                                   `tfsdk:"logit_bias" json:"logit_bias,optional"`
	Stop                          *[]types.String                                             `tfsdk:"stop" json:"stop,optional"`
	Tools                         *[]*ChatCompletionToolsModel                                `tfsdk:"tools" json:"tools,optional"`
	ContextLengthExceededBehavior types.String                                                `tfsdk:"context_length_exceeded_behavior" json:"context_length_exceeded_behavior,computed_optional"`
	ResponseFormat                customfield.NestedObject[ChatCompletionResponseFormatModel] `tfsdk:"response_format" json:"response_format,computed_optional"`
	Created                       types.Int64                                                 `tfsdk:"created" json:"created,computed"`
	Object                        types.String                                                `tfsdk:"object" json:"object,computed"`
	Choices                       customfield.NestedObjectList[ChatCompletionChoicesModel]    `tfsdk:"choices" json:"choices,computed"`
	Usage                         customfield.NestedObject[ChatCompletionUsageModel]          `tfsdk:"usage" json:"usage,computed"`
	Warnings                      customfield.NestedObjectList[ChatCompletionWarningsModel]   `tfsdk:"warnings" json:"warnings,computed"`
}

func (m ChatCompletionModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ChatCompletionModel) MarshalJSONForUpdate(state ChatCompletionModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ChatCompletionMessagesModel struct {
	Content      types.String                             `tfsdk:"content" json:"content,optional"`
	Role         types.String                             `tfsdk:"role" json:"role,required"`
	Name         types.String                             `tfsdk:"name" json:"name,optional"`
	FunctionCall *ChatCompletionMessagesFunctionCallModel `tfsdk:"function_call" json:"function_call,optional"`
	ToolCalls    *[]*ChatCompletionMessagesToolCallsModel `tfsdk:"tool_calls" json:"tool_calls,optional"`
	ToolCallID   types.String                             `tfsdk:"tool_call_id" json:"tool_call_id,optional"`
}

type ChatCompletionMessagesFunctionCallModel struct {
	Arguments types.String `tfsdk:"arguments" json:"arguments,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
}

type ChatCompletionMessagesToolCallsModel struct {
	ID       types.String                                  `tfsdk:"id" json:"id,required"`
	Function *ChatCompletionMessagesToolCallsFunctionModel `tfsdk:"function" json:"function,required"`
	Index    types.Float64                                 `tfsdk:"index" json:"index,required"`
	Type     types.String                                  `tfsdk:"type" json:"type,required"`
}

type ChatCompletionMessagesToolCallsFunctionModel struct {
	Arguments types.String `tfsdk:"arguments" json:"arguments,required"`
	Name      types.String `tfsdk:"name" json:"name,required"`
}

type ChatCompletionToolsModel struct {
	Function *ChatCompletionToolsFunctionModel `tfsdk:"function" json:"function,optional"`
	Type     types.String                      `tfsdk:"type" json:"type,optional"`
}

type ChatCompletionToolsFunctionModel struct {
	Description types.String                     `tfsdk:"description" json:"description,optional"`
	Name        types.String                     `tfsdk:"name" json:"name,optional"`
	Parameters  *map[string]jsontypes.Normalized `tfsdk:"parameters" json:"parameters,optional"`
}

type ChatCompletionResponseFormatModel struct {
	Type       types.String                                                          `tfsdk:"type" json:"type,required"`
	JsonSchema customfield.NestedObject[ChatCompletionResponseFormatJsonSchemaModel] `tfsdk:"json_schema" json:"json_schema,computed_optional"`
}

type ChatCompletionResponseFormatJsonSchemaModel struct {
	Name        types.String                     `tfsdk:"name" json:"name,required"`
	Description types.String                     `tfsdk:"description" json:"description,optional"`
	Schema      *map[string]jsontypes.Normalized `tfsdk:"schema" json:"schema,optional"`
	Strict      types.Bool                       `tfsdk:"strict" json:"strict,computed_optional"`
}

type ChatCompletionChoicesModel struct {
	FinishReason types.String                                                 `tfsdk:"finish_reason" json:"finish_reason,computed"`
	Index        types.Int64                                                  `tfsdk:"index" json:"index,computed"`
	Logprobs     customfield.NestedObject[ChatCompletionChoicesLogprobsModel] `tfsdk:"logprobs" json:"logprobs,computed"`
	Message      customfield.NestedObject[ChatCompletionChoicesMessageModel]  `tfsdk:"message" json:"message,computed"`
	Seed         types.Int64                                                  `tfsdk:"seed" json:"seed,computed"`
	Text         types.String                                                 `tfsdk:"text" json:"text,computed"`
}

type ChatCompletionChoicesLogprobsModel struct {
	TokenIDs      customfield.List[types.Float64] `tfsdk:"token_ids" json:"token_ids,computed"`
	TokenLogprobs customfield.List[types.Float64] `tfsdk:"token_logprobs" json:"token_logprobs,computed"`
	Tokens        customfield.List[types.String]  `tfsdk:"tokens" json:"tokens,computed"`
}

type ChatCompletionChoicesMessageModel struct {
	Content      types.String                                                             `tfsdk:"content" json:"content,computed"`
	Role         types.String                                                             `tfsdk:"role" json:"role,computed"`
	FunctionCall customfield.NestedObject[ChatCompletionChoicesMessageFunctionCallModel]  `tfsdk:"function_call" json:"function_call,computed"`
	Reasoning    types.String                                                             `tfsdk:"reasoning" json:"reasoning,computed"`
	ToolCalls    customfield.NestedObjectList[ChatCompletionChoicesMessageToolCallsModel] `tfsdk:"tool_calls" json:"tool_calls,computed"`
}

type ChatCompletionChoicesMessageFunctionCallModel struct {
	Arguments types.String `tfsdk:"arguments" json:"arguments,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
}

type ChatCompletionChoicesMessageToolCallsModel struct {
	ID       types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Function customfield.NestedObject[ChatCompletionChoicesMessageToolCallsFunctionModel] `tfsdk:"function" json:"function,computed"`
	Index    types.Float64                                                                `tfsdk:"index" json:"index,computed"`
	Type     types.String                                                                 `tfsdk:"type" json:"type,computed"`
}

type ChatCompletionChoicesMessageToolCallsFunctionModel struct {
	Arguments types.String `tfsdk:"arguments" json:"arguments,computed"`
	Name      types.String `tfsdk:"name" json:"name,computed"`
}

type ChatCompletionUsageModel struct {
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" json:"completion_tokens,computed"`
	PromptTokens     types.Int64 `tfsdk:"prompt_tokens" json:"prompt_tokens,computed"`
	TotalTokens      types.Int64 `tfsdk:"total_tokens" json:"total_tokens,computed"`
}

type ChatCompletionWarningsModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}
