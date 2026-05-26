// Copyright 2023 Google LLC All Rights Reserved.
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
	"context"
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type diskHandler struct {
	dir string
}

func NewDiskBlobHandler(dir string) BlobHandler {
	_ = "STUB: not implemented"
	return *new(BlobHandler)
}

func (m *diskHandler) blobHashPath(h v1.Hash) string { _ = "STUB: not implemented"; return "" }

func (m *diskHandler) Stat(_ context.Context, _ string, h v1.Hash) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *diskHandler) Get(_ context.Context, _ string, h v1.Hash) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (m *diskHandler) Put(_ context.Context, _ string, h v1.Hash, rc io.ReadCloser) error {
	_ = "STUB: not implemented"
	// Put the temp file in the same directory to avoid cross-device problems
	// during the os.Rename.  The filenames cannot conflict.
	return nil
}

func (m *diskHandler) Delete(_ context.Context, _ string, h v1.Hash) error {
	_ = "STUB: not implemented"
	return nil
}
