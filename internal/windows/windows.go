// Copyright 2021 Google LLC All Rights Reserved.
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

package windows

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// userOwnerAndGroupSID is a magic value needed to make the binary executable
// in a Windows container.
//
// owner: BUILTIN/Users group: BUILTIN/Users ($sddlValue="O:BUG:BU")
const userOwnerAndGroupSID = "AQAAgBQAAAAkAAAAAAAAAAAAAAABAgAAAAAABSAAAAAhAgAAAQIAAAAAAAUgAAAAIQIAAA=="

// Windows returns a Layer that is converted to be pullable on Windows.
func Windows(layer v1.Layer) (v1.Layer, error) {
	_ = "STUB: not implemented"
	// TODO: do this lazily.
	return *new(v1.Layer), nil
}

// Use a fixed Mode, so that this isn't sensitive to the directory and umask
// under which it was created. Additionally, windows can only set 0222,
// 0444, or 0666, none of which are executable.

// TODO: this seems to make the file executable on Windows;
// only do this if the file should be executable.

// gzip the contents, then create the layer
