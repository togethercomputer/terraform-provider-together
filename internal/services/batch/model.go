// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package batch

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type BatchModel struct {
	ID               types.String                            `tfsdk:"id" path:"id,optional"`
	Endpoint         types.String                            `tfsdk:"endpoint" json:"endpoint,required"`
	InputFileID      types.String                            `tfsdk:"input_file_id" json:"input_file_id,required"`
	CompletionWindow types.String                            `tfsdk:"completion_window" json:"completion_window,optional,no_refresh"`
	ModelID          types.String                            `tfsdk:"model_id" json:"model_id,optional"`
	Priority         types.Int64                             `tfsdk:"priority" json:"priority,optional,no_refresh"`
	CompletedAt      timetypes.RFC3339                       `tfsdk:"completed_at" json:"completed_at,computed" format:"date-time"`
	CreatedAt        timetypes.RFC3339                       `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Error            types.String                            `tfsdk:"error" json:"error,computed"`
	ErrorFileID      types.String                            `tfsdk:"error_file_id" json:"error_file_id,computed"`
	FileSizeBytes    types.Int64                             `tfsdk:"file_size_bytes" json:"file_size_bytes,computed"`
	JobDeadline      timetypes.RFC3339                       `tfsdk:"job_deadline" json:"job_deadline,computed" format:"date-time"`
	OutputFileID     types.String                            `tfsdk:"output_file_id" json:"output_file_id,computed"`
	Progress         types.Float64                           `tfsdk:"progress" json:"progress,computed"`
	Status           types.String                            `tfsdk:"status" json:"status,computed"`
	UserID           types.String                            `tfsdk:"user_id" json:"user_id,computed"`
	Warning          types.String                            `tfsdk:"warning" json:"warning,computed,no_refresh"`
	Job              customfield.NestedObject[BatchJobModel] `tfsdk:"job" json:"job,computed,no_refresh"`
}

func (m BatchModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BatchModel) MarshalJSONForUpdate(state BatchModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type BatchJobModel struct {
	ID            types.String      `tfsdk:"id" json:"id,computed"`
	CompletedAt   timetypes.RFC3339 `tfsdk:"completed_at" json:"completed_at,computed" format:"date-time"`
	CreatedAt     timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Endpoint      types.String      `tfsdk:"endpoint" json:"endpoint,computed"`
	Error         types.String      `tfsdk:"error" json:"error,computed"`
	ErrorFileID   types.String      `tfsdk:"error_file_id" json:"error_file_id,computed"`
	FileSizeBytes types.Int64       `tfsdk:"file_size_bytes" json:"file_size_bytes,computed"`
	InputFileID   types.String      `tfsdk:"input_file_id" json:"input_file_id,computed"`
	JobDeadline   timetypes.RFC3339 `tfsdk:"job_deadline" json:"job_deadline,computed" format:"date-time"`
	ModelID       types.String      `tfsdk:"model_id" json:"model_id,computed"`
	OutputFileID  types.String      `tfsdk:"output_file_id" json:"output_file_id,computed"`
	Progress      types.Float64     `tfsdk:"progress" json:"progress,computed"`
	Status        types.String      `tfsdk:"status" json:"status,computed"`
	UserID        types.String      `tfsdk:"user_id" json:"user_id,computed"`
}
