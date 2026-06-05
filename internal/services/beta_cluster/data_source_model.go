// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

type BetaClusterDataSourceModel struct {
	ClusterID                types.String                                                              `tfsdk:"cluster_id" path:"cluster_id,required"`
	AllocatedPreemptibleGPUs types.Int64                                                               `tfsdk:"allocated_preemptible_gpus" json:"allocated_preemptible_gpus,computed"`
	BillingType              types.String                                                              `tfsdk:"billing_type" json:"billing_type,computed"`
	CapacityPoolID           types.String                                                              `tfsdk:"capacity_pool_id" json:"capacity_pool_id,computed"`
	ClusterName              types.String                                                              `tfsdk:"cluster_name" json:"cluster_name,computed"`
	ClusterType              types.String                                                              `tfsdk:"cluster_type" json:"cluster_type,computed"`
	ControlPlaneReady        types.Bool                                                                `tfsdk:"control_plane_ready" json:"control_plane_ready,computed"`
	CreatedAt                timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	CudaVersion              types.String                                                              `tfsdk:"cuda_version" json:"cuda_version,computed"`
	DesiredPreemptibleGPUs   types.Int64                                                               `tfsdk:"desired_preemptible_gpus" json:"desired_preemptible_gpus,computed"`
	DurationHours            types.Int64                                                               `tfsdk:"duration_hours" json:"duration_hours,computed"`
	FirstReadyAt             timetypes.RFC3339                                                         `tfsdk:"first_ready_at" json:"first_ready_at,computed" format:"date-time"`
	GPUType                  types.String                                                              `tfsdk:"gpu_type" json:"gpu_type,computed"`
	InstallTraefik           types.Bool                                                                `tfsdk:"install_traefik" json:"install_traefik,computed"`
	IsInSubstrate            types.Bool                                                                `tfsdk:"is_in_substrate" json:"is_in_substrate,computed"`
	KubeConfig               types.String                                                              `tfsdk:"kube_config" json:"kube_config,computed"`
	MachineClusterID         types.String                                                              `tfsdk:"machine_cluster_id" json:"machine_cluster_id,computed"`
	NumCPUWorkers            types.Int64                                                               `tfsdk:"num_cpu_workers" json:"num_cpu_workers,computed"`
	NumGPUs                  types.Int64                                                               `tfsdk:"num_gpus" json:"num_gpus,computed"`
	NvidiaDriverVersion      types.String                                                              `tfsdk:"nvidia_driver_version" json:"nvidia_driver_version,computed"`
	NvidiaDriverVersionID    types.String                                                              `tfsdk:"nvidia_driver_version_id" json:"nvidia_driver_version_id,computed"`
	OsImage                  types.String                                                              `tfsdk:"os_image" json:"os_image,computed"`
	ProjectID                types.String                                                              `tfsdk:"project_id" json:"project_id,computed"`
	Region                   types.String                                                              `tfsdk:"region" json:"region,computed"`
	ReservationEndTime       timetypes.RFC3339                                                         `tfsdk:"reservation_end_time" json:"reservation_end_time,computed" format:"date-time"`
	ReservationStartTime     timetypes.RFC3339                                                         `tfsdk:"reservation_start_time" json:"reservation_start_time,computed" format:"date-time"`
	SlurmShmSizeGib          types.Int64                                                               `tfsdk:"slurm_shm_size_gib" json:"slurm_shm_size_gib,computed"`
	Status                   types.String                                                              `tfsdk:"status" json:"status,computed"`
	UmsOrgID                 types.String                                                              `tfsdk:"ums_org_id" json:"ums_org_id,computed"`
	UmsProjectID             types.String                                                              `tfsdk:"ums_project_id" json:"ums_project_id,computed"`
	AddOns                   customfield.NestedObjectList[BetaClusterAddOnsDataSourceModel]            `tfsdk:"add_ons" json:"add_ons,computed"`
	ClusterConfig            customfield.NestedObject[BetaClusterClusterConfigDataSourceModel]         `tfsdk:"cluster_config" json:"cluster_config,computed"`
	ControlPlaneNodes        customfield.NestedObjectList[BetaClusterControlPlaneNodesDataSourceModel] `tfsdk:"control_plane_nodes" json:"control_plane_nodes,computed"`
	GPUWorkerNodes           customfield.NestedObjectList[BetaClusterGPUWorkerNodesDataSourceModel]    `tfsdk:"gpu_worker_nodes" json:"gpu_worker_nodes,computed"`
	OidcConfig               customfield.NestedObject[BetaClusterOidcConfigDataSourceModel]            `tfsdk:"oidc_config" json:"oidc_config,computed"`
	PhaseTransitions         customfield.NestedObjectList[BetaClusterPhaseTransitionsDataSourceModel]  `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Volumes                  customfield.NestedObjectList[BetaClusterVolumesDataSourceModel]           `tfsdk:"volumes" json:"volumes,computed"`
}

type BetaClusterAddOnsDataSourceModel struct {
	AddOnType types.String                                                     `tfsdk:"add_on_type" json:"add_on_type,computed"`
	Config    customfield.NestedObject[BetaClusterAddOnsConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
	Name      types.String                                                     `tfsdk:"name" json:"name,computed"`
	State     customfield.NestedObject[BetaClusterAddOnsStateDataSourceModel]  `tfsdk:"state" json:"state,computed"`
}

type BetaClusterAddOnsConfigDataSourceModel struct {
	Dashboard customfield.NestedObject[BetaClusterAddOnsConfigDashboardDataSourceModel] `tfsdk:"dashboard" json:"dashboard,computed"`
	Ingress   customfield.NestedObject[BetaClusterAddOnsConfigIngressDataSourceModel]   `tfsdk:"ingress" json:"ingress,computed"`
}

type BetaClusterAddOnsConfigDashboardDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type BetaClusterAddOnsConfigIngressDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type BetaClusterAddOnsStateDataSourceModel struct {
	Dashboard customfield.NestedObject[BetaClusterAddOnsStateDashboardDataSourceModel] `tfsdk:"dashboard" json:"dashboard,computed"`
	Ingress   customfield.NestedObject[BetaClusterAddOnsStateIngressDataSourceModel]   `tfsdk:"ingress" json:"ingress,computed"`
}

type BetaClusterAddOnsStateDashboardDataSourceModel struct {
}

type BetaClusterAddOnsStateIngressDataSourceModel struct {
}

type BetaClusterClusterConfigDataSourceModel struct {
	LoadBalancer               types.String                                                                         `tfsdk:"load_balancer" json:"load_balancer,computed"`
	GPUOperatorVersion         types.String                                                                         `tfsdk:"gpu_operator_version" json:"gpu_operator_version,computed"`
	Ingress                    customfield.NestedObject[BetaClusterClusterConfigIngressDataSourceModel]             `tfsdk:"ingress" json:"ingress,computed"`
	JumphostEnabled            types.Bool                                                                           `tfsdk:"jumphost_enabled" json:"jumphost_enabled,computed"`
	KubernetesDashboardEnabled types.Bool                                                                           `tfsdk:"kubernetes_dashboard_enabled" json:"kubernetes_dashboard_enabled,computed"`
	Observability              customfield.NestedObject[BetaClusterClusterConfigObservabilityDataSourceModel]       `tfsdk:"observability" json:"observability,computed"`
	SlurmStartupScripts        customfield.NestedObject[BetaClusterClusterConfigSlurmStartupScriptsDataSourceModel] `tfsdk:"slurm_startup_scripts" json:"slurm_startup_scripts,computed"`
}

type BetaClusterClusterConfigIngressDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type BetaClusterClusterConfigObservabilityDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type BetaClusterClusterConfigSlurmStartupScriptsDataSourceModel struct {
	ControllerEpilog  types.String `tfsdk:"controller_epilog" json:"controller_epilog,computed"`
	ControllerProlog  types.String `tfsdk:"controller_prolog" json:"controller_prolog,computed"`
	ExtraSlurmConf    types.String `tfsdk:"extra_slurm_conf" json:"extra_slurm_conf,computed"`
	LoginInitScript   types.String `tfsdk:"login_init_script" json:"login_init_script,computed"`
	NodesetInitScript types.String `tfsdk:"nodeset_init_script" json:"nodeset_init_script,computed"`
	WorkerEpilog      types.String `tfsdk:"worker_epilog" json:"worker_epilog,computed"`
	WorkerProlog      types.String `tfsdk:"worker_prolog" json:"worker_prolog,computed"`
}

type BetaClusterControlPlaneNodesDataSourceModel struct {
	HostName         types.String                                                                              `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib        types.Float64                                                                             `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Network          types.String                                                                              `tfsdk:"network" json:"network,computed"`
	NodeID           types.String                                                                              `tfsdk:"node_id" json:"node_id,computed"`
	NumCPUCores      types.Int64                                                                               `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	PhaseTransitions customfield.NestedObjectList[BetaClusterControlPlaneNodesPhaseTransitionsDataSourceModel] `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Status           types.String                                                                              `tfsdk:"status" json:"status,computed"`
	PublicIpv4       types.String                                                                              `tfsdk:"public_ipv4" json:"public_ipv4,computed"`
}

type BetaClusterControlPlaneNodesPhaseTransitionsDataSourceModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterGPUWorkerNodesDataSourceModel struct {
	HostName               types.String                                                                           `tfsdk:"host_name" json:"host_name,computed"`
	MemoryGib              types.Float64                                                                          `tfsdk:"memory_gib" json:"memory_gib,computed"`
	Networks               customfield.List[types.String]                                                         `tfsdk:"networks" json:"networks,computed"`
	NodeID                 types.String                                                                           `tfsdk:"node_id" json:"node_id,computed"`
	NumCPUCores            types.Int64                                                                            `tfsdk:"num_cpu_cores" json:"num_cpu_cores,computed"`
	NumGPUs                types.Int64                                                                            `tfsdk:"num_gpus" json:"num_gpus,computed"`
	PhaseTransitions       customfield.NestedObjectList[BetaClusterGPUWorkerNodesPhaseTransitionsDataSourceModel] `tfsdk:"phase_transitions" json:"phase_transitions,computed"`
	Status                 types.String                                                                           `tfsdk:"status" json:"status,computed"`
	AutoRemediationEnabled types.Bool                                                                             `tfsdk:"auto_remediation_enabled" json:"auto_remediation_enabled,computed"`
	EphemeralStorage       types.String                                                                           `tfsdk:"ephemeral_storage" json:"ephemeral_storage,computed"`
	IbHcaCount             types.Int64                                                                            `tfsdk:"ib_hca_count" json:"ib_hca_count,computed"`
	IbHcaType              types.String                                                                           `tfsdk:"ib_hca_type" json:"ib_hca_type,computed"`
	InstanceID             types.String                                                                           `tfsdk:"instance_id" json:"instance_id,computed"`
	LatestRemediation      customfield.NestedObject[BetaClusterGPUWorkerNodesLatestRemediationDataSourceModel]    `tfsdk:"latest_remediation" json:"latest_remediation,computed"`
	MarkedForDeletion      types.Bool                                                                             `tfsdk:"marked_for_deletion" json:"marked_for_deletion,computed"`
	NvswitchCount          types.Int64                                                                            `tfsdk:"nvswitch_count" json:"nvswitch_count,computed"`
	NvswitchType           types.String                                                                           `tfsdk:"nvswitch_type" json:"nvswitch_type,computed"`
	PublicIpv4             types.String                                                                           `tfsdk:"public_ipv4" json:"public_ipv4,computed"`
	SlurmWorkerHostname    types.String                                                                           `tfsdk:"slurm_worker_hostname" json:"slurm_worker_hostname,computed"`
}

type BetaClusterGPUWorkerNodesPhaseTransitionsDataSourceModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterGPUWorkerNodesLatestRemediationDataSourceModel struct {
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

type BetaClusterOidcConfigDataSourceModel struct {
	ClientID       types.String `tfsdk:"client_id" json:"client_id,computed"`
	GroupClaim     types.String `tfsdk:"group_claim" json:"group_claim,computed"`
	GroupPrefix    types.String `tfsdk:"group_prefix" json:"group_prefix,computed"`
	IssuerURL      types.String `tfsdk:"issuer_url" json:"issuer_url,computed"`
	UsernameClaim  types.String `tfsdk:"username_claim" json:"username_claim,computed"`
	UsernamePrefix types.String `tfsdk:"username_prefix" json:"username_prefix,computed"`
	CaCert         types.String `tfsdk:"ca_cert" json:"ca_cert,computed"`
}

type BetaClusterPhaseTransitionsDataSourceModel struct {
	Phase          types.String      `tfsdk:"phase" json:"phase,computed"`
	TransitionTime timetypes.RFC3339 `tfsdk:"transition_time" json:"transition_time,computed" format:"date-time"`
}

type BetaClusterVolumesDataSourceModel struct {
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
