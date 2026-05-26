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

package random

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type randomIndex struct {
	images   map[v1.Hash]v1.Image
	manifest *v1.IndexManifest
}

// Index returns a pseudo-randomly generated ImageIndex with count images, each
// having the given number of layers of size byteSize.
func Index(byteSize, layers, count int64, options ...Option) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

func (i *randomIndex) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *randomIndex) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (i *randomIndex) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *randomIndex) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *randomIndex) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *randomIndex) Image(h v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (i *randomIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	// This is a single level index (for now?).
	return *new(v1.ImageIndex), nil
}
