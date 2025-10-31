// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/togetherai-terraform/internal/customfield"
)

type EndpointDataSourceModel struct {
	ID          types.String                                                 `tfsdk:"id" path:"endpointId,computed"`
	EndpointID  types.String                                                 `tfsdk:"endpoint_id" path:"endpointId,required"`
	CreatedAt   timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DisplayName types.String                                                 `tfsdk:"display_name" json:"display_name,computed"`
	Hardware    types.String                                                 `tfsdk:"hardware" json:"hardware,computed"`
	Model       types.String                                                 `tfsdk:"model" json:"model,computed"`
	Name        types.String                                                 `tfsdk:"name" json:"name,computed"`
	Object      types.String                                                 `tfsdk:"object" json:"object,computed"`
	Owner       types.String                                                 `tfsdk:"owner" json:"owner,computed"`
	State       types.String                                                 `tfsdk:"state" json:"state,computed"`
	Type        types.String                                                 `tfsdk:"type" json:"type,computed"`
	Autoscaling customfield.NestedObject[EndpointAutoscalingDataSourceModel] `tfsdk:"autoscaling" json:"autoscaling,computed"`
}

type EndpointAutoscalingDataSourceModel struct {
	MaxReplicas types.Int64 `tfsdk:"max_replicas" json:"max_replicas,computed"`
	MinReplicas types.Int64 `tfsdk:"min_replicas" json:"min_replicas,computed"`
}
