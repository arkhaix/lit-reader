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
		StatusCode: 200,
		StatusText: "Ok",
	}

	StatusBadRequest = &apicommon.Status{
		StatusCode: 400,
		StatusText: "Bad request",
	}

	StatusNotFound = &apicommon.Status{
		StatusCode: 404,
		StatusText: "Resource not found",
	}

	StatusInternalServerError = &apicommon.Status{
		StatusCode: 500,
		StatusText: "Internal server error",
	}
}
