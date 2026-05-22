// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*BetaClusterStorageResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "ID of the volume.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"volume_id": schema.StringAttribute{
				Description:   "ID of the volume.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"region": schema.StringAttribute{
				Description:   "Region name. Usable regions can be found from `clusters.list_regions()`",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"volume_name": schema.StringAttribute{
				Description:   "User provided name of the volume.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"is_lifecycle_independent": schema.BoolAttribute{
				Description:   "When true, the shared volume is not deleted when the cluster is decommissioned.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"size_tib": schema.Int64Attribute{
				Description: "Volume size in whole tebibytes (TiB).",
				Required:    true,
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
		},
	}
}

func (r *BetaClusterStorageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BetaClusterStorageResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
