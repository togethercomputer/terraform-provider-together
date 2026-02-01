# Changelog

## 0.1.0 (2026-02-01)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/togethercomputer/terraform-provider-togetherai/compare/v0.0.1...v0.1.0)

### ⚠ BREAKING CHANGES

* **api:** Change call signature for `audio.create` to `audio.speech.create` to match spec with python library and add space for future APIs
* **api:** Update method signature for reranking to `rerank.create()`
* **api:** For the TS SDK the `images.create` is now `images.generate`
* **api:** Access to the api for listing checkpoints has changed its name to `list_checkpoints`
* **api:** Access to fine tuning APIs namespace has changed from `fine_tune` to `fine_tuning`
* **api:** The default max retries for api calls has changed from 5 to 2. This may result in more frequent non-200 responses.

### Features

* **api:** Add fine_tune.delete API ([4920029](https://github.com/togethercomputer/terraform-provider-togetherai/commit/4920029c2665a84bb2f562fc32cea65bf84990de))
* **api:** api update ([92ccee4](https://github.com/togethercomputer/terraform-provider-togetherai/commit/92ccee4eaeaa12b9fed5a94f313993b5b07f46a3))
* **api:** api update ([f78427b](https://github.com/togethercomputer/terraform-provider-togetherai/commit/f78427b7d9182a753ccf4bc10e7c8430e2fcb98b))
* **api:** api update ([395f051](https://github.com/togethercomputer/terraform-provider-togetherai/commit/395f051420b0936a4ece0bf928334104bbfa9a87))
* **api:** api update ([2dea7b5](https://github.com/togethercomputer/terraform-provider-togetherai/commit/2dea7b5621c3ed092d251014de903a27598f861a))
* **api:** api update ([e636395](https://github.com/togethercomputer/terraform-provider-togetherai/commit/e63639515496d57a76a8f26e2a092a5c72624965))
* **api:** Change image creation signature to `images.generate` ([3cdd494](https://github.com/togethercomputer/terraform-provider-togetherai/commit/3cdd494c93dab22ae9ab91daf0110a9ddf75372c))
* **api:** Change rerank method signature ([c72be89](https://github.com/togethercomputer/terraform-provider-togetherai/commit/c72be89f4c718b8c7fda57a0a18a245f200168bb))
* **api:** Change the default max retries from 5 to 2 ([8d93523](https://github.com/togethercomputer/terraform-provider-togetherai/commit/8d935230d108a0d343f2dc223eca074cc58cbb42))
* **api:** Change TTS call signature ([a0fab47](https://github.com/togethercomputer/terraform-provider-togetherai/commit/a0fab47249e0264e4665744ef674d61efd9cf6ac))
* **api:** manual updates ([9c540e0](https://github.com/togethercomputer/terraform-provider-togetherai/commit/9c540e069412bdf3eeccac71e61784efdcaab123))
* **api:** manual updates ([37acadb](https://github.com/togethercomputer/terraform-provider-togetherai/commit/37acadbb3b9f6355f5c3201ac6503dc1049c4131))
* **api:** manual updates ([5551523](https://github.com/togethercomputer/terraform-provider-togetherai/commit/5551523c39bf15a3746845ebe3ad253ae2d1a974))
* **api:** Update Eval APIs ([a99ec0f](https://github.com/togethercomputer/terraform-provider-togetherai/commit/a99ec0f4c75b970b7c590b0e93c67a12de270cfd))


### Bug Fixes

* configurability of some attributes ([9efe090](https://github.com/togethercomputer/terraform-provider-togetherai/commit/9efe09005a3a1017ebce3e44c0b8db1f569f90d8))
* correctly mark a subset of fields shared between create and update calls as required ([55f8130](https://github.com/togethercomputer/terraform-provider-togetherai/commit/55f8130de9e47aaaaaae55342d30cb42946bf3cc))
* ensure derived request attribute schemas conform to the upstream configurability overrides ([0e981a6](https://github.com/togethercomputer/terraform-provider-togetherai/commit/0e981a6bd4fa6335c1bec615d538dcf3597fe23d))
* ensure dynamic values always yield valid container inner values ([a184800](https://github.com/togethercomputer/terraform-provider-togetherai/commit/a1848004449ce45aea2b1f35d943adcdd931f4cc))
* list style data sources should always have id value populated ([293e1a9](https://github.com/togethercomputer/terraform-provider-togetherai/commit/293e1a9714909a6e49f44215e4f75c4350ae0799))


### Chores

* Add Instant Clusters to OpenAPI spec ([efe8f80](https://github.com/togethercomputer/terraform-provider-togetherai/commit/efe8f80320602d7cc7d6b2cb8aa27cfe94bb2e67))
* **api:** Cleanup some exported types ([9ca7551](https://github.com/togethercomputer/terraform-provider-togetherai/commit/9ca755139fa946abbff1384a2d16045acf1bc6d0))
* **api:** Globally disable terraform module generation ([3c8455f](https://github.com/togethercomputer/terraform-provider-togetherai/commit/3c8455f8c51eab422cd2af882ddae9c597206b0b))
* **api:** Remove API that is not intended to be public. ([64af004](https://github.com/togethercomputer/terraform-provider-togetherai/commit/64af0044a56c83881e0031c53bc6f7c9734a8a8d))
* bump dependency version ([750f6f3](https://github.com/togethercomputer/terraform-provider-togetherai/commit/750f6f3f6a797fee561cd3be49c38493bd03467c))
* configure new SDK language ([3183562](https://github.com/togethercomputer/terraform-provider-togetherai/commit/3183562f6393254dd0c3c3e2efa1f45dcd44d9e7))
* ensure tests build as part of lint step ([1d73806](https://github.com/togethercomputer/terraform-provider-togetherai/commit/1d738065d2d960889eb7d735a011496d45a19b1a))
* **internal:** address linter warnings ([d319a72](https://github.com/togethercomputer/terraform-provider-togetherai/commit/d319a72e3fc655186a52835d24fb24872ebbc807))
* **internal:** codegen related update ([5b1effd](https://github.com/togethercomputer/terraform-provider-togetherai/commit/5b1effd2ff7fb2a891d46728718f82b9ee61cf74))
* **internal:** codegen related update ([6f55159](https://github.com/togethercomputer/terraform-provider-togetherai/commit/6f55159f92cca777f66322e1d80629a379b8bf33))
* **internal:** codegen related update ([5356b9d](https://github.com/togethercomputer/terraform-provider-togetherai/commit/5356b9d8b932ea0947787419fdcfa063c59d073a))
* **internal:** codegen related update ([826b17b](https://github.com/togethercomputer/terraform-provider-togetherai/commit/826b17bf28486e102f0c80d0a119f6c0ba6f39e1))
* **internal:** codegen related update ([a2a30f8](https://github.com/togethercomputer/terraform-provider-togetherai/commit/a2a30f89a3292b48b46bc57853ba6a85996bb9b8))
* **internal:** codegen related update ([2b2dc02](https://github.com/togethercomputer/terraform-provider-togetherai/commit/2b2dc02aa6f23982124bfedc6ade5e1c993bfa39))
* **internal:** codegen related update ([41a04a5](https://github.com/togethercomputer/terraform-provider-togetherai/commit/41a04a51a97268842c8d60156635726b912eaf0e))
* **internal:** codegen related update ([e3d18e2](https://github.com/togethercomputer/terraform-provider-togetherai/commit/e3d18e2ad0eb4616923f507760a9c33057477e5c))
* **internal:** codegen related update ([c792853](https://github.com/togethercomputer/terraform-provider-togetherai/commit/c792853aff4de447bb62720c38193ca8d40cb454))
* **internal:** codegen related update ([9a068a6](https://github.com/togethercomputer/terraform-provider-togetherai/commit/9a068a6fd64018372209bded87b0d36b5822d9f7))
* **internal:** update `actions/checkout` version ([4229996](https://github.com/togethercomputer/terraform-provider-togetherai/commit/4229996d2af1244d3dcda7b4cbabc295b2185938))
* Update cluster apis to reflect their new response shape ([0dc8895](https://github.com/togethercomputer/terraform-provider-togetherai/commit/0dc8895c2ee108089ae5cc49a757d61f21beb932))
* update Go SDK version ([7d98f53](https://github.com/togethercomputer/terraform-provider-togetherai/commit/7d98f5333ccb3fde4c976fb72ecaf2e47efeb765))
* Update terraform go sdk version ([60a225f](https://github.com/togethercomputer/terraform-provider-togetherai/commit/60a225feb3538baf38510b6f3cdc9c79d5c8d05b))


### Styles

* **api:** Change fine tuning method `retrieve_checkpoints` to `list_checkpoints` ([3af90b5](https://github.com/togethercomputer/terraform-provider-togetherai/commit/3af90b50a9ebd495935614e4b0f53f064e5cd0de))
* **api:** Change fine tuning namespace to `fine_tuning` ([c0e5fd2](https://github.com/togethercomputer/terraform-provider-togetherai/commit/c0e5fd222142f0fc21861ec41566cf3f930e69d2))
