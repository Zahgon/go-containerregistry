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

	"github.com/google/go-containerregistry/pkg/compression"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type layer struct {
	digest             v1.Hash
	diffID             v1.Hash
	size               int64
	compressedopener   Opener
	uncompressedopener Opener
	compression        compression.Compression
	compressionLevel   int
	annotations        map[string]string
	mediaType          types.MediaType
}

// Descriptor implements partial.withDescriptor.
func (l *layer) Descriptor() (*v1.Descriptor, error) { _ = "STUB: not implemented"; return nil, nil }

// Digest implements v1.Layer
func (l *layer) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// DiffID implements v1.Layer
	new(v1.Hash), nil
}

func (l *layer) DiffID() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *

	// Compressed implements v1.Layer
	new(v1.Hash), nil
}

func (l *layer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Uncompressed implements v1.Layer
func (l *layer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Size implements v1.Layer
func (l *layer) Size() (int64, error) {
	_ = "STUB: not implemented"

	// MediaType implements v1.Layer
	return 0, nil
}

func (l *layer) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *

	// LayerOption applies options to layer
	new(types.MediaType), nil
}

type LayerOption func(*layer)

// WithCompression is a functional option for overriding the default
// compression algorithm used for compressing uncompressed tarballs.
// Please note that WithCompression(compression.ZStd) should be used
// in conjunction with WithMediaType(types.OCILayerZStd)
func WithCompression(comp compression.Compression) LayerOption {
	_ = "STUB: not implemented"
	return *new(LayerOption)
}

// WithCompressionLevel is a functional option for overriding the default
// compression level used for compressing uncompressed tarballs.
func WithCompressionLevel(level int) LayerOption {
	_ = "STUB: not implemented"
	return *new(LayerOption)
}

// WithMediaType is a functional option for overriding the layer's media type.
func WithMediaType(mt types.MediaType) LayerOption {
	_ = "STUB: not implemented"
	return *new(LayerOption)
}

// WithCompressedCaching is a functional option that overrides the
// logic for accessing the compressed bytes to memoize the result
// and avoid expensive repeated gzips.
func WithCompressedCaching(l *layer) { _ = "STUB: not implemented"; return }

// WithEstargzOptions is a functional option that allow the caller to pass
// through estargz.Options to the underlying compression layer.  This is
// only meaningful when estargz is enabled.
//
// Deprecated: WithEstargz is deprecated; it is a no-op.
func WithEstargzOptions(...any) LayerOption {
	_ = "STUB: not implemented"
	return *

	// WithEstargz is a functional option that explicitly enables estargz support.
	//
	// Deprecated: WithEstargz is deprecated; it is a no-op.
	new(LayerOption)
}

func WithEstargz(*layer) {
	_ = "STUB: not implemented"

	// LayerFromFile returns a v1.Layer given a tarball
	return
}

func LayerFromFile(path string, opts ...LayerOption) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// LayerFromOpener returns a v1.Layer given an Opener function.
// The Opener may return either an uncompressed tarball (common),
// or a compressed tarball (uncommon).
//
// When using this in conjunction with something like remote.Write
// the uncompressed path may end up gzipping things multiple times:
//  1. Compute the layer SHA256
//  2. Upload the compressed layer.
//
// Since gzip can be expensive, we support an option to memoize the
// compression that can be passed here: tarball.WithCompressedCaching
func LayerFromOpener(opener Opener, opts ...LayerOption) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Warn if media type does not match compression

// LayerFromReader returns a v1.Layer given a io.Reader.
//
// The reader's contents are read and buffered to a temp file in the process.
//
// Deprecated: Use LayerFromOpener or stream.NewLayer instead, if possible.
func LayerFromReader(reader io.Reader, opts ...LayerOption) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func computeDigest(opener Opener) (v1.Hash, int64, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), 0, nil
}

func computeDiffID(opener Opener) (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}
