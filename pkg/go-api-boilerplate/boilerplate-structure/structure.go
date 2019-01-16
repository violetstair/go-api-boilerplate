package boilerplate_structure

type EOSInfo struct {
	Version       string `json:"version"`
	VersionString string `json:"version_string"`
}

type RESPONSE struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Items     interface{} `json:"items"`
}
