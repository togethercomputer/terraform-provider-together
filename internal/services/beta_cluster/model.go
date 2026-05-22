// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/apijson"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

type BetaClusterModel struct {
	ID                       types.String                                                    `tfsdk:"id" json:"-,computed"`
	ClusterID                types.String                                                    `tfsdk:"cluster_id" json:"cluster_id,computed"`
	ClusterName              types.String                                                    `tfsdk:"cluster_name" json:"cluster_name,required"`
	CudaVersion              types.String                                                    `tfsdk:"cuda_version" json:"cuda_version,required"`
	GPUType                  types.String                                                    `tfsdk:"gpu_type" json:"gpu_type,required"`
	NvidiaDriverVersion      types.String                                                    `tfsdk:"nvidia_driver_version" json:"nvidia_driver_version,required"`
	Region                   types.String                                                    `tfsdk:"region" json:"region,required"`
	AutoScale                types.Bool                                                      `tfsdk:"auto_scale" json:"auto_scale,optional,no_refresh"`
	AutoScaleMaxGPUs         types.Int64                                                     `tfsdk:"auto_scale_max_gpus" json:"auto_scale_max_gpus,optional,no_refresh"`
	CapacityPoolID           types.String                                                    `tfsdk:"capacity_pool_id" json:"capacity_pool_id,optional"`
	NumCapacityPoolGPUs      types.Int64                                                     `tfsdk:"num_capacity_pool_gpus" json:"num_capacity_pool_gpus,optional,no_refresh"`
	ProjectID                types.String                                                    `tfsdk:"project_id" json:"project_id,optional"`
	ReservationStartTime     timetypes.RFC3339                                               `tfsdk:"reservation_start_time" json:"reservation_start_time,optional" format:"date-time"`
	SlurmImage               types.String                                                    `tfsdk:"slurm_image" json:"slurm_image,optional,no_refresh"`
	SlurmShmSizeGib          types.Int64                                                     `tfsdk:"slurm_shm_size_gib" json:"slurm_shm_size_gib,optional"`
	VolumeID                 types.String                                                    `tfsdk:"volume_id" json:"volume_id,optional,no_refresh"`
	AcceptanceTestsParams    *BetaClusterAcceptanceTestsParamsModel                          `tfsdk:"acceptance_tests_params" json:"acceptance_tests_params,optional,no_refresh"`
	OidcConfig               *BetaClusterOidcConfigModel                                     `tfsdk:"oidc_config" json:"oidc_config,optional"`
	AutoScaled               types.Bool                                                      `tfsdk:"auto_scaled" json:"auto_scaled,computed_optional,no_refresh"`
	GPUNodeFailoverEnabled   types.Bool                                                      `tfsdk:"gpu_node_failover_enabled" json:"gpu_node_failover_enabled,computed_optional,no_refresh"`
	InstallTraefik           types.Bool                                                      `tfsdk:"install_traefik" json:"install_traefik,computed_optional"`
	NumGPUs                  types.Int64                                                     `tfsdk:"num_gpus" json:"num_gpus,required"`
	ClusterType              types.String                                                    `tfsdk:"cluster_type" json:"cluster_type,optional"`
	NumPreemptibleGPUs       types.Int64                                                     `tfsdk:"num_preemptible_gpus" json:"num_preemptible_gpus,optional,no_refresh"`
	NumReservedGPUs          types.Int64                                                     `tfsdk:"num_reserved_gpus" json:"num_reserved_gpus,optional,no_refresh"`
	ReservationEndTime       timetypes.RFC3339                                               `tfsdk:"reservation_end_time" json:"reservation_end_time,optional" format:"date-time"`
	AddOns                   *[]*BetaClusterAddOnsModel                                      `tfsdk:"add_ons" json:"add_ons,optional"`
	ClusterConfig            *BetaClusterClusterConfigModel                                  `tfsdk:"cluster_config" json:"cluster_config,optional"`
	AllocatedPreemptibleGPUs types.Int64                                                     `tfsdk:"allocated_preemptible_gpus" json:"allocated_preemptible_gpus,computed"`
	BillingType              types.String                                                    `tfsdk:"billing_type" json:"billing_type,computed"`
	CreatedAt                timetypes.RFC3339                                               `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DesiredPreemptibleGPUs   types.Int64                                                     `tfsdk:"desired_preemptible_gpus" json:"desired_preemptible_gpus,computed"`
	DurationHours            types.Int64                                                     `tfsdk:"duration_hours" json:"duration_hours,computed"`
	KubeConfig               types.String                                                    `tfsdk:"kube_config" json:"kube_config,computed"`
	NumCPUWorkers            types.Int64                                                     `tfsdk:"num_cpu_workers" json:"num_cpu_workers,computed"`
	Status                   types.String                                                    `tfsdk:"status" json:"status,computed"`
	ControlPlaneNodes        customfield.NestedObjectList[BetaClusterControlPlaneNodesModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes           customfield.NestedObjectList[BetaClusterGPUWorkerNodesModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	PhaseTransitions         customfield.NestedObjectList[BetaClusterPhaseTransitionsModel]  `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Volumes                  customfield.NestedObjectList[BetaClusterVolumesModel]           `tfsdk:"volumes" json:"volumes,computed"`
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

type BetaClusterAcceptanceTestsParamsModel struct {
	DcgmDiagLevel         types.String `tfsdk:"dcgm_diag_level" json:"dcgm_diag_level,optional"`
	DcgmDiagSkipped       types.Bool   `tfsdk:"dcgm_diag_skipped" json:"dcgm_diag_skipped,optional"`
	Enabled               types.Bool   `tfsdk:"enabled" json:"enabled,optional"`
	GPUBurnDuration       types.Int64  `tfsdk:"gpu_burn_duration" json:"gpu_burn_duration,optional"`
	GPUBurnSkipped        types.Bool   `tfsdk:"gpu_burn_skipped" json:"gpu_burn_skipped,optional"`
	NcclMultiNodeSkipped  types.Bool   `tfsdk:"nccl_multi_node_skipped" json:"nccl_multi_node_skipped,optional"`
	NcclSingleNodeSkipped types.Bool   `tfsdk:"nccl_single_node_skipped" json:"nccl_single_node_skipped,optional"`
}

type BetaClusterOidcConfigModel struct {
	ClientID       types.String `tfsdk:"client_id" json:"client_id,required"`
	GroupClaim     types.String `tfsdk:"group_claim" json:"group_claim,required"`
	GroupPrefix    types.String `tfsdk:"group_prefix" json:"group_prefix,required"`
	IssuerURL      types.String `tfsdk:"issuer_url" json:"issuer_url,required"`
	UsernameClaim  types.String `tfsdk:"username_claim" json:"username_claim,required"`
	UsernamePrefix types.String `tfsdk:"username_prefix" json:"username_prefix,required"`
	CaCert         types.String `tfsdk:"ca_cert" json:"ca_cert,optional"`
}

type BetaClusterAddOnsModel struct {
	AddOnType types.String                  `tfsdk:"add_on_type" json:"add_on_type,required"`
	Name      types.String                  `tfsdk:"name" json:"name,required"`
	Config    *BetaClusterAddOnsConfigModel `tfsdk:"config" json:"config,optional"`
}

type BetaClusterAddOnsConfigModel struct {
	Dashboard *BetaClusterAddOnsConfigDashboardModel `tfsdk:"dashboard" json:"dashboard,optional"`
	Ingress   *BetaClusterAddOnsConfigIngressModel   `tfsdk:"ingress" json:"ingress,optional"`
}

type BetaClusterAddOnsConfigDashboardModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type BetaClusterAddOnsConfigIngressModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type BetaClusterClusterConfigModel struct {
	LoadBalancer               types.String                                      `tfsdk:"load_balancer" json:"load_balancer,required"`
	GPUOperatorVersion         types.String                                      `tfsdk:"gpu_operator_version" json:"gpu_operator_version,optional"`
	Ingress                    *BetaClusterClusterConfigIngressModel             `tfsdk:"ingress" json:"ingress,optional"`
	JumphostEnabled            types.Bool                                        `tfsdk:"jumphost_enabled" json:"jumphost_enabled,optional"`
	KubernetesDashboardEnabled types.Bool                                        `tfsdk:"kubernetes_dashboard_enabled" json:"kubernetes_dashboard_enabled,optional"`
	Observability              *BetaClusterClusterConfigObservabilityModel       `tfsdk:"observability" json:"observability,optional"`
	SlurmStartupScripts        *BetaClusterClusterConfigSlurmStartupScriptsModel `tfsdk:"slurm_startup_scripts" json:"slurm_startup_scripts,optional"`
}

type BetaClusterClusterConfigIngressModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type BetaClusterClusterConfigObservabilityModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,optional"`
}

