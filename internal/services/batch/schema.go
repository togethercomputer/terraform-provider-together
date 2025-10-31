// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*BatchResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"endpoint": schema.StringAttribute{
				Description:   "The endpoint to use for batch processing",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"input_file_id": schema.StringAttribute{
				Description:   "ID of the uploaded input file containing batch requests",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"completion_window": schema.StringAttribute{
				Description:   "Time window for batch completion (optional)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"model_id": schema.StringAttribute{
				Description:   "Model to use for processing batch requests",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"priority": schema.Int64Attribute{
				Description:   "Priority for batch processing (optional)",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"completed_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"error": schema.StringAttribute{
				Computed: true,
			},
			"error_file_id": schema.StringAttribute{
				Computed: true,
			},
			"file_size_bytes": schema.Int64Attribute{
				Description: "Size of input file in bytes",
				Computed:    true,
			},
			"job_deadline": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"output_file_id": schema.StringAttribute{
				Computed: true,
			},
			"progress": schema.Float64Attribute{
				Description: "Completion progress (0.0 to 100)",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current status of the batch job\nAvailable values: \"VALIDATING\", \"IN_PROGRESS\", \"COMPLETED\", \"FAILED\", \"EXPIRED\", \"CANCELLED\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"VALIDATING",
						"IN_PROGRESS",
						"COMPLETED",
						"FAILED",
						"EXPIRED",
						"CANCELLED",
					),
				},
			},
			"user_id": schema.StringAttribute{
				Computed: true,
			},
			"warning": schema.StringAttribute{
				Computed: true,
			},
			"job": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[BatchJobModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
					"completed_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"created_at": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"endpoint": schema.StringAttribute{
						Computed: true,
					},
					"error": schema.StringAttribute{
						Computed: true,
					},
					"error_file_id": schema.StringAttribute{
						Computed: true,
					},
					"file_size_bytes": schema.Int64Attribute{
						Description: "Size of input file in bytes",
						Computed:    true,
					},
					"input_file_id": schema.StringAttribute{
						Computed: true,
					},
					"job_deadline": schema.StringAttribute{
						Computed:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"model_id": schema.StringAttribute{
						Description: "Model used for processing requests",
						Computed:    true,
					},
					"output_file_id": schema.StringAttribute{
						Computed: true,
					},
					"progress": schema.Float64Attribute{
						Description: "Completion progress (0.0 to 100)",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Current status of the batch job\nAvailable values: \"VALIDATING\", \"IN_PROGRESS\", \"COMPLETED\", \"FAILED\", \"EXPIRED\", \"CANCELLED\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"VALIDATING",
								"IN_PROGRESS",
								"COMPLETED",
								"FAILED",
								"EXPIRED",
								"CANCELLED",
							),
						},
					},
					"user_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *BatchResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BatchResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
