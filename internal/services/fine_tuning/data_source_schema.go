// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tuning

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
	"github.com/stainless-sdks/togetherai-terraform/internal/customvalidator"
)

var _ datasource.DataSourceWithConfigValidators = (*FineTuningDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
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
			"from_checkpoint": schema.StringAttribute{
				Computed: true,
			},
			"from_hf_model": schema.StringAttribute{
				Computed: true,
			},
			"hf_model_revision": schema.StringAttribute{
				Computed: true,
			},
			"job_id": schema.StringAttribute{
				Computed: true,
			},
			"learning_rate": schema.Float64Attribute{
				Computed: true,
			},
			"max_grad_norm": schema.Float64Attribute{
				Computed: true,
			},
			"model": schema.StringAttribute{
				Computed: true,
			},
			"model_output_name": schema.StringAttribute{
				Computed: true,
			},
			"model_output_path": schema.StringAttribute{
				Computed: true,
			},
			"n_checkpoints": schema.Int64Attribute{
				Computed: true,
			},
			"n_epochs": schema.Int64Attribute{
				Computed: true,
			},
			"n_evals": schema.Int64Attribute{
				Computed: true,
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
			"training_file": schema.StringAttribute{
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
			"validation_file": schema.StringAttribute{
				Computed: true,
			},
			"wandb_project_name": schema.StringAttribute{
				Computed: true,
			},
			"wandb_url": schema.StringAttribute{
				Computed: true,
			},
			"warmup_ratio": schema.Float64Attribute{
				Computed: true,
			},
			"weight_decay": schema.Float64Attribute{
				Computed: true,
			},
			"events": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[FineTuningEventsDataSourceModel](ctx),
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
			"lr_scheduler": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[FineTuningLrSchedulerDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"lr_scheduler_type": schema.StringAttribute{
						Description: `Available values: "linear", "cosine".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("linear", "cosine"),
						},
					},
					"lr_scheduler_args": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[FineTuningLrSchedulerLrSchedulerArgsDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"min_lr_ratio": schema.Float64Attribute{
								Description: "The ratio of the final learning rate to the peak learning rate",
								Computed:    true,
							},
							"num_cycles": schema.Float64Attribute{
								Description: "Number or fraction of cycles for the cosine learning rate scheduler",
								Computed:    true,
							},
						},
					},
				},
			},
			"training_method": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[FineTuningTrainingMethodDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"method": schema.StringAttribute{
						Description: `Available values: "sft", "dpo".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("sft", "dpo"),
						},
					},
					"train_on_inputs": schema.DynamicAttribute{
						Description: "Whether to mask the user messages in conversational data or prompts in instruction data.",
						Computed:    true,
						Validators: []validator.Dynamic{
							customvalidator.AllowedSubtypes(basetypes.BoolType{}, basetypes.StringType{}),
						},
						CustomType: customfield.NormalizedDynamicType{},
					},
					"dpo_beta": schema.Float64Attribute{
						Computed: true,
					},
					"dpo_normalize_logratios_by_length": schema.BoolAttribute{
						Computed: true,
					},
					"dpo_reference_free": schema.BoolAttribute{
						Computed: true,
					},
					"rpo_alpha": schema.Float64Attribute{
						Computed: true,
					},
					"simpo_gamma": schema.Float64Attribute{
						Computed: true,
					},
				},
			},
			"training_type": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[FineTuningTrainingTypeDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"type": schema.StringAttribute{
						Description: `Available values: "Full", "Lora".`,
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("Full", "Lora"),
						},
					},
					"lora_alpha": schema.Int64Attribute{
						Computed: true,
					},
					"lora_r": schema.Int64Attribute{
						Computed: true,
					},
					"lora_dropout": schema.Float64Attribute{
						Computed: true,
					},
					"lora_trainable_modules": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"batch_size": schema.DynamicAttribute{
				Computed: true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.Int64Type{}, basetypes.StringType{}),
				},
				CustomType: customfield.NormalizedDynamicType{},
			},
			"train_on_inputs": schema.DynamicAttribute{
				Computed: true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.BoolType{}, basetypes.StringType{}),
				},
				CustomType: customfield.NormalizedDynamicType{},
			},
		},
	}
}

func (d *FineTuningDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FineTuningDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
