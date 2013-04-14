// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hands

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{2, 4, 6}
	output := Map(input, func(i int) int {
		return i * 2
	})
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Map: Want %#v. Got %#v.", expected, output)
	}
}
