// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type VideoDataSourceModel struct {
	ID          types.String                                          `tfsdk:"id" path:"id,required"`
	CompletedAt types.Float64                                         `tfsdk:"completed_at" json:"completed_at,computed"`
	CreatedAt   types.Float64                                         `tfsdk:"created_at" json:"created_at,computed"`
	Model       types.String                                          `tfsdk:"model" json:"model,computed"`
	Object      types.String                                          `tfsdk:"object" json:"object,computed"`
	Seconds     types.String                                          `tfsdk:"seconds" json:"seconds,computed"`
	Size        types.String                                          `tfsdk:"size" json:"size,computed"`
	Status      types.String                                          `tfsdk:"status" json:"status,computed"`
	Error       customfield.NestedObject[VideoErrorDataSourceModel]   `tfsdk:"error" json:"error,computed"`
	Outputs     customfield.NestedObject[VideoOutputsDataSourceModel] `tfsdk:"outputs" json:"outputs,computed"`
}

type VideoErrorDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
	Code    types.String `tfsdk:"code" json:"code,computed"`
}

type VideoOutputsDataSourceModel struct {
	Cost     types.Int64  `tfsdk:"cost" json:"cost,computed"`
	VideoURL types.String `tfsdk:"video_url" json:"video_url,computed"`
}
