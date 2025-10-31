resource "togetherai_audio_translation" "example_audio_translation" {
  file = null
  language = "en"
  model = "openai/whisper-large-v3"
  prompt = "prompt"
  response_format = "json"
  temperature = 0
  timestamp_granularities = ["word", "segment"]
}
