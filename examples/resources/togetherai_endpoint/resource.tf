resource "togetherai_endpoint" "example_endpoint" {
  autoscaling = {
    max_replicas = 5
    min_replicas = 2
  }
  hardware = "1x_nvidia_a100_80gb_sxm"
  model = "meta-llama/Llama-3-8b-chat-hf"
  disable_prompt_cache = true
  disable_speculative_decoding = true
  display_name = "My Llama3 70b endpoint"
  inactive_timeout = 60
  state = "STARTED"
}
