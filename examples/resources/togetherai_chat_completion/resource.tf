resource "togetherai_chat_completion" "example_chat_completion" {
  messages = [{
    content = "content"
    role = "system"
    name = "name"
  }]
  model = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
  context_length_exceeded_behavior = "truncate"
  echo = true
  frequency_penalty = 0
  function_call = "none"
  logit_bias = {
    "105" = 21.4
    "1024" = -10.5
  }
  logprobs = 0
  max_tokens = 0
  min_p = 0
  n = 1
  presence_penalty = 0
  reasoning_effort = "medium"
  repetition_penalty = 0
  response_format = {
    schema = {
      foo = "bar"
    }
    type = "json"
  }
  safety_model = "safety_model_name"
  seed = 42
  stop = ["string"]
  stream = false
  temperature = 0
  tool_choice = "tool_name"
  tools = [{
    function = {
      description = "A description of the function."
      name = "function_name"
      parameters = {
        foo = "bar"
      }
    }
    type = "tool_type"
  }]
  top_k = 0
  top_p = 0
}
