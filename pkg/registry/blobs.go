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

package registry

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Returns whether this url should be handled by the blob handler
// This is complicated because blob is indicated by the trailing path, not the leading path.
// https://github.com/opencontainers/distribution-spec/blob/master/spec.md#pulling-a-layer
// https://github.com/opencontainers/distribution-spec/blob/master/spec.md#pushing-a-layer
func isBlob(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// BlobHandler represents a minimal blob storage backend, capable of serving
// blob contents.
type BlobHandler interface {
	// Get gets the blob contents, or errNotFound if the blob wasn't found.
	Get(ctx context.Context, repo string, h v1.Hash) (io.ReadCloser, error)
}

// BlobStatHandler is an extension interface representing a blob storage
// backend that can serve metadata about blobs.
type BlobStatHandler interface {
	// Stat returns the size of the blob, or errNotFound if the blob wasn't
	// found, or redirectError if the blob can be found elsewhere.
	Stat(ctx context.Context, repo string, h v1.Hash) (int64, error)
}

// BlobPutHandler is an extension interface representing a blob storage backend
// that can write blob contents.
type BlobPutHandler interface {
	// Put puts the blob contents.
	//
	// The contents will be verified against the expected size and digest
	// as the contents are read, and an error will be returned if these
	// don't match. Implementations should return that error, or a wrapper
	// around that error, to return the correct error when these don't match.
	Put(ctx context.Context, repo string, h v1.Hash, rc io.ReadCloser) error
}

// BlobDeleteHandler is an extension interface representing a blob storage
// backend that can delete blob contents.
type BlobDeleteHandler interface {
	// Delete the blob contents.
	Delete(ctx context.Context, repo string, h v1.Hash) error
}

// redirectError represents a signal that the blob handler doesn't have the blob
// contents, but that those contents are at another location which registry
// clients should redirect to.
type redirectError struct {
	// Location is the location to find the contents.
	Location string

	// Code is the HTTP redirect status code to return to clients.
	Code int
}

type bytesCloser struct {
	*bytes.Reader
}

func (r *bytesCloser) Close() error { _ = "STUB: not implemented"; return nil }

func (e redirectError) Error() string { _ = "STUB: not implemented"; return "" }

// errNotFound represents an error locating the blob.
var errNotFound = errors.New("not found")

type memHandler struct {
	m    map[string][]byte
	lock sync.Mutex
}

func NewInMemoryBlobHandler() BlobHandler { _ = "STUB: not implemented"; return *new(BlobHandler) }

func (m *memHandler) Stat(_ context.Context, _ string, h v1.Hash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *memHandler) Get(_ context.Context, _ string, h v1.Hash) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (m *memHandler) Put(_ context.Context, _ string, h v1.Hash, rc io.ReadCloser) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *memHandler) Delete(_ context.Context, _ string, h v1.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// blobs
type blobs struct {
	blobHandler BlobHandler

	// Each upload gets a unique id that writes occur to until finalized.
	uploads map[string][]byte
	lock    sync.Mutex
	log     *log.Logger
}

func (b *blobs) handle(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	return nil
}

// Must have a path of form /v2/{name}/blobs/{upload,sha256:}

// It is weird that this is "target" instead of "service", but
// that's how the index math works out above.

// OCI Distribution spec §10.5 requires 202 Accepted for chunk uploads.

// OCI Distribution spec §10.5 requires 202 Accepted for chunk uploads.
