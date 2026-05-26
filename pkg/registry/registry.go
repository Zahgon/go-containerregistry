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

// Package registry implements a docker V2 registry and the OCI distribution specification.
//
// It is designed to be used anywhere a low dependency container registry is needed, with an
// initial focus on tests.
//
// Its goal is to be standards compliant and its strictness will increase over time.
//
// This is currently a low flightmiles system. It's likely quite safe to use in tests; If you're using it
// in production, please let us know how and send us CL's for integration tests.
package registry

import (
	"log"
	"net/http"
)

type registry struct {
	log              *log.Logger
	blobs            blobs
	manifests        manifests
	referrersEnabled bool
	warnings         map[float64]string
}

// https://docs.docker.com/registry/spec/api/#api-version-check
// https://github.com/opencontainers/distribution-spec/blob/master/spec.md#api-version-check
func (r *registry) v2(resp http.ResponseWriter, req *http.Request) *regError {
	_ = "STUB: not implemented"
	return nil
}

func (r *registry) root(resp http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

// New returns a handler which implements the docker registry protocol.
// It should be registered at the site root.
func New(opts ...Option) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// Option describes the available options
// for creating the registry.
type Option func(r *registry)

// Logger overrides the logger used to record requests to the registry.
func Logger(l *log.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithReferrersSupport enables the referrers API endpoint (OCI 1.1+)
func WithReferrersSupport(enabled bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWarning(prob float64, msg string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBlobHandler(h BlobHandler) Option { _ = "STUB: not implemented"; return *new(Option) }
