// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*EvalDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "When the job was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"owner_id": schema.StringAttribute{
				Description: "ID of the job owner (admin only)",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current status of the job\nAvailable values: \"pending\", \"queued\", \"running\", \"completed\", \"error\", \"user_error\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"queued",
						"running",
						"completed",
						"error",
						"user_error",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of evaluation\nAvailable values: \"classify\", \"score\", \"compare\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"classify",
						"score",
						"compare",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the job was last updated",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"workflow_id": schema.StringAttribute{
				Description: "The evaluation job ID",
				Computed:    true,
			},
			"parameters": schema.MapAttribute{
				Description: "The parameters used for this evaluation",
				Computed:    true,
				CustomType:  customfield.NewMapType[jsontypes.Normalized](ctx),
				ElementType: jsontypes.NormalizedType{},
			},
			"results": schema.SingleNestedAttribute{
				Description: "Results of the evaluation (when completed)",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[EvalResultsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"generation_fail_count": schema.Float64Attribute{
						Description: "Number of failed generations.",
						Computed:    true,
					},
					"invalid_label_count": schema.Float64Attribute{
						Description: "Number of invalid labels",
						Computed:    true,
					},
					"judge_fail_count": schema.Float64Attribute{
						Description: "Number of failed judge generations",
						Computed:    true,
					},
					"label_counts": schema.StringAttribute{
						Description: "JSON string representing label counts",
						Computed:    true,
					},
					"pass_percentage": schema.Float64Attribute{
						Description: "Pecentage of pass labels.",
						Computed:    true,
					},
					"result_file_id": schema.StringAttribute{
						Description: "Data File ID",
						Computed:    true,
					},
					"aggregated_scores": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[EvalResultsAggregatedScoresDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"mean_score": schema.Float64Attribute{
								Computed: true,
							},
							"pass_percentage": schema.Float64Attribute{
								Computed: true,
							},
							"std_score": schema.Float64Attribute{
								Computed: true,
							},
						},
					},
					"failed_samples": schema.Float64Attribute{
						Description: "number of failed samples generated from model",
						Computed:    true,
					},
					"invalid_score_count": schema.Float64Attribute{
						Description: "number of invalid scores generated from model",
						Computed:    true,
					},
					"a_wins": schema.Int64Attribute{
						Description: "Number of times model A won",
						Computed:    true,
					},
					"b_wins": schema.Int64Attribute{
						Description: "Number of times model B won",
						Computed:    true,
					},
					"num_samples": schema.Int64Attribute{
						Description: "Total number of samples compared",
						Computed:    true,
					},
					"ties": schema.Int64Attribute{
						Description: "Number of ties",
						Computed:    true,
					},
					"error": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"status_updates": schema.ListNestedAttribute{
				Description: "History of status updates (admin only)",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EvalStatusUpdatesDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"message": schema.StringAttribute{
							Description: "Additional message for this update",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "The status at this update",
							Computed:    true,
						},
						"timestamp": schema.StringAttribute{
							Description: "When this update occurred",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *EvalDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EvalDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
