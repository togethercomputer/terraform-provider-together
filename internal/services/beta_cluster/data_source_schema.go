// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/togethercomputer/terraform-provider-together/internal/customfield"
)

var _ datasource.DataSourceWithConfigValidators = (*BetaClusterDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_id": schema.StringAttribute{
				Description: "The ID of the cluster to retrieve",
				Required:    true,
			},
			"allocated_preemptible_gpus": schema.Int64Attribute{
				Description: "Actual number of preemptible GPUs currently allocated to the cluster. Updated asynchronously by the fulfillment and reclamation workers; may be less than desired_preemptible_gpus when capacity is constrained.",
				Computed:    true,
			},
			"billing_type": schema.StringAttribute{
				Description: "Billing type for the cluster (RESERVED, ON_DEMAND, or SCHEDULED_CAPACITY).\nAvailable values: \"RESERVED\", \"ON_DEMAND\", \"SCHEDULED_CAPACITY\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"RESERVED",
						"ON_DEMAND",
						"SCHEDULED_CAPACITY",
					),
				},
			},
			"capacity_pool_id": schema.StringAttribute{
				Computed: true,
			},
			"cluster_name": schema.StringAttribute{
				Computed: true,
			},
			"cluster_type": schema.StringAttribute{
				Description: "Type of cluster.\nAvailable values: \"KUBERNETES\", \"SLURM\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("KUBERNETES", "SLURM"),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"cuda_version": schema.StringAttribute{
				Computed: true,
			},
			"desired_preemptible_gpus": schema.Int64Attribute{
				Description: "Customer's requested number of preemptible GPUs. Set on cluster create or update; persists until changed.",
				Computed:    true,
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
			"install_traefik": schema.BoolAttribute{
				Computed: true,
			},
			"kube_config": schema.StringAttribute{
				Computed: true,
			},
			"num_cpu_workers": schema.Int64Attribute{
				Description: "Number of CPU-only worker nodes in the cluster.",
				Computed:    true,
			},
			"num_gpus": schema.Int64Attribute{
				Computed: true,
			},
			"nvidia_driver_version": schema.StringAttribute{
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"reservation_end_time": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"reservation_start_time": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"slurm_shm_size_gib": schema.Int64Attribute{
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
			"add_ons": schema.ListNestedAttribute{
				Description: "Enabled add-ons on this cluster. Only add-ons with enabled=true in their config are returned.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BetaClusterAddOnsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"add_on_type": schema.StringAttribute{
							Computed: true,
						},
						"config": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsConfigDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"dashboard": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsConfigDashboardDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"ingress": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsConfigIngressDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"state": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsStateDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"dashboard": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsStateDashboardDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{},
								},
								"ingress": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[BetaClusterAddOnsStateIngressDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{},
								},
							},
						},
					},
				},
			},
			"cluster_config": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[BetaClusterClusterConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"load_balancer": schema.StringAttribute{
						Description: `Available values: "NONE", "TRAEFIK", "NGINX", "ISTIO".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"NONE",
								"TRAEFIK",
								"NGINX",
								"ISTIO",
							),
						},
					},
					"gpu_operator_version": schema.StringAttribute{
						Description: "NVIDIA GPU Operator chart/version for the tenant cluster (e.g. v24.6.2). When omitted, a service default is applied.",
						Computed:    true,
					},
					"ingress": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[BetaClusterClusterConfigIngressDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"jumphost_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"kubernetes_dashboard_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"observability": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[BetaClusterClusterConfigObservabilityDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"slurm_startup_scripts": schema.SingleNestedAttribute{
						Description: "SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog, init, extra conf).",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[BetaClusterClusterConfigSlurmStartupScriptsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"controller_epilog": schema.StringAttribute{
								Description: "Slurm controller epilog script.",
								Computed:    true,
							},
							"controller_prolog": schema.StringAttribute{
								Description: "Slurm controller prolog script.",
								Computed:    true,
							},
							"extra_slurm_conf": schema.StringAttribute{
								Description: "Additional slurm.conf fragments.",
								Computed:    true,
							},
							"login_init_script": schema.StringAttribute{
								Description: "Script run on Slurm login node init.",
								Computed:    true,
							},
							"nodeset_init_script": schema.StringAttribute{
								Description: "Script run on Slurm nodeset init.",
								Computed:    true,
							},
							"worker_epilog": schema.StringAttribute{
								Description: "Slurm worker node epilog script.",
								Computed:    true,
							},
							"worker_prolog": schema.StringAttribute{
								Description: "Slurm worker node prolog script.",
								Computed:    true,
							},
						},
					},
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
						"phase_transitions": schema.ListNestedAttribute{
							Description: "Phase transition history for this control plane node.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BetaClusterControlPlaneNodesPhaseTransitionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"phase": schema.StringAttribute{
										Description: "Node phase.\nAvailable values: \"NODE_PHASE_PENDING\", \"NODE_PHASE_SCHEDULING\", \"NODE_PHASE_BOOTING\", \"NODE_PHASE_BOOTSTRAPPING\", \"NODE_PHASE_RUNNING\", \"NODE_PHASE_SUCCEEDED\", \"NODE_PHASE_FAILED\", \"NODE_PHASE_PAUSED\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NODE_PHASE_PENDING",
												"NODE_PHASE_SCHEDULING",
												"NODE_PHASE_BOOTING",
												"NODE_PHASE_BOOTSTRAPPING",
												"NODE_PHASE_RUNNING",
												"NODE_PHASE_SUCCEEDED",
												"NODE_PHASE_FAILED",
												"NODE_PHASE_PAUSED",
											),
										},
									},
									"transition_time": schema.StringAttribute{
										Description: "Timestamp when the phase transition occurred.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
								},
							},
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
						"phase_transitions": schema.ListNestedAttribute{
							Description: "Phase transition history for this GPU worker node.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BetaClusterGPUWorkerNodesPhaseTransitionsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"phase": schema.StringAttribute{
										Description: "Node phase.\nAvailable values: \"NODE_PHASE_PENDING\", \"NODE_PHASE_SCHEDULING\", \"NODE_PHASE_BOOTING\", \"NODE_PHASE_BOOTSTRAPPING\", \"NODE_PHASE_RUNNING\", \"NODE_PHASE_SUCCEEDED\", \"NODE_PHASE_FAILED\", \"NODE_PHASE_PAUSED\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NODE_PHASE_PENDING",
												"NODE_PHASE_SCHEDULING",
												"NODE_PHASE_BOOTING",
												"NODE_PHASE_BOOTSTRAPPING",
												"NODE_PHASE_RUNNING",
												"NODE_PHASE_SUCCEEDED",
												"NODE_PHASE_FAILED",
												"NODE_PHASE_PAUSED",
											),
										},
									},
									"transition_time": schema.StringAttribute{
										Description: "Timestamp when the phase transition occurred.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
								},
							},
						},
						"status": schema.StringAttribute{
							Computed: true,
						},
						"instance_id": schema.StringAttribute{
							Computed: true,
						},
						"latest_remediation": schema.SingleNestedAttribute{
							Description: "Remediation represents a node remediation request for an instance.\nAn instance can have multiple remediations over time (e.g., failed attempts followed by retries).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[BetaClusterGPUWorkerNodesLatestRemediationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
								"cluster_id": schema.StringAttribute{
									Computed: true,
								},
								"instance_id": schema.StringAttribute{
									Computed: true,
								},
								"mode": schema.StringAttribute{
									Description: "Remediation mode specifies how the remediation should be performed.\n\n- `REMEDIATION_MODE_VM_ONLY`: Deletes the VM and provisions a new one on any available host.\n- `REMEDIATION_MODE_HOST_AWARE`: Cordons the host, deletes the VM, and provisions a new one on a different host.\nAvailable values: \"REMEDIATION_MODE_VM_ONLY\", \"REMEDIATION_MODE_HOST_AWARE\", \"REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT\", \"REMEDIATION_MODE_REBOOT_VM\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"REMEDIATION_MODE_VM_ONLY",
											"REMEDIATION_MODE_HOST_AWARE",
											"REMEDIATION_MODE_EVICT_WITHOUT_REPLACEMENT",
											"REMEDIATION_MODE_REBOOT_VM",
										),
									},
								},
								"state": schema.StringAttribute{
									Description: "RemediationState represents the lifecycle state of a remediation.\n\n- `PENDING_APPROVAL`: Awaiting approval before processing can begin.\n- `PENDING`: Approved and queued for processing.\n- `RUNNING`: Actively being processed.\n- `SUCCEEDED`: Successfully completed.\n- `FAILED`: Failed with an error.\n- `CANCELLED`: Cancelled by user or system.\n- `AUTO_RESOLVED`: The underlying issue was automatically resolved before processing.\nAvailable values: \"PENDING_APPROVAL\", \"PENDING\", \"RUNNING\", \"SUCCEEDED\", \"FAILED\", \"CANCELLED\", \"AUTO_RESOLVED\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"PENDING_APPROVAL",
											"PENDING",
											"RUNNING",
											"SUCCEEDED",
											"FAILED",
											"CANCELLED",
											"AUTO_RESOLVED",
										),
									},
								},
								"trigger": schema.StringAttribute{
									Description: "RemediationTrigger specifies how the remediation was triggered.\n\n- `REMEDIATION_TRIGGER_MANUAL`: A user-initiated remediation (either via web UI or API call).\n- `REMEDIATION_TRIGGER_AUTOMATED`: A system-initiated remediation that requires approval.\nAvailable values: \"REMEDIATION_TRIGGER_MANUAL\", \"REMEDIATION_TRIGGER_AUTOMATED\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("REMEDIATION_TRIGGER_MANUAL", "REMEDIATION_TRIGGER_AUTOMATED"),
									},
								},
								"active_health_check_run_id": schema.StringAttribute{
									Description: "Active health check run ID (UUID) that triggered this remediation.",
									Computed:    true,
								},
								"create_time": schema.StringAttribute{
									Description: "When the remediation was created.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"end_time": schema.StringAttribute{
									Description: "When the remediation completed.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"error_message": schema.StringAttribute{
									Description: "Error message if the remediation failed.",
									Computed:    true,
								},
								"passive_health_check_event_id": schema.StringAttribute{
									Description: "Passive health check event ID that triggered this remediation.",
									Computed:    true,
								},
								"reason": schema.StringAttribute{
									Description: "User-provided reason for the remediation.",
									Computed:    true,
								},
								"requested_by": schema.StringAttribute{
									Description: "Who requested the remediation.",
									Computed:    true,
								},
								"review_comment": schema.StringAttribute{
									Description: "Review comment.",
									Computed:    true,
								},
								"review_time": schema.StringAttribute{
									Description: "When the remediation was reviewed.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"reviewed_by": schema.StringAttribute{
									Description: "Who reviewed the remediation.",
									Computed:    true,
								},
								"start_time": schema.StringAttribute{
									Description: "When processing started.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"update_time": schema.StringAttribute{
									Description: "When the remediation was last updated.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
							},
						},
						"slurm_worker_hostname": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"oidc_config": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[BetaClusterOidcConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Description: "OIDC client ID for authentication.",
						Computed:    true,
					},
					"group_claim": schema.StringAttribute{
						Description: "JWT claim to use for user groups. For example, 'groups'",
						Computed:    true,
					},
					"group_prefix": schema.StringAttribute{
						Description: "Prefix to add to the group claim to form the final group name. For example, 'oidc:'",
						Computed:    true,
					},
					"issuer_url": schema.StringAttribute{
						Description: "OIDC issuer URL for authentication. For example, https://accounts.google.com",
						Computed:    true,
					},
					"username_claim": schema.StringAttribute{
						Description: "JWT claim to use as the username. For example, 'sub' or 'email'",
						Computed:    true,
					},
					"username_prefix": schema.StringAttribute{
						Description: "Prefix to add to the username claim to form the final username. For example, 'oidc:'",
						Computed:    true,
					},
					"ca_cert": schema.StringAttribute{
						Description: "CA certificate in PEM format to validate the OIDC issuer's TLS certificate. This field is optional but recommended if the issuer uses a private CA or self-signed certificate.",
						Computed:    true,
					},
				},
			},
			"phase_transitions": schema.ListNestedAttribute{
				Description: "Cluster-level phase transition history.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BetaClusterPhaseTransitionsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"phase": schema.StringAttribute{
							Description: "Cluster phase.\nAvailable values: \"CLUSTER_PHASE_QUEUED\", \"CLUSTER_PHASE_SCHEDULED\", \"CLUSTER_PHASE_WAITING_FOR_CONTROL_PLANE_NODES\", \"CLUSTER_PHASE_WAITING_FOR_DATA_PLANE_NODES\", \"CLUSTER_PHASE_WAITING_FOR_SUBNET\", \"CLUSTER_PHASE_WAITING_FOR_SHARED_VOLUME\", \"CLUSTER_PHASE_WAITING_FOR_AUTO_SCALER\", \"CLUSTER_PHASE_INSTALLING_DRIVERS\", \"CLUSTER_PHASE_RUNNING_ACCEPTANCE_TESTS\", \"CLUSTER_PHASE_ACCEPTANCE_TESTS_FAILED\", \"CLUSTER_PHASE_RUNNING_NCCL_TESTS\", \"CLUSTER_PHASE_NCCL_TESTS_FAILED\", \"CLUSTER_PHASE_READY\", \"CLUSTER_PHASE_PAUSED\", \"CLUSTER_PHASE_ON_DEMAND_COMPUTE_PAUSED\", \"CLUSTER_PHASE_DEGRADED\", \"CLUSTER_PHASE_DELETING\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"CLUSTER_PHASE_QUEUED",
									"CLUSTER_PHASE_SCHEDULED",
									"CLUSTER_PHASE_WAITING_FOR_CONTROL_PLANE_NODES",
									"CLUSTER_PHASE_WAITING_FOR_DATA_PLANE_NODES",
									"CLUSTER_PHASE_WAITING_FOR_SUBNET",
									"CLUSTER_PHASE_WAITING_FOR_SHARED_VOLUME",
									"CLUSTER_PHASE_WAITING_FOR_AUTO_SCALER",
									"CLUSTER_PHASE_INSTALLING_DRIVERS",
									"CLUSTER_PHASE_RUNNING_ACCEPTANCE_TESTS",
									"CLUSTER_PHASE_ACCEPTANCE_TESTS_FAILED",
									"CLUSTER_PHASE_RUNNING_NCCL_TESTS",
									"CLUSTER_PHASE_NCCL_TESTS_FAILED",
									"CLUSTER_PHASE_READY",
									"CLUSTER_PHASE_PAUSED",
									"CLUSTER_PHASE_ON_DEMAND_COMPUTE_PAUSED",
									"CLUSTER_PHASE_DEGRADED",
									"CLUSTER_PHASE_DELETING",
								),
							},
						},
						"transition_time": schema.StringAttribute{
							Description: "Timestamp when the phase transition occurred.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
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
							Description: "Size of the volume in TiB.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Current status of the volume.",
							Computed:    true,
						},
						"volume_id": schema.StringAttribute{
							Description: "ID of the volume.",
							Computed:    true,
						},
						"volume_name": schema.StringAttribute{
							Description: "User provided name of the volume.",
							Computed:    true,
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
