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

package registry

import (
	"net/http"
)

type regError struct {
	Status  int
	Code    string
	Message string
}

func (r *regError) Write(resp http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

// regErrInternal returns an internal server error.
func regErrInternal(err error) *regError { _ = "STUB: not implemented"; return nil }

var regErrBlobUnknown = &regError{
	Status:  http.StatusNotFound,
	Code:    "BLOB_UNKNOWN",
	Message: "Unknown blob",
}

var regErrUnsupported = &regError{
	Status:  http.StatusMethodNotAllowed,
	Code:    "UNSUPPORTED",
	Message: "Unsupported operation",
}

var regErrDigestMismatch = &regError{
	Status:  http.StatusBadRequest,
	Code:    "DIGEST_INVALID",
	Message: "digest does not match contents",
}

var regErrDigestInvalid = &regError{
	Status:  http.StatusBadRequest,
	Code:    "NAME_INVALID",
	Message: "invalid digest",
}
