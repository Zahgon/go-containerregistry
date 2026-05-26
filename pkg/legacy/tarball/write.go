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

package tarball

import (
	"archive/tar"
	"io"

	"github.com/google/go-containerregistry/pkg/legacy"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// repositoriesTarDescriptor represents the repositories file inside a `docker save` tarball.
type repositoriesTarDescriptor map[string]map[string]string

// v1Layer represents a layer with metadata needed by the v1 image spec https://github.com/moby/moby/blob/master/image/spec/v1.md.
type v1Layer struct {
	// config is the layer metadata.
	config *legacy.LayerConfigFile
	// layer is the v1.Layer object this v1Layer represents.
	layer v1.Layer
}

// json returns the raw bytes of the json metadata of the given v1Layer.
func (l *v1Layer) json() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// version returns the raw bytes of the "VERSION" file of the given v1Layer.
func (l *v1Layer) version() []byte { _ = "STUB: not implemented"; return nil }

// v1LayerID computes the v1 image format layer id for the given v1.Layer with the given v1 parent ID and raw image config.
func v1LayerID(layer v1.Layer, parentID string, rawConfig []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// newTopV1Layer creates a new v1Layer for a layer other than the top layer in a v1 image tarball.
func newV1Layer(layer v1.Layer, parent *v1Layer, history v1.History) (*v1Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newTopV1Layer creates a new v1Layer for the top layer in a v1 image tarball.
func newTopV1Layer(layer v1.Layer, parent *v1Layer, history v1.History, imgConfig *v1.ConfigFile, rawConfig []byte) (*v1Layer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:staticcheck // Field will be removed in next release

// splitTag splits the given tagged image name <registry>/<repository>:<tag>
// into <registry>/<repository> and <tag>.
func splitTag(name string) (string, string) {
	_ = "STUB: not implemented"
	// Split on ":"
	return "", ""
}

// Verify that we aren't confusing a tag for a hostname w/ port for the purposes of weak validation.

// addTags adds the given image tags to the given "repositories" file descriptor in a v1 image tarball.
func addTags(repos repositoriesTarDescriptor, tags []string, topLayerID string) {
	_ = "STUB: not implemented"
	return
}

// updateLayerSources updates the given layer digest to descriptor map with the descriptor of the given layer in the given image if it's an undistributable layer.
func updateLayerSources(layerSources map[v1.Hash]v1.Descriptor, layer v1.Layer, img v1.Image) error {
	_ = "STUB: not implemented"
	return nil
}

// Add to LayerSources if it's a foreign layer.

// Write is a wrapper to write a single image in V1 format and tag to a tarball.
func Write(ref name.Reference, img v1.Image, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// filterEmpty filters out the history corresponding to empty layers from the
// given history.
func filterEmpty(h []v1.History) []v1.History { _ = "STUB: not implemented"; return nil }

// MultiWrite writes the contents of each image to the provided reader, in the V1 image tarball format.
// The contents are written in the following format:
// One manifest.json file at the top level containing information about several images.
// One repositories file mapping from the image <registry>/<repo name> to <tag> to the id of the top most layer.
// For every layer, a directory named with the layer ID is created with the following contents:
//
//	layer.tar - The uncompressed layer tarball.
//	<layer id>.json- Layer metadata json.
//	VERSION- Schema version string. Always set to "1.0".
//
// One file for the config blob, named after its SHA.
func MultiWrite(refToImage map[name.Reference]v1.Image, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// Write the config.

// Store foreign layer info.

// Write the layers.

// Create a blank config history if the config didn't have a history.

// If the v1.Layer implements UncompressedSize efficiently, use that
// for the tar header. Otherwise, this iterates over Uncompressed().
// NOTE: If using a streaming layer, this may consume the layer.

// Generate the tar descriptor and write it.

// prev should be the top layer here. Use it to add the image tags
// to the tarball repositories file.

func dedupRefToImage(refToImage map[name.Reference]v1.Image) ([]v1.Image, map[v1.Image][]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Force specific order on tags

// Writes a file to the provided writer with a corresponding tar header
func writeTarEntry(tf *tar.Writer, path string, r io.Reader, size int64) error {
	_ = "STUB: not implemented"
	return nil
}
