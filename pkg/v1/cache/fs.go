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

package cache

import (
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type fscache struct {
	path string
}

// NewFilesystemCache returns a Cache implementation backed by files.
func NewFilesystemCache(path string) Cache { _ = "STUB: not implemented"; return *new(Cache) }

func (fs *fscache) Put(l v1.Layer) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

type layer struct {
	v1.Layer
	path           string
	digest, diffID v1.Hash
}

func (l *layer) create(h v1.Hash) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (l *layer) Compressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (l *layer) Uncompressed() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

type readcloser struct {
	t      io.Reader
	closes []func() error
}

func (rc *readcloser) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rc *readcloser) Close() error {
	_ = "STUB: not implemented"
	// Call all Close methods, even if any returned an error. Return the
	// first returned error.
	return nil
}

func (fs *fscache) Get(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Delete and return ErrNotFound because the layer was incomplete.

func (fs *fscache) Delete(h v1.Hash) error { _ = "STUB: not implemented"; return nil }

func cachepath(path string, h v1.Hash) string { _ = "STUB: not implemented"; return "" }
