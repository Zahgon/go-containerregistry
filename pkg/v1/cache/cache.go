// Copyright 2021 Google LLC All Rights Reserved.
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

// Package cache provides methods to cache layers.
package cache

import (
	"errors"
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Cache encapsulates methods to interact with cached layers.
type Cache interface {
	// Put writes the Layer to the Cache.
	//
	// The returned Layer should be used for future operations, since lazy
	// cachers might only populate the cache when the layer is actually
	// consumed.
	//
	// The returned layer can be consumed, and the cache entry populated,
	// by calling either Compressed or Uncompressed and consuming the
	// returned io.ReadCloser.
	Put(v1.Layer) (v1.Layer, error)

	// Get returns the Layer cached by the given Hash, or ErrNotFound if no
	// such layer was found.
	Get(v1.Hash) (v1.Layer, error)

	// Delete removes the Layer with the given Hash from the Cache.
	Delete(v1.Hash) error
}

// ErrNotFound is returned by Get when no layer with the given Hash is found.
var ErrNotFound = errors.New("layer was not found")

// Image returns a new Image which wraps the given Image, whose layers will be
// pulled from the Cache if they are found, and written to the Cache as they
// are read from the underlying Image.
func Image(i v1.Image, c Cache) v1.Image { _ = "STUB: not implemented"; return *new(v1.Image) }

type image struct {
	v1.Image
	c Cache
}

func (i *image) Layers() ([]v1.Layer, error) { _ = "STUB: not implemented"; return nil, nil }

type lazyLayer struct {
	inner v1.Layer
	c     Cache
}

func (l *lazyLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Layer found in the cache.

// Not cached, pull and return the real layer.

func (l *lazyLayer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Layer found in the cache.

// Not cached, pull and return the real layer.

func (l *lazyLayer) Size() (int64, error)     { _ = "STUB: not implemented"; return 0, nil }
func (l *lazyLayer) DiffID() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }
func (l *lazyLayer) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }
func (l *lazyLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *image) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Not cached, get it and write it.

func (i *image) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Not cached, get it and write it.

// ImageIndex returns a new ImageIndex which wraps the given ImageIndex's
// children with either Image(child, c) or ImageIndex(child, c) depending on type.
func ImageIndex(ii v1.ImageIndex, c Cache) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}

type imageIndex struct {
	inner v1.ImageIndex
	c     Cache
}

func (ii *imageIndex) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}
func (ii *imageIndex) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}
func (ii *imageIndex) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }
func (ii *imageIndex) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
func (ii *imageIndex) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ii *imageIndex) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (ii *imageIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}
