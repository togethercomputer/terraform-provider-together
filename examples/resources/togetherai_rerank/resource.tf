resource "togetherai_rerank" "example_rerank" {
  documents = [{
    title = "bar"
    text = "bar"
  }, {
    title = "bar"
    text = "bar"
  }, {
    title = "bar"
    text = "bar"
  }, {
    title = "bar"
    text = "bar"
  }]
  model = "Salesforce/Llama-Rank-V1"
  query = "What animals can I find near Peru?"
  rank_fields = ["title", "text"]
  return_documents = true
  top_n = 2
}
