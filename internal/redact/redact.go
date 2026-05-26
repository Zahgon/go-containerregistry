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

// Package redact contains a simple context signal for redacting requests.
package redact

import (
	"context"
	"net/url"
)

type contextKey string

var redactKey = contextKey("redact")

// NewContext creates a new ctx with the reason for redaction.
func NewContext(ctx context.Context, reason string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// FromContext returns the redaction reason, if any.
func FromContext(ctx context.Context) (bool, string) { _ = "STUB: not implemented"; return false, "" }

// Error redacts potentially sensitive query parameter values in the URL from the error's message.
//
// If the error is a *url.Error, this returns a *url.Error with the URL redacted.
// Any other error type, or nil, is returned unchanged.
func Error(err error) error {
	_ = "STUB: not implemented"
	// If the error is a url.Error, we can redact the URL.
	// Otherwise (including if err is nil), we can't redact.
	return nil
}

// If the URL can't be parsed, just return the original error.

// Update the URL to the redacted URL.

// The set of query string keys that we expect to send as part of the registry
// protocol. Anything else is potentially dangerous to leak, as it's probably
// from a redirect. These redirects often included tokens or signed URLs.
var paramAllowlist = map[string]struct{}{
	// Token exchange
	"scope":   {},
	"service": {},
	// Cross-repo mounting
	"mount": {},
	"from":  {},
	// Layer PUT
	"digest": {},
	// Listing tags and catalog
	"n":    {},
	"last": {},
}

// URL redacts potentially sensitive query parameter values from the URL's query string.
func URL(u *url.URL) string { _ = "STUB: not implemented"; return "" }

// key is not in the Allowlist
