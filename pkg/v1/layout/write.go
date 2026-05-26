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
	"os"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
)

var layoutFile = `{
    "imageLayoutVersion": "1.0.0"
}`

// renameMutex guards os.Rename calls in AppendImage on Windows only.
var renameMutex sync.Mutex

// AppendImage writes a v1.Image to the Path and updates
// the index.json to reference it.
func (l Path) AppendImage(img v1.Image, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AppendIndex writes a v1.ImageIndex to the Path and updates
// the index.json to reference it.
func (l Path) AppendIndex(ii v1.ImageIndex, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// AppendDescriptor adds a descriptor to the index.json of the Path.
func (l Path) AppendDescriptor(desc v1.Descriptor) error { _ = "STUB: not implemented"; return nil }

// ReplaceImage writes a v1.Image to the Path and updates
// the index.json to reference it, replacing any existing one that matches matcher, if found.
func (l Path) ReplaceImage(img v1.Image, matcher match.Matcher, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// ReplaceIndex writes a v1.ImageIndex to the Path and updates
// the index.json to reference it, replacing any existing one that matches matcher, if found.
func (l Path) ReplaceIndex(ii v1.ImageIndex, matcher match.Matcher, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// replaceDescriptor adds a descriptor to the index.json of the Path, replacing
// any one matching matcher, if found.
func (l Path) replaceDescriptor(appendable mutate.Appendable, matcher match.Matcher, options ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// RemoveDescriptors removes any descriptors that match the match.Matcher from the index.json of the Path.
func (l Path) RemoveDescriptors(matcher match.Matcher) error { _ = "STUB: not implemented"; return nil }

// WriteFile write a file with arbitrary data at an arbitrary location in a v1
// layout. Used mostly internally to write files like "oci-layout" and
// "index.json", also can be used to write other arbitrary files. Do *not* use
// this to write blobs. Use only WriteBlob() for that.
func (l Path) WriteFile(name string, data []byte, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteBlob copies a file to the blobs/ directory in the Path from the given ReadCloser at
// blobs/{hash.Algorithm}/{hash.Hex}.
func (l Path) WriteBlob(hash v1.Hash, r io.ReadCloser) error { _ = "STUB: not implemented"; return nil }

func (l Path) writeBlob(hash v1.Hash, size int64, rc io.ReadCloser, renamer func() (v1.Hash, error)) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if blob already exists and is the correct size

// If a renamer func was provided write to a temporary file

// Delete temp file if an error is encountered before renaming

// Write to file and exit if not renaming

// Always close reader before renaming, since Close computes the digest in
// the case of streaming layers. If Close is not called explicitly, it will
// occur in a goroutine that is not guaranteed to succeed before renamer is
// called. When renamer is the layer's Digest method, it can return
// ErrNotComputed.

// Always close file before renaming

// Rename file based on the final hash

// writeLayer writes the compressed layer to a blob. Unlike WriteBlob it will
// write to a temporary file (suffixed with .tmp) within the layout until the
// compressed reader is fully consumed and written to disk. Also unlike
// WriteBlob, it will not skip writing and exit without error when a blob file
// exists, but does not have the correct size. (The blob hash is not
// considered, because it may be expensive to compute.)
func (l Path) writeLayer(layer v1.Layer) error { _ = "STUB: not implemented"; return nil }

// Allow digest errors, since streams may not have calculated the hash
// yet. Instead, use an empty value, which will be transformed into a
// random file name with `os.CreateTemp` and the final digest will be
// calculated after writing to a temp file and before renaming to the
// final path.

// Allow size errors, since streams may not have calculated the size
// yet. Instead, use zero as a sentinel value meaning that no size
// comparison can be done and any sized blob file should be considered
// valid and not overwritten.
//
// TODO: Provide an option to always overwrite blobs.

// RemoveBlob removes a file from the blobs directory in the Path
// at blobs/{hash.Algorithm}/{hash.Hex}
// It does *not* remove any reference to it from other manifests or indexes, or
// from the root index.json.
func (l Path) RemoveBlob(hash v1.Hash) error { _ = "STUB: not implemented"; return nil }

// WriteImage writes an image, including its manifest, config and all of its
// layers, to the blobs directory. If any blob already exists, as determined by
// the hash filename, does not write it.
// This function does *not* update the `index.json` file. If you want to write the
// image and also update the `index.json`, call AppendImage(), which wraps this
// and also updates the `index.json`.
func (l Path) WriteImage(img v1.Image) error { _ = "STUB: not implemented"; return nil }

// Write the layers concurrently.

// Write the config.

// Write the img manifest.

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type withBlob interface {
	Blob(v1.Hash) (io.ReadCloser, error)
}

func (l Path) writeIndexToFile(indexFile string, ii v1.ImageIndex) error {
	_ = "STUB: not implemented"
	return nil
}

// Walk the descriptors and write any v1.Image or v1.ImageIndex that we find.
// If we come across something we don't expect, just write it as a blob.

// TODO: The layout could reference arbitrary things, which we should
// probably just pass through.

// Workaround for #819.

// WriteIndex writes an index to the blobs directory. Walks down the children,
// including its children manifests and/or indexes, and down the tree until all of
// config and all layers, have been written. If any blob already exists, as determined by
// the hash filename, does not write it.
// This function does *not* update the `index.json` file. If you want to write the
// index and also update the `index.json`, call AppendIndex(), which wraps this
// and also updates the `index.json`.
func (l Path) WriteIndex(ii v1.ImageIndex) error {
	_ = "STUB: not implemented"
	// Always just write oci-layout file, since it's small.
	return nil
}

// Write constructs a Path at path from an ImageIndex.
//
// The contents are written in the following format:
// At the top level, there is:
//
//	One oci-layout file containing the version of this image-layout.
//	One index.json file listing descriptors for the contained images.
//
// Under blobs/, there is, for each image:
//
//	One file for each layer, named after the layer's SHA.
//	One file for each config blob, named after its SHA.
//	One file for each manifest blob, named after its SHA.
func Write(path string, ii v1.ImageIndex) (Path, error) {
	_ = "STUB: not implemented"

	// Always just write oci-layout file, since it's small.
	return *new(Path), nil
}

// TODO create blobs/ in case there is a blobs file which would prevent the directory from being created
