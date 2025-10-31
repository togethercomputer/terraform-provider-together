// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package job

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type JobDataSourceModel struct {
	JobID         types.String                                                  `tfsdk:"job_id" path:"jobId,required"`
	CreatedAt     timetypes.RFC3339                                             `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Status        types.String                                                  `tfsdk:"status" json:"status,computed"`
	Type          types.String                                                  `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                             `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Args          customfield.NestedObject[JobArgsDataSourceModel]              `tfsdk:"args" json:"args,computed"`
	StatusUpdates customfield.NestedObjectList[JobStatusUpdatesDataSourceModel] `tfsdk:"status_updates" json:"status_updates,computed"`
}

type JobArgsDataSourceModel struct {
	Description types.String `tfsdk:"description" json:"description,computed"`
	ModelName   types.String `tfsdk:"model_name" json:"modelName,computed"`
	ModelSource types.String `tfsdk:"model_source" json:"modelSource,computed"`
}

type JobStatusUpdatesDataSourceModel struct {
	Message   types.String      `tfsdk:"message" json:"message,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
}
