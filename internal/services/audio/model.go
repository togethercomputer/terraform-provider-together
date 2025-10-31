// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
)

type AudioModel struct {
	Input            types.String  `tfsdk:"input" json:"input,required"`
	Model            types.String  `tfsdk:"model" json:"model,required"`
	Voice            types.String  `tfsdk:"voice" json:"voice,required"`
	Stream           types.Bool    `tfsdk:"stream" json:"stream,optional"`
	Language         types.String  `tfsdk:"language" json:"language,computed_optional"`
	ResponseEncoding types.String  `tfsdk:"response_encoding" json:"response_encoding,computed_optional"`
	ResponseFormat   types.String  `tfsdk:"response_format" json:"response_format,computed_optional"`
	SampleRate       types.Float64 `tfsdk:"sample_rate" json:"sample_rate,computed_optional"`
}

func (m AudioModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AudioModel) MarshalJSONForUpdate(state AudioModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
