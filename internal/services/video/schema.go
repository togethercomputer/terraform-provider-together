// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package video

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
	"github.com/stainless-sdks/togetherai-terraform/internal/customvalidator"
)

var _ resource.ResourceWithConfigValidators = (*VideoResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the video job.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The model to be used for the video creation request.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"fps": schema.Int64Attribute{
				Description:   "Frames per second. Defaults to 24.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"guidance_scale": schema.Int64Attribute{
				Description:   "Controls how closely the video generation follows your prompt. Higher values make the model adhere more strictly to your text description, while lower values allow more creative freedom. guidence_scale affects both visual content and temporal consistency.Recommended range is 6.0-10.0 for most video models. Values above 12 may cause over-guidance artifacts or unnatural motion patterns.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"height": schema.Int64Attribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"negative_prompt": schema.StringAttribute{
				Description:   "Similar to prompt, but specifies what to avoid instead of what to include",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"output_format": schema.StringAttribute{
				Description: "Specifies the format of the output video. Defaults to MP4.\nAvailable values: \"MP4\", \"WEBM\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("MP4", "WEBM"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"output_quality": schema.Int64Attribute{
				Description:   "Compression quality. Defaults to 20.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"prompt": schema.StringAttribute{
				Description:   "Text prompt that describes the video to generate.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"seconds": schema.StringAttribute{
				Description:   "Clip duration in seconds.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"seed": schema.Int64Attribute{
				Description:   "Seed to use in initializing the video generation.  Using the same seed allows deterministic video generation.  If not provided a random seed is generated for each request.",
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"steps": schema.Int64Attribute{
				Description: "The number of denoising steps the model performs during video generation. More steps typically result in higher quality output but require longer processing time.",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.Between(10, 50),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"width": schema.Int64Attribute{
				Optional:      true,
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"reference_images": schema.ListAttribute{
				Description:   "Unlike frame_images which constrain specific timeline positions, reference images guide the general appearance that should appear consistently across the video.",
				Optional:      true,
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"frame_images": schema.ListNestedAttribute{
				Description: "Array of images to guide video generation, similar to keyframes.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"input_image": schema.StringAttribute{
							Description: "URL path to hosted image that is used for a frame",
							Required:    true,
						},
						"frame": schema.DynamicAttribute{
							Description: "Optional param to specify where to insert the frame. If this is omitted, the following heuristics are applied:\n- frame_images size is one, frame is first.\n- If size is two, frames are first and last.\n- If size is larger, frames are first, last and evenly spaced between.",
							Optional:    true,
							Validators: []validator.Dynamic{
								customvalidator.AllowedSubtypes(basetypes.Float64Type{}, basetypes.StringType{}),
							},
							CustomType:    customfield.NormalizedDynamicType{},
							PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"completed_at": schema.Float64Attribute{
				Description: "Unix timestamp (seconds) for when the job completed, if finished.",
				Computed:    true,
			},
			"created_at": schema.Float64Attribute{
				Description: "Unix timestamp (seconds) for when the job was created.",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Description: "The object type, which is always video.\nAvailable values: \"video\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("video"),
				},
			},
			"size": schema.StringAttribute{
				Description: "The resolution of the generated video.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current lifecycle status of the video job.\nAvailable values: \"in_progress\", \"completed\", \"failed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"in_progress",
						"completed",
						"failed",
					),
				},
			},
			"error": schema.SingleNestedAttribute{
				Description: "Error payload that explains why generation failed, if applicable.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VideoErrorModel](ctx),
				Attributes: map[string]schema.Attribute{
					"message": schema.StringAttribute{
						Computed: true,
					},
					"code": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"outputs": schema.SingleNestedAttribute{
				Description: "Available upon completion, the outputs provides the cost charged and the hosted url to access the video",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[VideoOutputsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cost": schema.Int64Attribute{
						Description: "The cost of generated video charged to the owners account.",
						Computed:    true,
					},
					"video_url": schema.StringAttribute{
						Description: "URL hosting the generated video",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *VideoResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *VideoResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
