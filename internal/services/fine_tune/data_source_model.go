// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tune

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type FineTuneDataSourceModel struct {
	ID                   types.String                                                    `tfsdk:"id" path:"id,required"`
	CreatedAt            types.String                                                    `tfsdk:"created_at" json:"created_at,computed"`
	EpochsCompleted      types.Int64                                                     `tfsdk:"epochs_completed" json:"epochs_completed,computed"`
	EvalSteps            types.Int64                                                     `tfsdk:"eval_steps" json:"eval_steps,computed"`
	FromCheckpoint       types.String                                                    `tfsdk:"from_checkpoint" json:"from_checkpoint,computed"`
	FromHfModel          types.String                                                    `tfsdk:"from_hf_model" json:"from_hf_model,computed"`
	HfModelRevision      types.String                                                    `tfsdk:"hf_model_revision" json:"hf_model_revision,computed"`
	JobID                types.String                                                    `tfsdk:"job_id" json:"job_id,computed"`
	LearningRate         types.Float64                                                   `tfsdk:"learning_rate" json:"learning_rate,computed"`
	MaxGradNorm          types.Float64                                                   `tfsdk:"max_grad_norm" json:"max_grad_norm,computed"`
	Model                types.String                                                    `tfsdk:"model" json:"model,computed"`
	ModelOutputName      types.String                                                    `tfsdk:"model_output_name" json:"model_output_name,computed"`
	ModelOutputPath      types.String                                                    `tfsdk:"model_output_path" json:"model_output_path,computed"`
	NCheckpoints         types.Int64                                                     `tfsdk:"n_checkpoints" json:"n_checkpoints,computed"`
	NEpochs              types.Int64                                                     `tfsdk:"n_epochs" json:"n_epochs,computed"`
	NEvals               types.Int64                                                     `tfsdk:"n_evals" json:"n_evals,computed"`
	ParamCount           types.Int64                                                     `tfsdk:"param_count" json:"param_count,computed"`
	QueueDepth           types.Int64                                                     `tfsdk:"queue_depth" json:"queue_depth,computed"`
	Status               types.String                                                    `tfsdk:"status" json:"status,computed"`
	TokenCount           types.Int64                                                     `tfsdk:"token_count" json:"token_count,computed"`
	TotalPrice           types.Int64                                                     `tfsdk:"total_price" json:"total_price,computed"`
	TrainingFile         types.String                                                    `tfsdk:"training_file" json:"training_file,computed"`
	TrainingfileNumlines types.Int64                                                     `tfsdk:"trainingfile_numlines" json:"trainingfile_numlines,computed"`
	TrainingfileSize     types.Int64                                                     `tfsdk:"trainingfile_size" json:"trainingfile_size,computed"`
	UpdatedAt            types.String                                                    `tfsdk:"updated_at" json:"updated_at,computed"`
	ValidationFile       types.String                                                    `tfsdk:"validation_file" json:"validation_file,computed"`
	WandbProjectName     types.String                                                    `tfsdk:"wandb_project_name" json:"wandb_project_name,computed"`
	WandbURL             types.String                                                    `tfsdk:"wandb_url" json:"wandb_url,computed"`
	WarmupRatio          types.Float64                                                   `tfsdk:"warmup_ratio" json:"warmup_ratio,computed"`
	WeightDecay          types.Float64                                                   `tfsdk:"weight_decay" json:"weight_decay,computed"`
	Events               customfield.NestedObjectList[FineTuneEventsDataSourceModel]     `tfsdk:"events" json:"events,computed"`
	LrScheduler          customfield.NestedObject[FineTuneLrSchedulerDataSourceModel]    `tfsdk:"lr_scheduler" json:"lr_scheduler,computed"`
	TrainingMethod       customfield.NestedObject[FineTuneTrainingMethodDataSourceModel] `tfsdk:"training_method" json:"training_method,computed"`
	TrainingType         customfield.NestedObject[FineTuneTrainingTypeDataSourceModel]   `tfsdk:"training_type" json:"training_type,computed"`
	BatchSize            customfield.NormalizedDynamicValue                              `tfsdk:"batch_size" json:"batch_size,computed"`
	TrainOnInputs        customfield.NormalizedDynamicValue                              `tfsdk:"train_on_inputs" json:"train_on_inputs,computed"`
}

type FineTuneEventsDataSourceModel struct {
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

type FineTuneLrSchedulerDataSourceModel struct {
	LrSchedulerType types.String                                                                `tfsdk:"lr_scheduler_type" json:"lr_scheduler_type,computed"`
	LrSchedulerArgs customfield.NestedObject[FineTuneLrSchedulerLrSchedulerArgsDataSourceModel] `tfsdk:"lr_scheduler_args" json:"lr_scheduler_args,computed"`
}

type FineTuneLrSchedulerLrSchedulerArgsDataSourceModel struct {
	MinLrRatio types.Float64 `tfsdk:"min_lr_ratio" json:"min_lr_ratio,computed"`
	NumCycles  types.Float64 `tfsdk:"num_cycles" json:"num_cycles,computed"`
}

type FineTuneTrainingMethodDataSourceModel struct {
	Method                        types.String                       `tfsdk:"method" json:"method,computed"`
	TrainOnInputs                 customfield.NormalizedDynamicValue `tfsdk:"train_on_inputs" json:"train_on_inputs,computed"`
	DpoBeta                       types.Float64                      `tfsdk:"dpo_beta" json:"dpo_beta,computed"`
	DpoNormalizeLogratiosByLength types.Bool                         `tfsdk:"dpo_normalize_logratios_by_length" json:"dpo_normalize_logratios_by_length,computed"`
	DpoReferenceFree              types.Bool                         `tfsdk:"dpo_reference_free" json:"dpo_reference_free,computed"`
	RpoAlpha                      types.Float64                      `tfsdk:"rpo_alpha" json:"rpo_alpha,computed"`
	SimpoGamma                    types.Float64                      `tfsdk:"simpo_gamma" json:"simpo_gamma,computed"`
}

type FineTuneTrainingTypeDataSourceModel struct {
	Type                 types.String  `tfsdk:"type" json:"type,computed"`
	LoraAlpha            types.Int64   `tfsdk:"lora_alpha" json:"lora_alpha,computed"`
	LoraR                types.Int64   `tfsdk:"lora_r" json:"lora_r,computed"`
	LoraDropout          types.Float64 `tfsdk:"lora_dropout" json:"lora_dropout,computed"`
	LoraTrainableModules types.String  `tfsdk:"lora_trainable_modules" json:"lora_trainable_modules,computed"`
}
