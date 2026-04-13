resource "together_beta_cluster" "example_beta_cluster" {
  cluster_name = "cluster_name"
  cuda_version = "cuda_version"
  gpu_type = "H100_SXM"
  num_gpus = 0
  nvidia_driver_version = "nvidia_driver_version"
  region = "region"
  auto_scale_max_gpus = 0
  auto_scaled = true
  capacity_pool_id = "capacity_pool_id"
  cluster_type = "KUBERNETES"
  gpu_node_failover_enabled = true
  install_traefik = true
  reservation_end_time = "2019-12-27T18:11:19.117Z"
  reservation_start_time = "2019-12-27T18:11:19.117Z"
  slurm_image = "slurm_image"
  slurm_shm_size_gib = 0
  volume_id = "volume_id"
}
