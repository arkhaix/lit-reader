package common

import (
	apicommon "github.com/arkhaix/lit-reader/api/common"
)

var (
	// StatusOk represents a successful operation (http 200)
	StatusOk *apicommon.Status

	// StatusBadRequest represents a problem with the input parameters (http 400)
	StatusBadRequest *apicommon.Status

	// StatusNotFound represents a failure to locate the requested resource (http 404)
	StatusNotFound *apicommon.Status

	// StatusInternalServerError represents an internal failure (http 500)
	StatusInternalServerError *apicommon.Status
)

func init() {
	StatusOk = &apicommon.Status{
		Code: 200,
		Text: "Ok",
	}

	StatusBadRequest = &apicommon.Status{
		Code: 400,
		Text: "Bad request",
	}

	StatusNotFound = &apicommon.Status{
		Code: 404,
		Text: "Resource not found",
	}

	StatusInternalServerError = &apicommon.Status{
		Code: 500,
		Text: "Internal server error",
	}
}
