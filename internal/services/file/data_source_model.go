// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package file

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FileDataSourceModel struct {
	ID        types.String `tfsdk:"id" path:"id,required"`
	Bytes     types.Int64  `tfsdk:"bytes" json:"bytes,computed"`
	CreatedAt types.Int64  `tfsdk:"created_at" json:"created_at,computed"`
	FileType  types.String `tfsdk:"file_type" json:"FileType,computed"`
	Filename  types.String `tfsdk:"filename" json:"filename,computed"`
	LineCount types.Int64  `tfsdk:"line_count" json:"LineCount,computed"`
	Object    types.String `tfsdk:"object" json:"object,computed"`
	Processed types.Bool   `tfsdk:"processed" json:"Processed,computed"`
	Purpose   types.String `tfsdk:"purpose" json:"purpose,computed"`
}
