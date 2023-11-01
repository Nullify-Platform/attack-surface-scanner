package scan

type ScanResults struct {
	WithAuth     []ScanResult `json:"withAuth"`
	WithoutAuth  []ScanResult `json:"withoutAuth"`
	ErrorResults []ScanResult `json:"errors"`
}

type ScanResult struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Status int    `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}
