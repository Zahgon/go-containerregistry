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
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// UncompressedLayer represents the bare minimum interface a natively
// uncompressed layer must implement for us to produce a v1.Layer
type UncompressedLayer interface {
	// DiffID returns the Hash of the uncompressed layer.
	DiffID() (v1.Hash, error)

	// Uncompressed returns an io.ReadCloser for the uncompressed layer contents.
	Uncompressed() (io.ReadCloser, error)

	// Returns the mediaType for the compressed Layer
	MediaType() (types.MediaType, error)
}

// uncompressedLayerExtender implements v1.Image using the uncompressed base properties.
type uncompressedLayerExtender struct {
	UncompressedLayer
	// Memoize size/hash so that the methods aren't twice as
	// expensive as doing this manually.
	hash          v1.Hash
	size          int64
	hashSizeError error
	once          sync.Once
}

// Compressed implements v1.Layer
func (ule *uncompressedLayerExtender) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Digest implements v1.Layer
func (ule *uncompressedLayerExtender) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// Size implements v1.Layer
func (ule *uncompressedLayerExtender) Size() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ule *uncompressedLayerExtender) calcSizeHash() { _ = "STUB: not implemented"; return }

// UncompressedToLayer fills in the missing methods from an UncompressedLayer so that it implements v1.Layer
func UncompressedToLayer(ul UncompressedLayer) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// UncompressedImageCore represents the bare minimum interface a natively
// uncompressed image must implement for us to produce a v1.Image
type UncompressedImageCore interface {
	ImageCore

	// LayerByDiffID is a variation on the v1.Image method, which returns
	// an UncompressedLayer instead.
	LayerByDiffID(v1.Hash) (UncompressedLayer, error)
}

// UncompressedToImage fills in the missing methods from an UncompressedImageCore so that it implements v1.Image.
func UncompressedToImage(uic UncompressedImageCore) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// uncompressedImageExtender implements v1.Image by extending UncompressedImageCore with the
// appropriate methods computed from the minimal core.
type uncompressedImageExtender struct {
	UncompressedImageCore

	lock     sync.Mutex
	manifest *v1.Manifest
}

// Assert that our extender type completes the v1.Image interface
var _ v1.Image = (*uncompressedImageExtender)(nil)

// Digest implements v1.Image
func (i *uncompressedImageExtender) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"

	// Manifest implements v1.Image
	return *new(v1.Hash), nil
}

func (i *uncompressedImageExtender) Manifest() (*v1.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RawManifest implements v1.Image
func (i *uncompressedImageExtender) RawManifest() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Size implements v1.Image
		nil
}

func (i *uncompressedImageExtender) Size() (int64, error) {
	_ = "STUB: not implemented"

	// ConfigName implements v1.Image
	return 0, nil
}

func (i *uncompressedImageExtender) ConfigName() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// ConfigFile implements v1.Image
	new(v1.Hash), nil
}

func (i *uncompressedImageExtender) ConfigFile() (*v1.ConfigFile, error) {
	_ = "STUB: not implemented"
	return nil,

		// Layers implements v1.Image
		nil
}

func (i *uncompressedImageExtender) Layers() ([]v1.Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LayerByDiffID implements v1.Image
func (i *uncompressedImageExtender) LayerByDiffID(diffID v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// LayerByDigest implements v1.Image
func (i *uncompressedImageExtender) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}
