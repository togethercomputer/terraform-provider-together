resource "togetherai_eval" "example_eval" {
  parameters = {
    input_data_file_path = "file-1234-aefd"
    judge = {
      model_name = "meta-llama/Llama-3-70B-Instruct-Turbo"
      system_template = "Imagine you are a helpful assistant"
    }
    labels = ["yes", "no"]
    pass_labels = ["yes"]
    model_to_evaluate = "string"
  }
  type = "classify"
}
