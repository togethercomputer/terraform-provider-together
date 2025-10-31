resource "togetherai_fine_tune" "example_fine_tune" {
  model = "model"
  training_file = "training_file"
  batch_size = 0
  from_checkpoint = "from_checkpoint"
  from_hf_model = "from_hf_model"
  hf_api_token = "hf_api_token"
  hf_model_revision = "hf_model_revision"
  hf_output_repo_name = "hf_output_repo_name"
  learning_rate = 0
  lr_scheduler = {
    lr_scheduler_type = "linear"
    lr_scheduler_args = {
      min_lr_ratio = 0
    }
  }
  max_grad_norm = 0
  n_checkpoints = 0
  n_epochs = 0
  n_evals = 0
  suffix = "suffix"
  train_on_inputs = true
  training_method = {
    method = "sft"
    train_on_inputs = true
  }
  training_type = {
    type = "Full"
  }
  validation_file = "validation_file"
  wandb_api_key = "wandb_api_key"
  wandb_base_url = "wandb_base_url"
  wandb_name = "wandb_name"
  wandb_project_name = "wandb_project_name"
  warmup_ratio = 0
  weight_decay = 0
}
