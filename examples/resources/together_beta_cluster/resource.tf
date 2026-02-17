resource "together_beta_cluster" "example_beta_cluster" {
  cluster_name = "cluster_name"
  driver_version = "CUDA_12_5_555"
  gpu_type = "H100_SXM"
  num_gpus = 0
  region = "region"
  cluster_type = "KUBERNETES"
  volume_id = "volume_id"
}
