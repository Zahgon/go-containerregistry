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

package mutate

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// Rebase returns a new v1.Image where the oldBase in orig is replaced by newBase.
func Rebase(orig, oldBase, newBase v1.Image) (v1.Image, error) {
	_ = "STUB: not implemented"
	// Verify that oldBase's layers are present in orig, otherwise orig is
	// not based on oldBase at all.
	return *new(v1.Image), nil
}

// Stitch together an image that contains:
// - original image's config
// - new base image's os/arch properties
// - new base image's layers + top of original image's layers
// - new base image's history + top of original image's history

// Add new config properties from existing images.

// OS/Arch properties from new base

// Apply config properties to rebased.

// Get new base layers and config for history.

// Add new base layers.

// Add original layers above the old base.

// createAddendums makes a list of addendums from a history and layers starting from a specific history and layer
// indexes.
func createAddendums(startHistory, startLayer int, history []v1.History, layers []v1.Layer) []Addendum {
	_ = "STUB: not implemented"

	// History should be a superset of layers; empty layers (e.g. ENV statements) only exist in history.
	// They cannot be iterated identically but must be walked independently, only advancing the iterator for layers
	// when a history entry for a non-empty layer is seen.
	return nil
}

// In the event history was malformed or non-existent, append the remaining layers.
