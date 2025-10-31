// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package embedding

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type EmbeddingModel struct {
	Input  types.String                                     `tfsdk:"input" json:"input,required"`
	Model  types.String                                     `tfsdk:"model" json:"model,required"`
	Object types.String                                     `tfsdk:"object" json:"object,computed"`
	Data   customfield.NestedObjectList[EmbeddingDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m EmbeddingModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmbeddingModel) MarshalJSONForUpdate(state EmbeddingModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type EmbeddingDataModel struct {
	Embedding customfield.List[types.Float64] `tfsdk:"embedding" json:"embedding,computed"`
	Index     types.Int64                     `tfsdk:"index" json:"index,computed"`
	Object    types.String                    `tfsdk:"object" json:"object,computed"`
}
