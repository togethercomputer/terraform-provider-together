// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audio

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*AudioResource)(nil)

func (r *AudioResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
