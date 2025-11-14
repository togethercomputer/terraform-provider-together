// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio_transcription"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/audio_translation"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/batch"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/chat_completion"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/completion"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/embedding"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/endpoint"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/eval"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/file"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/fine_tuning"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/image"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/job"
	"github.com/stainless-sdks/togetherai-terraform/internal/services/video"
	"github.com/togethercomputer/together-go"
	"github.com/togethercomputer/together-go/option"
)

var _ provider.ProviderWithConfigValidators = (*TogetheraiProvider)(nil)

// TogetheraiProvider defines the provider implementation.
type TogetheraiProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// TogetheraiProviderModel describes the provider data model.
type TogetheraiProviderModel struct {
	BaseURL types.String `tfsdk:"base_url" json:"base_url,optional"`
	APIKey  types.String `tfsdk:"api_key" json:"api_key,optional"`
}

func (p *TogetheraiProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "togetherai"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *TogetheraiProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *TogetheraiProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data TogetheraiProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() && !data.BaseURL.IsUnknown() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	} else if o, ok := os.LookupEnv("TOGETHER_BASE_URL"); ok {
		opts = append(opts, option.WithBaseURL(o))
	}

	if !data.APIKey.IsNull() && !data.APIKey.IsUnknown() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	} else if o, ok := os.LookupEnv("TOGETHER_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	} else {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing api_key value",
			"The api_key field is required. Set it in provider configuration or via the \"TOGETHER_API_KEY\" environment variable.",
		)
		return
	}

	client := together.NewClient(
		opts...,
	)

	resp.DataSourceData = &client
	resp.ResourceData = &client
}

func (p *TogetheraiProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *TogetheraiProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		chat_completion.NewResource,
		completion.NewResource,
		embedding.NewResource,
		fine_tuning.NewResource,
		image.NewResource,
		video.NewResource,
		audio.NewResource,
		audio_transcription.NewResource,
		audio_translation.NewResource,
		endpoint.NewResource,
		batch.NewResource,
		eval.NewResource,
	}
}

func (p *TogetheraiProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		file.NewFileDataSource,
		fine_tuning.NewFineTuningDataSource,
		video.NewVideoDataSource,
		job.NewJobDataSource,
		endpoint.NewEndpointDataSource,
		batch.NewBatchDataSource,
		eval.NewEvalDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TogetheraiProvider{
			version: version,
		}
	}
}
