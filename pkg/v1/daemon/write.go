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

package daemon

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Tag adds a tag to an already existent image.
func Tag(src, dest name.Tag, options ...Option) error { _ = "STUB: not implemented"; return nil }

// Write saves the image into the daemon as the given tag.
func Write(tag name.Tag, img v1.Image, options ...Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If we already have this image by this image ID, we can skip loading it.

// If we already have this tag, we can skip tagging it.

// write the image in docker save format first, then load it
