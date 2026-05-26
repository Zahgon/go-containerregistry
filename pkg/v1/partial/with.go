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

package partial

import (
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// WithRawConfigFile defines the subset of v1.Image used by these helper methods
type WithRawConfigFile interface {
	// RawConfigFile returns the serialized bytes of this image's config file.
	RawConfigFile() ([]byte, error)
}

// ConfigFile is a helper for implementing v1.Image
func ConfigFile(i WithRawConfigFile) (*v1.ConfigFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigName is a helper for implementing v1.Image
func ConfigName(i WithRawConfigFile) (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

type configLayer struct {
	hash    v1.Hash
	content []byte
}

// Digest implements v1.Layer
func (cl *configLayer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// DiffID implements v1.Layer
	new(v1.Hash), nil
}

func (cl *configLayer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Uncompressed implements v1.Layer
	new(v1.Hash), nil
}

func (cl *configLayer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Compressed implements v1.Layer
func (cl *configLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Size implements v1.Layer
func (cl *configLayer) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (cl *configLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	// Defaulting this to OCIConfigJSON as it should remain
	// backwards compatible with DockerConfigJSON
	return *new(types.MediaType), nil
}

var _ v1.Layer = (*configLayer)(nil)

// withConfigLayer allows partial image implementations to provide a layer
// for their config file.
type withConfigLayer interface {
	ConfigLayer() (v1.Layer, error)
}

// ConfigLayer implements v1.Layer from the raw config bytes.
// This is so that clients (e.g. remote) can access the config as a blob.
//
// Images that want to return a specific layer implementation can implement
// withConfigLayer.
func ConfigLayer(i WithRawConfigFile) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// WithConfigFile defines the subset of v1.Image used by these helper methods
type WithConfigFile interface {
	// ConfigFile returns this image's config file.
	ConfigFile() (*v1.ConfigFile, error)
}

// DiffIDs is a helper for implementing v1.Image
func DiffIDs(i WithConfigFile) ([]v1.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

// RawConfigFile is a helper for implementing v1.Image
func RawConfigFile(i WithConfigFile) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// WithRawManifest defines the subset of v1.Image used by these helper methods
type WithRawManifest interface {
	// RawManifest returns the serialized bytes of this image's config file.
	RawManifest() ([]byte, error)
}

// Digest is a helper for implementing v1.Image
func Digest(i WithRawManifest) (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// Manifest is a helper for implementing v1.Image
func Manifest(i WithRawManifest) (*v1.Manifest, error) { _ = "STUB: not implemented"; return nil, nil }

// WithManifest defines the subset of v1.Image used by these helper methods
type WithManifest interface {
	// Manifest returns this image's Manifest object.
	Manifest() (*v1.Manifest, error)
}

// RawManifest is a helper for implementing v1.Image
func RawManifest(i WithManifest) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Size is a helper for implementing v1.Image
func Size(i WithRawManifest) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// FSLayers is a helper for implementing v1.Image
func FSLayers(i WithManifest) ([]v1.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

// BlobSize is a helper for implementing v1.Image
func BlobSize(i WithManifest, h v1.Hash) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// BlobDescriptor is a helper for implementing v1.Image
func BlobDescriptor(i WithManifest, h v1.Hash) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithManifestAndConfigFile defines the subset of v1.Image used by these helper methods
type WithManifestAndConfigFile interface {
	WithConfigFile

	// Manifest returns this image's Manifest object.
	Manifest() (*v1.Manifest, error)
}

// BlobToDiffID is a helper for mapping between compressed
// and uncompressed blob hashes.
func BlobToDiffID(i WithManifestAndConfigFile, h v1.Hash) (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// DiffIDToBlob is a helper for mapping between uncompressed
// and compressed blob hashes.
func DiffIDToBlob(wm WithManifestAndConfigFile, h v1.Hash) (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// WithDiffID defines the subset of v1.Layer for exposing the DiffID method.
type WithDiffID interface {
	DiffID() (v1.Hash, error)
}

// withDescriptor allows partial layer implementations to provide a layer
// descriptor to the partial image manifest builder. This allows partial
// uncompressed layers to provide foreign layer metadata like URLs to the
// uncompressed image manifest.
type withDescriptor interface {
	Descriptor() (*v1.Descriptor, error)
}

// Describable represents something for which we can produce a v1.Descriptor.
type Describable interface {
	Digest() (v1.Hash, error)
	MediaType() (types.MediaType, error)
	Size() (int64, error)
}

// Descriptor returns a v1.Descriptor given a Describable. It also encodes
// some logic for unwrapping things that have been wrapped by
// CompressedToLayer, UncompressedToLayer, CompressedToImage, or
// UncompressedToImage.
func Descriptor(d Describable) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	// If Describable implements Descriptor itself, return that.
	return nil, nil
}

// If all else fails, compute the descriptor from the individual methods.

// Failing to parse as a manifest should just be ignored.
// The manifest might not be valid, and that's okay.

type withArtifactType interface {
	ArtifactType() (string, error)
}

type withUncompressedSize interface {
	UncompressedSize() (int64, error)
}

// UncompressedSize returns the size of the Uncompressed layer. If the
// underlying implementation doesn't implement UncompressedSize directly,
// this will compute the uncompressedSize by reading everything returned
// by Compressed(). This is potentially expensive and may consume the contents
// for streaming layers.
func UncompressedSize(l v1.Layer) (int64, error) {
	_ = "STUB: not implemented"
	// If the layer implements UncompressedSize itself, return that.
	return 0, nil
}

// The layer doesn't implement UncompressedSize, we need to compute it.

type withExists interface {
	Exists() (bool, error)
}

// Exists checks to see if a layer exists. This is a hack to work around the
// mistakes of the partial package. Don't use this.
func Exists(l v1.Layer) (bool, error) {
	_ = "STUB: not implemented"
	// If the layer implements Exists itself, return that.
	return false, nil
}

// The layer doesn't implement Exists, so we hope that calling Compressed()
// is enough to trigger an error if the layer does not exist.

// We may want to try actually reading a single byte, but if we need to do
// that, we should just fix this hack.

// Recursively unwrap our wrappers so that we can check for the original implementation.
// We might want to expose this?
func unwrap(i any) any { _ = "STUB: not implemented"; return *new(any) }

// ArtifactType returns the artifact type for the given manifest.
//
// If the manifest reports its own artifact type, that's returned, otherwise
// the manifest is parsed and, if successful, its config.mediaType is returned.
func ArtifactType(w WithManifest) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Failing to parse as a manifest should just be ignored.
// The manifest might not be valid, and that's okay.
