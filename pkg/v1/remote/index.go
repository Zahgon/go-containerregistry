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

package remote

import (
	"context"
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

var acceptableIndexMediaTypes = []types.MediaType{
	types.DockerManifestList,
	types.OCIImageIndex,
}

// remoteIndex accesses an index from a remote registry
type remoteIndex struct {
	fetcher      fetcher
	ref          name.Reference
	ctx          context.Context
	manifestLock sync.Mutex // Protects manifest
	manifest     []byte
	mediaType    types.MediaType
	descriptor   *v1.Descriptor
}

// Index provides access to a remote index reference.
func Index(ref name.Reference, options ...Option) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

func (r *remoteIndex) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (r *remoteIndex) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (r *remoteIndex) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *remoteIndex) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NOTE(jonjohnsonjr): We should never get here because the public entrypoints
// do type-checking via remote.Descriptor. I've left this here for tests that
// directly instantiate a remoteIndex.

func (r *remoteIndex) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *remoteIndex) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Descriptor.Image will handle coercing nested indexes into an Image.

// Descriptor retains the original descriptor from an index manifest.
// See partial.Descriptor.
func (r *remoteIndex) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	// kind of a hack, but RawManifest does appropriate locking/memoization
	// and makes sure r.descriptor is populated.
	return nil, nil
}

func (r *remoteIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

// Workaround for #819.
func (r *remoteIndex) Layer(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (r *remoteIndex) imageByPlatform(platform v1.Platform) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Descriptor.Image will handle coercing nested indexes into an Image.

// This naively matches the first manifest with matching platform attributes.
//
// We should probably use this instead:
//
//	github.com/containerd/containerd/platforms
//
// But first we'd need to migrate to:
//
//	github.com/opencontainers/image-spec/specs-go/v1
func (r *remoteIndex) childByPlatform(platform v1.Platform) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If platform is missing from child descriptor, assume it's amd64/linux.

func (r *remoteIndex) childByHash(h v1.Hash) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert one of this index's child's v1.Descriptor into a remote.Descriptor, with the given platform option.
func (r *remoteIndex) childDescriptor(child v1.Descriptor, platform v1.Platform) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Failing to parse as a manifest should just be ignored.
// The manifest might not be valid, and that's okay.

// matchesPlatform checks if the given platform matches the required platforms.
// The given platform matches the required platform if
// - architecture and OS are identical.
// - OS version and variant are identical if provided.
// - features and OS features of the required platform are subsets of those of the given platform.
func matchesPlatform(given, required v1.Platform) bool {
	_ = "STUB: not implemented"
	// Required fields that must be identical.
	return false
}

// Optional fields that may be empty, but must be identical if provided.

// Verify required platform's features are a subset of given platform's features.

// isSubset checks if the required array of strings is a subset of the given lst.
func isSubset(lst, required []string) bool { _ = "STUB: not implemented"; return false }
