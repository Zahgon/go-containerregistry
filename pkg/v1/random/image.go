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
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// uncompressedLayer implements partial.UncompressedLayer from raw bytes.
type uncompressedLayer struct {
	diffID    v1.Hash
	mediaType types.MediaType
	content   []byte
}

// DiffID implements partial.UncompressedLayer
func (ul *uncompressedLayer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Uncompressed implements partial.UncompressedLayer
	new(v1.Hash), nil
}

func (ul *uncompressedLayer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// MediaType returns the media type of the layer
func (ul *uncompressedLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

var _ partial.UncompressedLayer = (*uncompressedLayer)(nil)

// Image returns a pseudo-randomly generated Image.
func Image(byteSize, layers int64, options ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Layer returns a layer with pseudo-randomly generated content.
func Layer(byteSize int64, mt types.MediaType, options ...Option) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

//nolint:gosec

// Hash the contents as we write it out to the buffer.

// Write a single file with a random name and random contents.
