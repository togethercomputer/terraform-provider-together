// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
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
			"cluster_name": schema.StringAttribute{
				Description:   "Name of the GPU cluster.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"cuda_version": schema.StringAttribute{
				Description:   "CUDA version for this cluster. For example, 12.5",
				Required:      true,
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
			"nvidia_driver_version": schema.StringAttribute{
				Description:   "Nvidia driver version for this cluster. For example, 550. Only some combination of cuda_version and nvidia_driver_version are supported.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description:   "Region to create the GPU cluster in. Usable regions can be found from `client.clusters.list_regions()`",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auto_scale_max_gpus": schema.Int64Attribute{
				Description:   "Maximum number of GPUs to which the cluster can be auto-scaled up. This field is required if auto_scaled is true.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"capacity_pool_id": schema.StringAttribute{
				Description:   "ID of the capacity pool to use for the cluster. This field is optional and only applicable if the cluster is created from a capacity pool.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"reservation_start_time": schema.StringAttribute{
				Description:   "Reservation start time of the cluster. This field is required for SCHEDULED billing to specify the reservation start time for the cluster. If not provided, the cluster provisions immediately.",
				Optional:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"slurm_image": schema.StringAttribute{
				Description:   "Custom Slurm image for Slurm clusters.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"slurm_shm_size_gib": schema.Int64Attribute{
				Description:   "Shared memory size in GiB for Slurm cluster. This field is required if cluster_type is SLURM.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"volume_id": schema.StringAttribute{
				Description:   "ID of an existing volume to use with the cluster creation.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"auto_scaled": schema.BoolAttribute{
				Description:   "Whether GPU cluster should be auto-scaled based on the workload. By default, it is not auto-scaled.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"gpu_node_failover_enabled": schema.BoolAttribute{
				Description:   "Whether automated GPU node failover should be enabled for this cluster. By default, it is disabled.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"install_traefik": schema.BoolAttribute{
				Description:   "Whether to install Traefik ingress controller in the cluster. This field is only applicable for Kubernetes clusters and is false by default.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"num_gpus": schema.Int64Attribute{
				Description: "Number of GPUs to allocate in the cluster. This must be multiple of 8. For example, 8, 16 or 24",
				Required:    true,
			},
			"cluster_type": schema.StringAttribute{
				Description: "Type of cluster to create.\nAvailable values: \"KUBERNETES\", \"SLURM\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("KUBERNETES", "SLURM"),
				},
			},
			"reservation_end_time": schema.StringAttribute{
				Description: "Reservation end time of the cluster. This field is required for SCHEDULED billing to specify the reservation end time for the cluster.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
						"instance_id": schema.StringAttribute{
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
