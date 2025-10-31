// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*EndpointDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"endpoint_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp when the endpoint was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"display_name": schema.StringAttribute{
				Description: "Human-readable name for the endpoint",
				Computed:    true,
			},
			"hardware": schema.StringAttribute{
				Description: "The hardware configuration used for this endpoint",
				Computed:    true,
			},
			"model": schema.StringAttribute{
				Description: "The model deployed on this endpoint",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "System name for the endpoint",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Description: "The type of object\nAvailable values: \"endpoint\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("endpoint"),
				},
			},
			"owner": schema.StringAttribute{
				Description: "The owner of this endpoint",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "Current state of the endpoint\nAvailable values: \"PENDING\", \"STARTING\", \"STARTED\", \"STOPPING\", \"STOPPED\", \"ERROR\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"PENDING",
						"STARTING",
						"STARTED",
						"STOPPING",
						"STOPPED",
						"ERROR",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of endpoint\nAvailable values: \"dedicated\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("dedicated"),
				},
			},
			"autoscaling": schema.SingleNestedAttribute{
				Description: "Configuration for automatic scaling of the endpoint",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[EndpointAutoscalingDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"max_replicas": schema.Int64Attribute{
						Description: "The maximum number of replicas to scale up to under load",
						Computed:    true,
					},
					"min_replicas": schema.Int64Attribute{
						Description: "The minimum number of replicas to maintain, even when there is no load",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *EndpointDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *EndpointDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
