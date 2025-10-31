resource "togetherai_video" "example_video" {
  model = "model"
  fps = 0
  frame_images = [{
    input_image = "input_image"
    frame = 0
  }]
  guidance_scale = 0
  height = 0
  negative_prompt = "negative_prompt"
  output_format = "MP4"
  output_quality = 0
  prompt = "x"
  reference_images = ["string"]
  seconds = "seconds"
  seed = 0
  steps = 10
  width = 0
}
