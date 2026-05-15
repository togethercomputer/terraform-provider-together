// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*BetaClusterStorageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"volume_id": schema.StringAttribute{
				Description: "The ID of the volume to retrieve",
				Required:    true,
			},
			"size_tib": schema.Int64Attribute{
				Description: "Size of the volume in TiB.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current status of the shared volume.\nAvailable values: \"scheduled\", \"available\", \"bound\", \"provisioning\", \"deleting\", \"failed\", \"access_revoked\", \"unknown\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"scheduled",
						"available",
						"bound",
						"provisioning",
						"deleting",
						"failed",
						"access_revoked",
						"unknown",
					),
				},
			},
			"volume_name": schema.StringAttribute{
				Description: "User provided name of the volume.",
				Computed:    true,
			},
		},
	}
}

func (d *BetaClusterStorageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BetaClusterStorageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
