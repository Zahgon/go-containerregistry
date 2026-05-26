// Copyright 2019 Google LLC All Rights Reserved.
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

package gcrane

import (
	"context"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

// Option is a functional option for gcrane operations.
type Option func(*options)

type options struct {
	jobs   int
	remote []remote.Option
	google []google.Option
	crane  []crane.Option
}

func makeOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

// WithJobs sets the number of concurrent jobs to run.
//
// The default number of jobs is GOMAXPROCS.
func WithJobs(jobs int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTransport is a functional option for overriding the default transport
// for remote operations.
func WithTransport(t http.RoundTripper) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithUserAgent adds the given string to the User-Agent header for any HTTP
// requests.
func WithUserAgent(ua string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContext is a functional option for setting the context.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithPlatform is a functional option for selecting a single platform from
// a multi-platform image. A nil platform copies the index unchanged.
func WithPlatform(platform *v1.Platform) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithKeychain is a functional option for overriding the default
// authenticator for remote operations, using an authn.Keychain to find
// credentials.
//
// By default, gcrane will use gcrane.Keychain.
func WithKeychain(keys authn.Keychain) Option {
	_ = "STUB: not implemented"
	return *

	// Replace the default keychain at position 0.
	new(Option)
}

// WithAuth is a functional option for overriding the default authenticator
// for remote operations.
//
// By default, gcrane will use gcrane.Keychain.
func WithAuth(auth authn.Authenticator) Option {
	_ = "STUB: not implemented"
	return *

	// Replace the default keychain at position 0.
	new(Option)
}
