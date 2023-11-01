package scan

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nullify-platform/attack-surface-scanner/internal/openapi"
	"github.com/nullify-platform/logger/pkg/logger"
	"gopkg.in/yaml.v3"
)

func loadSpecFile(path string) (*openapi.OpenAPI, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var openAPISpec openapi.OpenAPI

	if strings.HasSuffix(path, ".json") {
		err = json.NewDecoder(f).Decode(&openAPISpec)
	} else if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		err = yaml.NewDecoder(f).Decode(&openAPISpec)
	} else {
		err = fmt.Errorf("unsupported file type: %s", path)
	}

	if err != nil {
		return nil, err
	}

	logger.Debug(
		"loaded spec file",
		logger.Any("openapi", openAPISpec),
	)

	return &openAPISpec, nil
}
