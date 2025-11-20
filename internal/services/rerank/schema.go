// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rerank

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*RerankResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Request ID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The model to be used for the rerank request.<br> <br> [See all of Together AI's rerank models](https://docs.together.ai/docs/serverless-models#rerank-models)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"query": schema.StringAttribute{
				Description:   "The search query to be used for ranking.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"documents": schema.ListAttribute{
				Description: "List of documents, which can be either strings or objects.",
				Required:    true,
				ElementType: types.MapType{
					ElemType: jsontypes.NormalizedType{},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"return_documents": schema.BoolAttribute{
				Description:   "Whether to return supplied documents with the response.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"top_n": schema.Int64Attribute{
				Description:   "The number of top results to return.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"rank_fields": schema.ListAttribute{
				Description:   "List of keys in the JSON Object document to rank by. Defaults to use all supplied keys for ranking.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"object": schema.StringAttribute{
				Description: "Object type\nAvailable values: \"rerank\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("rerank"),
				},
			},
			"results": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[RerankResultsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"document": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[RerankResultsDocumentModel](ctx),
							Attributes: map[string]schema.Attribute{
								"text": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"relevance_score": schema.Float64Attribute{
							Computed: true,
						},
					},
				},
			},
			"usage": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[RerankUsageModel](ctx),
				Attributes: map[string]schema.Attribute{
					"completion_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"prompt_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"total_tokens": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (r *RerankResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *RerankResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
