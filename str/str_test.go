// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

import (
	"testing"
)

func TestMyStringReader(t *testing.T) {
	buf := make([]byte, 8)
	s := MyString("Francisco")
	n, err := s.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(buf) {
		t.Errorf("MyString.Read() returned wrong int. Want %d. Got %d.", len(buf), n)
	}
	if string(buf) != string(s[:8]) {
		t.Errorf("MyString.Read() wrote wrong content. Want %q. Got %q.", s[:8], string(buf))
	}
	buf = make([]byte, 8)
	s = MyString("Chico")
	n, err = s.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(s) {
		t.Errorf("MyString.Read() returned wrong int. Want %d. Got %d.", len(s), n)
	}
	if string(buf[:n]) != string(s) {
		t.Errorf("MyString.Read() wrote wrong content. Want %q. Got %q.", s, string(buf))
	}
}

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

func BenchmarkReverse(b *testing.B) {
	b.StopTimer()
	inputs := []string{
		"abjure", "abrogate", "abstemious", "acumen", "antebellum",
		"auspicious", "belie", "bellicose", "bowdlerize", "chicanery",
		"chromosome", "churlish", "circumlocution", "circumnavigate",
		"deciduous", "deleterious", "diffident", "enervate", "enfranchise",
		"epiphany", "equinox", "euro", "evanescent", "expurgate", "facetious",
		"fatuous", "feckless", "fiduciary", "filibuster", "gamete", "gauche",
		"gerrymander", "hegemony", "hemoglobin", "homogeneous", "hubris",
		"hypotenuse", "impeach", "incognito", "incontrovertible", "inculcate",
		"infrastructure", "interpolate", "irony", "jejune", "kinetic",
		"kowtow", "laissez faire", "lexicon", "loquacious",
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Reverse(inputs[i%len(inputs)])
	}
}
