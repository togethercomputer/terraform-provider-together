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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"cluster_id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
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
			"auto_scale": schema.BoolAttribute{
				Description:   "Whether to enable auto-scaling for the cluster. If true, the cluster will automatically scale the number of GPU worker nodes between num_gpus and auto_scale_max_gpus based on the workload.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
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
			"num_capacity_pool_gpus": schema.Int64Attribute{
				Description:   "Number of GPUs to allocate from the capacity pool. Must be a multiple of 8 and not exceed num_gpus.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"project_id": schema.StringAttribute{
				Description:   "Project ID for the cluster. If not set, the project from the request context is used.",
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
			"acceptance_tests_params": schema.SingleNestedAttribute{
				Description: "AcceptanceTestsParams groups all GPU acceptance test options when enabled is true.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"dcgm_diag_level": schema.StringAttribute{
						Description: "DCGM diagnostic depth. SHORT = readiness; MEDIUM = default; LONG = system validation; EXTENDED = memtest. An omitted value selects MEDIUM when enabled.\nAvailable values: \"DCGM_DIAG_LEVEL_SHORT\", \"DCGM_DIAG_LEVEL_MEDIUM\", \"DCGM_DIAG_LEVEL_LONG\", \"DCGM_DIAG_LEVEL_EXTENDED\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"DCGM_DIAG_LEVEL_SHORT",
								"DCGM_DIAG_LEVEL_MEDIUM",
								"DCGM_DIAG_LEVEL_LONG",
								"DCGM_DIAG_LEVEL_EXTENDED",
							),
						},
					},
					"dcgm_diag_skipped": schema.BoolAttribute{
						Description: "Skip DCGM diagnostics acceptance test.",
						Optional:    true,
					},
					"enabled": schema.BoolAttribute{
						Description: "Whether to run GPU acceptance tests during cluster bring-up.",
						Optional:    true,
					},
					"gpu_burn_duration": schema.Int64Attribute{
						Description: "GPU burn duration in seconds; 0 means use the default when enabled.",
						Optional:    true,
					},
					"gpu_burn_skipped": schema.BoolAttribute{
						Description: "Skip GPU burn acceptance test.",
						Optional:    true,
					},
					"nccl_multi_node_skipped": schema.BoolAttribute{
						Description: "Skip NCCL multi-node acceptance test.",
						Optional:    true,
					},
					"nccl_single_node_skipped": schema.BoolAttribute{
						Description: "Skip NCCL single-node acceptance test.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"oidc_config": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"client_id": schema.StringAttribute{
						Description: "OIDC client ID for authentication.",
						Required:    true,
					},
					"group_claim": schema.StringAttribute{
						Description: "JWT claim to use for user groups. For example, 'groups'",
						Required:    true,
					},
					"group_prefix": schema.StringAttribute{
						Description: "Prefix to add to the group claim to form the final group name. For example, 'oidc:'",
						Required:    true,
					},
					"issuer_url": schema.StringAttribute{
						Description: "OIDC issuer URL for authentication. For example, https://accounts.google.com",
						Required:    true,
					},
					"username_claim": schema.StringAttribute{
						Description: "JWT claim to use as the username. For example, 'sub' or 'email'",
						Required:    true,
					},
					"username_prefix": schema.StringAttribute{
						Description: "Prefix to add to the username claim to form the final username. For example, 'oidc:'",
						Required:    true,
					},
					"ca_cert": schema.StringAttribute{
						Description: "CA certificate in PEM format to validate the OIDC issuer's TLS certificate. This field is optional but recommended if the issuer uses a private CA or self-signed certificate.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"auto_scaled": schema.BoolAttribute{
				Description:        "Whether GPU cluster should be auto-scaled based on the workload. By default, it is not auto-scaled.",
				Computed:           true,
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
				PlanModifiers:      []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:            booldefault.StaticBool(false),
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
			"num_preemptible_gpus": schema.Int64Attribute{
				Description: "Number of preemptible GPUs to request alongside on-demand capacity. Must be a multiple of 8. Preemptible nodes are cheaper but may be reclaimed when on-demand capacity is needed elsewhere; the system fulfills this asynchronously and surfaces the actual count in allocated_preemptible_gpus.",
				Optional:    true,
			},
			"num_reserved_gpus": schema.Int64Attribute{
				Description: "Number of prepaid (PLG) reserved GPUs for this cluster. When omitted for RESERVED billing on create, the server defaults this to num_gpus.",
				Optional:    true,
			},
			"reservation_end_time": schema.StringAttribute{
				Description: "Reservation end time of the cluster. This field is required for SCHEDULED billing to specify the reservation end time for the cluster.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"add_ons": schema.ListNestedAttribute{
				Description: "Add-ons to enable on the cluster at creation time.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"add_on_type": schema.StringAttribute{
							Description: "Type of add-on. Valid values: 'dashboard', 'ingress'.",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "Human-readable name for this add-on instance.",
							Required:    true,
						},
						"config": schema.SingleNestedAttribute{
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"dashboard": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Optional: true,
										},
									},
								},
								"ingress": schema.SingleNestedAttribute{
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
			"cluster_config": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"load_balancer": schema.StringAttribute{
						Description: `Available values: "NONE", "TRAEFIK", "NGINX", "ISTIO".`,
						Required:    true,
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
						Optional:    true,
					},
					"ingress": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Optional: true,
							},
						},
					},
					"jumphost_enabled": schema.BoolAttribute{
						Optional: true,
					},
					"kubernetes_dashboard_enabled": schema.BoolAttribute{
						Optional: true,
					},
					"observability": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Optional: true,
							},
						},
					},
					"slurm_startup_scripts": schema.SingleNestedAttribute{
						Description: "SlurmStartupScripts carries optional Slurm lifecycle scripts (prolog/epilog, init, extra conf).",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"controller_epilog": schema.StringAttribute{
								Description: "Slurm controller epilog script.",
								Optional:    true,
							},
							"controller_prolog": schema.StringAttribute{
								Description: "Slurm controller prolog script.",
								Optional:    true,
							},
							"extra_slurm_conf": schema.StringAttribute{
								Description: "Additional slurm.conf fragments.",
								Optional:    true,
							},
							"login_init_script": schema.StringAttribute{
								Description: "Script run on Slurm login node init.",
								Optional:    true,
							},
							"nodeset_init_script": schema.StringAttribute{
								Description: "Script run on Slurm nodeset init.",
								Optional:    true,
							},
							"worker_epilog": schema.StringAttribute{
								Description: "Slurm worker node epilog script.",
								Optional:    true,
							},
							"worker_prolog": schema.StringAttribute{
								Description: "Slurm worker node prolog script.",
								Optional:    true,
							},
						},
					},
				},
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
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"desired_preemptible_gpus": schema.Int64Attribute{
				Description: "Customer's requested number of preemptible GPUs. Set on cluster create or update; persists until changed.",
				Computed:    true,
			},
			"duration_hours": schema.Int64Attribute{
				Computed: true,
			},
			"kube_config": schema.StringAttribute{
				Computed: true,
			},
			"num_cpu_workers": schema.Int64Attribute{
				Description: "Number of CPU-only worker nodes in the cluster.",
				Computed:    true,
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
						"phase_transitions": schema.ListNestedAttribute{
							Description: "Phase transition history for this control plane node.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BetaClusterControlPlaneNodesPhaseTransitionsModel](ctx),
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
						"phase_transitions": schema.ListNestedAttribute{
							Description: "Phase transition history for this GPU worker node.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[BetaClusterGPUWorkerNodesPhaseTransitionsModel](ctx),
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
							CustomType:  customfield.NewNestedObjectType[BetaClusterGPUWorkerNodesLatestRemediationModel](ctx),
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
								"instance_name": schema.StringAttribute{
									Description: "Display name of the targeted instance.",
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
			"phase_transitions": schema.ListNestedAttribute{
				Description: "Cluster-level phase transition history.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[BetaClusterPhaseTransitionsModel](ctx),
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
				CustomType: customfield.NewNestedObjectListType[BetaClusterVolumesModel](ctx),
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

func (r *BetaClusterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BetaClusterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
