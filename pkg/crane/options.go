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

package crane

import (
	"context"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

// Options hold the options that crane uses when calling other packages.
type Options struct {
	Name      []name.Option
	Remote    []remote.Option
	Platform  *v1.Platform
	Keychain  authn.Keychain
	Transport http.RoundTripper

	auth      authn.Authenticator
	insecure  bool
	jobs      int
	noclobber bool
	ctx       context.Context
}

// GetOptions exposes the underlying []remote.Option, []name.Option, and
// platform, based on the passed Option. Generally, you shouldn't need to use
// this unless you've painted yourself into a dependency corner as we have
// with the crane and gcrane cli packages.
func GetOptions(opts ...Option) Options { _ = "STUB: not implemented"; return *new(Options) }

func makeOptions(opts ...Option) Options { _ = "STUB: not implemented"; return *new(Options) }

// Allow for untrusted certificates if the user
// passed Insecure but no custom transport.

//nolint: gosec

// Option is a functional option for crane.
type Option func(*Options)

// WithTransport is a functional option for overriding the default transport
// for remote operations. Setting a transport will override the Insecure option's
// configuration allowing for image registries to use untrusted certificates.
func WithTransport(t http.RoundTripper) Option { _ = "STUB: not implemented"; return *new(Option) }

// Insecure is an Option that allows image references to be fetched without TLS.
// This will also allow for untrusted (e.g. self-signed) certificates in cases where
// the default transport is used (i.e. when WithTransport is not used).
func Insecure(o *Options) { _ = "STUB: not implemented"; return }

// WithPlatform is an Option to specify the platform.
func WithPlatform(platform *v1.Platform) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithAuthFromKeychain is a functional option for overriding the default
// authenticator for remote operations, using an authn.Keychain to find
// credentials.
//
// By default, crane will use authn.DefaultKeychain.
func WithAuthFromKeychain(keys authn.Keychain) Option {
	_ = "STUB: not implemented"
	return *

	// Replace the default keychain at position 0.
	new(Option)
}

// WithAuth is a functional option for overriding the default authenticator
// for remote operations.
//
// By default, crane will use authn.DefaultKeychain.
func WithAuth(auth authn.Authenticator) Option {
	_ = "STUB: not implemented"
	return *

	// Replace the default keychain at position 0.
	new(Option)
}

// WithUserAgent adds the given string to the User-Agent header for any HTTP
// requests.
func WithUserAgent(ua string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNondistributable is an option that allows pushing non-distributable
// layers.
func WithNondistributable() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContext is a functional option for setting the context.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithJobs sets the number of concurrent jobs to run.
//
// The default number of jobs is GOMAXPROCS.
func WithJobs(jobs int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoClobber modifies behavior to avoid overwriting existing tags, if possible.
func WithNoClobber(noclobber bool) Option { _ = "STUB: not implemented"; return *new(Option) }
