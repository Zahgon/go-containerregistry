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
	"context"
	"io"
	"sync"

	api "github.com/moby/moby/api/types/image"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/google/go-containerregistry/pkg/v1/types"
	specs "github.com/moby/docker-image-spec/specs-go/v1"
)

type image struct {
	ref          name.Reference
	opener       *imageOpener
	tarballImage v1.Image
	computed     bool
	id           *v1.Hash
	configFile   *v1.ConfigFile

	once sync.Once
	err  error
}

type imageOpener struct {
	ref name.Reference
	ctx context.Context

	bufferMode bufferMode
	client     Client

	once    sync.Once
	bytes   []byte
	tmpPath string
	err     error
}

func (i *imageOpener) saveImage() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (i *imageOpener) bufferedOpener() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	// Store the tarball in memory and return a new reader into the bytes each time we need to access something.
	return *new(io.ReadCloser), nil
}

// Wrap the bytes in a ReadCloser so it looks like an opened file.

func (i *imageOpener) fileBackedOpener() (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (i *imageOpener) opener() tarball.Opener {
	_ = "STUB: not implemented"
	return *new(tarball.Opener)
}

// Image provides access to an image reference from the Docker daemon,
// applying functional options to the underlying imageOpener before
// resolving the reference into a v1.Image.
func Image(ref name.Reference, options ...Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Eagerly fetch Image ID to ensure it actually exists.
// https://github.com/google/go-containerregistry/issues/1186

func (i *image) initialize() error {
	_ = "STUB: not implemented"
	// Don't re-initialize tarball if already initialized.
	return nil
}

func (i *image) compute() error {
	_ = "STUB: not implemented"
	// Don't re-compute if already computed.
	return nil
}

func (i *image) Layers() ([]v1.Layer, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *image) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}

func (i *image) Size() (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *image) ConfigName() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

func (i *image) ConfigFile() (*v1.ConfigFile, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *image) RawConfigFile() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// RawConfigFile cannot be generated from "docker inspect" because Docker Engine API returns serialized data,
// and formatting information of the raw config such as indent and prefix will be lost.

func (i *image) Digest() (v1.Hash, error) { _ = "STUB: not implemented"; return *new(v1.Hash), nil }

func (i *image) Manifest() (*v1.Manifest, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *image) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (i *image) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (i *image) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

func (i *image) configHistory(author string) ([]v1.History, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *image) diffIDs(rootFS api.RootFS) ([]v1.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *image) computeConfigFile(inspect api.InspectResponse) (*v1.ConfigFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (i *image) computeImageConfig(config *specs.DockerOCIImageConfig) v1.Config {
	_ = "STUB: not implemented"
	return *new(v1.Config)
}

//nolint:staticcheck // SA1019 this is erroneously deprecated, as windows uses it
