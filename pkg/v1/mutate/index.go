// Copyright 2019 Google LLC All Rights Reserved.
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

package mutate

import (
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

func computeDescriptor(ia IndexAddendum) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The IndexAddendum allows overriding Descriptor values.

type index struct {
	base v1.ImageIndex
	adds []IndexAddendum
	// remove is removed before adds
	remove match.Matcher

	computed    bool
	manifest    *v1.IndexManifest
	annotations map[string]string
	mediaType   *types.MediaType
	imageMap    map[v1.Hash]v1.Image
	indexMap    map[v1.Hash]v1.ImageIndex
	layerMap    map[v1.Hash]v1.Layer
	subject     *v1.Descriptor

	sync.Mutex
}

var _ v1.ImageIndex = (*index)(nil)

func (i *index) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *index) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *index) compute() error { _ = "STUB: not implemented"; return nil }

// Don't re-compute if already computed.

func (i *index) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (i *index) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

// Workaround for #819.
func (i *index) Layer(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Digest returns the sha256 of this image's manifest.
func (i *index) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

// Manifest returns this image's Manifest object.
func (i *index) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RawManifest returns the serialized bytes of Manifest()
func (i *index) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *index) Manifests() ([]partial.Describable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Index contains a streamable layer which has not yet been
// consumed. Just return the manifests we have in case the caller
// is going to consume the streamable layers.
