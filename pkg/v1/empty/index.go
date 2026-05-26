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

package empty

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Index is a singleton empty index, think: FROM scratch.
var Index = emptyIndex{}

type emptyIndex struct{}

func (i emptyIndex) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i emptyIndex) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

func (i emptyIndex) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i emptyIndex) IndexManifest() (*v1.IndexManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i emptyIndex) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i emptyIndex) Image(v1.Hash) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func (i emptyIndex) ImageIndex(v1.Hash) (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

func base() *v1.IndexManifest { _ = "STUB: not implemented"; return nil }
