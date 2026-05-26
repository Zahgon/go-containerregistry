// Copyright 2021 Google LLC All Rights Reserved.
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
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/spf13/cobra"
)

// NewCmdFlatten creates a new cobra.Command for the flatten subcommand.
func NewCmdFlatten(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// We need direct access to the underlying remote options because crane
// doesn't expose great facilities for working with an index (yet).

// Pull image and get config.

// If the new ref isn't provided, write over the original image.
// If that ref was provided by digest (e.g., output from
// another crane command), then strip that and push the
// mutated image by digest instead.

func flatten(ref name.Reference, repo name.Repository, use string, o crane.Options) (partial.Describable, error) {
	_ = "STUB: not implemented"
	return *new(partial.Describable), nil
}

func push(flat partial.Describable, ref name.Reference, o crane.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func flattenIndex(old v1.ImageIndex, repo name.Repository, use string, o crane.Options) (partial.Describable, error) {
	_ = "STUB: not implemented"
	return *new(partial.Describable), nil
}

// Keep the old descriptor (annotations and whatnot).

// Drop attestations (for now).
// https://github.com/google/go-containerregistry/issues/1622

// Retain any annotations from the original index.

// This is stupid, but some registries get mad if you try to push OCI media types that reference docker media types.

func flattenChild(old partial.Describable, repo name.Repository, use string, o crane.Options) (partial.Describable, error) {
	_ = "STUB: not implemented"
	return *new(partial.Describable), nil
}

func flattenImage(old v1.Image, repo name.Repository, use string, o crane.Options) (partial.Describable, error) {
	_ = "STUB: not implemented"
	return *new(partial.Describable), nil
}

// Clear layer-specific config file information.

// TODO: Make compression configurable?

// Retain any annotations from the original image.

// Propagate the original media type (e.g. OCI vs Docker) so that all
// manifests in an index use a consistent media type family. Without this,
// an OCI image index would reference Docker-typed image manifests, which
// confuses registries and tooling that assumes the index and its children
// share the same media-type convention.
