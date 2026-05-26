// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package transport

import (
	"net/http"
)

// Error implements error to support the following error specification:
// https://github.com/distribution/distribution/blob/aac2f6c8b7c5a6c60190848bab5cbeed2b5ba0a9/docs/spec/api.md#errors
type Error struct {
	Errors []Diagnostic `json:"errors,omitempty"`
	// The http status code returned.
	StatusCode int
	// The request that failed.
	Request *http.Request
	// The raw body if we couldn't understand it.
	rawBody string

	// Bit of a hack to make it easier to force a retry.
	temporary bool
}

// Check that Error implements error
var _ error = (*Error)(nil)

// Error implements error
func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) responseErr() string { _ = "STUB: not implemented"; return "" }

// Temporary returns whether the request that preceded the error is temporary.
func (e *Error) Temporary() bool { _ = "STUB: not implemented"; return false }

// Diagnostic represents a single error returned by a Docker registry interaction.
type Diagnostic struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message,omitempty"`
	Detail  any       `json:"detail,omitempty"`
}

// String stringifies the Diagnostic in the form: $Code: $Message[; $Detail]
func (d Diagnostic) String() string { _ = "STUB: not implemented"; return "" }

// ErrorCode is an enumeration of supported error codes.
type ErrorCode string

// The set of error conditions a registry may return:
// https://github.com/distribution/distribution/blob/aac2f6c8b7c5a6c60190848bab5cbeed2b5ba0a9/docs/spec/api.md#errors-2
// maxErrorBodySize limits HTTP error response body reads to prevent OOM when
// a registry returns an unexpectedly large error body.
const maxErrorBodySize = 64 * 1024 // 64 KiB

const (
	BlobUnknownErrorCode         ErrorCode = "BLOB_UNKNOWN"
	BlobUploadInvalidErrorCode   ErrorCode = "BLOB_UPLOAD_INVALID"
	BlobUploadUnknownErrorCode   ErrorCode = "BLOB_UPLOAD_UNKNOWN"
	DigestInvalidErrorCode       ErrorCode = "DIGEST_INVALID"
	ManifestBlobUnknownErrorCode ErrorCode = "MANIFEST_BLOB_UNKNOWN"
	ManifestInvalidErrorCode     ErrorCode = "MANIFEST_INVALID"
	ManifestUnknownErrorCode     ErrorCode = "MANIFEST_UNKNOWN"
	ManifestUnverifiedErrorCode  ErrorCode = "MANIFEST_UNVERIFIED"
	NameInvalidErrorCode         ErrorCode = "NAME_INVALID"
	NameUnknownErrorCode         ErrorCode = "NAME_UNKNOWN"
	SizeInvalidErrorCode         ErrorCode = "SIZE_INVALID"
	TagInvalidErrorCode          ErrorCode = "TAG_INVALID"
	UnauthorizedErrorCode        ErrorCode = "UNAUTHORIZED"
	DeniedErrorCode              ErrorCode = "DENIED"
	UnsupportedErrorCode         ErrorCode = "UNSUPPORTED"
	TooManyRequestsErrorCode     ErrorCode = "TOOMANYREQUESTS"
	UnknownErrorCode             ErrorCode = "UNKNOWN"

	// This isn't defined by either docker or OCI spec, but is defined by docker/distribution:
	// https://github.com/distribution/distribution/blob/6a977a5a754baa213041443f841705888107362a/registry/api/errcode/register.go#L60
	UnavailableErrorCode ErrorCode = "UNAVAILABLE"
)

// TODO: Include other error types.
var temporaryErrorCodes = map[ErrorCode]struct{}{
	BlobUploadInvalidErrorCode: {},
	TooManyRequestsErrorCode:   {},
	UnknownErrorCode:           {},
	UnavailableErrorCode:       {},
}

var temporaryStatusCodes = map[int]struct{}{
	http.StatusRequestTimeout:      {},
	http.StatusTooManyRequests:     {}, // matches TooManyRequestsErrorCode in temporaryErrorCodes
	http.StatusInternalServerError: {},
	http.StatusBadGateway:          {},
	http.StatusServiceUnavailable:  {},
	http.StatusGatewayTimeout:      {},
}

// CheckError returns a structured error if the response status is not in codes.
func CheckError(resp *http.Response, codes ...int) error { _ = "STUB: not implemented"; return nil }

// This is one of the supported status codes.

func makeError(resp *http.Response, body []byte) *Error {
	_ = "STUB: not implemented"
	// https://github.com/distribution/distribution/blob/aac2f6c8b7c5a6c60190848bab5cbeed2b5ba0a9/docs/spec/api.md#errors
	return nil
}

// This can fail if e.g. the response body is not valid JSON. That's fine,
// we'll construct an appropriate error string from the body and status code.

func retryError(resp *http.Response) error { _ = "STUB: not implemented"; return nil }

// Restore the body so that a subsequent CheckError call (after the
// retry loop exhausts its retries) can still read and parse the
// structured registry error from the response.
