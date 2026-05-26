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

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

var _ v1.ImageIndex = (*layoutIndex)(nil)

type layoutIndex struct {
	mediaType types.MediaType
	path      Path
	rawIndex  []byte
}

// ImageIndexFromPath is a convenience function which constructs a Path and returns its v1.ImageIndex.
func ImageIndexFromPath(path string) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

// ImageIndex returns a v1.ImageIndex for the Path.
func (l Path) ImageIndex() (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

func (i *layoutIndex) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *layoutIndex) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (i *layoutIndex) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *layoutIndex) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *layoutIndex) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *layoutIndex) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	// Look up the digest in our manifest first to return a better error.
	return *new(v1.Image), nil
}

func (i *layoutIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	// Look up the digest in our manifest first to return a better error.
	return *new(v1.ImageIndex), nil
}

func (i *layoutIndex) Blob(h v1.Hash) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (i *layoutIndex) findDescriptor(h v1.Hash) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: Pull this out into methods on types.MediaType? e.g. instead, have:
// * mt.IsIndex()
// * mt.IsImage()
func isExpectedMediaType(mt types.MediaType, expected ...types.MediaType) bool {
	_ = "STUB: not implemented"
	return false
}
