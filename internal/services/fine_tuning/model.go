// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tuning

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type FineTuningModel struct {
	ID                   types.String                                            `tfsdk:"id" json:"id,computed"`
	Model                types.String                                            `tfsdk:"model" json:"model,required"`
	TrainingFile         types.String                                            `tfsdk:"training_file" json:"training_file,required"`
	FromCheckpoint       types.String                                            `tfsdk:"from_checkpoint" json:"from_checkpoint,optional"`
	FromHfModel          types.String                                            `tfsdk:"from_hf_model" json:"from_hf_model,optional"`
	HfAPIToken           types.String                                            `tfsdk:"hf_api_token" json:"hf_api_token,optional,no_refresh"`
	HfModelRevision      types.String                                            `tfsdk:"hf_model_revision" json:"hf_model_revision,optional"`
	HfOutputRepoName     types.String                                            `tfsdk:"hf_output_repo_name" json:"hf_output_repo_name,optional,no_refresh"`
	Suffix               types.String                                            `tfsdk:"suffix" json:"suffix,optional,no_refresh"`
	ValidationFile       types.String                                            `tfsdk:"validation_file" json:"validation_file,optional"`
	WandbAPIKey          types.String                                            `tfsdk:"wandb_api_key" json:"wandb_api_key,optional,no_refresh"`
	WandbBaseURL         types.String                                            `tfsdk:"wandb_base_url" json:"wandb_base_url,optional,no_refresh"`
	WandbName            types.String                                            `tfsdk:"wandb_name" json:"wandb_name,optional,no_refresh"`
	WandbProjectName     types.String                                            `tfsdk:"wandb_project_name" json:"wandb_project_name,optional"`
	LearningRate         types.Float64                                           `tfsdk:"learning_rate" json:"learning_rate,computed_optional"`
	MaxGradNorm          types.Float64                                           `tfsdk:"max_grad_norm" json:"max_grad_norm,computed_optional"`
	NCheckpoints         types.Int64                                             `tfsdk:"n_checkpoints" json:"n_checkpoints,computed_optional"`
	NEpochs              types.Int64                                             `tfsdk:"n_epochs" json:"n_epochs,computed_optional"`
	NEvals               types.Int64                                             `tfsdk:"n_evals" json:"n_evals,computed_optional"`
	WarmupRatio          types.Float64                                           `tfsdk:"warmup_ratio" json:"warmup_ratio,computed_optional"`
	WeightDecay          types.Float64                                           `tfsdk:"weight_decay" json:"weight_decay,computed_optional"`
	LrScheduler          customfield.NestedObject[FineTuningLrSchedulerModel]    `tfsdk:"lr_scheduler" json:"lr_scheduler,computed_optional"`
	TrainingMethod       customfield.NestedObject[FineTuningTrainingMethodModel] `tfsdk:"training_method" json:"training_method,computed_optional"`
	TrainingType         customfield.NestedObject[FineTuningTrainingTypeModel]   `tfsdk:"training_type" json:"training_type,computed_optional"`
	BatchSize            customfield.NormalizedDynamicValue                      `tfsdk:"batch_size" json:"batch_size,computed_optional"`
	TrainOnInputs        customfield.NormalizedDynamicValue                      `tfsdk:"train_on_inputs" json:"train_on_inputs,computed_optional"`
	CreatedAt            types.String                                            `tfsdk:"created_at" json:"created_at,computed"`
	EpochsCompleted      types.Int64                                             `tfsdk:"epochs_completed" json:"epochs_completed,computed"`
	EvalSteps            types.Int64                                             `tfsdk:"eval_steps" json:"eval_steps,computed"`
	JobID                types.String                                            `tfsdk:"job_id" json:"job_id,computed"`
	ModelOutputName      types.String                                            `tfsdk:"model_output_name" json:"model_output_name,computed"`
	ModelOutputPath      types.String                                            `tfsdk:"model_output_path" json:"model_output_path,computed"`
	OwnerAddress         types.String                                            `tfsdk:"owner_address" json:"owner_address,computed,no_refresh"`
	ParamCount           types.Int64                                             `tfsdk:"param_count" json:"param_count,computed"`
	QueueDepth           types.Int64                                             `tfsdk:"queue_depth" json:"queue_depth,computed"`
	Status               types.String                                            `tfsdk:"status" json:"status,computed"`
	TokenCount           types.Int64                                             `tfsdk:"token_count" json:"token_count,computed"`
	TotalPrice           types.Int64                                             `tfsdk:"total_price" json:"total_price,computed"`
	TrainingfileNumlines types.Int64                                             `tfsdk:"trainingfile_numlines" json:"trainingfile_numlines,computed"`
	TrainingfileSize     types.Int64                                             `tfsdk:"trainingfile_size" json:"trainingfile_size,computed"`
	UpdatedAt            types.String                                            `tfsdk:"updated_at" json:"updated_at,computed"`
	UserID               types.String                                            `tfsdk:"user_id" json:"user_id,computed,no_refresh"`
	WandbURL             types.String                                            `tfsdk:"wandb_url" json:"wandb_url,computed"`
	Events               customfield.NestedObjectList[FineTuningEventsModel]     `tfsdk:"events" json:"events,computed"`
}

