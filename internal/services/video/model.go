// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type VideoModel struct {
	ID              types.String                                `tfsdk:"id" json:"id,computed"`
	Model           types.String                                `tfsdk:"model" json:"model,required"`
	Fps             types.Int64                                 `tfsdk:"fps" json:"fps,optional,no_refresh"`
	GuidanceScale   types.Int64                                 `tfsdk:"guidance_scale" json:"guidance_scale,optional,no_refresh"`
	Height          types.Int64                                 `tfsdk:"height" json:"height,optional,no_refresh"`
	NegativePrompt  types.String                                `tfsdk:"negative_prompt" json:"negative_prompt,optional,no_refresh"`
	OutputFormat    types.String                                `tfsdk:"output_format" json:"output_format,optional,no_refresh"`
	OutputQuality   types.Int64                                 `tfsdk:"output_quality" json:"output_quality,optional,no_refresh"`
	Prompt          types.String                                `tfsdk:"prompt" json:"prompt,optional,no_refresh"`
	Seconds         types.String                                `tfsdk:"seconds" json:"seconds,optional"`
	Seed            types.Int64                                 `tfsdk:"seed" json:"seed,optional,no_refresh"`
	Steps           types.Int64                                 `tfsdk:"steps" json:"steps,optional,no_refresh"`
	Width           types.Int64                                 `tfsdk:"width" json:"width,optional,no_refresh"`
	ReferenceImages *[]types.String                             `tfsdk:"reference_images" json:"reference_images,optional,no_refresh"`
	FrameImages     *[]*VideoFrameImagesModel                   `tfsdk:"frame_images" json:"frame_images,optional,no_refresh"`
	CompletedAt     types.Float64                               `tfsdk:"completed_at" json:"completed_at,computed"`
	CreatedAt       types.Float64                               `tfsdk:"created_at" json:"created_at,computed"`
	Object          types.String                                `tfsdk:"object" json:"object,computed"`
	Size            types.String                                `tfsdk:"size" json:"size,computed"`
	Status          types.String                                `tfsdk:"status" json:"status,computed"`
	Error           customfield.NestedObject[VideoErrorModel]   `tfsdk:"error" json:"error,computed"`
	Outputs         customfield.NestedObject[VideoOutputsModel] `tfsdk:"outputs" json:"outputs,computed"`
}

func (m VideoModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m VideoModel) MarshalJSONForUpdate(state VideoModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type VideoFrameImagesModel struct {
	InputImage types.String                       `tfsdk:"input_image" json:"input_image,required"`
	Frame      customfield.NormalizedDynamicValue `tfsdk:"frame" json:"frame,optional"`
}

type VideoErrorModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
	Code    types.String `tfsdk:"code" json:"code,computed"`
}

type VideoOutputsModel struct {
	Cost     types.Int64  `tfsdk:"cost" json:"cost,computed"`
	VideoURL types.String `tfsdk:"video_url" json:"video_url,computed"`
}
