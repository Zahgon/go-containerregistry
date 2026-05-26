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

package validate

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// Index validates that idx does not violate any invariants of the index format.
func Index(idx v1.ImageIndex, opt ...Option) error { _ = "STUB: not implemented"; return nil }

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

func validateChildren(idx v1.ImageIndex, opt ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Workaround for #819.

type withMediaType interface {
	MediaType() (types.MediaType, error)
}

func validateMediaType(i withMediaType, want types.MediaType) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIndexManifest(idx v1.ImageIndex) error { _ = "STUB: not implemented"; return nil }

func validatePlatform(img v1.Image, want *v1.Platform) error { _ = "STUB: not implemented"; return nil }

// If we got here, some features might be mismatched. Just add those...
