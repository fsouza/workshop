// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hands

func Map(values []int, fn func(int) int) []int {
	result := make([]int, len(values))
	for i, v := range values {
		result[i] = fn(v)
	}
	return result
}
