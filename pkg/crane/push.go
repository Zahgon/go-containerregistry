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

package crane

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Load reads the tarball at path as a v1.Image.
func Load(path string, opt ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// LoadTag reads a tag from the tarball at path as a v1.Image.
// If tag is "", will attempt to read the tarball as a single image.
func LoadTag(path, tag string, opt ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Push pushes the v1.Image img to a registry as dst.
func Push(img v1.Image, dst string, opt ...Option) error { _ = "STUB: not implemented"; return nil }

// Upload pushes the v1.Layer to a given repo.
func Upload(layer v1.Layer, repo string, opt ...Option) error {
	_ = "STUB: not implemented"
	return nil
}
