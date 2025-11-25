// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type EvalModel struct {
	ID            types.String                                         `tfsdk:"id" path:"id,optional"`
	Type          types.String                                         `tfsdk:"type" json:"type,required"`
	Parameters    *EvalParametersModel                                 `tfsdk:"parameters" json:"parameters,required,no_refresh"`
	CreatedAt     timetypes.RFC3339                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	OwnerID       types.String                                         `tfsdk:"owner_id" json:"owner_id,computed"`
	Status        types.String                                         `tfsdk:"status" json:"status,computed"`
	UpdatedAt     timetypes.RFC3339                                    `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	WorkflowID    types.String                                         `tfsdk:"workflow_id" json:"workflow_id,computed"`
	Results       customfield.NestedObject[EvalResultsModel]           `tfsdk:"results" json:"results,computed"`
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

type EvalResultsModel struct {
	GenerationFailCount types.Float64                                              `tfsdk:"generation_fail_count" json:"generation_fail_count,computed"`
	InvalidLabelCount   types.Float64                                              `tfsdk:"invalid_label_count" json:"invalid_label_count,computed"`
	JudgeFailCount      types.Float64                                              `tfsdk:"judge_fail_count" json:"judge_fail_count,computed"`
	LabelCounts         types.String                                               `tfsdk:"label_counts" json:"label_counts,computed"`
	PassPercentage      types.Float64                                              `tfsdk:"pass_percentage" json:"pass_percentage,computed"`
	ResultFileID        types.String                                               `tfsdk:"result_file_id" json:"result_file_id,computed"`
	AggregatedScores    customfield.NestedObject[EvalResultsAggregatedScoresModel] `tfsdk:"aggregated_scores" json:"aggregated_scores,computed"`
	FailedSamples       types.Float64                                              `tfsdk:"failed_samples" json:"failed_samples,computed"`
	InvalidScoreCount   types.Float64                                              `tfsdk:"invalid_score_count" json:"invalid_score_count,computed"`
	AWins               types.Int64                                                `tfsdk:"a_wins" json:"A_wins,computed"`
	BWins               types.Int64                                                `tfsdk:"b_wins" json:"B_wins,computed"`
	NumSamples          types.Int64                                                `tfsdk:"num_samples" json:"num_samples,computed"`
	Ties                types.Int64                                                `tfsdk:"ties" json:"Ties,computed"`
	Error               types.String                                               `tfsdk:"error" json:"error,computed"`
}

type EvalResultsAggregatedScoresModel struct {
	MeanScore      types.Float64 `tfsdk:"mean_score" json:"mean_score,computed"`
	PassPercentage types.Float64 `tfsdk:"pass_percentage" json:"pass_percentage,computed"`
	StdScore       types.Float64 `tfsdk:"std_score" json:"std_score,computed"`
}

type EvalStatusUpdatesModel struct {
	Message   types.String      `tfsdk:"message" json:"message,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
}
