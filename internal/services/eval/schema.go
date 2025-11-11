// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package eval

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*EvalResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"type": schema.StringAttribute{
				Description: "The type of evaluation to perform\nAvailable values: \"classify\", \"score\", \"compare\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"classify",
						"score",
						"compare",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"parameters": schema.SingleNestedAttribute{
				Description: "Type-specific parameters for the evaluation",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"input_data_file_path": schema.StringAttribute{
						Description: "Data file ID",
						Required:    true,
					},
					"judge": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"model_name": schema.StringAttribute{
								Description: "Name of the judge model",
								Required:    true,
							},
							"system_template": schema.StringAttribute{
								Description: "System prompt template for the judge",
								Required:    true,
							},
						},
					},
					"labels": schema.ListAttribute{
						Description: "List of possible classification labels",
						Optional:    true,
						ElementType: types.StringType,
					},
					"pass_labels": schema.ListAttribute{
						Description: "List of labels that are considered passing",
						Optional:    true,
						ElementType: types.StringType,
					},
					"model_to_evaluate": schema.StringAttribute{
						Description: "Field name in the input data",
						Optional:    true,
					},
					"max_score": schema.Float64Attribute{
						Description: "Maximum possible score",
						Optional:    true,
					},
					"min_score": schema.Float64Attribute{
						Description: "Minimum possible score",
						Optional:    true,
					},
					"pass_threshold": schema.Float64Attribute{
						Description: "Score threshold for passing",
						Optional:    true,
					},
					"model_a": schema.StringAttribute{
						Description: "Field name in the input data",
						Optional:    true,
					},
					"model_b": schema.StringAttribute{
						Description: "Field name in the input data",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"error": schema.StringAttribute{
				Description: "Error message when status is 'error' or 'user_error'",
				Optional:    true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "completed", "error", "user_error", "running", "queued", "pending".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"completed",
						"error",
						"user_error",
						"running",
						"queued",
						"pending",
					),
				},
			},
			"results": schema.StringAttribute{
				Description: "The results of the evaluation job. The concrete structure depends on the type of evaluation job",
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
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
			"updated_at": schema.StringAttribute{
				Description: "When the job was last updated",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"workflow_id": schema.StringAttribute{
				Description: "The evaluation job ID",
				Computed:    true,
			},
			"status_updates": schema.ListNestedAttribute{
				Description: "History of status updates (admin only)",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[EvalStatusUpdatesModel](ctx),
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

func (r *EvalResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EvalResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
