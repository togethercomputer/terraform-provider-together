// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_storage

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/sjson"
	"github.com/togethercomputer/terraform-provider-togetherai/internal/apijson"
)

type BetaClusterStorageModel struct {
	ID         types.String `tfsdk:"id" json:"-,computed"`
	VolumeID   types.String `tfsdk:"volume_id" json:"volume_id,computed"`
	Region     types.String `tfsdk:"region" json:"region,required,no_refresh"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,required"`
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,required"`
	Status     types.String `tfsdk:"status" json:"status,computed"`
}

func (m BetaClusterStorageModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m BetaClusterStorageModel) MarshalJSONForUpdate(state BetaClusterStorageModel) (data []byte, err error) {
	data, err = apijson.MarshalForUpdate(m, state)
	if err != nil {
		return nil, err
	}
	return sjson.SetBytes(data, "volume_id", m.VolumeID.ValueString())
}
