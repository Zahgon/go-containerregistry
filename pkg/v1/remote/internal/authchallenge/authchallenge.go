// Copyright 2014 Docker, Inc.
// Copyright 2021-2026 The Distribution contributors
// Copyright 2026 Google LLC All Rights Reserved.
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

package authchallenge

import (
	"net/http"
	"strings"
)

// Octet types from RFC 2616.
type octetType byte

var octetTypes [256]octetType

const (
	isToken octetType = 1 << iota
	isSpace
)

func init() {
	// OCTET      = <any 8-bit sequence of data>
	// CHAR       = <any US-ASCII character (octets 0 - 127)>
	// CTL        = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
	// CR         = <US-ASCII CR, carriage return (13)>
	// LF         = <US-ASCII LF, linefeed (10)>
	// SP         = <US-ASCII SP, space (32)>
	// HT         = <US-ASCII HT, horizontal-tab (9)>
	// <">        = <US-ASCII double-quote mark (34)>
	// CRLF       = CR LF
	// LWS        = [CRLF] 1*( SP | HT )
	// TEXT       = <any OCTET except CTLs, but including LWS>
	// separators = "(" | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\" | <">
	//              | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP | HT
	// token      = 1*<any CHAR except CTLs or separators>
	// qdtext     = <any TEXT except <">>

	for c := range 256 {
		var t octetType
		isCtl := c <= 31 || c == 127
		isChar := 0 <= c && c <= 127
		isSeparator := strings.ContainsRune(" \t\"(),/:;<=>?@[]\\{}", rune(c))
		if strings.ContainsRune(" \t\r\n", rune(c)) {
			t |= isSpace
		}
		if isChar && !isCtl && !isSeparator {
			t |= isToken
		}
		octetTypes[c] = t
	}
}

// Challenge carries information from a WWW-Authenticate response header.
// See RFC 2617.
type Challenge struct {
	// Scheme is the auth-scheme according to RFC 2617
	Scheme string

	// Parameters are the auth-params according to RFC 2617
	Parameters map[string]string
}

// ResponseChallenges returns a list of authorization challenges
// for the given http Response. Challenges are only checked if
// the response status code was a 401.
func ResponseChallenges(resp *http.Response) []Challenge { _ = "STUB: not implemented"; return nil }

// Parse the WWW-Authenticate Header and store the challenges
// on this endpoint object.

func parseAuthHeader(header http.Header) []Challenge { _ = "STUB: not implemented"; return nil }

func parseValueAndParams(header string) (value string, params map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

func skipSpace(s string) (rest string) { _ = "STUB: not implemented"; return "" }

func expectToken(s string) (token, rest string) { _ = "STUB: not implemented"; return "", "" }

func expectTokenOrQuoted(s string) (value string, rest string) {
	_ = "STUB: not implemented"
	return "", ""
}
