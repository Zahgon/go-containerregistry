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

import (
	"encoding"
	"encoding/json"
	"regexp"
)

// Detect more complex forms of localhost references.
var reLocal = regexp.MustCompile(`.*\.localhost(?::\d{1,5})?$`)

// Detect the loopback IP (127.0.0.1)
var reLoopback = regexp.MustCompile(regexp.QuoteMeta("127.0.0.1"))

// Detect the loopback IPV6 (::1)
var reipv6Loopback = regexp.MustCompile(regexp.QuoteMeta("::1"))

// Registry stores a docker registry name in a structured form.
type Registry struct {
	insecure bool
	registry string
}

var _ encoding.TextMarshaler = (*Registry)(nil)
var _ encoding.TextUnmarshaler = (*Registry)(nil)
var _ json.Marshaler = (*Registry)(nil)
var _ json.Unmarshaler = (*Registry)(nil)

// RegistryStr returns the registry component of the Registry.
func (r Registry) RegistryStr() string {
	_ = "STUB: not implemented"

	// Name returns the name from which the Registry was derived.
	return ""
}

func (r Registry) Name() string { _ = "STUB: not implemented"; return "" }

func (r Registry) String() string {
	_ = "STUB: not implemented"

	// Repo returns a Repository in the Registry with the given name.
	return ""
}

func (r Registry) Repo(repo ...string) Repository {
	_ = "STUB: not implemented"
	return *new(Repository)
}

// Scope returns the scope required to access the registry.
func (r Registry) Scope(string) string {
	_ = "STUB: not implemented"
	// The only resource under 'registry' is 'catalog'. http://goo.gl/N9cN9Z
	return ""
}

func (r Registry) isRFC1918() bool { _ = "STUB: not implemented"; return false }

// Scheme returns https scheme for all the endpoints except localhost or when explicitly defined.
func (r Registry) Scheme() string { _ = "STUB: not implemented"; return "" }

func checkRegistry(name string) error {
	_ = "STUB: not implemented"
	// Per RFC 3986, registries (authorities) are required to be prefixed with "//"
	// url.Host == hostname[:port] == authority
	return nil
}

// NewRegistry returns a Registry based on the given name.
// Strict validation requires explicit, valid RFC 3986 URI authorities to be given.
func NewRegistry(name string, opts ...Option) (Registry, error) {
	_ = "STUB: not implemented"
	return *new(Registry), nil
}

// Rewrite "docker.io" to "index.docker.io".
// See: https://github.com/google/go-containerregistry/issues/68

// NewInsecureRegistry returns an Insecure Registry based on the given name.
//
// Deprecated: Use the Insecure Option with NewRegistry instead.
func NewInsecureRegistry(name string, opts ...Option) (Registry, error) {
	_ = "STUB: not implemented"
	return *new(Registry), nil
}

// MarshalJSON formats the Registry into a string for JSON serialization.
func (r Registry) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a JSON string into a Registry.
func (r *Registry) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText formats the registry into a string for text serialization.
func (r Registry) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText parses a text string into a Registry.
func (r *Registry) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }
