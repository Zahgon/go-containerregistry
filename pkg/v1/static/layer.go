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

package static

import (
	"io"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// NewLayer returns a layer containing the given bytes, with the given mediaType.
//
// Contents will not be compressed.
func NewLayer(b []byte, mt types.MediaType) v1.Layer {
	_ = "STUB: not implemented"
	return *new(v1.Layer)
}

type staticLayer struct {
	b  []byte
	mt types.MediaType

	once sync.Once
	h    v1.Hash
}

func (l *staticLayer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"

	// Only calculate digest the first time we're asked.
	return *new(v1.Hash), nil
}

func (l *staticLayer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (l *staticLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (l *staticLayer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (l *staticLayer) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *staticLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}
