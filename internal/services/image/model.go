// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type ImageModel struct {
	ID                   types.String                                 `tfsdk:"id" json:"id,computed"`
	Model                types.String                                 `tfsdk:"model" json:"model,required"`
	Prompt               types.String                                 `tfsdk:"prompt" json:"prompt,required"`
	DisableSafetyChecker types.Bool                                   `tfsdk:"disable_safety_checker" json:"disable_safety_checker,optional"`
	ImageURL             types.String                                 `tfsdk:"image_url" json:"image_url,optional"`
	NegativePrompt       types.String                                 `tfsdk:"negative_prompt" json:"negative_prompt,optional"`
	ResponseFormat       types.String                                 `tfsdk:"response_format" json:"response_format,optional"`
	Seed                 types.Int64                                  `tfsdk:"seed" json:"seed,optional"`
	ImageLoras           *[]*ImageImageLorasModel                     `tfsdk:"image_loras" json:"image_loras,optional"`
	GuidanceScale        types.Float64                                `tfsdk:"guidance_scale" json:"guidance_scale,computed_optional"`
	Height               types.Int64                                  `tfsdk:"height" json:"height,computed_optional"`
	N                    types.Int64                                  `tfsdk:"n" json:"n,computed_optional"`
	OutputFormat         types.String                                 `tfsdk:"output_format" json:"output_format,computed_optional"`
	Steps                types.Int64                                  `tfsdk:"steps" json:"steps,computed_optional"`
	Width                types.Int64                                  `tfsdk:"width" json:"width,computed_optional"`
	Object               types.String                                 `tfsdk:"object" json:"object,computed"`
	Data                 customfield.NestedObjectList[ImageDataModel] `tfsdk:"data" json:"data,computed"`
}

func (m ImageModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ImageModel) MarshalJSONForUpdate(state ImageModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ImageImageLorasModel struct {
	Path  types.String  `tfsdk:"path" json:"path,required"`
	Scale types.Float64 `tfsdk:"scale" json:"scale,required"`
}

type ImageDataModel struct {
	B64Json types.String `tfsdk:"b64_json" json:"b64_json,computed"`
	Index   types.Int64  `tfsdk:"index" json:"index,computed"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
	URL     types.String `tfsdk:"url" json:"url,computed"`
}
