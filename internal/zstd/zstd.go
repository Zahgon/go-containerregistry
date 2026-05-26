// Copyright 2022 Google LLC All Rights Reserved.
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

// Package zstd provides helper functions for interacting with zstd streams.
package zstd

import (
	"io"
)

// MagicHeader is the start of zstd files.
var MagicHeader = []byte{'\x28', '\xb5', '\x2f', '\xfd'}

// ReadCloser reads uncompressed input data from the io.ReadCloser and
// returns an io.ReadCloser from which compressed data may be read.
// This uses zstd level 1 for the compression.
func ReadCloser(r io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

// ReadCloserLevel reads uncompressed input data from the io.ReadCloser and
// returns an io.ReadCloser from which compressed data may be read.
func ReadCloserLevel(r io.ReadCloser, level int) io.ReadCloser {
	_ = "STUB: not implemented"
	return *

	// For highly compressible layers, zstd.Writer will output a very small
	// number of bytes per Write(). This is normally fine, but when pushing
	// to a registry, we want to ensure that we're taking full advantage of
	// the available bandwidth instead of sending tons of tiny writes over
	// the wire.
	// 64K ought to be small enough for anybody.
	new(io.ReadCloser)
}

// Returns err so we can pw.CloseWithError(err)

// TODO(go1.14): Just defer {pw,zw,r}.Close like you'd expect.
// Context: https://golang.org/issue/24283

// Close zstd writer to Flush it and write zstd trailers.

// Flush bufio writer to ensure we write out everything.

// We don't really care if these fail.

// UnzipReadCloser reads compressed input data from the io.ReadCloser and
// returns an io.ReadCloser from which uncompressed data may be read.
func UnzipReadCloser(r io.ReadCloser) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// If the unzip fails, then this seems to return the same
// error as the read.  We don't want this to interfere with
// us closing the main ReadCloser, since this could leave
// an open file descriptor (fails on Windows).

// Is detects whether the input stream is compressed.
func Is(r io.Reader) (bool, error) { _ = "STUB: not implemented"; return false, nil }
