resource "togetherai_beta_cluster" "example_beta_cluster" {
  billing_type = "RESERVED"
  cluster_name = "cluster_name"
  driver_version = "CUDA_12_5_555"
  gpu_type = "H100_SXM"
  num_gpus = 0
  region = "us-central-8"
  cluster_type = "KUBERNETES"
  duration_days = 0
  shared_volume = {
    region = "region"
    size_tib = 0
    volume_name = "volume_name"
  }
  volume_id = "volume_id"
}
