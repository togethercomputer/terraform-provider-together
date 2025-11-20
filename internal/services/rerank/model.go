// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rerank

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type RerankModel struct {
	ID              types.String                                     `tfsdk:"id" json:"id,computed"`
	Model           types.String                                     `tfsdk:"model" json:"model,required"`
	Query           types.String                                     `tfsdk:"query" json:"query,required"`
	Documents       *[]*map[string]jsontypes.Normalized              `tfsdk:"documents" json:"documents,required"`
	ReturnDocuments types.Bool                                       `tfsdk:"return_documents" json:"return_documents,optional"`
	TopN            types.Int64                                      `tfsdk:"top_n" json:"top_n,optional"`
	RankFields      *[]types.String                                  `tfsdk:"rank_fields" json:"rank_fields,optional"`
	Object          types.String                                     `tfsdk:"object" json:"object,computed"`
	Results         customfield.NestedObjectList[RerankResultsModel] `tfsdk:"results" json:"results,computed"`
	Usage           customfield.NestedObject[RerankUsageModel]       `tfsdk:"usage" json:"usage,computed"`
}

func (m RerankModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m RerankModel) MarshalJSONForUpdate(state RerankModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type RerankResultsModel struct {
	Document       customfield.NestedObject[RerankResultsDocumentModel] `tfsdk:"document" json:"document,computed"`
	Index          types.Int64                                          `tfsdk:"index" json:"index,computed"`
	RelevanceScore types.Float64                                        `tfsdk:"relevance_score" json:"relevance_score,computed"`
}

type RerankResultsDocumentModel struct {
	Text types.String `tfsdk:"text" json:"text,computed"`
}

type RerankUsageModel struct {
	CompletionTokens types.Int64 `tfsdk:"completion_tokens" json:"completion_tokens,computed"`
	PromptTokens     types.Int64 `tfsdk:"prompt_tokens" json:"prompt_tokens,computed"`
	TotalTokens      types.Int64 `tfsdk:"total_tokens" json:"total_tokens,computed"`
}
