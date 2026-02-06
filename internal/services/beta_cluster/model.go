// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/apijson"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

type BetaClusterModel struct {
	ID                types.String                                                    `tfsdk:"id" json:"-,computed"`
	ClusterID         types.String                                                    `tfsdk:"cluster_id" json:"cluster_id,computed"`
	BillingType       types.String                                                    `tfsdk:"billing_type" json:"billing_type,required,no_refresh"`
	ClusterName       types.String                                                    `tfsdk:"cluster_name" json:"cluster_name,required"`
	DriverVersion     types.String                                                    `tfsdk:"driver_version" json:"driver_version,required"`
	GPUType           types.String                                                    `tfsdk:"gpu_type" json:"gpu_type,required"`
	Region            types.String                                                    `tfsdk:"region" json:"region,required"`
	DurationDays      types.Int64                                                     `tfsdk:"duration_days" json:"duration_days,optional,no_refresh"`
	VolumeID          types.String                                                    `tfsdk:"volume_id" json:"volume_id,optional,no_refresh"`
	SharedVolume      *BetaClusterSharedVolumeModel                                   `tfsdk:"shared_volume" json:"shared_volume,optional,no_refresh"`
	NumGPUs           types.Int64                                                     `tfsdk:"num_gpus" json:"num_gpus,required"`
	ClusterType       types.String                                                    `tfsdk:"cluster_type" json:"cluster_type,optional"`
	DurationHours     types.Int64                                                     `tfsdk:"duration_hours" json:"duration_hours,computed"`
	KubeConfig        types.String                                                    `tfsdk:"kube_config" json:"kube_config,computed"`
	Status            types.String                                                    `tfsdk:"status" json:"status,computed"`
	ControlPlaneNodes customfield.NestedObjectList[BetaClusterControlPlaneNodesModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes    customfield.NestedObjectList[BetaClusterGPUWorkerNodesModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	Volumes           customfield.NestedObjectList[BetaClusterVolumesModel]           `tfsdk:"volumes" json:"volumes,computed"`
}

func (m BetaClusterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

// <CustomCode>
// For Terraform to be stateful with Instant Clusters, we force the billing type to be ON_DEMAND.
type InternalBetaClusterModel struct {
	BetaClusterModel
	BillingType       types.String                                                    `tfsdk:"billing_type" json:"billing_type,required,no_refresh"`
}

func (m InternalBetaClusterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}
// </CustomCode>

func (m BetaClusterModel) MarshalJSONForUpdate(state BetaClusterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type BetaClusterSharedVolumeModel struct {
	Region     types.String `tfsdk:"region" json:"region,required"`
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,required"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,required"`
}

type BetaClusterControlPlaneNodesModel struct {
	HostName    types.String  `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib   types.Float64 `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Network     types.String  `tfsdk:"network" json:"network,computed"`
	NodeID      types.String  `tfsdk:"node_id" json:"node_id,computed"`
	NodeName    types.String  `tfsdk:"node_name" json:"node_name,computed"`
	NumCPUCores types.Int64   `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	Status      types.String  `tfsdk:"status" json:"status,computed"`
}

type BetaClusterGPUWorkerNodesModel struct {
	HostName    types.String                   `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib   types.Float64                  `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Networks    customfield.List[types.String] `tfsdk:"networks" json:"networks,computed"`
	NodeID      types.String                   `tfsdk:"node_id" json:"node_id,computed"`
	NodeName    types.String                   `tfsdk:"node_name" json:"node_name,computed"`
	NumCPUCores types.Int64                    `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	NumGPUs     types.Int64                    `tfsdk:"num_gpus" json:"num_gpus,computed"`
	Status      types.String                   `tfsdk:"status" json:"status,computed"`
}

type BetaClusterVolumesModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
