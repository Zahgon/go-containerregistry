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
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/spf13/cobra"
)

// NewCmdMutate creates a new cobra.Command for the mutate subcommand.
func NewCmdMutate(options *[]crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

// Pull image and get config.

// Set labels.

// set envvars if specified

// Set entrypoint.

// This matches Docker's behavior.

// Set cmd.

// Set user.

// Set workdir.

// Set ports

// Set platform

// Mutate and write image.

// If the new ref isn't provided, write over the original image.
// If that ref was provided by digest (e.g., output from
// another crane command), then strip that and push the
// mutated image by digest instead.

// Using "set-platform" to avoid clobbering "platform" persistent flag.

// validateKeyVals ensures no values are empty, returns error if they are
func validateKeyVals(kvPairs map[string]string) error { _ = "STUB: not implemented"; return nil }

// setEnvVars override envvars in a config
func setEnvVars(cfg *v1.ConfigFile, envVars keyToValue) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep the old values.

// Override in place to keep ordering of original env.

// Remove this from eMap so we don't add it twice.

// Append the new values.

// If we come across a value not in eMap, it means we replaced the
// old env in-place and deleted it from eMap, so we can skip adding.

type env struct {
	key   string
	value string
}

type keyToValue struct {
	values  []env
	changed bool
	mapped  map[string]string
}

func (o *keyToValue) Set(val string) error { _ = "STUB: not implemented"; return nil }

func (o *keyToValue) Type() string { _ = "STUB: not implemented"; return "" }

func (o *keyToValue) String() string { _ = "STUB: not implemented"; return "" }

func (o *keyToValue) Map() map[string]string { _ = "STUB: not implemented"; return nil }
