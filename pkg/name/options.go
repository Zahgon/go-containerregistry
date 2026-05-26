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

package name

const (
	// DefaultRegistry is the registry name that will be used if no registry
	// provided and the default is not overridden.
	DefaultRegistry      = "index.docker.io"
	defaultRegistryAlias = "docker.io"

	// DefaultTag is the tag name that will be used if no tag provided and the
	// default is not overridden.
	DefaultTag = "latest"
)

type options struct {
	strict          bool // weak by default
	insecure        bool // secure by default
	defaultRegistry string
	defaultTag      string
}

func makeOptions(opts ...Option) options { _ = "STUB: not implemented"; return *new(options) }

// Option is a functional option for name parsing.
type Option func(*options)

// StrictValidation is an Option that requires image references to be fully
// specified; i.e. no defaulting for registry (dockerhub), repo (library),
// or tag (latest).
func StrictValidation(opts *options) {
	_ = "STUB: not implemented"

	// WeakValidation is an Option that sets defaults when parsing names, see
	// StrictValidation.
	return
}

func WeakValidation(opts *options) { _ = "STUB: not implemented"; return }

// Insecure is an Option that allows image references to be fetched without TLS.
func Insecure(opts *options) { _ = "STUB: not implemented"; return }

// OptionFn is a function that returns an option.
type OptionFn func() Option

// WithDefaultRegistry sets the default registry that will be used if one is not
// provided.
func WithDefaultRegistry(r string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithDefaultTag sets the default tag that will be used if one is not provided.
func WithDefaultTag(t string) Option { _ = "STUB: not implemented"; return *new(Option) }
