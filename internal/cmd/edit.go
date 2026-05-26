// Copyright 2022 Google LLC All Rights Reserved.
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

package cmd

import (
	"archive/tar"
	"context"
	"io"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/spf13/cobra"
)

// NewCmdEdit creates a new cobra.Command for the edit subcommand.
//
// This is currently hidden until we're happy with the interface and can test
// it on different operating systems and editors.
func NewCmdEdit(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdConfig creates a new cobra.Command for the config subcommand.
func NewCmdEditConfig(options *[]crane.Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// NewCmdManifest creates a new cobra.Command for the manifest subcommand.
func NewCmdEditManifest(options *[]crane.Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// NewCmdExport creates a new cobra.Command for the export subcommand.
func NewCmdEditFs(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

func interactive(in io.Reader, out io.Writer) bool { _ = "STUB: not implemented"; return false }

func interactiveFile(i any) bool { _ = "STUB: not implemented"; return false }

func editConfig(ctx context.Context, in io.Reader, out io.Writer, src, dst string, options ...crane.Option) (name.Reference, error) {
	_ = "STUB: not implemented"
	return *new(name.Reference), nil
}

// We want to omit Layers in certain situations, so we don't use v1.Image.Manifest() here.
// Instead, we treat the manifest as a map[string]any and just manipulate the config desc.

// this has to happen before we modify the descriptor (so we can use verify.Descriptor to validate whether m.Config.Data matches m.Config.Digest/Size)

// https://github.com/google/go-containerregistry/issues/1552#issuecomment-1452653875
// "if data is non-empty and correct, we should update it"

func editManifest(in io.Reader, out io.Writer, src string, dst string, mt string, options ...crane.Option) (name.Reference, error) {
	_ = "STUB: not implemented"
	return *new(name.Reference), nil
}

// If --media-type is unset, use Content-Type by default.

// If document contains mediaType, default to that.

func editFile(in io.Reader, out io.Writer, src, file, dst string, options ...crane.Option) (name.Reference, error) {
	_ = "STUB: not implemented"
	return *new(name.Reference), nil
}

// If stdin has content, read it in and use that for the file.
// Otherwise, scran through the image and open that file in an editor.

func findFile(img v1.Image, name string) (io.Reader, *tar.Header, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil, nil
}

// If we don't find the file, we should create a new one.

func blankHeader(name string) *tar.Header { _ = "STUB: not implemented"; return nil }

// Use a fixed Mode, so that this isn't sensitive to the directory and umask
// under which it was created. Additionally, windows can only set 0222,
// 0444, or 0666, none of which are executable.

func normalize(name string) string { _ = "STUB: not implemented"; return "" }

type withMediaType struct {
	MediaType string `json:"mediaType,omitempty"`
}

type rawManifest struct {
	body      []byte
	mediaType types.MediaType
}

func (r *rawManifest) RawManifest() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *rawManifest) MediaType() (types.MediaType, error) {
	_ = "STUB: not implemented"
	return *new(types.MediaType), nil
}
