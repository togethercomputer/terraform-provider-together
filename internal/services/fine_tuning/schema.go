// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tuning

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
	"github.com/stainless-sdks/togetherai-terraform/internal/customvalidator"
)

var _ resource.ResourceWithConfigValidators = (*FineTuningResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the fine-tune job",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "Name of the base model to run fine-tune job on",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"training_file": schema.StringAttribute{
				Description:   "File-ID of a training file uploaded to the Together API",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"from_checkpoint": schema.StringAttribute{
				Description:   "The checkpoint identifier to continue training from a previous fine-tuning job. Format is `{$JOB_ID}` or `{$OUTPUT_MODEL_NAME}` or `{$JOB_ID}:{$STEP}` or `{$OUTPUT_MODEL_NAME}:{$STEP}`. The step value is optional; without it, the final checkpoint will be used.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"from_hf_model": schema.StringAttribute{
				Description:   "The Hugging Face Hub repo to start training from. Should be as close as possible to the base model (specified by the `model` argument) in terms of architecture and size.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hf_api_token": schema.StringAttribute{
				Description:   "The API token for the Hugging Face Hub.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hf_model_revision": schema.StringAttribute{
				Description:   "The revision of the Hugging Face Hub model to continue training from. E.g., hf_model_revision=main (default, used if the argument is not provided) or hf_model_revision='607a30d783dfa663caf39e06633721c8d4cfcd7e' (specific commit).",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"hf_output_repo_name": schema.StringAttribute{
				Description:   "The name of the Hugging Face repository to upload the fine-tuned model to.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"suffix": schema.StringAttribute{
				Description:   "Suffix that will be added to your fine-tuned model name",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"validation_file": schema.StringAttribute{
				Description:   "File-ID of a validation file uploaded to the Together API",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"wandb_api_key": schema.StringAttribute{
				Description:   "Integration key for tracking experiments and model metrics on W&B platform",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"wandb_base_url": schema.StringAttribute{
				Description:   "The base URL of a dedicated Weights & Biases instance.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"wandb_name": schema.StringAttribute{
				Description:   "The Weights & Biases name for your run.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"wandb_project_name": schema.StringAttribute{
				Description:   "The Weights & Biases project for your run. If not specified, will use `together` as the project name.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"learning_rate": schema.Float64Attribute{
				Description:   "Controls how quickly the model adapts to new information (too high may cause instability, too low may slow convergence)",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(0.00001),
			},
			"max_grad_norm": schema.Float64Attribute{
				Description:   "Max gradient norm to be used for gradient clipping. Set to 0 to disable.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(1),
			},
			"n_checkpoints": schema.Int64Attribute{
				Description:   "Number of intermediate model versions saved during training for evaluation",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(1),
			},
			"n_epochs": schema.Int64Attribute{
				Description:   "Number of complete passes through the training dataset (higher values may improve results but increase cost and risk of overfitting)",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(1),
			},
			"n_evals": schema.Int64Attribute{
				Description:   "Number of evaluations to be run on a given validation set during training",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(0),
			},
			"warmup_ratio": schema.Float64Attribute{
				Description:   "The percent of steps at the start of training to linearly increase the learning rate.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(0),
			},
			"weight_decay": schema.Float64Attribute{
				Description:   "Weight decay. Regularization parameter for the optimizer.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(0),
			},
			"lr_scheduler": schema.SingleNestedAttribute{
				Description: "The learning rate scheduler to use. It specifies how the learning rate is adjusted during training.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[FineTuningLrSchedulerModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lr_scheduler_type": schema.StringAttribute{
						Description: `Available values: "linear", "cosine".`,
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("linear", "cosine"),
						},
					},
					"lr_scheduler_args": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[FineTuningLrSchedulerLrSchedulerArgsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"min_lr_ratio": schema.Float64Attribute{
								Description: "The ratio of the final learning rate to the peak learning rate",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(0),
							},
							"num_cycles": schema.Float64Attribute{
								Description: "Number or fraction of cycles for the cosine learning rate scheduler",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(0.5),
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"training_method": schema.SingleNestedAttribute{
				Description: "The training method to use. 'sft' for Supervised Fine-Tuning or 'dpo' for Direct Preference Optimization.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[FineTuningTrainingMethodModel](ctx),
				Attributes: map[string]schema.Attribute{
					"method": schema.StringAttribute{
						Description: `Available values: "sft", "dpo".`,
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("sft", "dpo"),
						},
					},
					"train_on_inputs": schema.DynamicAttribute{
						Description: "Whether to mask the user messages in conversational data or prompts in instruction data.",
						Computed:    true,
						Optional:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.BoolType{}, basetypes.StringType{}),
						},
						CustomType:    customfield.NormalizedDynamicType{},
						PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
					},
					"dpo_beta": schema.Float64Attribute{
						Computed: true,
						Optional: true,
						Default:  float64default.StaticFloat64(0.1),
					},
					"dpo_normalize_logratios_by_length": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"dpo_reference_free": schema.BoolAttribute{
						Computed: true,
						Optional: true,
						Default:  booldefault.StaticBool(false),
					},
					"rpo_alpha": schema.Float64Attribute{
						Computed: true,
						Optional: true,
						Default:  float64default.StaticFloat64(0),
					},
					"simpo_gamma": schema.Float64Attribute{
						Computed: true,
						Optional: true,
						Default:  float64default.StaticFloat64(0),
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"training_type": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[FineTuningTrainingTypeModel](ctx),
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "Full", "Lora".`,
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("Full", "Lora"),
						},
					},
					"lora_alpha": schema.Int64Attribute{
						Optional: true,
					},
					"lora_r": schema.Int64Attribute{
						Optional: true,
					},
					"lora_dropout": schema.Float64Attribute{
						Computed: true,
						Optional: true,
						Default:  float64default.StaticFloat64(0),
					},
					"lora_trainable_modules": schema.StringAttribute{
						Computed: true,
						Optional: true,
						Default:  stringdefault.StaticString("all-linear"),
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"batch_size": schema.DynamicAttribute{
				Description: `Number of training examples processed together (larger batches use more memory but may train faster). Defaults to "max". We use training optimizations like packing, so the effective batch size may be different than the value you set.`,
				Computed:    true,
				Optional:    true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
				},
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"train_on_inputs": schema.DynamicAttribute{
				Description:        "Whether to mask the user messages in conversational data or prompts in instruction data.",
				Computed:           true,
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.BoolType{}, basetypes.StringType{}),
				},
				CustomType:    customfield.NormalizedDynamicType{},
				PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"epochs_completed": schema.Int64Attribute{
				Computed: true,
			},
			"eval_steps": schema.Int64Attribute{
				Computed: true,
			},
			"job_id": schema.StringAttribute{
				Computed: true,
			},
			"model_output_name": schema.StringAttribute{
				Computed: true,
			},
			"model_output_path": schema.StringAttribute{
				Computed: true,
			},
			"owner_address": schema.StringAttribute{
				Description: "Owner address information",
				Computed:    true,
			},
			"param_count": schema.Int64Attribute{
				Computed: true,
			},
			"queue_depth": schema.Int64Attribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Description: `Available values: "pending", "queued", "running", "compressing", "uploading", "cancel_requested", "cancelled", "error", "completed".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending",
						"queued",
						"running",
						"compressing",
						"uploading",
						"cancel_requested",
						"cancelled",
						"error",
						"completed",
					),
				},
			},
			"token_count": schema.Int64Attribute{
				Computed: true,
			},
			"total_price": schema.Int64Attribute{
				Computed: true,
			},
			"trainingfile_numlines": schema.Int64Attribute{
				Computed: true,
			},
			"trainingfile_size": schema.Int64Attribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"user_id": schema.StringAttribute{
				Description: "Identifier for the user who created the job",
				Computed:    true,
			},
			"wandb_url": schema.StringAttribute{
				Computed: true,
			},
			"events": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[FineTuningEventsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"checkpoint_path": schema.StringAttribute{
							Computed: true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"hash": schema.StringAttribute{
							Computed: true,
						},
						"message": schema.StringAttribute{
							Computed: true,
						},
						"model_path": schema.StringAttribute{
							Computed: true,
						},
						"object": schema.StringAttribute{
							Description: `Available values: "fine-tune-event".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("fine-tune-event"),
							},
						},
						"param_count": schema.Int64Attribute{
							Computed: true,
						},
						"step": schema.Int64Attribute{
							Computed: true,
						},
						"token_count": schema.Int64Attribute{
							Computed: true,
						},
						"total_steps": schema.Int64Attribute{
							Computed: true,
						},
						"training_offset": schema.Int64Attribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: `Available values: "job_pending", "job_start", "job_stopped", "model_downloading", "model_download_complete", "training_data_downloading", "training_data_download_complete", "validation_data_downloading", "validation_data_download_complete", "wandb_init", "training_start", "checkpoint_save", "billing_limit", "epoch_complete", "training_complete", "model_compressing", "model_compression_complete", "model_uploading", "model_upload_complete", "job_complete", "job_error", "cancel_requested", "job_restarted", "refund", "warning".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"job_pending",
									"job_start",
									"job_stopped",
									"model_downloading",
									"model_download_complete",
									"training_data_downloading",
									"training_data_download_complete",
									"validation_data_downloading",
									"validation_data_download_complete",
									"wandb_init",
									"training_start",
									"checkpoint_save",
									"billing_limit",
									"epoch_complete",
									"training_complete",
									"model_compressing",
									"model_compression_complete",
									"model_uploading",
									"model_upload_complete",
									"job_complete",
									"job_error",
									"cancel_requested",
									"job_restarted",
									"refund",
									"warning",
								),
							},
						},
						"wandb_url": schema.StringAttribute{
							Computed: true,
						},
						"level": schema.StringAttribute{
							Description: `Available values: "info", "warning", "error", "legacy_info", "legacy_iwarning", "legacy_ierror".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"info",
									"warning",
									"error",
									"legacy_info",
									"legacy_iwarning",
									"legacy_ierror",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *FineTuningResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *FineTuningResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
