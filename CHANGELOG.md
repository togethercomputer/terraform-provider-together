# Changelog

## 0.1.0 (2026-05-22)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/togethercomputer/terraform-provider-together/compare/v0.0.1...v0.1.0)

### ⚠ BREAKING CHANGES

* **api:** Change call signature for `audio.create` to `audio.speech.create` to match spec with python library and add space for future APIs
* **api:** Update method signature for reranking to `rerank.create()`
* **api:** For the TS SDK the `images.create` is now `images.generate`
* **api:** Access to the api for listing checkpoints has changed its name to `list_checkpoints`
* **api:** Access to fine tuning APIs namespace has changed from `fine_tune` to `fine_tuning`
* **api:** The default max retries for api calls has changed from 5 to 2. This may result in more frequent non-200 responses.

### Features

* add per-resource api permissions to schema description ([3b0b159](https://github.com/togethercomputer/terraform-provider-together/commit/3b0b159525381e57916b2d2bc45c4381d74a7b5f))
* **api:** add auto-scale/scheduling/OIDC/capacity params, cuda/driver versions to beta_cluster ([f7cec17](https://github.com/togethercomputer/terraform-provider-together/commit/f7cec17dac16f6bceb8b5e75d47f0b4c55a5dbe0))
* **api:** Add fine_tune.delete API ([4920029](https://github.com/togethercomputer/terraform-provider-together/commit/4920029c2665a84bb2f562fc32cea65bf84990de))
* **api:** add instance_name field to beta_cluster remediation response ([247e1f8](https://github.com/togethercomputer/terraform-provider-together/commit/247e1f84025c35943e6dae11bfb4636f4d10771c))
* **api:** add oidc/add-ons/remediation/phase tracking/cluster config, update beta_cluster_storage ([5e390a4](https://github.com/togethercomputer/terraform-provider-together/commit/5e390a4cc014bf200e3def29017af31c9683acd0))
* **api:** api update ([92ccee4](https://github.com/togethercomputer/terraform-provider-together/commit/92ccee4eaeaa12b9fed5a94f313993b5b07f46a3))
* **api:** api update ([f78427b](https://github.com/togethercomputer/terraform-provider-together/commit/f78427b7d9182a753ccf4bc10e7c8430e2fcb98b))
* **api:** api update ([395f051](https://github.com/togethercomputer/terraform-provider-together/commit/395f051420b0936a4ece0bf928334104bbfa9a87))
* **api:** api update ([2dea7b5](https://github.com/togethercomputer/terraform-provider-together/commit/2dea7b5621c3ed092d251014de903a27598f861a))
* **api:** api update ([e636395](https://github.com/togethercomputer/terraform-provider-together/commit/e63639515496d57a76a8f26e2a092a5c72624965))
* **api:** Change image creation signature to `images.generate` ([3cdd494](https://github.com/togethercomputer/terraform-provider-together/commit/3cdd494c93dab22ae9ab91daf0110a9ddf75372c))
* **api:** Change rerank method signature ([c72be89](https://github.com/togethercomputer/terraform-provider-together/commit/c72be89f4c718b8c7fda57a0a18a245f200168bb))
* **api:** Change the default max retries from 5 to 2 ([8d93523](https://github.com/togethercomputer/terraform-provider-together/commit/8d935230d108a0d343f2dc223eca074cc58cbb42))
* **api:** Change TTS call signature ([a0fab47](https://github.com/togethercomputer/terraform-provider-together/commit/a0fab47249e0264e4665744ef674d61efd9cf6ac))
* **api:** manual updates ([9c540e0](https://github.com/togethercomputer/terraform-provider-together/commit/9c540e069412bdf3eeccac71e61784efdcaab123))
* **api:** manual updates ([37acadb](https://github.com/togethercomputer/terraform-provider-together/commit/37acadbb3b9f6355f5c3201ac6503dc1049c4131))
* **api:** manual updates ([5551523](https://github.com/togethercomputer/terraform-provider-together/commit/5551523c39bf15a3746845ebe3ad253ae2d1a974))
* **api:** remove node_name field from beta_cluster control_plane_nodes/gpu_worker_nodes ([080c53b](https://github.com/togethercomputer/terraform-provider-together/commit/080c53b26c33538f3ed6b16ee9ce5891356ee4bd))
* **api:** Update Eval APIs ([a99ec0f](https://github.com/togethercomputer/terraform-provider-together/commit/a99ec0f4c75b970b7c590b0e93c67a12de270cfd))


### Bug Fixes

* align RL forward-backward loss ([b6bb427](https://github.com/togethercomputer/terraform-provider-together/commit/b6bb427db6292b6b6b59b1c086d5e9b198d52ca9))
* **ci:** in custom setup-go, pass through go-version and cache-dependency-path ([004f531](https://github.com/togethercomputer/terraform-provider-together/commit/004f531d59d72478b6a27c980eaf188a034620b0))
* configurability of some attributes ([9efe090](https://github.com/togethercomputer/terraform-provider-together/commit/9efe09005a3a1017ebce3e44c0b8db1f569f90d8))
* correctly mark a subset of fields shared between create and update calls as required ([55f8130](https://github.com/togethercomputer/terraform-provider-together/commit/55f8130de9e47aaaaaae55342d30cb42946bf3cc))
* ensure derived request attribute schemas conform to the upstream configurability overrides ([0e981a6](https://github.com/togethercomputer/terraform-provider-together/commit/0e981a6bd4fa6335c1bec615d538dcf3597fe23d))
* ensure dynamic values always yield valid container inner values ([a184800](https://github.com/togethercomputer/terraform-provider-together/commit/a1848004449ce45aea2b1f35d943adcdd931f4cc))
* fall back to main branch if linking fails in CI ([aa2130d](https://github.com/togethercomputer/terraform-provider-together/commit/aa2130d50be7419cbbf28b9d1606646e65f2136e))
* fix for failing to drop invalid module replace in link script ([658720b](https://github.com/togethercomputer/terraform-provider-together/commit/658720bc96603371dc669a9ded18f81c1d0bf5fd))
* fix quoting typo ([c7c19bb](https://github.com/togethercomputer/terraform-provider-together/commit/c7c19bb0c94f6cbda31472dd68be07a6bd0e6bb2))
* improve linking behavior when developing on a branch not in the Go SDK ([58ebc99](https://github.com/togethercomputer/terraform-provider-together/commit/58ebc9913bbb89faaca798a8dd4877be6f6360b1))
* improved workflow for developing on branches ([8bb6eda](https://github.com/togethercomputer/terraform-provider-together/commit/8bb6edad8e21fc1be9a0c65725fbeedeb118206d))
* list style data sources should always have id value populated ([293e1a9](https://github.com/togethercomputer/terraform-provider-together/commit/293e1a9714909a6e49f44215e4f75c4350ae0799))
* no longer require an API key when building on production repos ([e60dd34](https://github.com/togethercomputer/terraform-provider-together/commit/e60dd34670d9ad0f16c940cc511e5224fa6f6943))
* patch style requests should never send empty json body for objects ([beb5c72](https://github.com/togethercomputer/terraform-provider-together/commit/beb5c727190756e9db4b57ce1f53f1571aec493a))
* spurious update plans for float attributes after import ([b1ddc5b](https://github.com/togethercomputer/terraform-provider-together/commit/b1ddc5ba0f2227559b67c53b16ce0f067658e331))
* **tests:** update hc-install to fix PGP key mismatch ([4a70745](https://github.com/togethercomputer/terraform-provider-together/commit/4a70745965635b2fe7f258f203459d185dfd65f1))
* **types:** add enum validators to status field in beta_cluster_storage ([1735faf](https://github.com/togethercomputer/terraform-provider-together/commit/1735faff9a5213c6cd334fae41b9c0312a2a11b9))


### Chores

* Add Instant Clusters to OpenAPI spec ([efe8f80](https://github.com/togethercomputer/terraform-provider-together/commit/efe8f80320602d7cc7d6b2cb8aa27cfe94bb2e67))
* add local tmpfile directory ([bc6d152](https://github.com/togethercomputer/terraform-provider-together/commit/bc6d152a0757fedb7f819939860e2796fde53587))
* **api:** Cleanup some exported types ([9ca7551](https://github.com/togethercomputer/terraform-provider-together/commit/9ca755139fa946abbff1384a2d16045acf1bc6d0))
* **api:** Globally disable terraform module generation ([3c8455f](https://github.com/togethercomputer/terraform-provider-together/commit/3c8455f8c51eab422cd2af882ddae9c597206b0b))
* **api:** Remove API that is not intended to be public. ([64af004](https://github.com/togethercomputer/terraform-provider-together/commit/64af0044a56c83881e0031c53bc6f7c9734a8a8d))
* bump dependency version ([750f6f3](https://github.com/togethercomputer/terraform-provider-together/commit/750f6f3f6a797fee561cd3be49c38493bd03467c))
* Configure fields to be omitted from terraform cluster create ([a4cd81b](https://github.com/togethercomputer/terraform-provider-together/commit/a4cd81b5da438d1e9a784d19823e2288769ff800))
* configure new SDK language ([3183562](https://github.com/togethercomputer/terraform-provider-together/commit/3183562f6393254dd0c3c3e2efa1f45dcd44d9e7))
* **docs:** update terraform-plugin-docs to v0.24.0 ([728996a](https://github.com/togethercomputer/terraform-provider-together/commit/728996af117c1d123281c9b647c1ef09df8455b7))
* ensure tests build as part of lint step ([1d73806](https://github.com/togethercomputer/terraform-provider-together/commit/1d738065d2d960889eb7d735a011496d45a19b1a))
* Fix various docstrings ([7a56885](https://github.com/togethercomputer/terraform-provider-together/commit/7a56885e44ffae8b0c86e8c515551576d7505c0a))
* Force cluster creation billing type to be ON_DEMAND ([#5](https://github.com/togethercomputer/terraform-provider-together/issues/5)) ([648e7f7](https://github.com/togethercomputer/terraform-provider-together/commit/648e7f7170c073f931ff112d4ac682fb45898e92))
* **internal:** address linter warnings ([d319a72](https://github.com/togethercomputer/terraform-provider-together/commit/d319a72e3fc655186a52835d24fb24872ebbc807))
* **internal:** codegen related update ([4e877ee](https://github.com/togethercomputer/terraform-provider-together/commit/4e877eeff6d436b5c3941226a6cb7747e487710c))
* **internal:** codegen related update ([eb6a17b](https://github.com/togethercomputer/terraform-provider-together/commit/eb6a17bca5b181a4e0dc9b36606d74c9d15e4988))
* **internal:** codegen related update ([a09e0bf](https://github.com/togethercomputer/terraform-provider-together/commit/a09e0bff28d61624bb75cfb6b95319bee058ff36))
* **internal:** codegen related update ([2d03d96](https://github.com/togethercomputer/terraform-provider-together/commit/2d03d96a99b6c66196ec0401758c16803edde4b7))
* **internal:** codegen related update ([5b1effd](https://github.com/togethercomputer/terraform-provider-together/commit/5b1effd2ff7fb2a891d46728718f82b9ee61cf74))
* **internal:** codegen related update ([6f55159](https://github.com/togethercomputer/terraform-provider-together/commit/6f55159f92cca777f66322e1d80629a379b8bf33))
* **internal:** codegen related update ([5356b9d](https://github.com/togethercomputer/terraform-provider-together/commit/5356b9d8b932ea0947787419fdcfa063c59d073a))
* **internal:** codegen related update ([826b17b](https://github.com/togethercomputer/terraform-provider-together/commit/826b17bf28486e102f0c80d0a119f6c0ba6f39e1))
* **internal:** codegen related update ([a2a30f8](https://github.com/togethercomputer/terraform-provider-together/commit/a2a30f89a3292b48b46bc57853ba6a85996bb9b8))
* **internal:** codegen related update ([2b2dc02](https://github.com/togethercomputer/terraform-provider-together/commit/2b2dc02aa6f23982124bfedc6ade5e1c993bfa39))
* **internal:** codegen related update ([41a04a5](https://github.com/togethercomputer/terraform-provider-together/commit/41a04a51a97268842c8d60156635726b912eaf0e))
* **internal:** codegen related update ([e3d18e2](https://github.com/togethercomputer/terraform-provider-together/commit/e3d18e2ad0eb4616923f507760a9c33057477e5c))
* **internal:** codegen related update ([c792853](https://github.com/togethercomputer/terraform-provider-together/commit/c792853aff4de447bb62720c38193ca8d40cb454))
* **internal:** codegen related update ([9a068a6](https://github.com/togethercomputer/terraform-provider-together/commit/9a068a6fd64018372209bded87b0d36b5822d9f7))
* **internal:** more robust bootstrap script ([8b47491](https://github.com/togethercomputer/terraform-provider-together/commit/8b47491f1ac3e467939eb6d8bf2e7952a6be1b8a))
* **internal:** regenerate SDK with no functional changes ([c965edd](https://github.com/togethercomputer/terraform-provider-together/commit/c965eddd79b55113c7a65ce3b8a33b174b056e9d))
* **internal:** regenerate SDK with no functional changes ([c7b5c9d](https://github.com/togethercomputer/terraform-provider-together/commit/c7b5c9d1833c6e002f88e3955263bb45874c9a91))
* **internal:** tweak CI branches ([d4c14c8](https://github.com/togethercomputer/terraform-provider-together/commit/d4c14c8700d18595a17c60ff140700eaa9b08f0f))
* **internal:** update `actions/checkout` version ([4229996](https://github.com/togethercomputer/terraform-provider-together/commit/4229996d2af1244d3dcda7b4cbabc295b2185938))
* **internal:** update gitignore ([e401993](https://github.com/togethercomputer/terraform-provider-together/commit/e401993f37d00e6abd5ac5b92a1d799483d7d711))
* **internal:** update multipart form array serialization ([e1dd165](https://github.com/togethercomputer/terraform-provider-together/commit/e1dd165db0692796090cfed51a9c28e054f15faa))
* pin go releaser version ([eb53437](https://github.com/togethercomputer/terraform-provider-together/commit/eb53437f81f94a1d949e91d3803cae68724fafe1))
* **test:** do not count install time for mock server timeout ([1173e88](https://github.com/togethercomputer/terraform-provider-together/commit/1173e88db854a75c7b9b899802da70ad8a676e66))
* **tests:** bump steady to v0.19.4 ([ea679af](https://github.com/togethercomputer/terraform-provider-together/commit/ea679af7303cd480bc12fa8e60f384db449cbabd))
* **tests:** bump steady to v0.19.5 ([7f66fb1](https://github.com/togethercomputer/terraform-provider-together/commit/7f66fb120ed327c2b6520623088f5a10e4535e07))
* **tests:** bump steady to v0.19.6 ([eea0d57](https://github.com/togethercomputer/terraform-provider-together/commit/eea0d57fd2e7c6a6938779ca606d074fc5670180))
* **tests:** bump steady to v0.19.7 ([8e17ea7](https://github.com/togethercomputer/terraform-provider-together/commit/8e17ea73ffb32ac9bdb4a6b5ad8a8e4242e1582f))
* **tests:** bump steady to v0.20.1 ([50c7e70](https://github.com/togethercomputer/terraform-provider-together/commit/50c7e70ac761809c2874abd2cf1833cb46003a30))
* **tests:** bump steady to v0.20.2 ([0acfd39](https://github.com/togethercomputer/terraform-provider-together/commit/0acfd39f95c7c30b7bdf67bbd7ab7bd5c03391a4))
* **tests:** bump steady to v0.22.1 ([66a3311](https://github.com/togethercomputer/terraform-provider-together/commit/66a33116be46bc2b5235348fb1e1ef5bf215bff2))
* Update cluster apis to reflect their new response shape ([0dc8895](https://github.com/togethercomputer/terraform-provider-together/commit/0dc8895c2ee108089ae5cc49a757d61f21beb932))
* Update descriptions for jig queue methods and properties ([7016581](https://github.com/togethercomputer/terraform-provider-together/commit/7016581b9bd929568497387daca961ec686655db))
* update Go SDK version ([7d98f53](https://github.com/togethercomputer/terraform-provider-together/commit/7d98f5333ccb3fde4c976fb72ecaf2e47efeb765))
* Update reference to terraform repo name ([0aad161](https://github.com/togethercomputer/terraform-provider-together/commit/0aad161042b0668b2317a3d3e78962991f57797d))
* Update terraform go sdk version ([60a225f](https://github.com/togethercomputer/terraform-provider-together/commit/60a225feb3538baf38510b6f3cdc9c79d5c8d05b))
* update terraform-plugin-framework to v1.19.0 ([91d9c68](https://github.com/togethercomputer/terraform-provider-together/commit/91d9c687d4e2a6b303d41606c6962aff3fa9d964))


### Documentation

* add field descriptions to beta_cluster and beta_cluster_storage resources ([28eccd7](https://github.com/togethercomputer/terraform-provider-together/commit/28eccd78838f53e495fa10fdf25fecfe11b7e3ae))
* **api:** update reservation_start_time description in beta_cluster ([5ab9ea8](https://github.com/togethercomputer/terraform-provider-together/commit/5ab9ea8cf370041c18cead69d157ffb130ea6460))


### Styles

* **api:** Change fine tuning method `retrieve_checkpoints` to `list_checkpoints` ([3af90b5](https://github.com/togethercomputer/terraform-provider-together/commit/3af90b50a9ebd495935614e4b0f53f064e5cd0de))
* **api:** Change fine tuning namespace to `fine_tuning` ([c0e5fd2](https://github.com/togethercomputer/terraform-provider-together/commit/c0e5fd222142f0fc21861ec41566cf3f930e69d2))


### Refactors

* **tests:** switch from prism to steady ([e2c59b4](https://github.com/togethercomputer/terraform-provider-together/commit/e2c59b46cdea1798cb9476b7c1fdc4cf3227016a))
