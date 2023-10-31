package scan

import (
	"net/http"
	"slices"
	"strings"

	"github.com/nullify-platform/attack-surface-scanner/internal/openapi"
	"github.com/nullify-platform/logger/pkg/logger"
)

var authStatusCodes = []int{
	http.StatusUnauthorized,
	http.StatusForbidden,
	http.StatusNotFound,
}

func executeScan(targetHost string, spec *openapi.OpenAPI) *ScanResults {
	results := ScanResults{
		WithAuth:     []ScanResult{},
		WithoutAuth:  []ScanResult{},
		ErrorResults: []ScanResult{},
	}

	for path, methods := range spec.Paths {
		for method := range methods {
			scanResult := ScanResult{
				Method: method,
				Path:   path,
			}

			method = strings.ToUpper(method)

			req, err := http.NewRequest(method, targetHost+path, nil)
			if err != nil {
				scanResult.Error = err.Error()
				results.ErrorResults = append(results.ErrorResults, scanResult)
				continue
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				scanResult.Error = err.Error()
				results.ErrorResults = append(results.ErrorResults, scanResult)
				continue
			}

			logger.Info(
				"api response",
				logger.String("method", method),
				logger.String("path", path),
				logger.Int("status", res.StatusCode),
			)

			scanResult.Status = res.StatusCode

			if slices.Contains(authStatusCodes, res.StatusCode) {
				results.WithAuth = append(results.WithAuth, scanResult)
			} else if res.StatusCode >= 500 {
				results.ErrorResults = append(results.ErrorResults, scanResult)
			} else {
				results.WithoutAuth = append(results.WithoutAuth, scanResult)
			}
		}
	}

	return &results
}
