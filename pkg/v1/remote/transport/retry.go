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
	"time"

	"github.com/google/go-containerregistry/internal/retry"
)

// Sleep for 0.1 then 0.3 seconds. This should cover networking blips.
var defaultBackoff = retry.Backoff{
	Duration: 100 * time.Millisecond,
	Factor:   3.0,
	Jitter:   0.1,
	Steps:    3,
}

var _ http.RoundTripper = (*retryTransport)(nil)

// retryTransport wraps a RoundTripper and retries temporary network errors.
type retryTransport struct {
	inner     http.RoundTripper
	backoff   retry.Backoff
	predicate retry.Predicate
	codes     []int
}

// Option is a functional option for retryTransport.
type Option func(*options)

type options struct {
	backoff   retry.Backoff
	predicate retry.Predicate
	codes     []int
}

// Backoff is an alias of retry.Backoff to expose this configuration option to consumers of this lib
type Backoff = retry.Backoff

// WithRetryBackoff sets the backoff for retry operations.
func WithRetryBackoff(backoff Backoff) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithRetryPredicate sets the predicate for retry operations.
func WithRetryPredicate(predicate func(error) bool) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithRetryStatusCodes sets which http response codes will be retried.
func WithRetryStatusCodes(codes ...int) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewRetry returns a transport that retries errors.
func NewRetry(inner http.RoundTripper, opts ...Option) http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func (t *retryTransport) RoundTrip(in *http.Request) (out *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
