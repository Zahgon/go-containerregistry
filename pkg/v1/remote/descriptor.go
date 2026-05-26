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

package remote

import (
	"context"
	"errors"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

var allManifestMediaTypes = append(append([]types.MediaType{
	types.DockerManifestSchema1,
	types.DockerManifestSchema1Signed,
}, acceptableImageMediaTypes...), acceptableIndexMediaTypes...)

// ErrSchema1 indicates that we received a schema1 manifest from the registry.
// This library doesn't have plans to support this legacy image format:
// https://github.com/google/go-containerregistry/issues/377
var ErrSchema1 = errors.New("see https://github.com/google/go-containerregistry/issues/377")

// newErrSchema1 returns an ErrSchema1 with the unexpected MediaType.
func newErrSchema1(schema types.MediaType) error { _ = "STUB: not implemented"; return nil }

// Descriptor provides access to metadata about remote artifact and accessors
// for efficiently converting it into a v1.Image or v1.ImageIndex.
type Descriptor struct {
	fetcher fetcher
	v1.Descriptor

	ref      name.Reference
	Manifest []byte
	ctx      context.Context

	// So we can share this implementation with Image.
	platform v1.Platform
}

func (d *Descriptor) toDesc() v1.Descriptor {
	_ = "STUB: not implemented"
	return *

	// RawManifest exists to satisfy the Taggable interface.
	new(v1.Descriptor)
}

func (d *Descriptor) RawManifest() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Get returns a remote.Descriptor for the given reference. The response from
		// the registry is left un-interpreted, for the most part. This is useful for
		// querying what kind of artifact a reference represents.
		//
		// See Head if you don't need the response body.
		nil
}

func Get(ref name.Reference, options ...Option) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Head returns a v1.Descriptor for the given reference by issuing a HEAD
// request.
//
// Note that the server response will not have a body, so any errors encountered
// should be retried with Get to get more details.
func Head(ref name.Reference, options ...Option) (*v1.Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle options and fetch the manifest with the acceptable MediaTypes in the
// Accept header.
func get(ref name.Reference, acceptable []types.MediaType, options ...Option) (*Descriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Image converts the Descriptor into a v1.Image.
//
// If the fetched artifact is already an image, it will just return it.
//
// If the fetched artifact is an index, it will attempt to resolve the index to
// a child image with the appropriate platform.
//
// See WithPlatform to set the desired platform.
func (d *Descriptor) Image() (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// We don't care to support schema 1 images:
// https://github.com/google/go-containerregistry/issues/377

// We want an image but the registry has an index, resolve it to an image.

// These are expected. Enumerated here to allow a default case.

// We could just return an error here, but some registries (e.g. static
// registries) don't set the Content-Type headers correctly, so instead...

// Wrap the v1.Layers returned by this v1.Image in a hint for downstream
// remote.Write calls to facilitate cross-repo "mounting".

// Schema1 converts the Descriptor into a v1.Image for v2 schema 1 media types.
//
// The v1.Image returned by this method does not implement the entire interface because it would be inefficient.
// This exists mostly to make it easier to copy schema 1 images around or look at their filesystems.
// This is separate from Image() to avoid a backward incompatible change for callers expecting ErrSchema1.
func (d *Descriptor) Schema1() (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// ImageIndex converts the Descriptor into a v1.ImageIndex.
func (d *Descriptor) ImageIndex() (v1.ImageIndex, error) {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex), nil
}

// We don't care to support schema 1 images:
// https://github.com/google/go-containerregistry/issues/377

// We want an index but the registry has an image, nothing we can do.

// These are expected.

// We could just return an error here, but some registries (e.g. static
// registries) don't set the Content-Type headers correctly, so instead...

func (d *Descriptor) remoteImage() *remoteImage { _ = "STUB: not implemented"; return nil }

func (d *Descriptor) remoteIndex() *remoteIndex { _ = "STUB: not implemented"; return nil }
