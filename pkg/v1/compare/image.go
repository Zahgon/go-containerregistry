// Copyright 2019 Google LLC All Rights Reserved.
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

package compare

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Images compares the given images to each other and returns an error if they
// differ.
func Images(a, b v1.Image) error { _ = "STUB: not implemented"; return nil }

// If we have fewer layers than the first image, abort with an error so we don't panic.

// Compare each layer.

// Wrap the error in newlines to delineate layer errors.
