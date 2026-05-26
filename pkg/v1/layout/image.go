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

package layout

import (
	"io"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type layoutImage struct {
	path         Path
	desc         v1.Descriptor
	manifestLock sync.Mutex // Protects rawManifest
	rawManifest  []byte
}

var _ partial.CompressedImageCore = (*layoutImage)(nil)

// Image reads a v1.Image with digest h from the Path.
func (l Path) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (li *layoutImage) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// Implements WithManifest for partial.Blobset.
func (li *layoutImage) Manifest() (*v1.Manifest, error) { _ = "STUB: not implemented"; return nil, nil }

func (li *layoutImage) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (li *layoutImage) RawConfigFile() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (li *layoutImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	_ = "STUB: not implemented"
	return *new(partial.CompressedLayer), nil
}

type compressedBlob struct {
	path Path
	desc v1.Descriptor
}

func (b *compressedBlob) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (b *compressedBlob) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (b *compressedBlob) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *compressedBlob) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// Descriptor implements partial.withDescriptor.
func (b *compressedBlob) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil,

		// See partial.Exists.
		nil
}

func (b *compressedBlob) Exists() (bool, error) { _ = "STUB: not implemented"; return false, nil }
