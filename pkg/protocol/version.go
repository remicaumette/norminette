package protocol

type VersionRequest struct {
	Action	string	`json:"action"`
}

func NewVersionRequest() *VersionRequest {
	return &VersionRequest{Action: "version"}
}

type VersionResponse struct {
	Version string `json:"display"`
	Stop    bool   `json:"stop"`
}
