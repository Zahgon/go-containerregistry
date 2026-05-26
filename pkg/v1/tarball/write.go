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

package tarball

import (
	"archive/tar"
	"io"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

// WriteToFile writes in the compressed format to a tarball, on disk.
// This is just syntactic sugar wrapping tarball.Write with a new file.
func WriteToFile(p string, ref name.Reference, img v1.Image, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiWriteToFile writes in the compressed format to a tarball, on disk.
// This is just syntactic sugar wrapping tarball.MultiWrite with a new file.
func MultiWriteToFile(p string, tagToImage map[name.Tag]v1.Image, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiRefWriteToFile writes in the compressed format to a tarball, on disk.
// This is just syntactic sugar wrapping tarball.MultiRefWrite with a new file.
func MultiRefWriteToFile(p string, refToImage map[name.Reference]v1.Image, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Write is a wrapper to write a single image and tag to a tarball.
func Write(ref name.Reference, img v1.Image, w io.Writer, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiWrite writes the contents of each image to the provided writer, in the compressed format.
// The contents are written in the following format:
// One manifest.json file at the top level containing information about several images.
// One file for each layer, named after the layer's SHA.
// One file for the config blob, named after its SHA.
func MultiWrite(tagToImage map[name.Tag]v1.Image, w io.Writer, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	return nil
}

// MultiRefWrite writes the contents of each image to the provided writer, in the compressed format.
// The contents are written in the following format:
// One manifest.json file at the top level containing information about several images.
// One file for each layer, named after the layer's SHA.
// One file for the config blob, named after its SHA.
func MultiRefWrite(refToImage map[name.Reference]v1.Image, w io.Writer, opts ...WriteOption) error {
	_ = "STUB: not implemented"
	// process options
	return nil
}

// sendUpdateReturn return the passed in error message, also sending on update channel, if it exists
func sendUpdateReturn(o *writeOptions, err error) error { _ = "STUB: not implemented"; return nil }

// sendProgressWriterReturn return the passed in error message, also sending on update channel, if it exists, along with downloaded information
func sendProgressWriterReturn(pw *progressWriter, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// writeImagesToTar writes the images to the tarball
func writeImagesToTar(imageToTags map[v1.Image][]string, m []byte, size int64, w io.Writer, o *writeOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// we only calculate the sizes and use a progressWriter if we were provided
// an option with a progress channel

// Write the config.

// Write the layers.

// Munge the file name to appease ancient technology.
//
// tar assumes anything with a colon is a remote tape drive:
// https://www.gnu.org/software/tar/manual/html_section/tar_45.html
// Drop the algorithm prefix, e.g. "sha256:"

// gunzip expects certain file extensions:
// https://www.gnu.org/software/gzip/manual/html_node/Overview.html

// be sure to close the tar writer so everything is flushed out before we send our EOF

// send an EOF to indicate finished on the channel, but nil as our return error

// calculateManifest calculates the manifest and optionally the size of the tar file
func calculateManifest(imageToTags map[v1.Image][]string) (m Manifest, err error) {
	_ = "STUB: not implemented"
	return *new(Manifest), nil
}

// Store foreign layer info.

// Write the layers.

// Munge the file name to appease ancient technology.
//
// tar assumes anything with a colon is a remote tape drive:
// https://www.gnu.org/software/tar/manual/html_section/tar_45.html
// Drop the algorithm prefix, e.g. "sha256:"

// gunzip expects certain file extensions:
// https://www.gnu.org/software/gzip/manual/html_node/Overview.html

// Add to LayerSources if it's a foreign layer.

// Generate the tar descriptor and write it.

// sort by name of the repotags so it is consistent. Alternatively, we could sort by hash of the
// descriptor, but that would make it hard for humans to process

// CalculateSize calculates the expected complete size of the output tar file
func CalculateSize(refToImage map[name.Reference]v1.Image) (size int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func getSizeAndManifest(imageToTags map[v1.Image][]string) (int64, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// calculateTarballSize calculates the size of the tar file
func calculateTarballSize(imageToTags map[v1.Image][]string, mBytes []byte) (size int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// add the manifest

// add the two padding blocks that indicate end of a tar file

func dedupRefToImage(refToImage map[name.Reference]v1.Image) map[v1.Image][]string {
	_ = "STUB: not implemented"
	return nil
}

// Docker cannot load tarballs without an explicit tag:
// https://github.com/google/go-containerregistry/issues/890
//
// We can't use the fully qualified tag.Name() because of rules_docker:
// https://github.com/google/go-containerregistry/issues/527
//
// If the tag is "latest", but tag.String() doesn't end in ":latest",
// just append it. Kind of gross, but should work for now.

// writeTarEntry writes a file to the provided writer with a corresponding tar header
func writeTarEntry(tf *tar.Writer, path string, r io.Reader, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

// writeLayer streams a layer's compressed blob into the tar writer and closes
// the reader before returning. The close releases any pull-limiter slot held
// by a remote-backed layer (remote.WithJobs); leaving it open would deadlock
// the write loop after defaultJobs layers.
func writeLayer(tf *tar.Writer, name string, l v1.Layer) error {
	_ = "STUB: not implemented"
	return nil
}

// ComputeManifest get the manifest.json that will be written to the tarball
// for multiple references
func ComputeManifest(refToImage map[name.Reference]v1.Image) (Manifest, error) {
	_ = "STUB: not implemented"
	return *new(Manifest), nil
}

// WriteOption a function option to pass to Write()
type WriteOption func(*writeOptions) error
type writeOptions struct {
	updates chan<- v1.Update
}

// WithProgress create a WriteOption for passing to Write() that enables
// a channel to receive updates as they are downloaded and written to disk.
func WithProgress(updates chan<- v1.Update) WriteOption {
	_ = "STUB: not implemented"
	return *new(WriteOption)
}

// progressWriter is a writer which will send the download progress
type progressWriter struct {
	w              io.Writer
	updates        chan<- v1.Update
	size, complete int64
}

func (pw *progressWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (pw *progressWriter) Error(err error) error { _ = "STUB: not implemented"; return nil }

func (pw *progressWriter) Close() error { _ = "STUB: not implemented"; return nil }

// calculateSingleFileInTarSize calculate the size a file will take up in a tar archive,
// given the input data. Provided by rounding up to nearest whole block (512)
// and adding header 512
func calculateSingleFileInTarSize(in int64) (out int64) {
	_ = "STUB: not implemented"
	// doing this manually, because math.Round() works with float64
	return 0
}
