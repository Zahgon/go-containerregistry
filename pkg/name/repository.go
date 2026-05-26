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
	defaultNamespace = "library"
	repositoryChars  = "abcdefghijklmnopqrstuvwxyz0123456789_-./"
	regRepoDelimiter = "/"
)

// Repository stores a docker repository name in a structured form.
type Repository struct {
	Registry
	repository string
}

var _ encoding.TextMarshaler = (*Repository)(nil)
var _ encoding.TextUnmarshaler = (*Repository)(nil)
var _ json.Marshaler = (*Repository)(nil)
var _ json.Unmarshaler = (*Repository)(nil)

// See https://docs.docker.com/docker-hub/official_repos
func hasImplicitNamespace(repo string, reg Registry) bool { _ = "STUB: not implemented"; return false }

// RepositoryStr returns the repository component of the Repository.
func (r Repository) RepositoryStr() string { _ = "STUB: not implemented"; return "" }

// Name returns the name from which the Repository was derived.
func (r Repository) Name() string { _ = "STUB: not implemented"; return "" }

// TODO: As far as I can tell, this is unreachable.

func (r Repository) String() string {
	_ = "STUB: not implemented"

	// Scope returns the scope required to perform the given action on the registry.
	// TODO(jonjohnsonjr): consider moving scopes to a separate package.
	return ""
}

func (r Repository) Scope(action string) string { _ = "STUB: not implemented"; return "" }

func checkRepository(repository string) error { _ = "STUB: not implemented"; return nil }

// NewRepository returns a new Repository representing the given name, according to the given strictness.
func NewRepository(name string, opts ...Option) (Repository, error) {
	_ = "STUB: not implemented"
	return *new(Repository), nil
}

// The first part of the repository is treated as the registry domain
// iff it contains a '.' or ':' character, otherwise it is all repository
// and the domain defaults to Docker Hub.

// Tag returns a Tag in this Repository.
func (r Repository) Tag(identifier string) Tag { _ = "STUB: not implemented"; return *new(Tag) }

// Digest returns a Digest in this Repository.
func (r Repository) Digest(identifier string) Digest {
	_ = "STUB: not implemented"
	return *new(Digest)
}

// MarshalJSON formats the Repository into a string for JSON serialization.
func (r Repository) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON parses a JSON string into a Repository.
func (r *Repository) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText formats the repository name into a string for text serialization.
func (r Repository) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText parses a text string into a Repository.
func (r *Repository) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }
