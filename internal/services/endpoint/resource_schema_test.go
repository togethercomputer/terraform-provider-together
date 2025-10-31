// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/endpoint"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestEndpointModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*endpoint.EndpointModel)(nil)
	schema := endpoint.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
