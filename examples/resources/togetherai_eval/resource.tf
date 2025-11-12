resource "togetherai_eval" "example_eval" {
  parameters = {
    input_data_file_path = "file-1234-aefd"
    judge = {
      model = "meta-llama/Llama-3-70B-Instruct-Turbo"
      model_source = "serverless"
      system_template = "Imagine you are a helpful assistant"
      external_api_token = "external_api_token"
      external_base_url = "external_base_url"
    }
    labels = ["yes", "no"]
    pass_labels = ["yes"]
    model_to_evaluate = "string"
  }
  type = "classify"
}
