package common

import (
	api "github.com/arkhaix/lit-reader/api/common"
)

// Status is returned with all http responses
type Status struct {
	Code int
	Text string
}

// NewStatusFromProto creates a Status struct from the api/common/status proto message
func NewStatusFromProto(status *api.Status) Status {
	return Status{
		Code: int(status.GetStatusCode()),
		Text: status.GetStatusText(),
	}
}
