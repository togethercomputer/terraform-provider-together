resource "togetherai_image" "example_image" {
  model = "black-forest-labs/FLUX.1-schnell"
  prompt = "cat floating in space, cinematic"
  disable_safety_checker = true
  guidance_scale = 0
  height = 0
  image_loras = [{
    path = "path"
    scale = 0
  }]
  image_url = "image_url"
  n = 0
  negative_prompt = "negative_prompt"
  output_format = "jpeg"
  response_format = "base64"
  seed = 0
  steps = 0
  width = 0
}
