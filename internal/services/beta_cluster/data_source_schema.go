// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BetaClusterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				Required: true,
			},
			"cluster_name": schema.StringAttribute{
				Computed: true,
			},
			"cluster_type": schema.StringAttribute{
				Description: `Available values: "KUBERNETES", "SLURM".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("KUBERNETES", "SLURM"),
				},
			},
			"driver_version": schema.StringAttribute{
				Description: `Available values: "CUDA_12_5_555", "CUDA_12_6_560", "CUDA_12_6_565", "CUDA_12_8_570".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"CUDA_12_5_555",
						"CUDA_12_6_560",
						"CUDA_12_6_565",
						"CUDA_12_8_570",
					),
				},
			},
			"duration_hours": schema.Int64Attribute{
				Computed: true,
			},
			"gpu_type": schema.StringAttribute{
				Description: `Available values: "H100_SXM", "H200_SXM", "RTX_6000_PCI", "L40_PCIE", "B200_SXM", "H100_SXM_INF".`,
				Computed:    true,
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
			},
			"kube_config": schema.StringAttribute{
				Computed: true,
			},
			"num_gpus": schema.Int64Attribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
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
				CustomType: customfield.NewNestedObjectListType[BetaClusterControlPlaneNodesDataSourceModel](ctx),
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
				CustomType: customfield.NewNestedObjectListType[BetaClusterGPUWorkerNodesDataSourceModel](ctx),
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
				CustomType: customfield.NewNestedObjectListType[BetaClusterVolumesDataSourceModel](ctx),
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

func (d *BetaClusterDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *BetaClusterDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
