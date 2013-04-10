// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"Francisco", "ocsicnarF"},
		{"ovo", "ovo"},
		{"a", "a"},
		{"soma", "amos"},
	}
	for _, tt := range tests {
		got := Reverse(tt.input)
		if got != tt.want {
			t.Errorf("Reverse(%q): Want %q. Got %q.", tt.input, tt.want, got)
		}
	}
}
