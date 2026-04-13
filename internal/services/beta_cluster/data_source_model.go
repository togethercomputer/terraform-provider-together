// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

type BetaClusterDataSourceModel struct {
	ClusterID            types.String                                                              `tfsdk:"cluster_id" path:"cluster_id,required"`
	CapacityPoolID       types.String                                                              `tfsdk:"capacity_pool_id" json:"capacity_pool_id,computed"`
	ClusterName          types.String                                                              `tfsdk:"cluster_name" json:"cluster_name,computed"`
	ClusterType          types.String                                                              `tfsdk:"cluster_type" json:"cluster_type,computed"`
	CreatedAt            timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CudaVersion          types.String                                                              `tfsdk:"cuda_version" json:"cuda_version,computed"`
	DurationHours        types.Int64                                                               `tfsdk:"duration_hours" json:"duration_hours,computed"`
	GPUType              types.String                                                              `tfsdk:"gpu_type" json:"gpu_type,computed"`
	InstallTraefik       types.Bool                                                                `tfsdk:"install_traefik" json:"install_traefik,computed"`
	KubeConfig           types.String                                                              `tfsdk:"kube_config" json:"kube_config,computed"`
	NumGPUs              types.Int64                                                               `tfsdk:"num_gpus" json:"num_gpus,computed"`
	NvidiaDriverVersion  types.String                                                              `tfsdk:"nvidia_driver_version" json:"nvidia_driver_version,computed"`
	Region               types.String                                                              `tfsdk:"region" json:"region,computed"`
	ReservationEndTime   timetypes.RFC3339                                                         `tfsdk:"reservation_end_time" json:"reservation_end_time,computed" format:"date-time"`
	ReservationStartTime timetypes.RFC3339                                                         `tfsdk:"reservation_start_time" json:"reservation_start_time,computed" format:"date-time"`
	SlurmShmSizeGib      types.Int64                                                               `tfsdk:"slurm_shm_size_gib" json:"slurm_shm_size_gib,computed"`
	Status               types.String                                                              `tfsdk:"status" json:"status,computed"`
	ControlPlaneNodes    customfield.NestedObjectList[BetaClusterControlPlaneNodesDataSourceModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes       customfield.NestedObjectList[BetaClusterGPUWorkerNodesDataSourceModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	Volumes              customfield.NestedObjectList[BetaClusterVolumesDataSourceModel]           `tfsdk:"volumes" json:"volumes,computed"`
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
	InstanceID  types.String                   `tfsdk:"instance_id" json:"instance_id,computed"`
}

type BetaClusterVolumesDataSourceModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
