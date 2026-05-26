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

package cmd

import (
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Version can be set via:
// -ldflags="-X 'github.com/google/go-containerregistry/cmd/crane/cmd.Version=$TAG'"
var Version string

func init() {
	if Version == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		Version = i.Main.Version
	}
}

// NewCmdVersion creates a new cobra.Command for the version subcommand.
func NewCmdVersion() *cobra.Command { _ = "STUB: not implemented"; return nil }
