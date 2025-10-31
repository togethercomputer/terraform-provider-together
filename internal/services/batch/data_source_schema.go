// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*BatchDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
	}
}

func (d *BatchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BatchDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
