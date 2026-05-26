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
)

// Image validates that img does not violate any invariants of the image format.
func Image(img v1.Image, opt ...Option) error { _ = "STUB: not implemented"; return nil }

func validateConfig(img v1.Image) error { _ = "STUB: not implemented"; return nil }

func validateLayers(img v1.Image, opt ...Option) error { _ = "STUB: not implemented"; return nil }

// Errored while reading tar content of layer because a header or
// content section was not the correct length. This is most likely
// due to an incomplete download or otherwise interrupted process.

// Compute all of these first before we call Config() and Manifest() to allow
// for lazy access e.g. for stream.Layer.

func validateManifest(img v1.Image) error { _ = "STUB: not implemented"; return nil }

func layersExist(layers []v1.Layer) error { _ = "STUB: not implemented"; return nil }
