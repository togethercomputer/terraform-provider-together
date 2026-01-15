// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*BetaClusterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cluster_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"billing_type": schema.StringAttribute{
				Description: `Available values: "RESERVED", "ON_DEMAND".`,
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("RESERVED", "ON_DEMAND"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cluster_name": schema.StringAttribute{
				Description:   "Name of the GPU cluster.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"driver_version": schema.StringAttribute{
				Description: "NVIDIA driver version to use in the cluster.\nAvailable values: \"CUDA_12_5_555\", \"CUDA_12_6_560\", \"CUDA_12_6_565\", \"CUDA_12_8_570\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"CUDA_12_5_555",
						"CUDA_12_6_560",
						"CUDA_12_6_565",
						"CUDA_12_8_570",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"gpu_type": schema.StringAttribute{
				Description: "Type of GPU to use in the cluster\nAvailable values: \"H100_SXM\", \"H200_SXM\", \"RTX_6000_PCI\", \"L40_PCIE\", \"B200_SXM\", \"H100_SXM_INF\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"H100_SXM",
						"H200_SXM",
						"RTX_6000_PCI",
						"L40_PCIE",
						"B200_SXM",
						"H100_SXM_INF",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description: "Region to create the GPU cluster in. Valid values are us-central-8 and us-central-4.\nAvailable values: \"us-central-8\", \"us-central-4\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("us-central-8", "us-central-4"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"duration_days": schema.Int64Attribute{
				Description:   "Duration in days to keep the cluster running.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"volume_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"shared_volume": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"region": schema.StringAttribute{
						Description: "Region name. Usable regions can be found from `client.clusters.list_regions()`",
						Required:    true,
					},
					"size_tib": schema.Int64Attribute{
						Description: "Volume size in whole tebibytes (TiB).",
						Required:    true,
					},
					"volume_name": schema.StringAttribute{
						Required: true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"num_gpus": schema.Int64Attribute{
				Description: "Number of GPUs to allocate in the cluster. This must be multiple of 8. For example, 8, 16 or 24",
				Required:    true,
			},
			"cluster_type": schema.StringAttribute{
				Description: `Available values: "KUBERNETES", "SLURM".`,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("KUBERNETES", "SLURM"),
				},
			},
			"duration_hours": schema.Int64Attribute{
				Computed: true,
			},
			"kube_config": schema.StringAttribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: "Current status of the GPU cluster.\nAvailable values: \"WaitingForControlPlaneNodes\", \"WaitingForDataPlaneNodes\", \"WaitingForSubnet\", \"WaitingForSharedVolume\", \"InstallingDrivers\", \"RunningAcceptanceTests\", \"Paused\", \"OnDemandComputePaused\", \"Ready\", \"Degraded\", \"Deleting\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"WaitingForControlPlaneNodes",
						"WaitingForDataPlaneNodes",
						"WaitingForSubnet",
						"WaitingForSharedVolume",
						"InstallingDrivers",
						"RunningAcceptanceTests",
						"Paused",
						"OnDemandComputePaused",
						"Ready",
						"Degraded",
						"Deleting",
					),
				},
			},
			"control_plane_nodes": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[BetaClusterControlPlaneNodesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							Computed: true,
						},
						"memory_gib": schema.Float64Attribute{
							Computed: true,
						},
						"network": schema.StringAttribute{
							Computed: true,
						},
						"node_id": schema.StringAttribute{
							Computed: true,
						},
						"node_name": schema.StringAttribute{
							Computed: true,
						},
						"num_cpu_cores": schema.Int64Attribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"gpu_worker_nodes": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[BetaClusterGPUWorkerNodesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"host_name": schema.StringAttribute{
							Computed: true,
						},
						"memory_gib": schema.Float64Attribute{
							Computed: true,
						},
						"networks": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"node_id": schema.StringAttribute{
							Computed: true,
						},
						"node_name": schema.StringAttribute{
							Computed: true,
						},
						"num_cpu_cores": schema.Int64Attribute{
							Computed: true,
						},
						"num_gpus": schema.Int64Attribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"volumes": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[BetaClusterVolumesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"size_tib": schema.Int64Attribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"volume_id": schema.StringAttribute{
							Computed: true,
						},
						"volume_name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *BetaClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BetaClusterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
