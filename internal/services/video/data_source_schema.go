// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*VideoDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"completed_at": schema.Float64Attribute{
				Description: "Unix timestamp (seconds) for when the job completed, if finished.",
				Computed:    true,
			},
			"created_at": schema.Float64Attribute{
				Description: "Unix timestamp (seconds) for when the job was created.",
				Computed:    true,
			},
			"model": schema.StringAttribute{
				Description: "The video generation model that produced the job.",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Description: "The object type, which is always video.\nAvailable values: \"video\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("video"),
				},
			},
			"seconds": schema.StringAttribute{
				Description: "Duration of the generated clip in seconds.",
				Computed:    true,
			},
			"size": schema.StringAttribute{
				Description: "The resolution of the generated video.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current lifecycle status of the video job.\nAvailable values: \"in_progress\", \"completed\", \"failed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"in_progress",
						"completed",
						"failed",
					),
				},
			},
			"error": schema.SingleNestedAttribute{
				Description: "Error payload that explains why generation failed, if applicable.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VideoErrorDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"message": schema.StringAttribute{
						Computed: true,
					},
					"code": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"outputs": schema.SingleNestedAttribute{
				Description: "Available upon completion, the outputs provides the cost charged and the hosted url to access the video",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VideoOutputsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cost": schema.Int64Attribute{
						Description: "The cost of generated video charged to the owners account.",
						Computed:    true,
					},
					"video_url": schema.StringAttribute{
						Description: "URL hosting the generated video",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *VideoDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *VideoDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
