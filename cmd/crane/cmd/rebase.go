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

package cmd

import (
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/spf13/cobra"
)

// NewCmdRebase creates a new cobra.Command for the rebase subcommand.
func NewCmdRebase(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// If the new ref isn't provided, write over the original image.
// If that ref was provided by digest (e.g., output from
// another crane command), then strip that and push the
// mutated image by digest instead.

// Stupid hack to support insecure flag.

// rebaseImage parses the references and uses them to perform a rebase on the
// original image.
//
// If oldBase or newBase are "", rebaseImage attempts to derive them using
// annotations in the original image. If those annotations are not found,
// rebaseImage returns an error.
//
// If rebasing is successful, base image annotations are set on the resulting
// image to facilitate implicit rebasing next time.
func rebaseImage(orig v1.Image, oldBase, newBase string, opt ...crane.Option) (v1.Image, error) {
	_ = "STUB: not implemented"
	return *new(v1.Image), nil
}

// NB: if newBase is an index, we need to grab the index's digest to
// annotate the resulting image, even though we pull the
// platform-specific image to rebase.
// crane.Digest will pull a platform-specific image, so use crane.Head
// here instead.

// Update base image annotations for the new image manifest.
