// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package file

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*FileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"bytes": schema.Int64Attribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed: true,
			},
			"file_type": schema.StringAttribute{
				Description: "The type of the file\nAvailable values: \"csv\", \"jsonl\", \"parquet\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"csv",
						"jsonl",
						"parquet",
					),
				},
			},
			"filename": schema.StringAttribute{
				Computed: true,
			},
			"line_count": schema.Int64Attribute{
				Computed: true,
			},
			"object": schema.StringAttribute{
				Computed: true,
			},
			"processed": schema.BoolAttribute{
				Computed: true,
			},
			"purpose": schema.StringAttribute{
				Description: "The purpose of the file\nAvailable values: \"fine-tune\", \"eval\", \"eval-sample\", \"eval-output\", \"eval-summary\", \"batch-generated\", \"batch-api\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"fine-tune",
						"eval",
						"eval-sample",
						"eval-output",
						"eval-summary",
						"batch-generated",
						"batch-api",
					),
				},
			},
		},
	}
}

func (d *FileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *FileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
