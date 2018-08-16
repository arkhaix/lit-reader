package common

import (
	"errors"
	"net"
	"net/url"
)

// ScraperError aggregates the underlying error with a helper to indicate retryability
// This is done so that users don't need to inspect underlying error types and repeat logic
type ScraperError struct {
	Err error
}

func (e ScraperError) Error() string {
	return e.Err.Error()
}

// NewScraperError returns a ScraperError which wraps the specified error
func NewScraperError(err error) ScraperError {
	return ScraperError{
		Err: err,
	}
}

// NewScraperErrorString returns a ScraperError from a string
func NewScraperErrorString(s string) ScraperError {
	return ScraperError{
		Err: errors.New(s),
	}
}

// CanRetry returns true if the scraper error is temporary and can be retried
func (e ScraperError) CanRetry() bool {
	// net.Error
	if netErr, ok := e.Err.(net.Error); ok {
		if netErr.Temporary() || netErr.Timeout() {
			return true
		}
	}

	// url.Error
	if urlErr, ok := e.Err.(*url.Error); ok {
		if urlErr.Temporary() || urlErr.Timeout() {
			return true
		}
	}

	return false
}
