// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fine_tune

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*FineTuneResource)(nil)

func (r *FineTuneResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
