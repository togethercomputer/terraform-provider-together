// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/apijson"
)

type EndpointModel struct {
	ID                         types.String              `tfsdk:"id" json:"id,computed"`
	Hardware                   types.String              `tfsdk:"hardware" json:"hardware,required"`
	Model                      types.String              `tfsdk:"model" json:"model,required"`
	AvailabilityZone           types.String              `tfsdk:"availability_zone" json:"availability_zone,optional,no_refresh"`
	DisablePromptCache         types.Bool                `tfsdk:"disable_prompt_cache" json:"disable_prompt_cache,computed_optional,no_refresh"`
	DisableSpeculativeDecoding types.Bool                `tfsdk:"disable_speculative_decoding" json:"disable_speculative_decoding,computed_optional,no_refresh"`
	Autoscaling                *EndpointAutoscalingModel `tfsdk:"autoscaling" json:"autoscaling,required"`
	DisplayName                types.String              `tfsdk:"display_name" json:"display_name,optional"`
	InactiveTimeout            types.Int64               `tfsdk:"inactive_timeout" json:"inactive_timeout,optional,no_refresh"`
	State                      types.String              `tfsdk:"state" json:"state,computed_optional"`
	CreatedAt                  timetypes.RFC3339         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Name                       types.String              `tfsdk:"name" json:"name,computed"`
	Object                     types.String              `tfsdk:"object" json:"object,computed"`
	Owner                      types.String              `tfsdk:"owner" json:"owner,computed"`
	Type                       types.String              `tfsdk:"type" json:"type,computed"`
}

func (m EndpointModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EndpointModel) MarshalJSONForUpdate(state EndpointModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type EndpointAutoscalingModel struct {
	MaxReplicas types.Int64 `tfsdk:"max_replicas" json:"max_replicas,required"`
	MinReplicas types.Int64 `tfsdk:"min_replicas" json:"min_replicas,required"`
}
