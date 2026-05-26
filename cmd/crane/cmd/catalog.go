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
	"context"
	"io"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
)

// NewCmdCatalog creates a new cobra.Command for the catalog subcommand.
func NewCmdCatalog(options *[]crane.Option, _ ...string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func catalog(ctx context.Context, w io.Writer, src string, fullRef bool, o crane.Options) error {
	_ = "STUB: not implemented"
	return nil
}
