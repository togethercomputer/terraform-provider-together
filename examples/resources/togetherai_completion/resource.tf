resource "togetherai_completion" "example_completion" {
  model = "mistralai/Mixtral-8x7B-Instruct-v0.1"
  prompt = "<s>[INST] What is the capital of France? [/INST]"
  echo = true
  frequency_penalty = 0
  logit_bias = {
    "105" = 21.4
    "1024" = -10.5
  }
  logprobs = 0
  max_tokens = 0
  min_p = 0
  n = 1
  presence_penalty = 0
  repetition_penalty = 0
  safety_model = "safety_model_name"
  seed = 42
  stop = ["string"]
  stream = false
  temperature = 0
  top_k = 0
  top_p = 0
}
