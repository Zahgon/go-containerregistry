// Copyright 2020 Google LLC All Rights Reserved.
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
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
)

// NewCmdAuth creates a new cobra.Command for the auth subcommand.
func NewCmdAuth(options []crane.Option, argv ...string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

func NewCmdAuthToken(options []crane.Option) *cobra.Command { _ = "STUB: not implemented"; return nil }

type credentials struct {
	Username string `json:"Username,omitempty"`
	Secret   string `json:"Secret,omitempty"`
}

// https://github.com/docker/cli/blob/2291f610ae73533e6e0749d4ef1e360149b1e46b/cli/config/credentials/native_store.go#L100-L109
func toCreds(config *authn.AuthConfig) credentials {
	_ = "STUB: not implemented"
	return *new(credentials)
}

// NewCmdAuthGet creates a new `crane auth get` command.
func NewCmdAuthGet(options []crane.Option, argv ...string) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// If we don't find any credentials, there's a magic error to return:
//
// https://github.com/docker/docker-credential-helpers/blob/f78081d1f7fef6ad74ad6b79368de6348386e591/credentials/error.go#L4-L6
// https://github.com/docker/docker-credential-helpers/blob/f78081d1f7fef6ad74ad6b79368de6348386e591/credentials/credentials.go#L61-L63

// Convert back to a form that credential helpers can parse so that this
// can act as a meta credential helper.

// NewCmdAuthLogin creates a new `crane auth login` command.
func NewCmdAuthLogin(argv ...string) *cobra.Command { _ = "STUB: not implemented"; return nil }

type loginOptions struct {
	serverAddress string
	user          string
	password      string
	passwordStdin bool
}

func login(opts loginOptions) error { _ = "STUB: not implemented"; return nil }

// NewCmdAuthLogout creates a new `crane auth logout` command.
func NewCmdAuthLogout(argv ...string) *cobra.Command { _ = "STUB: not implemented"; return nil }
