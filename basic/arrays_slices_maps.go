// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main() {
	var values []int
	values = []int{10, 15, 18}
	values = values[:2]

	var array [16]int
	values = array[:]

	var count map[string]int
	count = make(map[string]int)
	count["x"] = 1
	count["y"] = 1
	count["x"] += 1
}