type BetaClusterClusterConfigSlurmStartupScriptsModel struct {
	ControllerEpilog  types.String `tfsdk:"controller_epilog" json:"controller_epilog,optional"`
	ControllerProlog  types.String `tfsdk:"controller_prolog" json:"controller_prolog,optional"`
	ExtraSlurmConf    types.String `tfsdk:"extra_slurm_conf" json:"extra_slurm_conf,optional"`
	LoginInitScript   types.String `tfsdk:"login_init_script" json:"login_init_script,optional"`
	NodesetInitScript types.String `tfsdk:"nodeset_init_script" json:"nodeset_init_script,optional"`
	WorkerEpilog      types.String `tfsdk:"worker_epilog" json:"worker_epilog,optional"`
	WorkerProlog      types.String `tfsdk:"worker_prolog" json:"worker_prolog,optional"`
}

type BetaClusterControlPlaneNodesModel struct {
	HostName         types.String                                                                    `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib        types.Float64                                                                   `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Network          types.String                                                                    `tfsdk:"network" json:"network,computed"`
	NodeID           types.String                                                                    `tfsdk:"node_id" json:"node_id,computed"`
	NumCPUCores      types.Int64                                                                     `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	PhaseTransitions customfield.NestedObjectList[BetaClusterControlPlaneNodesPhaseTransitionsModel] `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Status           types.String                                                                    `tfsdk:"status" json:"status,computed"`
}

