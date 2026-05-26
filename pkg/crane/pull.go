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

// Tag applied to images that were pulled by digest. This denotes that the
// image was (probably) never tagged with this, but lets us avoid applying the
// ":latest" tag which might be misleading.
const iWasADigestTag = "i-was-a-digest"

// Pull returns a v1.Image of the remote image src.
func Pull(src string, opt ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Save writes the v1.Image img as a tarball at path with tag src.
func Save(img v1.Image, src, path string) error { _ = "STUB: not implemented"; return nil }

// MultiSave writes collection of v1.Image img with tag as a tarball.
func MultiSave(imgMap map[string]v1.Image, path string, opt ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteToFile wants a tag to write to the tarball, but we might have
// been given a digest.
// If the original ref was a tag, use that. Otherwise, if it was a
// digest, tag the image with :i-was-a-digest instead.

// no progress channel (for now)

// PullLayer returns the given layer from a registry.
func PullLayer(ref string, opt ...Option) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// SaveLegacy writes the v1.Image img as a legacy tarball at path with tag src.
func SaveLegacy(img v1.Image, src, path string) error { _ = "STUB: not implemented"; return nil }

// MultiSaveLegacy writes collection of v1.Image img with tag as a legacy tarball.
func MultiSaveLegacy(imgMap map[string]v1.Image, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// SaveOCI writes the v1.Image img as an OCI Image Layout at path. If a layout
// already exists at that path, it will add the image to the index.
func SaveOCI(img v1.Image, path string) error { _ = "STUB: not implemented"; return nil }

// MultiSaveOCI writes collection of v1.Image img as an OCI Image Layout at path. If a layout
// already exists at that path, it will add the image to the index.
func MultiSaveOCI(imgMap map[string]v1.Image, path string) error {
	_ = "STUB: not implemented"
	return nil
}
