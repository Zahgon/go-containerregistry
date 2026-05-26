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

// Package depcheck defines a test utility for ensuring certain packages don't
// take on heavy dependencies.
//
// This is forked from https://pkg.go.dev/knative.dev/pkg/depcheck
package depcheck

import (
	"testing"
)

type node struct {
	importpath string
	consumers  map[string]struct{}
}

type graph map[string]node

func (g graph) contains(name string) bool { _ = "STUB: not implemented"; return false }

func (g graph) order() []string { _ = "STUB: not implemented"; return nil }

// path constructs an examplary path that looks something like:
//
//	knative.dev/pkg/apis/duck
//	knative.dev/pkg/apis  # Also: [knative.dev/pkg/kmeta knative.dev/pkg/tracker]
//	k8s.io/api/core/v1
func (g graph) path(name string) []string {
	_ = "STUB: not implemented"

	// Base case.
	return nil
}

// Inductive step.

// Don't decorate the first entry, which is always an entrypoint.

// Attach other consumers to the last entry in base.

func buildGraph(importpath string, buildFlags ...string) (graph, error) {
	_ = "STUB: not implemented"
	return *new(graph), nil
}

// StdlibPackages returns the list of all standard library packages, including
// some golang.org/x/ dependencies.
func StdlibPackages() []string {
	_ = "STUB: not implemented"
	// pkg/registry is allowed to depend on any stdlib package, so collect
	// all of those -- this also includes golang.org/x/ packages.
	return nil
}

// CheckNoDependency checks that the given import paths (ip) does not
// depend (transitively) on certain banned imports.
func CheckNoDependency(ip string, banned []string, buildFlags ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// AssertNoDependency checks that the given import paths (the keys) do not
// depend (transitively) on certain banned imports (the values)
func AssertNoDependency(t *testing.T, banned map[string][]string, buildFlags ...string) {
	_ = "STUB: not implemented"
	return
}

// AssertOnlyDependencies checks that the given import paths (the keys) only
// depend (transitively) on certain allowed imports (the values).
// Note: while perhaps counterintuitive we allow the value to be a superset
// of the actual imports to that folks can use a constant that holds blessed
// import paths.
func AssertOnlyDependencies(t *testing.T, allowed map[string][]string, buildFlags ...string) {
	_ = "STUB: not implemented"
	return
}

// Always include our own package in the set of allowed dependencies.

// CheckOnlyDependencies checks that the given import path only
// depends (transitively) on certain allowed imports.
// Note: while perhaps counterintuitive we allow the value to be a superset
// of the actual imports to that folks can use a constant that holds blessed
// import paths.
func CheckOnlyDependencies(ip string, allowed map[string]struct{}, buildFlags ...string) error {
	_ = "STUB: not implemented"
	return nil
}
