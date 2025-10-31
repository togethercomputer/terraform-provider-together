// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package completion

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*CompletionResource)(nil)

func (r *CompletionResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
