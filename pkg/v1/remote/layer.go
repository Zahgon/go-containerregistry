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

package remote

import (
	"context"
	"io"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// remoteImagelayer implements partial.CompressedLayer
type remoteLayer struct {
	ctx     context.Context
	fetcher fetcher
	digest  v1.Hash
}

// Compressed implements partial.CompressedLayer
func (rl *remoteLayer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	// We don't want to log binary layers -- this can break terminals.
	return *new(io.ReadCloser), nil
}

// Compressed implements partial.CompressedLayer
func (rl *remoteLayer) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Digest implements partial.CompressedLayer
func (rl *remoteLayer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// MediaType implements v1.Layer
	new(v1.Hash), nil
}

func (rl *remoteLayer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

// See partial.Exists.
func (rl *remoteLayer) Exists() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Layer reads the given blob reference from a registry as a Layer. A blob
// reference here is just a punned name.Digest where the digest portion is the
// digest of the blob to be read and the repository portion is the repo where
// that blob lives.
func Layer(ref name.Digest, options ...Option) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}