func (m FineTuningModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FineTuningModel) MarshalJSONForUpdate(state FineTuningModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type FineTuningLrSchedulerModel struct {
	LrSchedulerType types.String                               `tfsdk:"lr_scheduler_type" json:"lr_scheduler_type,required"`
	LrSchedulerArgs *FineTuningLrSchedulerLrSchedulerArgsModel `tfsdk:"lr_scheduler_args" json:"lr_scheduler_args,optional"`
}

type FineTuningLrSchedulerLrSchedulerArgsModel struct {
	MinLrRatio types.Float64 `tfsdk:"min_lr_ratio" json:"min_lr_ratio,computed_optional"`
	NumCycles  types.Float64 `tfsdk:"num_cycles" json:"num_cycles,computed_optional"`
}

type FineTuningTrainingMethodModel struct {
	Method                        types.String                       `tfsdk:"method" json:"method,required"`
	TrainOnInputs                 customfield.NormalizedDynamicValue `tfsdk:"train_on_inputs" json:"train_on_inputs,computed_optional"`
	DpoBeta                       types.Float64                      `tfsdk:"dpo_beta" json:"dpo_beta,computed_optional"`
	DpoNormalizeLogratiosByLength types.Bool                         `tfsdk:"dpo_normalize_logratios_by_length" json:"dpo_normalize_logratios_by_length,computed_optional"`
	DpoReferenceFree              types.Bool                         `tfsdk:"dpo_reference_free" json:"dpo_reference_free,computed_optional"`
	RpoAlpha                      types.Float64                      `tfsdk:"rpo_alpha" json:"rpo_alpha,computed_optional"`
	SimpoGamma                    types.Float64                      `tfsdk:"simpo_gamma" json:"simpo_gamma,computed_optional"`
}

type FineTuningTrainingTypeModel struct {
	Type                 types.String  `tfsdk:"type" json:"type,required"`
	LoraAlpha            types.Int64   `tfsdk:"lora_alpha" json:"lora_alpha,optional"`
	LoraR                types.Int64   `tfsdk:"lora_r" json:"lora_r,optional"`
	LoraDropout          types.Float64 `tfsdk:"lora_dropout" json:"lora_dropout,computed_optional"`
	LoraTrainableModules types.String  `tfsdk:"lora_trainable_modules" json:"lora_trainable_modules,computed_optional"`
}

type FineTuningEventsModel struct {
	CheckpointPath types.String `tfsdk:"checkpoint_path" json:"checkpoint_path,computed"`
	CreatedAt      types.String `tfsdk:"created_at" json:"created_at,computed"`
	Hash           types.String `tfsdk:"hash" json:"hash,computed"`
	Message        types.String `tfsdk:"message" json:"message,computed"`
	ModelPath      types.String `tfsdk:"model_path" json:"model_path,computed"`
	Object         types.String `tfsdk:"object" json:"object,computed"`
	ParamCount     types.Int64  `tfsdk:"param_count" json:"param_count,computed"`
	Step           types.Int64  `tfsdk:"step" json:"step,computed"`
	TokenCount     types.Int64  `tfsdk:"token_count" json:"token_count,computed"`
	TotalSteps     types.Int64  `tfsdk:"total_steps" json:"total_steps,computed"`
	TrainingOffset types.Int64  `tfsdk:"training_offset" json:"training_offset,computed"`
	Type           types.String `tfsdk:"type" json:"type,computed"`
	WandbURL       types.String `tfsdk:"wandb_url" json:"wandb_url,computed"`
	Level          types.String `tfsdk:"level" json:"level,computed"`
}
