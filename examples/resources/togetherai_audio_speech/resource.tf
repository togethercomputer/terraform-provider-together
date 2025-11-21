resource "togetherai_audio_speech" "example_audio_speech" {
  input = "input"
  model = "canopylabs/orpheus-3b-0.1-ft"
  voice = "voice"
  language = "en"
  response_encoding = "pcm_f32le"
  response_format = "mp3"
  sample_rate = 0
  stream = false
}
