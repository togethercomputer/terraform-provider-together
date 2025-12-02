resource "togetherai_audio_transcription" "example_audio_transcription" {
  file = null
  diarize = true
  language = "en"
  max_speakers = 0
  min_speakers = 0
  model = "openai/whisper-large-v3"
  prompt = "prompt"
  response_format = "json"
  temperature = 0
  timestamp_granularities = ["word", "segment"]
}
