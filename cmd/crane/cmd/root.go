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
	"net/http"
	"sync"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/spf13/cobra"
)

const (
	use   = "crane"
	short = "Crane is a tool for managing container images"
)

var Root = New(use, short, []crane.Option{})

// New returns a top-level command for crane. This is mostly exposed
// to share code with gcrane.
func New(use, short string, options []crane.Option) *cobra.Command {
	_ = "STUB: not implemented"
	return nil
}

// TODO(jonjohnsonjr): crane.Verbose option?

//nolint: gosec

// Add any http headers if they are set in the config file.

// Inject our warning-collecting transport.

// Report any collected warnings.

// headerTransport sets headers on outgoing requests.
type headerTransport struct {
	httpHeaders map[string]string
	inner       http.RoundTripper
}

// RoundTrip implements http.RoundTripper.
func (ht *headerTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Docker sets this, which is annoying, since we're not docker.
// We might want to revisit completely ignoring this.

type warnTransport struct {
	mu    sync.Mutex
	warns map[string]struct{}
	inner http.RoundTripper
}

func (wt *warnTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Warning response headers are supposed to have
// warn-code 299 and warn-agent "-"; discard these.

func (wt *warnTransport) Report() { _ = "STUB: not implemented"; return }

// TODO: Consider using logs.Warn here if we move this out of crane.

func nocolor() bool {
	_ = "STUB: not implemented"
	// These adapted from https://github.com/kubernetes/kubernetes/blob/fe91bc257b505eb6057eb50b9c550a7c63e9fb91/staging/src/k8s.io/kubectl/pkg/util/term/term.go
	return false
}

// https://en.wikipedia.org/wiki/Computer_terminal#Dumb_terminals

// https://no-color.org/

// On Windows WT_SESSION is set by the modern terminal component.
// Older terminals have poor support for UTF-8, VT escape codes, etc.
