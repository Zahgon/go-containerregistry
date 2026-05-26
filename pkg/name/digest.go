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
	// nolint: depguard
	_ "crypto/sha256" // Recommended by go-digest.
	"encoding"
	"encoding/json"
)

const digestDelim = "@"

// Digest stores a digest name in a structured form.
type Digest struct {
	Repository
	digest   string
	original string
}

var _ Reference = (*Digest)(nil)
var _ encoding.TextMarshaler = (*Digest)(nil)
var _ encoding.TextUnmarshaler = (*Digest)(nil)
var _ json.Marshaler = (*Digest)(nil)
var _ json.Unmarshaler = (*Digest)(nil)

// Context implements Reference.
func (d Digest) Context() Repository {
	_ = "STUB: not implemented"
	return *

	// Identifier implements Reference.
	new(Repository)
}

func (d Digest) Identifier() string { _ = "STUB: not implemented"; return "" }

// DigestStr returns the digest component of the Digest.
func (d Digest) DigestStr() string {
	_ = "STUB: not implemented"

	// Name returns the name from which the Digest was derived.
	return ""
}

func (d Digest) Name() string { _ = "STUB: not implemented"; return "" }

// String returns the original input string.
func (d Digest) String() string {
	_ = "STUB: not implemented"

	// MarshalJSON formats the digest into a string for JSON serialization.
	return ""
}

func (d Digest) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a JSON string into a Digest.
func (d *Digest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText formats the digest into a string for text serialization.
func (d Digest) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText parses a text string into a Digest.
func (d *Digest) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// NewDigest returns a new Digest representing the given name.
func NewDigest(name string, opts ...Option) (Digest, error) {
	_ = "STUB: not implemented"
	// Split on "@"
	return *new(Digest), nil
}
