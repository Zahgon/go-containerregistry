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
	"context"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/spf13/cobra"
)

// NewCmdGc creates a new cobra.Command for the gc subcommand.
func NewCmdGc() *cobra.Command { _ = "STUB: not implemented"; return nil }

func gc(ctx context.Context, root string, recursive bool) error {
	_ = "STUB: not implemented"
	return nil
}

func printUntaggedImages(repo name.Repository, tags *google.Tags, err error) error {
	_ = "STUB: not implemented"
	return nil
}
