// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type EvalModel struct {
	ID            types.String                                         `tfsdk:"id" path:"id,optional"`
	Type          types.String                                         `tfsdk:"type" json:"type,required"`
	Parameters    *EvalParametersModel                                 `tfsdk:"parameters" json:"parameters,required,no_refresh"`
	Error         types.String                                         `tfsdk:"error" json:"error,optional,no_refresh"`
	Status        types.String                                         `tfsdk:"status" json:"status,optional"`
	Results       jsontypes.Normalized                                 `tfsdk:"results" json:"results,optional,no_refresh"`
	CreatedAt     timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	OwnerID       types.String                                         `tfsdk:"owner_id" json:"owner_id,computed"`
	UpdatedAt     timetypes.RFC3339                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	WorkflowID    types.String                                         `tfsdk:"workflow_id" json:"workflow_id,computed"`
	StatusUpdates customfield.NestedObjectList[EvalStatusUpdatesModel] `tfsdk:"status_updates" json:"status_updates,computed"`
}

func (m EvalModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EvalModel) MarshalJSONForUpdate(state EvalModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type EvalParametersModel struct {
	InputDataFilePath types.String              `tfsdk:"input_data_file_path" json:"input_data_file_path,required"`
	Judge             *EvalParametersJudgeModel `tfsdk:"judge" json:"judge,required"`
	Labels            *[]types.String           `tfsdk:"labels" json:"labels,optional"`
	PassLabels        *[]types.String           `tfsdk:"pass_labels" json:"pass_labels,optional"`
	ModelToEvaluate   types.String              `tfsdk:"model_to_evaluate" json:"model_to_evaluate,optional"`
	MaxScore          types.Float64             `tfsdk:"max_score" json:"max_score,optional"`
	MinScore          types.Float64             `tfsdk:"min_score" json:"min_score,optional"`
	PassThreshold     types.Float64             `tfsdk:"pass_threshold" json:"pass_threshold,optional"`
	ModelA            types.String              `tfsdk:"model_a" json:"model_a,optional"`
	ModelB            types.String              `tfsdk:"model_b" json:"model_b,optional"`
}

type EvalParametersJudgeModel struct {
	Model            types.String `tfsdk:"model" json:"model,required"`
	ModelSource      types.String `tfsdk:"model_source" json:"model_source,required"`
	SystemTemplate   types.String `tfsdk:"system_template" json:"system_template,required"`
	ExternalAPIToken types.String `tfsdk:"external_api_token" json:"external_api_token,optional"`
	ExternalBaseURL  types.String `tfsdk:"external_base_url" json:"external_base_url,optional"`
}

type EvalStatusUpdatesModel struct {
	Message   types.String      `tfsdk:"message" json:"message,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
}
