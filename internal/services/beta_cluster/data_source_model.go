// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-togetherai/internal/customfield"
)

type BetaClusterDataSourceModel struct {
	ClusterID         types.String                                                              `tfsdk:"cluster_id" path:"cluster_id,required"`
	ClusterName       types.String                                                              `tfsdk:"cluster_name" json:"cluster_name,computed"`
	ClusterType       types.String                                                              `tfsdk:"cluster_type" json:"cluster_type,computed"`
	DriverVersion     types.String                                                              `tfsdk:"driver_version" json:"driver_version,computed"`
	DurationHours     types.Int64                                                               `tfsdk:"duration_hours" json:"duration_hours,computed"`
	GPUType           types.String                                                              `tfsdk:"gpu_type" json:"gpu_type,computed"`
	KubeConfig        types.String                                                              `tfsdk:"kube_config" json:"kube_config,computed"`
	NumGPUs           types.Int64                                                               `tfsdk:"num_gpus" json:"num_gpus,computed"`
	Region            types.String                                                              `tfsdk:"region" json:"region,computed"`
	Status            types.String                                                              `tfsdk:"status" json:"status,computed"`
	ControlPlaneNodes customfield.NestedObjectList[BetaClusterControlPlaneNodesDataSourceModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes    customfield.NestedObjectList[BetaClusterGPUWorkerNodesDataSourceModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	Volumes           customfield.NestedObjectList[BetaClusterVolumesDataSourceModel]           `tfsdk:"volumes" json:"volumes,computed"`
}

type BetaClusterControlPlaneNodesDataSourceModel struct {
	HostName    types.String  `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib   types.Float64 `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Network     types.String  `tfsdk:"network" json:"network,computed"`
	NodeID      types.String  `tfsdk:"node_id" json:"node_id,computed"`
	NodeName    types.String  `tfsdk:"node_name" json:"node_name,computed"`
	NumCPUCores types.Int64   `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	Status      types.String  `tfsdk:"status" json:"status,computed"`
}

type BetaClusterGPUWorkerNodesDataSourceModel struct {
	HostName    types.String                   `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib   types.Float64                  `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Networks    customfield.List[types.String] `tfsdk:"networks" json:"networks,computed"`
	NodeID      types.String                   `tfsdk:"node_id" json:"node_id,computed"`
	NodeName    types.String                   `tfsdk:"node_name" json:"node_name,computed"`
	NumCPUCores types.Int64                    `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	NumGPUs     types.Int64                    `tfsdk:"num_gpus" json:"num_gpus,computed"`
	Status      types.String                   `tfsdk:"status" json:"status,computed"`
}

type BetaClusterVolumesDataSourceModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
