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
	"os"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

func isWindows(img v1.Image) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// Append reads a layer from path and appends it the the v1.Image base.
//
// If the base image is a Windows base image (i.e., its config.OS is
// "windows"), the contents of the tarballs will be modified to be suitable for
// a Windows container image.`,
func Append(base v1.Image, paths ...string) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

func getLayer(path string, layerType types.MediaType) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

// This is dumb but the tarball package assumes things about mediaTypes that aren't true
// and doesn't have enough context to know what the right default is.

// If we're dealing with a named pipe, trying to open it multiple times will
// fail, so we need to do a streaming upload.
//
// returns nil, nil for non-streaming files
func streamFile(path string) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }
