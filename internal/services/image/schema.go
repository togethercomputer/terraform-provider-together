// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package image

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

var _ resource.ResourceWithConfigValidators = (*ImageResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The model to use for image generation.<br> <br> [See all of Together AI's image models](https://docs.together.ai/docs/serverless-models#image-models)",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"prompt": schema.StringAttribute{
				Description:   "A description of the desired images. Maximum length varies by model.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"disable_safety_checker": schema.BoolAttribute{
				Description:   "If true, disables the safety checker for image generation.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"image_url": schema.StringAttribute{
				Description:   "URL of an image to use for image models that support it.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"negative_prompt": schema.StringAttribute{
				Description:   "The prompt or prompts not to guide the image generation.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"response_format": schema.StringAttribute{
				Description: "Format of the image response. Can be either a base64 string or a URL.\nAvailable values: \"base64\", \"url\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("base64", "url"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"seed": schema.Int64Attribute{
				Description:   "Seed used for generation. Can be used to reproduce image generations.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"image_loras": schema.ListNestedAttribute{
				Description: "An array of objects that define LoRAs (Low-Rank Adaptations) to influence the generated image.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"path": schema.StringAttribute{
							Description: "The URL of the LoRA to apply (e.g. https://huggingface.co/strangerzonehf/Flux-Midjourney-Mix2-LoRA).",
							Required:    true,
						},
						"scale": schema.Float64Attribute{
							Description: "The strength of the LoRA's influence. Most LoRA's recommend a value of 1.",
							Required:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"guidance_scale": schema.Float64Attribute{
				Description:   "Adjusts the alignment of the generated image with the input prompt. Higher values (e.g., 8-10) make the output more faithful to the prompt, while lower values (e.g., 1-5) encourage more creative freedom.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplaceIfConfigured()},
				Default:       float64default.StaticFloat64(3.5),
			},
			"height": schema.Int64Attribute{
				Description:   "Height of the image to generate in number of pixels.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(1024),
			},
			"n": schema.Int64Attribute{
				Description:   "Number of image results to generate.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(1),
			},
			"output_format": schema.StringAttribute{
				Description: "The format of the image response. Can be either be `jpeg` or `png`. Defaults to `jpeg`.\nAvailable values: \"jpeg\", \"png\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("jpeg", "png"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("jpeg"),
			},
			"steps": schema.Int64Attribute{
				Description:   "Number of generation steps.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(20),
			},
			"width": schema.Int64Attribute{
				Description:   "Width of the image to generate in number of pixels.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplaceIfConfigured()},
				Default:       int64default.StaticInt64(1024),
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
				CustomType: customfield.NewNestedObjectListType[ImageDataModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"b64_json": schema.StringAttribute{
							Computed: true,
						},
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Description: `Available values: "b64_json", "url".`,
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("b64_json", "url"),
							},
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *ImageResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ImageResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
