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

package tarball

import (
	"io"
	"sync"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type image struct {
	opener        Opener
	manifest      *Manifest
	config        []byte
	imgDescriptor *Descriptor

	tag *name.Tag
}

type uncompressedImage struct {
	*image
}

type compressedImage struct {
	*image
	manifestLock sync.Mutex // Protects manifest
	manifest     *v1.Manifest
}

var _ partial.UncompressedImageCore = (*uncompressedImage)(nil)
var _ partial.CompressedImageCore = (*compressedImage)(nil)

// Opener is a thunk for opening a tar file.
type Opener func() (io.ReadCloser, error)

func pathOpener(path string) Opener { _ = "STUB: not implemented"; return *new(Opener) }

// ImageFromPath returns a v1.Image from a tarball located on path.
func ImageFromPath(path string, tag *name.Tag) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// LoadManifest load manifest
func LoadManifest(opener Opener) (Manifest, error) {
	_ = "STUB: not implemented"
	return *new(Manifest), nil
}

// Image exposes an image from the tarball at the provided path.
func Image(opener Opener, tag *name.Tag) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Peek at the first layer and see if it's compressed.

func (i *image) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// Descriptor stores the manifest data for a single image inside a `docker save` tarball.
type Descriptor struct {
	Config   string
	RepoTags []string
	Layers   []string

	// Tracks foreign layer info. Key is DiffID.
	LayerSources map[v1.Hash]v1.Descriptor `json:",omitempty"`
}

// Manifest represents the manifests of all images as the `manifest.json` file in a `docker save` tarball.
type Manifest []Descriptor

func (m Manifest) findDescriptor(tag *name.Tag) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compare the resolved names, since there are several ways to specify the same tag.

func (i *image) areLayersCompressed() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (i *image) loadTarDescriptorAndConfig() error { _ = "STUB: not implemented"; return nil }

func (i *image) RawConfigFile() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// tarFile represents a single file inside a tar. Closing it closes the tar itself.
		nil
}

type tarFile struct {
	io.Reader
	io.Closer
}

func extractFileFromTar(opener Opener, filePath string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func followLinks(opener Opener, filePath string, visited map[string]bool) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// uncompressedLayerFromTarball implements partial.UncompressedLayer
type uncompressedLayerFromTarball struct {
	diffID    v1.Hash
	mediaType types.MediaType
	opener    Opener
	filePath  string
}

// foreignUncompressedLayer implements partial.UncompressedLayer but returns
// a custom descriptor. This allows the foreign layer URLs to be included in
// the generated image manifest for uncompressed layers.
type foreignUncompressedLayer struct {
	uncompressedLayerFromTarball
	desc v1.Descriptor
}

func (fl *foreignUncompressedLayer) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil,

		// DiffID implements partial.UncompressedLayer
		nil
}

func (ulft *uncompressedLayerFromTarball) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Uncompressed implements partial.UncompressedLayer
	new(v1.Hash), nil
}

func (ulft *uncompressedLayerFromTarball) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (ulft *uncompressedLayerFromTarball) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *uncompressedImage) LayerByDiffID(h v1.Hash) (partial.UncompressedLayer, error) {
	_ = "STUB: not implemented"
	return *new(partial.UncompressedLayer), nil
}

// Technically the media type should be 'application/tar' but given that our
// v1.Layer doesn't force consumers to care about whether the layer is compressed
// we should be fine returning the DockerLayer media type

// This is janky, but we don't want to implement Descriptor for
// uncompressed layers because it breaks a bunch of assumptions in partial.
// See https://github.com/google/go-containerregistry/issues/1870

// Overwrite the mediaType for foreign layers.

// Intentional fall through.

func (c *compressedImage) Manifest() (*v1.Manifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If it's a foreign layer, just append the descriptor so we can avoid
// reading the entire file.

func (c *compressedImage) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// compressedLayerFromTarball implements partial.CompressedLayer
type compressedLayerFromTarball struct {
	desc     v1.Descriptor
	opener   Opener
	filePath string
}

// Digest implements partial.CompressedLayer
func (clft *compressedLayerFromTarball) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

// Compressed implements partial.CompressedLayer
func (clft *compressedLayerFromTarball) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// MediaType implements partial.CompressedLayer
func (clft *compressedLayerFromTarball) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// Size implements partial.CompressedLayer
func (clft *compressedLayerFromTarball) Size() (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *compressedImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	_ = "STUB: not implemented"
	return *new(partial.CompressedLayer), nil
}
