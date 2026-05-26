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

// This is an EXPERIMENTAL package, and may change in arbitrary ways without notice.
package layout

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// GarbageCollect removes unreferenced blobs from the oci-layout
//
//	This is an experimental api, and not subject to any stability guarantees
//	We may abandon it at any time, without prior notice.
//	Deprecated: Use it at your own risk!
func (l Path) GarbageCollect() ([]v1.Hash, error) { _ = "STUB: not implemented"; return nil, nil }

func (l Path) garbageCollectImageIndex(index v1.ImageIndex, blobsToKeep map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (l Path) garbageCollectImage(image v1.Image, blobsToKeep map[string]bool) error {
	_ = "STUB: not implemented"
	return nil
}
