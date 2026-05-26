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
	"io"
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

var acceptableImageMediaTypes = []types.MediaType{
	types.DockerManifestSchema2,
	types.OCIManifestSchema1,
}

// remoteImage accesses an image from a remote registry
type remoteImage struct {
	fetcher      fetcher
	ref          name.Reference
	ctx          context.Context
	manifestLock sync.Mutex // Protects manifest
	manifest     []byte
	configLock   sync.Mutex // Protects config
	config       []byte
	mediaType    types.MediaType
	descriptor   *v1.Descriptor
}

func (r *remoteImage) ArtifactType() (string, error) {
	_ = "STUB: not implemented"
	// kind of a hack, but RawManifest does appropriate locking/memoization
	// and makes sure r.descriptor is populated.
	return "", nil
}

var _ partial.CompressedImageCore = (*remoteImage)(nil)

// Image provides access to a remote image reference.
func Image(ref name.Reference, options ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (r *remoteImage) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (r *remoteImage) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NOTE(jonjohnsonjr): We should never get here because the public entrypoints
// do type-checking via remote.Descriptor. I've left this here for tests that
// directly instantiate a remoteImage.

func (r *remoteImage) RawConfigFile() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Descriptor retains the original descriptor from an index manifest.
// See partial.Descriptor.
func (r *remoteImage) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	// kind of a hack, but RawManifest does appropriate locking/memoization
	// and makes sure r.descriptor is populated.
	return nil, nil
}

func (r *remoteImage) ConfigLayer() (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// remoteImageLayer implements partial.CompressedLayer
type remoteImageLayer struct {
	ri     *remoteImage
	ctx    context.Context
	digest v1.Hash
}

// Digest implements partial.CompressedLayer
func (rl *remoteImageLayer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Compressed implements partial.CompressedLayer
	new(v1.Hash), nil
}

func (rl *remoteImageLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Add alternative layer sources from URLs (usually none).

// We don't want to log binary layers -- this can break terminals.

// The lastErr for most pulls will be the same (the first error), but for
// foreign layers we'll want to surface the last one, since we try to pull
// from the registry first, which would often fail.
// TODO: Maybe we don't want to try pulling from the registry first?

// Manifest implements partial.WithManifest so that we can use partial.BlobSize below.
func (rl *remoteImageLayer) Manifest() (*v1.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// MediaType implements v1.Layer
}

func (rl *remoteImageLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// Size implements partial.CompressedLayer
func (rl *remoteImageLayer) Size() (int64, error) {
	_ = "STUB: not implemented"
	// Look up the size of this digest in the manifest to avoid a request.
	return 0, nil
}

// ConfigFile implements partial.WithManifestAndConfigFile so that we can use partial.BlobToDiffID below.
func (rl *remoteImageLayer) ConfigFile() (*v1.ConfigFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DiffID implements partial.WithDiffID so that we don't recompute a DiffID that we already have
// available in our ConfigFile.
func (rl *remoteImageLayer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// Descriptor retains the original descriptor from an image manifest.
// See partial.Descriptor.
func (rl *remoteImageLayer) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See partial.Exists.
func (rl *remoteImageLayer) Exists() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// LayerByDigest implements partial.CompressedLayer
func (r *remoteImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	_ = "STUB: not implemented"
	return *new(partial.CompressedLayer), nil
}
