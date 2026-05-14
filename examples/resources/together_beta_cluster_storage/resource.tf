resource "together_beta_cluster_storage" "example_beta_cluster_storage" {
  region = "region"
  size_tib = 0
  volume_name = "volume_name"
  is_lifecycle_independent = true
}
