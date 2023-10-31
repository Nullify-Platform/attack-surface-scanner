package scan

import (
	"testing"

	"github.com/nullify-platform/attack-surface-scanner/internal/openapi"
	"github.com/stretchr/testify/require"
)

func TestSpecfile(t *testing.T) {
	spec, err := loadSpecFile("test/openapi.yml")
	require.NoError(t, err)

	require.Equal(t, "3.0.0", spec.OpenAPI)

	paths := map[string]map[string]openapi.Operation{
		"/users": {
			"get": {
				Tags:        []string(nil),
				Summary:     "Returns a list of users.",
				Description: "Optional extended description in CommonMark or HTML.",
				OperationID: "",
				Parameters:  []openapi.Parameter(nil),
				Responses: map[string]openapi.Schema{
					"200": {
						Type: "",
					},
				},
			},
		},
	}

	require.Equal(t, paths, spec.Paths)
}
