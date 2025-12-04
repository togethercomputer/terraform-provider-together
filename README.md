# Together Terraform Provider

The [Together Terraform provider](https://registry.terraform.io/providers/stainless-sdks/togetherai/latest/docs) provides convenient access to
the [Together REST API](https://docs.together.ai/) from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

```hcl
# Declare the provider and version
terraform {
  required_providers {
    togetherai = {
      source  = "stainless-sdks/togetherai"
      version = "~> 0.0.1"
    }
  }
}

# Initialize the provider
provider "togetherai" {
  api_key = "My API Key" # or set TOGETHER_API_KEY env variable
}

# Configure a resource
resource "togetherai_chat_completion" "example_chat_completion" {
  messages = [{
    content = "Say this is a test"
    role = "user"
    name = "name"
  }]
  model = "meta-llama/Meta-Llama-3.1-8B-Instruct-Turbo"
  context_length_exceeded_behavior = "truncate"
  echo = true
  frequency_penalty = 0
  function_call = "none"
  logit_bias = {
    "105" = 21.4
    "1024" = -10.5
  }
  logprobs = 0
  max_tokens = 0
  min_p = 0
  n = 1
  presence_penalty = 0
  reasoning_effort = "medium"
  repetition_penalty = 0
  response_format = {
    type = "text"
  }
  safety_model = "safety_model_name"
  seed = 42
  stop = ["string"]
  stream = false
  temperature = 0
  tool_choice = "tool_name"
  tools = [{
    function = {
      description = "A description of the function."
      name = "function_name"
      parameters = {
        foo = "bar"
      }
    }
    type = "tool_type"
  }]
  top_k = 0
  top_p = 0
}
```

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/stainless-sdks/togetherai/latest/docs).

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

We are keen for your feedback; please open an [issue](https://www.github.com/stainless-sdks/togetherai-terraform/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
