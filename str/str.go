// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package str

// Reverse returns the given string in the reverse order.
func Reverse(input string) string {
	other := make([]byte, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		other[i] = input[len(input)-1-i]
	}
	return string(other)
}
