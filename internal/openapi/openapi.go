package openapi

type OpenAPI struct {
	OpenAPI string                          `json:"openapi" yaml:"openapi"`
	Info    Info                            `json:"info"    yaml:"info"`
	Servers []Server                        `json:"servers" yaml:"servers"`
	Paths   map[string]map[string]Operation `json:"paths"   yaml:"paths"`
}

type Server struct {
	URL string `json:"url" yaml:"url"`
}

type Info struct {
	Title   string `json:"title"   yaml:"title"`
	Version string `json:"version" yaml:"version"`
}

type Operation struct {
	Tags        []string          `json:"tags"        yaml:"tags"`
	Summary     string            `json:"summary"     yaml:"summary"`
	Description string            `json:"description" yaml:"description"`
	OperationID string            `json:"operationId" yaml:"operationId"`
	Parameters  []Parameter       `json:"parameters"  yaml:"parameters"`
	Responses   map[string]Schema `json:"responses"   yaml:"responses"`
}

type Parameter struct {
	Name        string `json:"name"        yaml:"name"`
	In          string `json:"in"          yaml:"in"`
	Description string `json:"description" yaml:"description"`
	Required    bool   `json:"required"    yaml:"required"`
	Schema      Schema `json:"schema"      yaml:"schema"`
}

type Schema struct {
	Type string `json:"type" yaml:"type"`
}
