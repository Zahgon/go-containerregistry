// Copyright 2023 Google LLC All Rights Reserved.
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

package remote

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Referrers returns a list of descriptors that refer to the given manifest digest.
//
// The subject manifest doesn't have to exist in the registry for there to be descriptors that refer to it.
func Referrers(d name.Digest, options ...Option) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

// https://github.com/opencontainers/distribution-spec/blob/main/spec.md#referrers-tag-schema
func fallbackTag(d name.Digest) name.Tag { _ = "STUB: not implemented"; return *new(name.Tag) }

func (f *fetcher) fetchReferrers(ctx context.Context, filter map[string]string, d name.Digest) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	// Check the Referrers API endpoint first.
	return *new(v1.ImageIndex), nil
}

// The registry doesn't support the Referrers API endpoint, so we'll use the fallback tag scheme.

// Not found just means there are no attachments yet. Start with an empty manifest.

// If filter applied, filter out by artifactType.
// See https://github.com/opencontainers/distribution-spec/blob/main/spec.md#listing-referrers
func filterReferrersResponse(filter map[string]string, in v1.ImageIndex) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}
