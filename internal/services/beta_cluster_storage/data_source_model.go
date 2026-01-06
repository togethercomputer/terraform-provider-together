// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package beta_cluster_storage

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type BetaClusterStorageDataSourceModel struct {
	VolumeID   types.String `tfsdk:"volume_id" path:"volume_id,required"`
	SizeTib    types.Int64  `tfsdk:"size_tib" json:"size_tib,computed"`
	VolumeName types.String `tfsdk:"volume_name" json:"volume_name,computed"`
}
