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
)

const (
	// TODO(dekkagaijin): use the docker/distribution regexes for validation.
	tagChars = "abcdefghijklmnopqrstuvwxyz0123456789_-.ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tagDelim = ":"
)

// Tag stores a docker tag name in a structured form.
type Tag struct {
	Repository
	tag      string
	original string
}

var _ Reference = (*Tag)(nil)
var _ encoding.TextMarshaler = (*Tag)(nil)
var _ encoding.TextUnmarshaler = (*Tag)(nil)
var _ json.Marshaler = (*Tag)(nil)
var _ json.Unmarshaler = (*Tag)(nil)

// Context implements Reference.
func (t Tag) Context() Repository {
	_ = "STUB: not implemented"
	return *

	// Identifier implements Reference.
	new(Repository)
}

func (t Tag) Identifier() string {
	_ = "STUB: not implemented"

	// TagStr returns the tag component of the Tag.
	return ""
}

func (t Tag) TagStr() string {
	_ = "STUB: not implemented"

	// Name returns the name from which the Tag was derived.
	return ""
}

func (t Tag) Name() string { _ = "STUB: not implemented"; return "" }

// String returns the original input string.
func (t Tag) String() string {
	_ = "STUB: not implemented"

	// Scope returns the scope required to perform the given action on the tag.
	return ""
}

func (t Tag) Scope(action string) string { _ = "STUB: not implemented"; return "" }

func checkTag(name string) error { _ = "STUB: not implemented"; return nil }

// NewTag returns a new Tag representing the given name, according to the given strictness.
func NewTag(name string, opts ...Option) (Tag, error) {
	_ = "STUB: not implemented"
	return *new(Tag), nil
}

// Split on ":"

// Verify that we aren't confusing a tag for a hostname w/ port for the purposes of weak validation.

// We don't require a tag, but if we get one check it's valid,
// even when not being strict.
// If we are being strict, we want to validate the tag regardless in case
// it's empty.

// MarshalJSON formats the Tag into a string for JSON serialization.
func (t Tag) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a JSON string into a Tag.
func (t *Tag) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText formats the tag into a string for text serialization.
func (t Tag) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText parses a text string into a Tag.
func (t *Tag) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }
