# Together Terraform Provider

The [Together Terraform provider](https://registry.terraform.io/providers/togethercomputer/together/latest/docs) provides convenient access to
the [Together REST API](https://docs.together.ai/) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    together = {
      source  = "togethercomputer/together"
      version = "~> 0.0.1"
    }
  }
}

# Initialize the provider
provider "together" {
  api_key = "My API Key" # or set TOGETHER_API_KEY env variable
}

# Configure a resource
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
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/togethercomputer/together/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property | Environment variable | Required | Default value |
| -------- | -------------------- | -------- | ------------- |
| api_key  | `TOGETHER_API_KEY`   | true     | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/togethercomputer/terraform-provider-together/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
