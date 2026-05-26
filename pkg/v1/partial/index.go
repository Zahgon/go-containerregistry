// Copyright 2020 Google LLC All Rights Reserved.
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

package partial

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

// FindManifests given a v1.ImageIndex, find the manifests that fit the matcher.
func FindManifests(index v1.ImageIndex, matcher match.Matcher) ([]v1.Descriptor, error) {
	_ = "STUB: not implemented"
	// get the actual manifest list
	return nil, nil
}

// try to get the root of our image

// FindImages given a v1.ImageIndex, find the images that fit the matcher. If a Descriptor
// matches the provider Matcher, but the referenced item is not an Image, ignores it.
// Only returns those that match the Matcher and are images.
func FindImages(index v1.ImageIndex, matcher match.Matcher) ([]v1.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if it is not an image, ignore it

// FindIndexes given a v1.ImageIndex, find the indexes that fit the matcher. If a Descriptor
// matches the provider Matcher, but the referenced item is not an Index, ignores it.
// Only returns those that match the Matcher and are indexes.
func FindIndexes(index v1.ImageIndex, matcher match.Matcher) ([]v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if it is not an index, ignore it

type withManifests interface {
	Manifests() ([]Describable, error)
}

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type describable struct {
	desc v1.Descriptor
}

func (d describable) Digest() (v1.Hash, error) {
	_ = "STUB: not implemented"
	return *new(v1.Hash), nil
}

func (d describable) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (d describable) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (d describable) Descriptor() (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil,

		// Manifests is analogous to v1.Image.Layers in that it allows values in the
		// returned list to be lazily evaluated, which enables an index to contain
		// an image that contains a streaming layer.
		//
		// This should have been part of the v1.ImageIndex interface, but wasn't.
		// It is instead usable through this extension interface.
		nil
}

func Manifests(idx v1.ImageIndex) ([]Describable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ComputeManifests provides a fallback implementation for Manifests.
func ComputeManifests(idx v1.ImageIndex) ([]Describable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
