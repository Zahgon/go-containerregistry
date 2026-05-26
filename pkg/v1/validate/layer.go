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

package validate

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Layer validates that the values return by its methods are consistent with the
// contents returned by Compressed and Uncompressed.
func Layer(layer v1.Layer, opt ...Option) error { _ = "STUB: not implemented"; return nil }

type computedLayer struct {
	// Calculated from Compressed stream.
	digest v1.Hash
	size   int64
	diffid v1.Hash

	// Calculated from Uncompressed stream.
	uncompressedDiffid v1.Hash
	uncompressedSize   int64
}

func computeLayer(layer v1.Layer) (*computedLayer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Keep track of compressed digest.

// Everything read from compressed is written to digester to compute digest.

// Call io.Copy to write from the layer Reader through to the tarReader on
// the other side of the pipe.

// Now close the compressed reader, to flush the gzip stream
// and calculate digest/diffID/size. This will cause pr to
// return EOF which will cause readers of the Compressed stream
// to finish reading.

// Read the bytes through gzip.Reader to compute the DiffID.

// Ensure there aren't duplicate file paths.

// Discard any trailing padding that the tar.Reader doesn't consume.
