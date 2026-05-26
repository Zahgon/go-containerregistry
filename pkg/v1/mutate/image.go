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
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type image struct {
	base v1.Image
	adds []Addendum

	computed        bool
	configFile      *v1.ConfigFile
	manifest        *v1.Manifest
	annotations     map[string]string
	mediaType       *types.MediaType
	configMediaType *types.MediaType
	diffIDMap       map[v1.Hash]v1.Layer
	digestMap       map[v1.Hash]v1.Layer
	subject         *v1.Descriptor

	sync.Mutex
}

var _ v1.Image = (*image)(nil)

func (i *image) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// isImageConfig reports whether the media type is a Docker or OCI image config.
func isImageConfig(mt types.MediaType) bool { _ = "STUB: not implemented"; return false }

func (i *image) compute() error { _ = "STUB: not implemented"; return nil }

// Don't re-compute if already computed.

// Determine effective config media type (user override takes precedence).

// For image configs, update RootFS.DiffIDs and History from added layers.
// For artifacts, skip this: the config has no rootfs or history fields.

// For image configs, re-marshal the config and update the manifest digest.
// For artifacts, preserve the original config blob as-is to avoid
// corrupting the digest via re-marshaling.

// Layers returns the ordered collection of filesystem layers that comprise this image.
// The order of the list is oldest/base layer first, and most-recent/top layer last.
func (i *image) Layers() ([]v1.Layer, error) { _ = "STUB: not implemented"; return nil, nil }

// Stream not yet consumed, or non-image OCI artifact (RootFS.DiffIDs
// is empty so partial.DiffIDs returns nothing). Fall back to the base
// layers plus any added layers.

// Walk manifest layer descriptors by digest rather than rootfs diff
// IDs. Two layers can legitimately share a diff ID — same uncompressed
// content, different compression — and produce distinct digests. The
// manifest preserves the per-occurrence digest; LayerByDiffID does not,
// which previously caused duplicate-diff-ID layers to collapse to a
// single entry in the returned slice and break blob upload for
// downstream pushers (see #2034).

// ConfigName returns the hash of the image's config file.
func (i *image) ConfigName() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

// ConfigFile returns this image's config file.
func (i *image) ConfigFile() (*v1.ConfigFile, error) { _ = "STUB: not implemented"; return nil, nil }

// RawConfigFile returns the serialized bytes of ConfigFile().
// For non-image OCI artifacts, returns the original raw config to preserve
// the config blob digest.
func (i *image) RawConfigFile() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// If the manifest config is not a standard image config, return the
// original raw bytes to avoid corrupting the digest via re-marshaling.

// Digest returns the sha256 of this image's manifest.
func (i *image) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

// Size implements v1.Image.
func (i *image) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Manifest returns this image's Manifest object.
func (i *image) Manifest() (*v1.Manifest, error) { _ = "STUB: not implemented"; return nil, nil }

// RawManifest returns the serialized bytes of Manifest()
func (i *image) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// LayerByDigest returns a Layer for interacting with a particular layer of
// the image, looking it up by "digest" (the compressed hash).
func (i *image) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// LayerByDiffID is an analog to LayerByDigest, looking up by "diff id"
// (the uncompressed hash).
func (i *image) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func validate(adds []Addendum) error { _ = "STUB: not implemented"; return nil }
