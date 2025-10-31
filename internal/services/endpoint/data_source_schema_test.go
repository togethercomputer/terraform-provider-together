// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package endpoint_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/togetherai-terraform/internal/services/endpoint"
	"github.com/stainless-sdks/togetherai-terraform/internal/test_helpers"
)

func TestEndpointDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*endpoint.EndpointDataSourceModel)(nil)
	schema := endpoint.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
