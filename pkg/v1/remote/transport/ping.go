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
	"context"
	"net/http"
	"time"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote/internal/authchallenge"
)

// 300ms is the default fallback period for go's DNS dialer but we could make this configurable.
var fallbackDelay = 300 * time.Millisecond

type Challenge struct {
	Scheme string

	// Following the challenge there are often key/value pairs
	// e.g. Bearer service="gcr.io",realm="https://auth.gcr.io/v36/tokenz"
	Parameters map[string]string

	// Whether we had to use http to complete the Ping.
	Insecure bool
}

// Ping does a GET /v2/ against the registry and returns the response.
func Ping(ctx context.Context, reg name.Registry, t http.RoundTripper) (*Challenge, error) {
	_ = "STUB: not implemented"
	// This first attempts to use "https" for every request, falling back to http
	// if the registry matches our localhost heuristic or if it is intentionally
	// set to insecure via name.NewInsecureRegistry.
	return nil, nil
}

func pingSingle(ctx context.Context, reg name.Registry, t http.RoundTripper, scheme string) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// By draining the body, make sure to reuse the connection made by
// the ping for the following access to the registry

// If resp.Request is set, we may have followed a redirect,
// so we want to prefer resp.Request.URL.Scheme (if it's set)
// falling back to the original request's scheme.

// If we get a 200, then no authentication is needed.

// If we hit more than one, let's try to find one that we know how to handle.

// Otherwise, just return the challenge without parameters.

// Based on the golang happy eyeballs dialParallel impl in net/dial.go.
func pingParallel(ctx context.Context, reg name.Registry, t http.RoundTripper, schemes []string) (*Challenge, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Primary failed and we haven't started the fallback,
// reset time to start fallback immediately.

func pickFromMultipleChallenges(challenges []authchallenge.Challenge) authchallenge.Challenge {
	_ = "STUB: not implemented"
	// It might happen there are multiple www-authenticate headers, e.g. `Negotiate` and `Basic`.
	// Picking simply the first one could result eventually in `unrecognized challenge` error,
	// that's why we're looping through the challenges in search for one that can be handled.
	return *new(authchallenge.Challenge)
}

type multierrs []error

func (m multierrs) Error() string { _ = "STUB: not implemented"; return "" }

func (m multierrs) As(target any) bool { _ = "STUB: not implemented"; return false }

func (m multierrs) Is(target error) bool { _ = "STUB: not implemented"; return false }
