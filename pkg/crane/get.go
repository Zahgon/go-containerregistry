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
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func getImage(r string, opt ...Option) (v1.Image, name.Reference, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), *new(name.Reference), nil
}

func getManifest(r string, opt ...Option) (*remote.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get calls remote.Get and returns an uninterpreted response.
func Get(r string, opt ...Option) (*remote.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Head performs a HEAD request for a manifest and returns a content descriptor
	// based on the registry's response.
}

func Head(r string, opt ...Option) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
