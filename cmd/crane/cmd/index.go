// Copyright 2023 Google LLC All Rights Reserved.
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
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/spf13/cobra"
)

// NewCmdIndex creates a new cobra.Command for the index subcommand.
func NewCmdIndex(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// NewCmdIndexList creates a new cobra.Command for the index list subcommand.
func NewCmdIndexList(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Print header

// NewCmdIndexFilter creates a new cobra.Command for the index filter subcommand.
func NewCmdIndexFilter(options *[]crane.Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// If we can't get the image index, we just return nil to silently stop.

// Consider reusing the persistent flag for this, it's separate so we can have multiple values.

// NewCmdIndexAppend creates a new cobra.Command for the index append subcommand.
func NewCmdIndexAppend(options *[]crane.Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func appendLocal(baseRef string, adds []mutate.IndexAddendum) error {
	_ = "STUB: not implemented"
	return nil
}

// If the path exists, try to open it as an OCI layout.

// If the path does not exist, initialize a new OCI layout.

func appendRemote(cmd *cobra.Command, baseRef, newTag string, adds []mutate.IndexAddendum, o crane.Options, dockerEmptyBase bool) error {
	_ = "STUB: not implemented"
	return nil
}

func filterIndex(idx v1.ImageIndex, platforms []v1.Platform) v1.ImageIndex {
	_ = "STUB: not implemented"
	return *new(v1.ImageIndex)
}

func satisfiesPlatforms(platforms []v1.Platform) match.Matcher {
	_ = "STUB: not implemented"
	return *new(match.Matcher)
}

func not(in match.Matcher) match.Matcher { _ = "STUB: not implemented"; return *new(match.Matcher) }

// resolveManifest resolves a reference to either a local OCI layout or a remote manifest.
// It returns the manifest (Image or ImageIndex) and its descriptor.
func resolveManifest(ref string, o crane.Options) (partial.WithRawManifest, v1.Descriptor, error) {
	_ = "STUB: not implemented"
	// Try loading as local image first. If this fails, we fall back to remote.
	return *new(partial.WithRawManifest), *new(v1.Descriptor), nil
}

// If it's an image, try to get platform from config.
// Platform info is not in the manifest, but is required for the index descriptor.

// Fallback to remote

// Populate platform info from the config blob for the index descriptor.

// collectAddendums resolves a list of manifest references (local or remote) and
// returns a slice of IndexAddendums that can be appended to an index.
// If flatten is true, any referenced indices will have their individual child
// manifests added to the result rather than the index itself.
func collectAddendums(manifests []string, o crane.Options, flatten bool) ([]mutate.IndexAddendum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// isLocalReference checks if the reference is a local file path.
// It returns true if the path exists on the filesystem.
// If the path does NOT exist, it still returns true if it looks like a path
// (starts with . or / or \ on Windows). This is necessary because
// 'crane index append' can create new local OCI layout directories that
// do not yet exist, and we need to distinguish these from remote references.
func isLocalReference(ref string) bool { _ = "STUB: not implemented"; return false }
