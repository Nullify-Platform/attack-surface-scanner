package scan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nullify-platform/logger/pkg/logger"
)

func Scan(specPath string, targetHost string) error {
	logger.Info(
		"running scan",
		logger.String("specPath", specPath),
		logger.String("targetHost", targetHost),
	)

	// remove trailing slash from target host
	targetHost = strings.TrimRight(targetHost, "/")

	spec, err := loadSpecFile(specPath)
	if err != nil {
		return fmt.Errorf("failed to load openapi spec file: %w", err)
	}

	results := executeScan(targetHost, spec)

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(results)
	if err != nil {
		return fmt.Errorf("failed to encode scan results: %w", err)
	}

	return nil
}