type BetaClusterControlPlaneNodesPhaseTransitionsModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterGPUWorkerNodesModel struct {
	HostName            types.String                                                                 `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib           types.Float64                                                                `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Networks            customfield.List[types.String]                                               `tfsdk:"networks" json:"networks,computed"`
	NodeID              types.String                                                                 `tfsdk:"node_id" json:"node_id,computed"`
	NumCPUCores         types.Int64                                                                  `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	NumGPUs             types.Int64                                                                  `tfsdk:"num_gpus" json:"num_gpus,computed"`
	PhaseTransitions    customfield.NestedObjectList[BetaClusterGPUWorkerNodesPhaseTransitionsModel] `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Status              types.String                                                                 `tfsdk:"status" json:"status,computed"`
	InstanceID          types.String                                                                 `tfsdk:"instance_id" json:"instance_id,computed"`
	LatestRemediation   customfield.NestedObject[BetaClusterGPUWorkerNodesLatestRemediationModel]    `tfsdk:"latest_remediation" json:"latest_remediation,computed"`
	SlurmWorkerHostname types.String                                                                 `tfsdk:"slurm_worker_hostname" json:"slurm_worker_hostname,computed"`
}

type BetaClusterGPUWorkerNodesPhaseTransitionsModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterGPUWorkerNodesLatestRemediationModel struct {
	ID                        types.String      `tfsdk:"id" json:"id,computed"`
	ClusterID                 types.String      `tfsdk:"cluster_id" json:"cluster_id,computed"`
	InstanceID                types.String      `tfsdk:"instance_id" json:"instance_id,computed"`
	Mode                      types.String      `tfsdk:"mode" json:"mode,computed"`
	State                     types.String      `tfsdk:"state" json:"state,computed"`
	Trigger                   types.String      `tfsdk:"trigger" json:"trigger,computed"`
	ActiveHealthCheckRunID    types.String      `tfsdk:"active_health_check_run_id" json:"active_health_check_run_id,computed"`
	CreateTime                timetypes.RFC3339 `tfsdk:"create_time" json:"create_time,computed" format:"date-time"`
	EndTime                   timetypes.RFC3339 `tfsdk:"end_time" json:"end_time,computed" format:"date-time"`
	ErrorMessage              types.String      `tfsdk:"error_message" json:"error_message,computed"`
	InstanceName              types.String      `tfsdk:"instance_name" json:"instance_name,computed"`
	PassiveHealthCheckEventID types.String      `tfsdk:"passive_health_check_event_id" json:"passive_health_check_event_id,computed"`
	Reason                    types.String      `tfsdk:"reason" json:"reason,computed"`
	RequestedBy               types.String      `tfsdk:"requested_by" json:"requested_by,computed"`
	ReviewComment             types.String      `tfsdk:"review_comment" json:"review_comment,computed"`
	ReviewTime                timetypes.RFC3339 `tfsdk:"review_time" json:"review_time,computed" format:"date-time"`
	ReviewedBy                types.String      `tfsdk:"reviewed_by" json:"reviewed_by,computed"`
	StartTime                 timetypes.RFC3339 `tfsdk:"start_time" json:"start_time,computed" format:"date-time"`
	UpdateTime                timetypes.RFC3339 `tfsdk:"update_time" json:"update_time,computed" format:"date-time"`
}

type BetaClusterPhaseTransitionsModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterVolumesModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
