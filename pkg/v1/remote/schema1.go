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

package remote

import (
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type schema1 struct {
	ref        name.Reference
	ctx        context.Context
	fetcher    fetcher
	manifest   []byte
	mediaType  types.MediaType
	descriptor *v1.Descriptor
}

func (s *schema1) Layers() ([]v1.Layer, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *schema1) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (s *schema1) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *schema1) ConfigName() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (s *schema1) ConfigFile() (*v1.ConfigFile, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *schema1) RawConfigFile() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *schema1) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

func (s *schema1) Manifest() (*v1.Manifest, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *schema1) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *schema1) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (s *schema1) LayerByDiffID(v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

type fslayer struct {
	BlobSum string `json:"blobSum"`
}

type schema1Manifest struct {
	FSLayers []fslayer `json:"fsLayers"`
}
