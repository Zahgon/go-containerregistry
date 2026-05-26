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

package v1

import (
	"encoding"
	"encoding/json"
	"hash"
	"io"
)

// Hash is an unqualified digest of some content, e.g. sha256:deadbeef
type Hash struct {
	// Algorithm holds the algorithm used to compute the hash.
	Algorithm string

	// Hex holds the hex portion of the content hash.
	Hex string
}

var _ encoding.TextMarshaler = (*Hash)(nil)
var _ encoding.TextUnmarshaler = (*Hash)(nil)
var _ json.Marshaler = (*Hash)(nil)
var _ json.Unmarshaler = (*Hash)(nil)

// String reverses NewHash returning the string-form of the hash.
func (h Hash) String() string { _ = "STUB: not implemented"; return "" }

// NewHash validates the input string is a hash and returns a strongly type Hash object.
func NewHash(s string) (Hash, error) { _ = "STUB: not implemented"; return *new(Hash), nil }

// MarshalJSON implements json.Marshaler
func (h Hash) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements json.Unmarshaler
func (h *Hash) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText implements encoding.TextMarshaler. This is required to use
// v1.Hash as a key in a map when marshalling JSON.
func (h Hash) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler. This is required to use
// v1.Hash as a key in a map when unmarshalling JSON.
func (h *Hash) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// Hasher returns a hash.Hash for the named algorithm (e.g. "sha256")
func Hasher(name string) (hash.Hash, error) { _ = "STUB: not implemented"; return *new(hash.Hash), nil }

func (h *Hash) parse(unquoted string) error { _ = "STUB: not implemented"; return nil }

// Compare the hex to the expected size (2 hex characters per byte)

// SHA256 computes the Hash of the provided io.Reader's content.
func SHA256(r io.Reader) (Hash, int64, error) { _ = "STUB: not implemented"; return *new(Hash), 0, nil }
