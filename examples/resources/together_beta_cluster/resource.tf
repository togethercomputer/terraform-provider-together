resource "together_beta_cluster" "example_beta_cluster" {
  cluster_name = "cluster_name"
  cuda_version = "cuda_version"
  gpu_type = "H100_SXM"
  num_gpus = 0
  nvidia_driver_version = "nvidia_driver_version"
  region = "region"
  acceptance_tests_params = {
    dcgm_diag_level = "DCGM_DIAG_LEVEL_SHORT"
    dcgm_diag_skipped = true
    enabled = true
    gpu_burn_duration = 0
    gpu_burn_skipped = true
    nccl_multi_node_skipped = true
    nccl_single_node_skipped = true
  }
  add_ons = [{
    add_on_type = "add_on_type"
    name = "name"
    config = {
      dashboard = {
        enabled = true
      }
      ingress = {
        enabled = true
      }
    }
  }]
  auto_scale = true
  auto_scale_max_gpus = 0
  auto_scaled = true
  capacity_pool_id = "capacity_pool_id"
  cluster_config = {
    load_balancer = "NONE"
    gpu_operator_version = "gpu_operator_version"
    ingress = {
      enabled = true
    }
    jumphost_enabled = true
    kubernetes_dashboard_enabled = true
    observability = {
      enabled = true
    }
    slurm_startup_scripts = {
      controller_epilog = "controller_epilog"
      controller_prolog = "controller_prolog"
      extra_slurm_conf = "extra_slurm_conf"
      login_init_script = "login_init_script"
      nodeset_init_script = "nodeset_init_script"
      worker_epilog = "worker_epilog"
      worker_prolog = "worker_prolog"
    }
  }
  cluster_type = "KUBERNETES"
  gpu_node_failover_enabled = true
  install_traefik = true
  num_capacity_pool_gpus = 0
  num_preemptible_gpus = 0
  num_reserved_gpus = 0
  oidc_config = {
    client_id = "client_id"
    group_claim = "group_claim"
    group_prefix = "group_prefix"
    issuer_url = "issuer_url"
    username_claim = "username_claim"
    username_prefix = "username_prefix"
    ca_cert = "ca_cert"
  }
  project_id = "project_id"
  reservation_end_time = "2019-12-27T18:11:19.117Z"
  reservation_start_time = "2019-12-27T18:11:19.117Z"
  slurm_image = "slurm_image"
  slurm_shm_size_gib = 0
  volume_id = "volume_id"
}
