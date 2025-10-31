// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package embedding

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*EmbeddingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"input": schema.StringAttribute{
				Description:   "A string providing the text for the model to embed.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The name of the embedding model to use.<br> <br> [See all of Together AI's embedding models](https://docs.together.ai/docs/serverless-models#embedding-models)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"object": schema.StringAttribute{
				Description: `Available values: "list".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("list"),
				},
			},
			"data": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[EmbeddingDataModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"embedding": schema.ListAttribute{
							Computed:    true,
							CustomType:  customfield.NewListType[types.Float64](ctx),
							ElementType: types.Float64Type,
						},
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"object": schema.StringAttribute{
							Description: `Available values: "embedding".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("embedding"),
							},
						},
					},
				},
			},
		},
	}
}

func (r *EmbeddingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EmbeddingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
