// Copyright 2013 Francisco Souza. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

func main() {
	var values []int
	values = []int{10, 15, 18}
	values = values[:2]
	for i := 0; i < len(values); i++ {
		fmt.Printf("values[%d]: %d.\n", i, values[i])
	}

	for i, value := range values {
		fmt.Printf("values[%d]: %d.\n", i, value)
	}
	values = make([]int, 3)
	values[0] = 10
	values[1] = 15
	values[2] = 18
	values = append(values, 20, 30, 40, 50)
	for i, value := range values {
		fmt.Printf("values[%d]: %d.\n", i, value)
	}

	var array [16]int
	values = array[:]
	for i, value := range array {
		fmt.Printf("array[%d]: %d.\n", i, value)
	}

	var count map[string]int
	count = make(map[string]int)
	count["x"] = 1
	count["y"] = 1
	count["x"] += 1
	for key, value := range count {
		fmt.Printf("count[%q]: %d.\n", key, value)
	}


	values = make([]int, 0, 6)
	values = append(values, 10, 2, 4, 55, 82, 3)
	values = append(values, 2)
}
