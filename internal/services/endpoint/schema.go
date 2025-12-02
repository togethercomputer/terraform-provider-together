// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*EndpointResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier for the endpoint",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"hardware": schema.StringAttribute{
				Description:   "The hardware configuration to use for this endpoint",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"model": schema.StringAttribute{
				Description:   "The model to deploy on this endpoint",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"availability_zone": schema.StringAttribute{
				Description:   "Create the endpoint in a specified availability zone (e.g., us-central-4b)",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"disable_prompt_cache": schema.BoolAttribute{
				Description:   "Whether to disable the prompt cache for this endpoint",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"disable_speculative_decoding": schema.BoolAttribute{
				Description:   "Whether to disable speculative decoding for this endpoint",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplaceIfConfigured()},
				Default:       booldefault.StaticBool(false),
			},
			"autoscaling": schema.SingleNestedAttribute{
				Description: "Configuration for automatic scaling of the endpoint",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"max_replicas": schema.Int64Attribute{
						Description: "The maximum number of replicas to scale up to under load",
						Required:    true,
					},
					"min_replicas": schema.Int64Attribute{
						Description: "The minimum number of replicas to maintain, even when there is no load",
						Required:    true,
					},
				},
			},
			"display_name": schema.StringAttribute{
				Description: "A human-readable name for the endpoint",
				Optional:    true,
			},
			"inactive_timeout": schema.Int64Attribute{
				Description: "The number of minutes of inactivity after which the endpoint will be automatically stopped. Set to null, omit or set to 0 to disable automatic timeout.",
				Optional:    true,
			},
			"state": schema.StringAttribute{
				Description: "The desired state of the endpoint\nAvailable values: \"STARTED\", \"STOPPED\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("STARTED", "STOPPED"),
				},
				Default: stringdefault.StaticString("STARTED"),
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp when the endpoint was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "System name for the endpoint",
				Computed:    true,
			},
			"object": schema.StringAttribute{
				Description: "The type of object\nAvailable values: \"endpoint\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("endpoint"),
				},
			},
			"owner": schema.StringAttribute{
				Description: "The owner of this endpoint",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of endpoint\nAvailable values: \"dedicated\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("dedicated"),
				},
			},
		},
	}
}

func (r *EndpointResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *EndpointResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
