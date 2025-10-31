// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type EvalDataSourceModel struct {
	ID            types.String                                                   `tfsdk:"id" path:"id,required"`
	CreatedAt     timetypes.RFC3339                                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	OwnerID       types.String                                                   `tfsdk:"owner_id" json:"owner_id,computed"`
	Status        types.String                                                   `tfsdk:"status" json:"status,computed"`
	Type          types.String                                                   `tfsdk:"type" json:"type,computed"`
	UpdatedAt     timetypes.RFC3339                                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	WorkflowID    types.String                                                   `tfsdk:"workflow_id" json:"workflow_id,computed"`
	Parameters    customfield.Map[jsontypes.Normalized]                          `tfsdk:"parameters" json:"parameters,computed"`
	Results       customfield.NestedObject[EvalResultsDataSourceModel]           `tfsdk:"results" json:"results,computed"`
	StatusUpdates customfield.NestedObjectList[EvalStatusUpdatesDataSourceModel] `tfsdk:"status_updates" json:"status_updates,computed"`
}

type EvalResultsDataSourceModel struct {
	GenerationFailCount types.Float64                                                        `tfsdk:"generation_fail_count" json:"generation_fail_count,computed"`
	InvalidLabelCount   types.Float64                                                        `tfsdk:"invalid_label_count" json:"invalid_label_count,computed"`
	JudgeFailCount      types.Float64                                                        `tfsdk:"judge_fail_count" json:"judge_fail_count,computed"`
	LabelCounts         types.String                                                         `tfsdk:"label_counts" json:"label_counts,computed"`
	PassPercentage      types.Float64                                                        `tfsdk:"pass_percentage" json:"pass_percentage,computed"`
	ResultFileID        types.String                                                         `tfsdk:"result_file_id" json:"result_file_id,computed"`
	AggregatedScores    customfield.NestedObject[EvalResultsAggregatedScoresDataSourceModel] `tfsdk:"aggregated_scores" json:"aggregated_scores,computed"`
	FailedSamples       types.Float64                                                        `tfsdk:"failed_samples" json:"failed_samples,computed"`
	InvalidScoreCount   types.Float64                                                        `tfsdk:"invalid_score_count" json:"invalid_score_count,computed"`
	AWins               types.Int64                                                          `tfsdk:"a_wins" json:"A_wins,computed"`
	BWins               types.Int64                                                          `tfsdk:"b_wins" json:"B_wins,computed"`
	NumSamples          types.Int64                                                          `tfsdk:"num_samples" json:"num_samples,computed"`
	Ties                types.Int64                                                          `tfsdk:"ties" json:"Ties,computed"`
	Error               types.String                                                         `tfsdk:"error" json:"error,computed"`
}

type EvalResultsAggregatedScoresDataSourceModel struct {
	MeanScore      types.Float64 `tfsdk:"mean_score" json:"mean_score,computed"`
	PassPercentage types.Float64 `tfsdk:"pass_percentage" json:"pass_percentage,computed"`
	StdScore       types.Float64 `tfsdk:"std_score" json:"std_score,computed"`
}

type EvalStatusUpdatesDataSourceModel struct {
	Message   types.String      `tfsdk:"message" json:"message,computed"`
	Status    types.String      `tfsdk:"status" json:"status,computed"`
	Timestamp timetypes.RFC3339 `tfsdk:"timestamp" json:"timestamp,computed" format:"date-time"`
}
