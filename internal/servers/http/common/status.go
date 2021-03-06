package common

import (
	api "github.com/arkhaix/lit-reader/api/common"
)

// Status is returned with all http responses
type Status struct {
	Code int    `json:"Code"`
	Text string `json:"Text"`
}

// NewStatusFromProto creates a Status struct from the api/common/status proto message
func NewStatusFromProto(status *api.Status) Status {
	return Status{
		Code: int(status.GetCode()),
		Text: status.GetText(),
	}
}
