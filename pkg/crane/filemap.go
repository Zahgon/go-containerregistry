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

// Layer creates a layer from a single file map. These layers are reproducible and consistent.
// A filemap is a path -> file content map representing a file system.
func Layer(filemap map[string][]byte) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// Return a new copy of the buffer each time it's opened.

// Image creates a image with the given filemaps as its contents. These images are reproducible and consistent.
// A filemap is a path -> file content map representing a file system.
func Image(filemap map[string][]byte) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}
