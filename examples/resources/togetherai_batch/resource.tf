resource "togetherai_batch" "example_batch" {
  endpoint = "/v1/chat/completions"
  input_file_id = "file-abc123def456ghi789"
  completion_window = "24h"
  model_id = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
  priority = 1
}
