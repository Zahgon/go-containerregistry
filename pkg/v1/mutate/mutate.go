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
	"archive/tar"
	"io"
	"time"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

const whiteoutPrefix = ".wh."

// Addendum contains layers and history to be appended
// to a base image
type Addendum struct {
	Layer       v1.Layer
	History     v1.History
	URLs        []string
	Annotations map[string]string
	MediaType   types.MediaType
}

// AppendLayers applies layers to a base image.
func AppendLayers(base v1.Image, layers ...v1.Layer) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Append will apply the list of addendums to the base image
func Append(base v1.Image, adds ...Addendum) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Appendable is an interface that represents something that can be appended
// to an ImageIndex. We need to be able to construct a v1.Descriptor in order
// to append something, and this is the minimum required information for that.
type Appendable interface {
	MediaType() (types.MediaType, error)
	Digest() (v1.Hash, error)
	Size() (int64, error)
}

// IndexAddendum represents an appendable thing and all the properties that
// we may want to override in the resulting v1.Descriptor.
type IndexAddendum struct {
	Add Appendable
	v1.Descriptor
}

// AppendManifests appends a manifest to the ImageIndex.
func AppendManifests(base v1.ImageIndex, adds ...IndexAddendum) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}

// RemoveManifests removes any descriptors that match the match.Matcher.
func RemoveManifests(base v1.ImageIndex, matcher match.Matcher) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}

// Config mutates the provided v1.Image to have the provided v1.Config
func Config(base v1.Image, cfg v1.Config) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Subject mutates the subject on an image or index manifest.
//
// The input is expected to be a v1.Image or v1.ImageIndex, and
// returns the same type. You can type-assert the result like so:
//
//	img := Subject(empty.Image, subj).(v1.Image)
//
// Or for an index:
//
//	idx := Subject(empty.Index, subj).(v1.ImageIndex)
//
// If the input is not an Image or ImageIndex, the result will
// attempt to lazily annotate the raw manifest.
func Subject(f partial.WithRawManifest, subject v1.Descriptor) partial.WithRawManifest {
	_ = "STUB: not implemented"
	return *new(partial.WithRawManifest)
}

// Annotations mutates the annotations on an annotatable image or index manifest.
//
// The annotatable input is expected to be a v1.Image or v1.ImageIndex, and
// returns the same type. You can type-assert the result like so:
//
//	img := Annotations(empty.Image, map[string]string{
//	    "foo": "bar",
//	}).(v1.Image)
//
// Or for an index:
//
//	idx := Annotations(empty.Index, map[string]string{
//	    "foo": "bar",
//	}).(v1.ImageIndex)
//
// If the input Annotatable is not an Image or ImageIndex, the result will
// attempt to lazily annotate the raw manifest.
func Annotations(f partial.WithRawManifest, anns map[string]string) partial.WithRawManifest {
	_ = "STUB: not implemented"
	return *new(partial.WithRawManifest)
}

type arbitraryRawManifest struct {
	a       partial.WithRawManifest
	anns    map[string]string
	subject *v1.Descriptor
}

func (a arbitraryRawManifest) RawManifest() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConfigFile mutates the provided v1.Image to have the provided v1.ConfigFile
func ConfigFile(base v1.Image, cfg *v1.ConfigFile) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// CreatedAt mutates the provided v1.Image to have the provided v1.Time
func CreatedAt(base v1.Image, created v1.Time) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// Extract takes an image and returns an io.ReadCloser containing the image's
// flattened filesystem.
//
// Callers can read the filesystem contents by passing the reader to
// tar.NewReader, or io.Copy it directly to some output.
//
// If a caller doesn't read the full contents, they should Close it to free up
// resources used during extraction.
func Extract(img v1.Image) io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

// Close the writer with any errors encountered during
// extraction. These errors will be returned by the reader end
// on subsequent reads. If err == nil, the reader will return
// EOF.

// Adapted from https://github.com/google/containerregistry/blob/da03b395ccdc4e149e34fbb540483efce962dc64/client/v2_2/docker_image_.py#L816
func extract(img v1.Image, w io.Writer) error { _ = "STUB: not implemented"; return nil }

// we iterate through the layers in reverse order because it makes handling
// whiteout layers more efficient, since we can just keep track of the removed
// files as we see .wh. layers and ignore those in previous layers.

func extractLayer(tarWriter *tar.Writer, fileMap map[string]bool, layer v1.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// Some tools prepend everything with "./", so if we don't Clean the
// name, we may have duplicate entries, which angers tar-split.

// Reject relative symlinks and hardlinks whose targets escape the
// image rootfs. Relative targets are resolved against the symlink's
// own directory: if the clean result starts with ".." the link would
// leave the rootfs. Relative symlinks that stay within the rootfs
// (common for glibc, C toolchains, etc.) are preserved unchanged.
// Absolute targets are left as-is; see #2238 for ongoing discussion
// on whether they should be pruned.

//nolint:gosec // G305: path is only used for validation, not file I/O

// force PAX format to remove Name/Linkname length limit of 100 characters
// required by USTAR and to not depend on internal tar package guess which
// prefers USTAR over PAX

// check if we have seen value before
// if we're checking a directory, don't filepath.Join names

// check for a whited out parent directory

// mark file as handled. non-directory implicitly tombstones
// any entries with a matching (or child) name

// Drain any bytes the tar.Reader did not consume (trailing data after the
// end-of-archive marker) so the underlying verifying reader reaches io.EOF
// and the layer's digest is verified. Without this, a layer whose contents
// do not match the manifest's layer digest is extracted without error.
// pkg/v1/validate/layer.go performs the same drain.

func inWhiteoutDir(fileMap map[string]bool, file string) bool {
	_ = "STUB: not implemented"
	return false
}

// Time sets all timestamps in an image to the given timestamp.
func Time(img v1.Image, t time.Time) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// try to search for the history entry that corresponds to this layer

// if it's an EmptyLayer, do not set the Layer and have the Addendum with just the History
// and move on to the next History entry

// otherwise, we can exit from the cycle

// add all leftover History entries

// Copy basic config over

// Strip away timestamps from the config file

// Explicitly ignore Author field; which hinders reproducibility

func layerTime(layer v1.Layer, t time.Time) (v1.Layer, error) {
	_ = "STUB: not implemented"
	return *new(v1.Layer), nil
}

//PAX and GNU Format support additional timestamps in the header

// TODO(#1168): This should be lazy, and not buffer the entire layer contents.

// Drain trailing bytes so the underlying verifying reader reaches io.EOF
// and the layer digest is verified (see extractLayer).

// gzip the contents, then create the layer

// Canonical is a helper function to combine Time and configFile
// to remove any randomness during a docker build.
func Canonical(img v1.Image) (v1.Image, error) {
	_ = "STUB: not implemented"
	// Set all timestamps to 0
	return *new(v1.Image), nil
}

// Get rid of host-dependent random config

//nolint:staticcheck // Field will be removed in next release

// MediaType modifies the MediaType() of the given image.
func MediaType(img v1.Image, mt types.MediaType) v1.Image {
	_ = "STUB: not implemented"
	return *new(v1.Image)
}

// ConfigMediaType modifies the MediaType() of the given image's Config.
//
// If !mt.IsConfig(), this will be the image's artifactType in any indexes it's a part of.
func ConfigMediaType(img v1.Image, mt types.MediaType) v1.Image {
	_ = "STUB: not implemented"
	return *new(v1.Image)
}

// IndexMediaType modifies the MediaType() of the given index.
func IndexMediaType(idx v1.ImageIndex, mt types.MediaType) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}
