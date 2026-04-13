// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/apijson"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

type BetaClusterModel struct {
	ID                     types.String                                                    `tfsdk:"id" json:"-,computed"`
	ClusterID              types.String                                                    `tfsdk:"cluster_id" json:"cluster_id,computed"`
	ClusterName            types.String                                                    `tfsdk:"cluster_name" json:"cluster_name,required"`
	CudaVersion            types.String                                                    `tfsdk:"cuda_version" json:"cuda_version,required"`
	GPUType                types.String                                                    `tfsdk:"gpu_type" json:"gpu_type,required"`
	NvidiaDriverVersion    types.String                                                    `tfsdk:"nvidia_driver_version" json:"nvidia_driver_version,required"`
	Region                 types.String                                                    `tfsdk:"region" json:"region,required"`
	AutoScaleMaxGPUs       types.Int64                                                     `tfsdk:"auto_scale_max_gpus" json:"auto_scale_max_gpus,optional,no_refresh"`
	CapacityPoolID         types.String                                                    `tfsdk:"capacity_pool_id" json:"capacity_pool_id,optional"`
	ReservationStartTime   timetypes.RFC3339                                               `tfsdk:"reservation_start_time" json:"reservation_start_time,optional" format:"date-time"`
	SlurmImage             types.String                                                    `tfsdk:"slurm_image" json:"slurm_image,optional,no_refresh"`
	SlurmShmSizeGib        types.Int64                                                     `tfsdk:"slurm_shm_size_gib" json:"slurm_shm_size_gib,optional"`
	VolumeID               types.String                                                    `tfsdk:"volume_id" json:"volume_id,optional,no_refresh"`
	AutoScaled             types.Bool                                                      `tfsdk:"auto_scaled" json:"auto_scaled,computed_optional,no_refresh"`
	GPUNodeFailoverEnabled types.Bool                                                      `tfsdk:"gpu_node_failover_enabled" json:"gpu_node_failover_enabled,computed_optional,no_refresh"`
	InstallTraefik         types.Bool                                                      `tfsdk:"install_traefik" json:"install_traefik,computed_optional"`
	NumGPUs                types.Int64                                                     `tfsdk:"num_gpus" json:"num_gpus,required"`
	ClusterType            types.String                                                    `tfsdk:"cluster_type" json:"cluster_type,optional"`
	ReservationEndTime     timetypes.RFC3339                                               `tfsdk:"reservation_end_time" json:"reservation_end_time,optional" format:"date-time"`
	CreatedAt              timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DurationHours          types.Int64                                                     `tfsdk:"duration_hours" json:"duration_hours,computed"`
	KubeConfig             types.String                                                    `tfsdk:"kube_config" json:"kube_config,computed"`
	Status                 types.String                                                    `tfsdk:"status" json:"status,computed"`
	ControlPlaneNodes      customfield.NestedObjectList[BetaClusterControlPlaneNodesModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes         customfield.NestedObjectList[BetaClusterGPUWorkerNodesModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	Volumes                customfield.NestedObjectList[BetaClusterVolumesModel]           `tfsdk:"volumes" json:"volumes,computed"`
}

func (m BetaClusterModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BetaClusterModel) MarshalJSONForUpdate(state BetaClusterModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
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
	InstanceID  types.String                   `tfsdk:"instance_id" json:"instance_id,computed"`
}

type BetaClusterVolumesModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
