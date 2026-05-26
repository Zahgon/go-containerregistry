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
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type platformsValue struct {
	platforms []v1.Platform
}

func (ps *platformsValue) Set(platform string) error { _ = "STUB: not implemented"; return nil }

func (ps *platformsValue) String() string { _ = "STUB: not implemented"; return "" }

func (ps *platformsValue) Type() string { _ = "STUB: not implemented"; return "" }

type platformValue struct {
	platform *v1.Platform
}

func (pv *platformValue) Set(platform string) error { _ = "STUB: not implemented"; return nil }

func (pv *platformValue) String() string { _ = "STUB: not implemented"; return "" }

func (pv *platformValue) Type() string { _ = "STUB: not implemented"; return "" }

func platformToString(p *v1.Platform) string { _ = "STUB: not implemented"; return "" }

func parsePlatform(platform string) (*v1.Platform, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
